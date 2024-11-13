package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron"
)

func TestCronJob(t *testing.T) {
	t.Log("Create new cron")

	c := cron.New()
	// Define the Cron job schedule
	c.AddFunc("* * * * *", func() {
		fmt.Println("Hello world!")
	})
	c.Start()

	time.Sleep(10 * time.Second)
}
