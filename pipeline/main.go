package main

import (
	"fmt"
	"time"
)

type Item struct {
	price    int
	category string
}

func main() {
	start := time.Now().UTC()
	c := gen(
		Item{10, "shirt"},
		Item{50, "shoe"},
		Item{13, "drink"},
		Item{17, "shoe"},
	)
	out := discount(c)
	for processed := range out {
		fmt.Println("Category:", processed.category, "Price:", processed.price)
	}
	fmt.Println("Took:", time.Since(start))
}

func discount(items <-chan Item) <-chan Item {
	out := make(chan Item)
	go func() {
		defer close(out)
		for item := range items {
			time.Sleep(time.Second)
			if item.category == "shoe" {
				item.price /= 2
			}
			out <- item
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
