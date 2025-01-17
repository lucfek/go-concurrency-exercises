//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweetChan chan<- *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tweetChan)
			return
		}
		tweetChan <- tweet
	}
}

func consumer(tweetChan <-chan *Tweet, done chan<- struct{}) {
	for {
		t, more := <-tweetChan
		if more {
			if t.IsTalkingAboutGo() {
				fmt.Println(t.Username, "\ttweets about golang")
			} else {
				fmt.Println(t.Username, "\tdoes not tweet about golang")
			}
		} else {
			done <- struct{}{}
			return
		}
	}
}

func main() {
	tweetChan := make(chan *Tweet)
	done := make(chan struct{})
	start := time.Now()
	stream := GetMockStream()

	// Producer
	go producer(stream, tweetChan)

	// Consumer
	go consumer(tweetChan, done)
	<-done
	fmt.Printf("Process took %s\n", time.Since(start))
}
