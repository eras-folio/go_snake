package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type Point struct {
	x, y int
}

type Snake struct {
	body  []Point
	dir   Point
	grow  bool
}

var (
	width, height int
	food          Point
	snake         Snake
)

func initGame() {
	width, height = termbox.Size()
	snake = Snake{
		body: []Point{
			{width / 2, height / 2},
		},
		dir: Point{1, 0},
	}
	placeFood()
}

func placeFood() {
	food = Point{
		x: rand.Intn(width),
		y: rand.Intn(height),
	}
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for _, p := range snake.body {
		termbox.SetCell(p.x, p.y, 'O', termbox.ColorGreen, termbox.ColorDefault)
	}
	termbox.SetCell(food.x, food.y, 'X', termbox.ColorRed, termbox.ColorDefault)
	termbox.Flush()
}

func update() {
	head := snake.body[0]
	newHead := Point{head.x + snake.dir.x, head.y + snake.dir.y}

	if newHead.x < 0 || newHead.y < 0 || newHead.x >= width || newHead.y >= height {
		termbox.Close()
		panic("Game Over!")
	}

	for _, p := range snake.body {
		if p == newHead {
			termbox.Close()
			panic("Game Over!")
		}
	}

	if newHead == food {
		snake.grow = true
		placeFood()
	}

	snake.body = append([]Point{newHead}, snake.body...)
	if !snake.grow {
		snake.body = snake.body[:len(snake.body)-1]
	} else {
		snake.grow = false
	}
}

func handleInput() {
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowUp:
			if snake.dir != (Point{0, 1}) {
				snake.dir = Point{0, -1}
			}
		case termbox.KeyArrowDown:
			if snake.dir != (Point{0, -1}) {
				snake.dir = Point{0, 1}
			}
		case termbox.KeyArrowLeft:
			if snake.dir != (Point{1, 0}) {
				snake.dir = Point{-1, 0}
			}
		case termbox.KeyArrowRight:
			if snake.dir != (Point{-1, 0}) {
				snake.dir = Point{1, 0}
			}
		case termbox.KeyEsc:
			termbox.Close()
			panic("Game Over!")
		}
	case termbox.EventResize:
		width, height = termbox.Size()
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	rand.Seed(time.Now().UnixNano())
	initGame()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			update()
			draw()
		default:
			handleInput()
		}
	}
}
