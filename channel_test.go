package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Fahril Hadi"
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string)  {
	time.Sleep(2 * time.Second)
	channel <- "Fahril Hadi"
}

func TestChannelAsParameter(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string)  {
	time.Sleep(2 * time.Second)
	channel <- "Fahril Hadi"
}

func OnlyOut(channel <-chan string)  {
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T)  {
	channel := make(chan string, 1)
	defer close(channel)

	channel <- "Fahril"

	fmt.Println(<- channel)

	fmt.Println("Selesai")
}