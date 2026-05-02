package snake

import (
	"fmt"
	"time"
)

const (
	targetFrame     = 60
	targetFrameTime = time.Second / targetFrame
)

type Game struct {
	m Map

	running bool
}

func NewGame(m Map) *Game {
	fmt.Print("\033[2J")   // Clear screen
	fmt.Print("\033[?25l") // Hide cursor
	fmt.Print("\033[H")    // Move cursor to top-left

	return &Game{
		m:       m,
		running: true,
	}
}

func (g *Game) Run() {
	for g.running {
		start := time.Now()

		g.m.Draw()

		elapsed := time.Since(start)
		remaining := targetFrameTime - elapsed

		if remaining > 0 {
			time.Sleep(remaining)
		}
	}
}
