package crawler

import (
	"context"
	"fmt"
	"log"
	"magical-crwler/models"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
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
	requestCount atomic.Int64
}

type Task struct {
	Link string
}

type Result struct {
	TimeSpent time.Duration
	RAMUsage  int
	CPUUsage  int
	Err       error
	Ad        *Ad
}

func (r Result) String() string {
	if r.Err == nil {
		return fmt.Sprintf("Successful crawl: %s ( Time-Spent: %s, RAM-Usage: %dKB ) ", r.Ad.Link, r.TimeSpent, r.RAMUsage)
	} else {
		return fmt.Sprintf("Error in crawl: %s ( Err: %s ) ", r.Ad.Link, r.Err.Error())
	}

}

// Dispatcher: enqueues tasks into the jobs queue and starts the workers.
func (wp *WorkerPoll) dispatcher(ctx context.Context) {
	var wg sync.WaitGroup

	// Start worker goroutines.
	for i := 1; i <= wp.numWorkers; i++ {
		wg.Add(1)
		go wp.worker(i, ctx, &wg)
	}

	// Enqueue tasks into the jobs queue.
	for i := 0; i < len(wp.links); i++ {
		wp.jobsQueue <- Task{Link: wp.links[i]}
	}
	close(wp.jobsQueue) // Signal no more jobs will be added.

	wg.Wait()              // Wait for all workers to complete
	close(wp.resultsQueue) // Signal no more results will be sent.
}
func (wp *WorkerPoll) worker(id int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range wp.jobsQueue { // Process jobs from the queue
		log.Printf("Worker %d started crawl page %s\n", id, task.Link)
		start := time.Now()
		var memStatsStart, memStatsEnd runtime.MemStats
		runtime.ReadMemStats(&memStatsStart)
		crawlData, requestCount, err := wp.crawler.CrawlPageUrl(ctx, task.Link)
		wp.requestCount.Add(int64(requestCount))
		runtime.ReadMemStats(&memStatsEnd)
		ramUsage := 0

		if memStatsEnd.HeapAlloc > memStatsStart.HeapAlloc {
			ramUsage = int((memStatsEnd.HeapAlloc - memStatsStart.HeapAlloc) / (1024)) // Convert to KB
		}
		fmt.Printf("crawled data :%v\n", crawlData)
		fmt.Printf("error crawl :%v\n", err)

		result := Result{
			Ad:        crawlData,
			Err:       err,
			TimeSpent: time.Since(start),
			RAMUsage:  ramUsage,
			CPUUsage:  0,
		}

		wp.resultsQueue <- result // Send result to the collector
	}
}

// ResultsCollector: collects results of crawling pages
func (wp *WorkerPoll) resultsCollector(done chan bool) {
	for result := range wp.resultsQueue {
		log.Println(result)
		if result.Err != nil {
			wp.errors = append(wp.errors, result)
		} else {
			wp.results = append(wp.results, result)
		}
	}
	done <- true // Signal that result collection is complete
}

func (wp *WorkerPoll) Start(ctx context.Context) {
	log.Printf("start worker-pool with %d worker", wp.numWorkers)
	links, requestCount, err := wp.crawler.CrawlAdsLinks(ctx, wp.mainLink)
	if err != nil {
		log.Println("Error in crawl main page with crawler ", err)
		return
	}
	log.Printf("from main link %s fetch %d sub-link", wp.mainLink, len(links))
	wp.links = links

	go wp.resultsCollector(wp.done)
	// Start the dispatcher to crawl pages and start workers
	wp.dispatcher(ctx)
	<-wp.done
	wp.requestCount.Add(int64(requestCount))
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

func (wp *WorkerPoll) GetResults() []Result {
	return wp.results
}

func (wp *WorkerPoll) GetErrors() []Result {
	return wp.errors
}

// generate one crawl metrics
func (wp *WorkerPoll) GetCrawlerFunctionalityReport() (models.CrawlerFunctionality, error) {
	totalTimeSpent := 0
	totalRamUsage := 0
	totalCpuUsage := 0

	for index := range wp.results {
		result := wp.results[index]
		totalTimeSpent += int(result.TimeSpent.Seconds())
		totalCpuUsage += result.CPUUsage
		totalRamUsage += result.RAMUsage
	}

	for index := range wp.errors {
		err := wp.errors[index]
		totalTimeSpent += int(err.TimeSpent.Seconds())
		totalCpuUsage += err.CPUUsage
		totalRamUsage += err.RAMUsage
	}

	cf := models.CrawlerFunctionality{
		Date:               time.Now(),
		Duration:           totalTimeSpent,
		CPUUsage:           totalCpuUsage,
		RAMUsage:           totalRamUsage,
		TotalRequests:      int(wp.requestCount.Load()),
		SuccessfulRequests: len(wp.results),
		FailedRequests:     len(wp.errors),
	}
	return cf, nil
}
