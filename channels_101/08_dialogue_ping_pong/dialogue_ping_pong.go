package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

type Ball int

/*
Will block, until the next on writes.
That way we can enforce that we will PING AND PONG :D
*/

func Play(playerName string, table chan Ball) {
	var lastValue Ball = -1
	for {
		ball := <-table
		log.Println(playerName, ball)
		ball = Ball(rand.Intn(11))
		if ball == 5 && lastValue != -1 {
			log.Println(playerName, " WINS")
			os.Exit(0)
		}
		lastValue = ball
		table <- ball
		if playerName == "A"{
			//even player A being more fast will happen one at the time
			time.Sleep(time.Second/3)
		}else {
			time.Sleep(time.Second)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	table := make(chan Ball)
	go func() {
		table <- 1
	}()
	go Play("A", table)
	Play("B", table)
}
