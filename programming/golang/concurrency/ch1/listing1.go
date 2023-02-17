package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(job *Job, goGroup *sync.WaitGroup) {
	for job.i < job.max {
		time.Sleep(1 * time.Millisecond)
		job.i += 1
		fmt.Println(job.text)
	}
	goGroup.Done()
}

func main() {
	hello := new(Job)
	world := new(Job)
	hello.text = "hello"
	hello.max = 3

	world.text = "world"
	world.max = 5

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go outputText(hello, wg)
	go outputText(world, wg)
	time.Sleep(1 * time.Second)
	wg.Wait()
	// for {
	// }

}
