package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 70
	var remainingTickets = 70

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}
