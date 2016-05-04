package main

import (
	"fmt"
	"time"
)

/*
Problem of concurrency: shared data
traditional languages' solutions: lock, mutex, semiphore

!!! Go's solution:
  Real world example: buying and eating eggs,
  In my own house, both me and my wife will bug eggs when find out no more eggs left, so it might end with both of us buying duplicate eggs.
  Use the solution in Go, I could say that I'll control the eggs. I'll be the one who looks at the eggs.
  I could tell my wife don't ever look at the eggs and decide that we're going to get eggs.
  I'll keep track of the eggs since I eat most of the eggs and I'm the one using them.
  I'll be the one who owns the eggs data. So I'll look at the egg and then, when we run low on eggs,
  and I know that you're at the store, I'm going to go ahead and use my cell phone and I'm going to call you and I'll ask you,
  can you get some eggs? This way you're never looking at the eggs and not knowing if I'm going to purchase the eggs and we don't end up with duplicate purchases of eggs.
  If I want to go out and get eggs because I see its low, then I'll do that and you won't have to worry about it.
  But, if I know that you're out at the store and we need more eggs, I'll let you know.
  So what I'm doing here is I'm basically taking control of that data. I'm no longer sharing the data.
  Instead what I'm doing is I'm using a communication channel to share information about that data versus using the data itself to communicate.
  And there's a big difference here when you think about that. It's really about who is in charge or what is controlling the actions.
  Before, the data itself was controlling the action, so the count of the eggs in the refrigerator was controlling my actions and my wife's actions as far as purchasing the eggs.
  But I changed that around and said, look, this communication is going to control the actions instead.

  How Go solves it:
  we're not going to try and share data. That's too messy.
  If you have multiple threads accessing the same data, it's bad.
  Instead of sharing the data at all, what we're going to do is we're going to communicate data so the data is not going to be the thing that we share.
  Instead, we're going to share a communication channel and we're going to talk about the data and I'm going to signal one thread to the other some information that is going to determine how that data gets manipulated.
  One thread in particular will own a piece of data and they'll communicate instead by a channel instead of using just the data to communicate.
  One thing that's often said of the Go language is that it's designed to enforce this principal.
  Do not communicate by sharing memory; instead, share memory by communicating.
  And the emphasis is where the communication happens or what drives that communication, not the data.
  Instead we're going to share the data by communicating.

  thread1(Goroutine1) <- Channel -> thread2(Goroutine2)
     ↓
    Data

  Goroutines:
    - Lightweight thread
    - Managed by Go runtime

  Channel:
    - basic: only one thing in the channel at a specific time
    - buffered: # of buffered things in the channel
*/

type ChannelData struct {
	id   int
	name string
}

func PrintChannelDataInfo(data []ChannelData, customMsg string) {
	for _, t := range data {
		fmt.Printf("%d: %s (%s)\n", t.id, t.name, customMsg)
	}
}

// Channel with range
type ChannelDatas []ChannelData

/*
Think of a channel as a message queue.
If the channel is on the right of the left arrow (<-) operator, it means to dequeue an entry. Saving the entry in a variable is optional
	e <- q
If the channel is on the left of the left arrow operator, it means to enqueue an entry.
	q <- e
*/
func (channelDatas ChannelDatas) PrintChannelDatas(c chan ChannelData) {
	// async filling up the channel with data
	// can be done in different coroutines, too.
	for _, s := range channelDatas {
		c <- s
	}
	// without closing channel, it'll raise "fatal error: all goroutines are asleep - deadlock!",
	// since the loop in TestChannelWithRange will keep waiting
	close(c)
}
func TestChannelWithRange() {
	c := make(chan ChannelData)

	channelDatas := ChannelDatas{
		{11, "Zi"},
		{22, "Z"},
		{33, "Zee"},
	}

	go channelDatas.PrintChannelDatas(c)
	for s := range c {
		fmt.Println(s.name)
	}
}

// select statement in channels
/* like a switch, but no break/fallthrough, rules are:
Execute case the is "ready", means has data in the channel
If more than one is "ready" execute one at random
If none are "ready", block unless has defined something in default case
*/
func TestSelectInChannels() {
	channelDatas := ChannelDatas{
		{111, "Zi"},
		{222, "Z"},
		{333, "Zee"},
	}

	c1 := make(chan ChannelData)
	c2 := make(chan ChannelData)
	go channelDatas.PrintChannelDatas(c1)
	go channelDatas.PrintChannelDatas(c2)

	for { // infinite loop
		select { // check the following 3 cases to see which one is doable
		case s1, ok := <-c1:
			if ok {
				fmt.Println(s1, ":channel 1")
			} else {
				fmt.Println("Nothing in channel 1")
				return
			}
		case s2, ok := <-c2:
			if ok {
				fmt.Println(s2, ":channel 2")
			} else {
				fmt.Println("Nothing in channel 2")
				return
			}
		default:
			fmt.Println("waiting...")
		}
	}
}

func main_concurrency() {
	data := []ChannelData{
		{1, "Zi"},
		{2, "Z"},
		{3, "Zee"},
	}

	// basic channel
	done := make(chan bool) // only allows to pass bool data

	// buffered channel
	bufferedDone := make(chan bool, 2)

	// buffered channel to show race condition
	badBufferedDone := make(chan bool, 2)

	// go PrintChannelDataInfo(data, "<C>") // concurrent
	go func() {
		PrintChannelDataInfo(data, "<C>")
		done <- true // return result to the done channel in a sepreate thread
		/* !!! it's a basic buffer, so if we add the following line,
		   the second "done <- true" since the first one is in the channel,
		   when main thread reaches "res := <-done", the first one will be gone,
		   and the second goes into the channle, but now the program already exits,
		   so we won't see "Done in a seperate thead." printed out.
		*/
		done <- true
		fmt.Println("Done in a seperate thead 1.")
	}()

	go func() {
		PrintChannelDataInfo(data, "<C>")
		bufferedDone <- true
		bufferedDone <- true
		// "Done in a seperate thead 2." will always be printed out because bufferedDone allows at most 2 in the channel at the same time
		fmt.Println("Done in a seperate thead 2.")
	}()

	go func() {
		PrintChannelDataInfo(data, "<C>")
		badBufferedDone <- true
		time.Sleep(100 * time.Millisecond) // "Done in a seperate thead 3." won't be printed out since the following line executes after main thread exiting
		// need to add the following infinite while loop at the end of main function, which doesn't make sense
		/*
		   for {
		     time.Sleep(100 * time.Millisecond)
		   }
		*/
		badBufferedDone <- true
		fmt.Println("Done in a seperate thead 3.")
	}()

	PrintChannelDataInfo(data, "None")
	// time.Sleep(100 * time.Millisecond) // without this line, "<C>" messages won't have time to be printed out

	// get the result from the done channel in the main thread
	res := <-done // only "<- done" will just block the main thread and wait for the result
	fmt.Println(res)

	<-bufferedDone

	TestChannelWithRange()

	TestSelectInChannels()
}
