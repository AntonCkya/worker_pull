package main

import (
	"fmt"
	"reflect"
	"time"
)

type Job struct {
	Id   int
	Job  func([]interface{}) ([]interface{}, error)
	Argv []interface{}
}

func (w *Job) Work() ([]interface{}, error) {
	res, err := w.Job(w.Argv)

	if err != nil {
		return nil, err
	}
	return res, nil
}

type Result struct {
	Value    []interface{}
	Type     reflect.Type
	WorkerId int
	JobId    int
}

func worker(id int, jobs <-chan Job, res chan<- Result, quit <-chan bool) {
	defer fmt.Printf("worker_id: %v done\n", id)
	for {
		var job Job
		select {
		case job = <-jobs:
			currRes, err := job.Work()
			if err != nil {
				fmt.Printf("ERROR in worker_id: %v, job_id: %v, err: %v\n", id, job.Id, err.Error())
				time.Sleep(time.Second * 10)
				continue
			}
			res <- Result{
				Value:    currRes,
				Type:     reflect.TypeOf(currRes),
				WorkerId: id,
				JobId:    job.Id,
			}
			time.Sleep(time.Second * 10)
		case <-quit:
			return
		}
	}
}

func write_result(res <-chan Result) {
	for r := range res {
		fmt.Printf("worker_id: %v, job_id: %v, res: %v\n", r.WorkerId, r.JobId, r.Value)
	}
}

func worker_pool(f func([]interface{}) ([]interface{}, error), channel_buffer int, start_workers int) {
	Jobs := make(chan Job, channel_buffer)
	Res := make(chan Result, channel_buffer)
	Quit := make(chan bool, channel_buffer)
	jobId := 0

	for i := 0; i < 3; i++ {
		go worker(i+1, Jobs, Res, Quit)
	}
	workersAll := start_workers
	currWorkers := start_workers

	go write_result(Res)

	for true {
		var command string
		fmt.Scan(&command)
		if command == "job" {
			var jobString string
			fmt.Scan(&jobString)
			Jobs <- Job{
				Id:   jobId,
				Job:  f,
				Argv: append(make([]interface{}, 0), jobString),
			}
			jobId++
		} else if command == "add" {
			var addWorkers int
			fmt.Scan(&addWorkers)
			for i := 0; i < addWorkers; i++ {
				go worker(workersAll+i+1, Jobs, Res, Quit)
			}
			workersAll += addWorkers
			currWorkers += addWorkers
		} else if command == "del" {
			var delWorkers int
			fmt.Scan(&delWorkers)
			for i := 0; i < delWorkers; i++ {
				Quit <- true
				currWorkers--
				if currWorkers < 1 {
					fmt.Println("WARNING deleted workers is more than exist, the remaining workers will be deleted on the next creation")
				}
			}
		}
	}
}

func main() {
	f := func(argv []interface{}) ([]interface{}, error) {
		str, _ := argv[0].(string)
		newStrByte := make([]byte, len(str))
		for i := 0; i < len(str); i++ {
			newStrByte[i] = str[len(str)-i-1]
		}
		newStr := string(newStrByte)
		res := make([]interface{}, 1)
		res[0] = newStr
		return res, nil
	}
	worker_pool(f, 1000, 3)
}
