package main

import (
	"flag"
	"fmt"
)

func main() {
	channel := flag.String("channelName", "philipkrueck", "Your YouTube channel name.")
	flag.Parse()
	fmt.Println(*channel)
}
