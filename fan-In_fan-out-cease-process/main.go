package main

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	price    int
	category string
}

func main() {
	done := make(chan bool)
	defer close(done)
	start := time.Now().UTC()
	c := gen(
		Item{10, "shirt"},
		Item{50, "shoe"},
		Item{13, "drink"},
		Item{17, "shoe"},
		Item{5645, "table"},
		Item{100, "lamp"},
		Item{1, "candy"},
		Item{50, "cap"},

	)
	c1 := discount(done, c)
	c2 := discount(done, c)
	out := FanIn(done, c1, c2)

	for item := range out {
		fmt.Println("Category:", item.category, "Price:", item.price)
	}
	//item := <-out
	//fmt.Println("Category:", item.category, "Price:", item.price)
	//item = <-out
	//fmt.Println("Category:", item.category, "Price:", item.price)
	fmt.Println()

	fmt.Println("Took:", time.Since(start))
}

func FanIn(done <-chan bool, channels ...<-chan Item) <-chan Item {
	// we need a wg to sync everything up
	var wg sync.WaitGroup
	// this chan will unite all other chan that we might have
	out := make(chan Item)
	output := func(c <-chan Item) {
		// just to control th wg
		defer wg.Done()
		// it will iterate over the chan that we receive and will put everything
		// in our OUT chan. WILL JOIN MORE THAN ONE CHAN INTO ONE
		for item := range c {

			select {
			case out <- item:
			case <-done:
				return
			}
		}
	}



	// will add the number of parameters to WG
	wg.Add(len(channels))

	// for every channel that we receive it will try to join to ON UNIQUE CHAN
	for _, c := range channels {
		// call our join func
		go output(c)
	}

	// Here is a bit tricky, this go func will block, because have the WAIT inside
	// of it. This will ensure that all data receive is stored int that chan
	go func() {
		wg.Wait()
		close(out)
	}()

	// return our read-only chan
	return out
}

func discount(done chan bool, items <-chan Item) <-chan Item {
	out := make(chan Item)
	go func() {
		defer close(out)
		for item := range items {
			time.Sleep(time.Second)
			if item.category == "shoe" {
				item.price /= 2
			}
			select {
			case out <- item:
			case <-done:
				return
			}
		}
	}()
	return out
}

// gen converts list of items to a read-only chan
func gen(items ...Item) <-chan Item {
	// buffered chan
	out := make(chan Item, len(items))

	// iterate over all items to put them into the chan
	for _, item := range items {
		out <- item
	}

	close(out)
	return out
}
