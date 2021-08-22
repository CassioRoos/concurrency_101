package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

const BarSize = 10

type Customer int
type Customers chan Customer


/*
THIS IS NOT GOOD, but may work in some specific cases.

BE CAREFUL, think well and be mindful,
the process won't be complete but the channel will be empty
*/
func serverCustomer(customers Customers) {
	for c := range customers {
		log.Print("->Customer ", c, " enters the bar ")
		d := time.Second * time.Duration(2+rand.Intn(6))
		log.Print("--Customer ", c," will take ", d, " secs ")
		time.Sleep(d)
		log.Print("<-CUSTOMER ", c, " LEFT" )
		// set service ad done
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	customers := make(Customers, BarSize)

	//Just to have some fun, we will wait a signal to stop the process
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// create 10 consumers,they will be showing the customers to their seats
	for i := 0; i < 10; i++ {
		go serverCustomer(customers)
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
		customers <- Customer(infiniteCustomers)
	}
	for{
		l := len(customers)
		fmt.Println(l)
		if l ==0 {
			break
		}
		time.Sleep(time.Second/10)
	}

	// it's time to close the bar
	close(customers)
	log.Println("All done. Bye!")
}
