package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-co-op/gocron/v2"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("weather sss"))
		if err != nil {
			log.Printf("error writing response: %v", err)
		}
	})
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("starting server")
		err := http.ListenAndServe(":3000", r)
		if err != nil {
			panic(err)
		}
	}()

	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	jobs, err := initJob(s)
	if err != nil {
		panic(err)
	}
	_ = jobs

	go func() {
		defer wg.Done()
		fmt.Printf("starting scheduler %v\n", jobs[0].ID())
		s.Start()
	}()

	wg.Wait() // Keep main running
}

func initCron() (gocron.Job, error) {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	// add a job to the scheduler
	j, err := s.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			func() {
				fmt.Println("cron job")
			},
		),
	)
	if err != nil {
		return nil, err
	}
	// each job has a unique id
	return j, nil
}

func initJob(scheduler gocron.Scheduler) ([]gocron.Job, error) {
	j, err := scheduler.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			func(a string, b int) {
				fmt.Printf("cron job executed: %s %d\n", a, b)
			},
			"hello",
			1,
		),
	)
	if err != nil {
		return nil, err
	}
	return []gocron.Job{j}, nil
}
