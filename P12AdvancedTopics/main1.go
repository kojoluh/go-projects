package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkTofuPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var itemPrice = rand.Float32() * 20
		fmt.Println("itemTofuPrice ", itemPrice)
		if itemPrice <= MAX_TOFU_PRICE {
			c <- website
			break
		}
	}
}

func checkChickenPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var itemPrice = rand.Float32() * 20
		fmt.Println("itemChickenPrice ", itemPrice)
		if itemPrice <= MAX_CHICKEN_PRICE {
			c <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nText Sent: Found deal on chicken at %v.\n", website)
	case website := <-tofuChannel:
		fmt.Printf("\nEmail Sent: Found deal on tofu at %v.\n", website)
	}
}
