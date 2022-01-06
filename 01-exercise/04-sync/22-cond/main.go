package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	cd := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		cd.L.Lock()
		for len(sharedRsc) < 1 {
			cd.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		cd.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		cd.L.Lock()
		for len(sharedRsc) < 2 {
			cd.Wait()
		}

		fmt.Println(sharedRsc["rsc2"])
		cd.L.Unlock()
	}()
	cd.L.Lock()
	// writes changes to sharedRsc
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	cd.Broadcast()
	cd.L.Unlock()
	wg.Wait()
}
