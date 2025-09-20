package main

import (
	"L2-8/pkg/ntp"
	"fmt"
	"log"
)

func main() {
	ntp := ntp.New("pool.ntp.org")

	now, err := ntp.GetTime()
	if err != nil {
		log.Fatalf("failed to get time: %s", err)
	}

	fmt.Println(now)
}
