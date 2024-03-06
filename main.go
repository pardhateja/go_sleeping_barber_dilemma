package main

import (
	"fmt"
	"math/rand"
	barbershop "sleeping_barber_dilemma/sleeping_barber/sleeping_barber.go"
	"strconv"
	"time"
)

func main() {
	var availableSeats, timePerHairCut, incomingRate, shopOpenTime int

	fmt.Printf("enter the total available seats: ")

	_, err := fmt.Scanf("%d ", &availableSeats)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("enter the avg time taken for each haircut in milliseconds: ")

	_, err = fmt.Scanf("%d ", &timePerHairCut)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("enter the incoming rate of the customers to the barber shop: ")

	_, err = fmt.Scanf("%d ", &incomingRate)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("enter the total time the shop will be open in secs: ")

	_, err = fmt.Scanf("%d ", &shopOpenTime)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	runOperations(
		availableSeats,
		time.Millisecond*time.Duration(timePerHairCut),
		incomingRate,
		time.Second*time.Duration(shopOpenTime),
	)
}

func runOperations(
	availableSeats int,
	timePerHairCut time.Duration,
	incomingRate int,
	shopOpenTime time.Duration,
) {
	shop := barbershop.New(availableSeats, timePerHairCut)
	shop.AddBarber("Barber")
	i := 1

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(shopOpenTime)
		shopClosing <- true
		shop.Close()
		closed <- true
	}()

	go func() {
		for {
			// Get a random number with average arrival rate as specified
			randomMilliseconds := rand.Int() % (2 * incomingRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.AddClient(strconv.Itoa(i))
				i++
			}
		}
	}()

	<-closed
}
