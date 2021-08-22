package main

import (
	"log"
)

/*
A select block with one default branch and only one case branch is called a try-send
or try-receive channel operation, depending on whether the channel operation
following the case keyword is a channel send or receive operation.
	If the operation following the case keyword is a send operation, then the select
		block is called as try-send operation. If the send operation would block,
		then the default branch will get executed (fail to send), otherwise, the send
		succeeds and the only case branch will get executed.
	If the operation following the case keyword is a receive operation, then the
		select block is called as try-receive operation. If the receive operation
		would block, then the default branch will get executed (fail to receive),
		otherwise, the receive succeeds and the only case branch will get executed.
Try-send and try-receive operations never block.

The standard Go compiler makes special optimizations for try-send and try-receive
select blocks, their execution efficiencies are much higher than multi-case select
blocks.

The following is an example which shows how try-send and try-receive work.
*/

type Book struct {
	ID int
}

func main() {
	bookShelf := make(chan Book,3)

	for i := 0; i < cap(bookShelf)*2; i++ {
		select {
		case bookShelf<- Book{ID: i}:
			log.Println("success to put book", i)
		default:
			log.Println("failed to put book")
		}
	}

	for i := 0; i < cap(bookShelf)*2; i++ {
		select {
		case book:= <- bookShelf:
			log.Println("succeeded to get boo", book.ID)
		default:
			log.Println("failed to get book")
		}
	}
}
