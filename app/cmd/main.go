package main

import (
	"context"
	"log"

	"dish/app/services/gamesvc"
)

func main() {
	if _, err := gamesvc.GameSVC().CreateGame(context.TODO(), "asdf"); err != nil {
		log.Fatal(err)
	}
}
