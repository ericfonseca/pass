package main

import (
	"fmt"
	"strings"
	"time"
)

type message struct {
	ID uint64
	To uint64
	From uint64
	Body string
	CreatedAt time.Time
}

var upstream chan message
var downstream chan message 

func processMessage(msg message) {
	msg.Body = strings.Replace(msg.Body, "alot", " a lot", -1)
	downstream <- msg

}

func upstreamReceive() {
	for {
		msg := <-upstream
		processMessage(msg)
	}
}

func downstreamReceive() {
	for {
		msg := <-downstream
		fmt.Printf("downstream recv msg id: %d, delay: %v\n", msg.ID, time.Now().Sub(msg.CreatedAt))
	}
}

func main() {
	upstream = make(chan message, 100)
	downstream = make(chan message, 100)

	go upstreamReceive()
	go downstreamReceive()

	upstream <- message{ID: 1, To: 5578923893, From: 90123109, Body: "hey do you have a minute", CreatedAt: time.Now()}
	upstream <- message{ID: 2, To: 90123109, From: 5578923893, Body: "sure, whats up?", CreatedAt: time.Now()}
	upstream <- message{ID: 3, To: 123098, From: 90123109, Body: "oh man who brought garlic bread", CreatedAt: time.Now()}
	upstream <- message{ID: 4, To: 83473848, From: 127655123, Body: "i dont really agree with this", CreatedAt: time.Now()}
	upstream <- message{ID: 5, To: 3092, From: 5, Body: "steve!", CreatedAt: time.Now()}
	upstream <- message{ID: 6, To: 66912398, From: 9823989, Body: "sure, let me wrap up the testing and ill do that next", CreatedAt: time.Now()}
	upstream <- message{ID: 7, To: 5578923893, From: 90123109, Body: "we need to discuss the new design, conf room?", CreatedAt: time.Now()}
	upstream <- message{ID: 8, To: 90123109, From: 123098, Body: "i did, you like? its gilroy garlic!", CreatedAt: time.Now()}
	upstream <- message{ID: 9, To: 123098, From: 90123109, Body: "yes i like, i like alot", CreatedAt: time.Now()}
	upstream <- message{ID: 10, To: 90123109, From: 5578923893, Body: "c ya there", CreatedAt: time.Now()}
 
	for{}
}