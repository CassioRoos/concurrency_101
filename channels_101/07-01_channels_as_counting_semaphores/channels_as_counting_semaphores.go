package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

const BarSize = 10

type Customer int
type Customers chan Customer


/*
IN THIS EXAMPLE WE DELETED SEATS, WHICH WAS NOT NECESSARY, AS YOU CAN SEE,
THE RESULT IS THE SAME
*/
func serverCustomer(customers Customers,wg *sync.WaitGroup ) {
	for c := range customers {
		log.Print("->Customer ", c, " enters the bar ")
		d := time.Second * time.Duration(2+rand.Intn(6))
		log.Print("--Customer ", c," will take ", d, " secs ")
		time.Sleep(d)
		log.Print("<-CUSTOMER ", c, " LEFT" )
		// set service ad done
		wg.Done()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	customers := make(Customers, BarSize)
	wg := sync.WaitGroup{}

	//Just to have some fun, we will wait a signal to stop the process
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// create 10 consumers,they will be showing the customers to their seats
	for i := 0; i < 10; i++ {
		go serverCustomer(customers, &wg)
	}

	// this is a trick to break the loop instead of breaking the select
loop:
	for infiniteCustomers := 0; ; infiniteCustomers++ {
		// it is an ID
		select {
		case <-c:
			log.Println()
			log.Println()
			log.Println("********************LAST CALL********************")
			log.Println()
			log.Println()
			break loop

		default:
		}
		// we need to increment the number of items in the WG
		wg.Add(1)
		customers <- Customer(infiniteCustomers)
	}
	// this will block until all is done
	wg.Wait()
	// it's time to close the bar
	close(customers)
	log.Println("All done. Bye!")
}
