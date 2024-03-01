package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type ChannelMessage struct {
	chats   []string
	friends []string
}

func main() {
	userName := "kwame"
	userId := getUserByName(userName)
	fmt.Printf("userId for %v: %v\n", userName, userId)
	now := time.Now()

	wg := &sync.WaitGroup{}
	ch := make(chan *ChannelMessage, 2)

	wg.Add(2)
	go getUserChats(userId, ch, wg)
	go getUserFriends(userId, ch, wg)

	wg.Wait()
	close(ch)

	for msg := range ch {
		log.Printf("\nresults:\n chats: %v\n friends: %v\n", msg.chats, msg.friends)
	}
	log.Printf("Time used: %v", time.Since(now))
}

func getUserChats(userId string, ch chan<- *ChannelMessage, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)

	ch <- &ChannelMessage{
		chats: []string{
			"Akwasi",
			"Abena",
		},
	}
	wg.Done()
}

func getUserFriends(userId string, ch chan<- *ChannelMessage, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 3)

	ch <- &ChannelMessage{
		friends: []string{
			"Akwasi",
			"Ama",
			"Abena",
			"Kojo",
		},
	}
	wg.Done()
}

func getUserByName(userName string) string {
	time.Sleep(time.Second * 1)

	return fmt.Sprintf("%s-11", userName)
}
