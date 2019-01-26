package routine

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go func() {
		channel <- "hello"
		fmt.Println("write \"hello\" done!")

		channel <- "world"
		fmt.Println("write \"World\" done!")

		fmt.Println("Write go sleep...")
		time.Sleep(3*time.Second)
		channel <- "channel"
		fmt.Println("write \"channel\" done!")
		}()

	time.Sleep(10*time.Second)
	fmt.Println("Reader Wake up...")

	msg := <-channel
	fmt.Println("Reader: ", msg)

	msg = <-channel
	fmt.Println("Reader: ", msg)

	msg = <-channel //Writer在Sleep，这里在阻塞
	fmt.Println("Reader: ", msg)
}