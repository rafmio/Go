package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var Password = secret{password: "myPassword"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	time.Sleep(time.Millisecond * 100)
	c.password = pass
	c.RWM.Unlock()
}

func show(c *secret) string {
	c.RWM.RLock()
	fmt.Println("Show")
	time.Sleep(time.Millisecond * 300)
	defer c.RWM.RUnlock()
	return c.password
}

func showWithLock(c *secret) string {
	c.RWM.Lock()
	fmt.Println("ShowWithLock")
	time.Sleep(time.Millisecond * 300)
	defer c.RWM.Unlock()
	return c.password
}

func main() {
	var showFunction = func(c *secret) string { return "" }
	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex!")
		showFunction = show
	} else {
		fmt.Println("Using sync.Mutex!")
		showFunction = showWithLock
	}

	var wg sync.WaitGroup

	fmt.Println("Pass:", showFunction(&Password))

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Go Pass:", showFunction(&Password))
		}()
	}

	go func() {
		wg.Add(1)
		defer wg.Done()
		Change(&Password, "1234567")
	}()

	wg.Wait()
	fmt.Println("Pass:", showFunction(&Password))
}
