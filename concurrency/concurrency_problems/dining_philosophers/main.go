package main

import (
	"fmt"
	"sync"
	"time"
)

// creating philosopher struct with his name, forks and his hunger
type Philosopher struct{
  Name string
  RightFork Fork
  LeftFork Fork
  Hunger int
}

// fork struct with its number and mutex to lock and unlock it
type Fork struct{
  Number int
  ForkMutex *sync.Mutex
}
// wait group to wait for all philosophers to finish eating
var dining sync.WaitGroup

// main function where philosopher eats and locks forks, then unlocks them and thinking. Repeat untill
// his is full
func (philosopher Philosopher) Eat(){
  fmt.Println(philosopher.Name,"Seated")

  for h := philosopher.Hunger; h > 0; h--{
    fmt.Println(philosopher.Name,"is Hungry")
    philosopher.LeftFork.ForkMutex.Lock() 
    philosopher.RightFork.ForkMutex.Lock() 
    fmt.Println(philosopher.Name,"is Eating")
    fmt.Println("---------------------")
    time.Sleep(time.Millisecond * 2000)
    philosopher.LeftFork.ForkMutex.Unlock()
    philosopher.RightFork.ForkMutex.Unlock()
    fmt.Println(philosopher.Name,"Thinking")
    fmt.Println("---------------------")
    time.Sleep(time.Millisecond * 2000)
  }
    fmt.Println(philosopher.Name,"isnt hungry anymore")
    dining.Done()
    fmt.Println(philosopher.Name,"has left the table")
}

// main function, adding 2 items in a wait group to wait for 2 philosophers to finish 
// then creating forks for first philosopher and create him
// create 1 more fork and the second philosopher, theyhave the same fork
// so there are 3 forks for 2 philosophers, they are waiting for each other to finish eating
func main(){
  dining.Add(2)
  var forkMutex0 sync.Mutex
  var forkMutex1 sync.Mutex
  fork0 := &Fork{0,&forkMutex0}
  fork1 := &Fork{1,&forkMutex1}
  phil0 := &Philosopher{"First",*fork0,*fork1,5}

  var forkMutex2 sync.Mutex
  fork2 := &Fork{0,&forkMutex2}
  phil1 := &Philosopher{"Second",*fork0,*fork2,5}

  go phil0.Eat()
  phil1.Eat()
  // .Wait to wait for all to finish eating and "say" .Done for the sync.WaitGroup
  dining.Wait()
}

