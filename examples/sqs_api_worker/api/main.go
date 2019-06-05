package main

import (
	"log"

	"github.com/vmihailenco/taskq/examples/sqs_api_worker"
)

func main() {
	go sqs_api_worker.LogStats()

	go func() {
		for {
			err := sqs_api_worker.CountTask.Call()
			if err != nil {
				log.Fatal(err)
			}
			sqs_api_worker.IncrLocalCounter()
		}
	}()

	sig := sqs_api_worker.WaitSignal()
	log.Println(sig.String())
}