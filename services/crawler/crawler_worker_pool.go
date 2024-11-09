package crawler

import (
	"log"
	"sync"
)

type WorkerPoll struct {
	mainLink     string
	crawler      CrawlerInterface
	links        []string
	numWorkers   int
	results      []Result
	errors       []Result
	done         chan bool
	jobsQueue    chan Task
	resultsQueue chan Result
}

type Task struct {
	Link string
}

type Result struct {
	error error
	ad    *Ad
}

// Dispatcher: enqueues tasks into the jobs queue and starts the workers.
func (wp *WorkerPoll) dispatcher() {
	var wg sync.WaitGroup

	// Start worker goroutines.
	for i := 1; i <= wp.numWorkers; i++ {
		wg.Add(1)
		go wp.worker(i, &wg)
	}

	// Enqueue tasks into the jobs queue.
	for i := 0; i < len(wp.links); i++ {
		wp.jobsQueue <- Task{Link: wp.links[i]}
	}
	close(wp.jobsQueue) // Signal no more jobs will be added.

	wg.Wait()              // Wait for all workers to complete
	close(wp.resultsQueue) // Signal no more results will be sent.
}

// Worker Goroutines: function to process page.
func (wp *WorkerPoll) worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range wp.jobsQueue { // Process jobs from the queue
		log.Printf("Worker %d started crawl page %s\n", id, task.Link)
		result, error := wp.crawler.CrawlPageUrl(task.Link)
		log.Printf("Worker %d finished crawl page %s\n", id, task.Link)
		wp.resultsQueue <- Result{ad: result, error: error} // Send result to the collector
	}
}

// ResultsCollector: collects results of crawling pages
func (wp *WorkerPoll) resultsCollector(done chan bool) {
	for result := range wp.resultsQueue {
		if result.error != nil {
			wp.errors = append(wp.errors, result)
		} else {
			wp.results = append(wp.results, result)
		}
	}
	done <- true // Signal that result collection is complete
}

func (wp *WorkerPoll) Start() {
	log.Printf("start worker-pool with %d worker", wp.numWorkers)
	links, err := wp.crawler.CrawlAdsLinks(wp.mainLink)
	if err != nil {
		log.Println("Error in crawl main page with crawler ")
		return
	}
	log.Printf("from main link %s fetch %d sub-link", wp.mainLink, len(links))
	wp.links = links

	go wp.resultsCollector(wp.done)
	// Start the dispatcher to crawl pages and start workers
	wp.dispatcher()
	<-wp.done
}

func NewWorkerPool(mainLink string, numWorkers int, crawler CrawlerInterface) *WorkerPoll {
	return &WorkerPoll{
		mainLink:     mainLink,
		crawler:      crawler,
		numWorkers:   numWorkers,
		links:        make([]string, 0),
		results:      make([]Result, 0),
		errors:       make([]Result, 0),
		done:         make(chan bool),
		jobsQueue:    make(chan Task),
		resultsQueue: make(chan Result),
	}
}

func (wp WorkerPoll) GetResults() []Result {
	return wp.results
}

func (wp WorkerPoll) GetErrors() []Result {
	return wp.errors
}
