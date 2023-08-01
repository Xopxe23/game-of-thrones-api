package main

import (
	"fmt"
	"log"
	"time"

	gameofthrones "github.com/Xopxe23/game-of-thrones-api/game-of-thrones"
)

func main() {
	client, err := gameofthrones.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}
	houses, err := client.GetHouses()
	if err != nil {
		log.Fatal(err)
	}
	for _, house := range houses {
		fmt.Println(house.Info())
	}
}
