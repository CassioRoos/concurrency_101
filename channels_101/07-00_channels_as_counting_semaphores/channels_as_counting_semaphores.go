package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

const BarSize = 10

type Seat int
type Bar chan Seat
type Customer int
type Customers chan Customer

func (bar Bar) serverCustomer(customers Customers,wg *sync.WaitGroup ) {
	for c := range customers {
		log.Print("->Customer ", c, " enters the bar ")
		// this might be confusing but, in this case seat IS A CHAN ON SEAT (type Bar chan Seat)
		seat := <-bar
		log.Print("++Customer ", c, " drinks at seat ", seat)
		d := time.Second * time.Duration(2+rand.Intn(6))
		log.Print("--Customer ", c," will take ", d, " secs ")
		time.Sleep(d)

		log.Print("<-CUSTOMER ", c, " LEFT SEAT ", seat)
		bar <- seat // free seat and leave the bar
		// set service ad done
		wg.Done()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// The bar has 10 seats and do not close
	bar24x7 := make(Bar, BarSize)
	customers := make(Customers, BarSize)
	wg := sync.WaitGroup{}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// place seats in a bar
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		// None of the sends will block
		bar24x7 <- Seat(seatId)
	}

	// create 10 consumers,they will be showing the customers to their seats
	for i := 0; i < 10; i++ {
		go bar24x7.serverCustomer(customers, &wg)
	}

	// this is a trick to break the loop instead of breaking the select
loop:
	for infiniteCustomers := 0; ; infiniteCustomers++ {
		// it is an ID
		select {
		case <-c:
			fmt.Println()
			fmt.Println()
			fmt.Println("********************STOPPED********************")
			fmt.Println()
			fmt.Println()
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
	close(bar24x7)
	fmt.Println("All done. Bye!")
}
