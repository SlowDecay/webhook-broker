package main

import (
	"fmt"
	"sync"
)

// Message represents the message to be delivered to
type Message struct {
	Payload string
}

// Job represents the job to be run
type Job struct {
	Data     Message
	Priority int
}

// jobQueue is a A buffered channel that we can send work requests on.
var jobQueue chan Job
var jobQueueInitializer sync.Once

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

// NewWorker creates a Worker
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				fmt.Println("HOLA! " + job.Data.Payload)

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

// Dispatcher is responsible for dispatching job
type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
	// Max number of workers
	MaxWorkers int
	// JobQueue to fetch jobs from
	JobQueue chan Job
	Workers  []*Worker
}

// MaxWorkersConfig represents the max number of works to spin up
type MaxWorkersConfig int

// MaxQueuesConfig represents the max number of jobs to process before blocking
type MaxQueuesConfig int

// NewMaxWorkersConfig retrieves the configuration for max workers
func NewMaxWorkersConfig() MaxWorkersConfig {
	return 50
}

// NewMaxQueuesConfig retrieves the configuration for max jobs to process at once
func NewMaxQueuesConfig() MaxQueuesConfig {
	return 1000000
}

// NewJobQueue ensures a initialized job queue is retrieved
func NewJobQueue(maxQueues MaxQueuesConfig) chan Job {
	jobQueueInitializer.Do(func() {
		jobQueue = make(chan Job, maxQueues)
	})
	return jobQueue
}

// NewDispatcher creates a new Dispatcher
func NewDispatcher(maxWorkers MaxWorkersConfig, jobQueue chan Job) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, MaxWorkers: int(maxWorkers), JobQueue: jobQueue, Workers: make([]*Worker, maxWorkers)}
}

// Run starts and preps the workers
func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
		d.Workers[i] = &worker
	}

	go d.dispatch()
}

// Stop stops the workers of the dispatcher
func (d *Dispatcher) Stop() {
	for i := 0; i < d.MaxWorkers; i++ {
		d.Workers[i].Stop()
	}
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
