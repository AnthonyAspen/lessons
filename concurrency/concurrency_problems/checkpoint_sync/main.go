package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


// creating worker struct which contains name and how long it takes for a worker to assemble his part
type Worker struct {
  Name string
  WorkTime time.Duration
}
// worker.Work(*) function that simulates worker's work
func (worker Worker) Work(wg *sync.WaitGroup,part string){
  fmt.Printf("%s begins to assemble his part %s\n",worker.Name,part)
  time.Sleep(worker.WorkTime)
  fmt.Printf("%s finished his part\n",worker.Name)
  wg.Done()
}
// function to create a worker
func createWorker(name string,workTime time.Duration) *Worker{
  return &Worker{name,workTime}
}

func main(){
  // creating a wait group, to wait for all workers to finish 
  var wg sync.WaitGroup
  wg.Add(3)
  // list of names of workers
  names := []string{"John","Adam","Apple"}
  // list of names of parts
  parts := []string{"A","B","C"}
  // creating a slice of workers
  workers := make([]Worker,3,3)
  
  // filling the slice of workers with names and random time duration that takes for individual worker to finish his job
  for _,name := range names{
    time := time.Duration(rand.Int63n(1e6))
    worker := Worker{name,time}
    workers = append(workers,worker)
  }

  // goroutines go!
  for i,part := range parts{
    go workers[i].Work(&wg,part)
  }
  // waiting for all to finish 
  wg.Wait()
  
  fmt.Println("cycle completed")





}
