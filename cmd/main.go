package main

import (
	"log"

	"github.com/marcem7D0/snake"
	"golang.org/x/term"
)

func run() error {
	width, height, err := term.GetSize(0)
	if err != nil {
		return err
	}

	snakeMap := snake.NewMap(width, height)
	game := snake.NewGame(snakeMap)

	game.Run()

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
