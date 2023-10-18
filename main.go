package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	width  = 40
	height = 10
)

type paddle struct {
	position int
}

type ball struct {
	x, y     int
	xVel, yVel int
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // For Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func draw(paddle *paddle, ball *ball) {
	clearScreen()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 || i == height-1 {
				fmt.Print("=")
			} else if j == 0 || j == width-1 {
				fmt.Print("|")
			} else if i == paddle.position && j >= 1 && j <= 3 {
				fmt.Print("P")
			} else if i == ball.y && j == ball.x {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func update(paddle *paddle, ball *ball) {
	ball.x += ball.xVel
	ball.y += ball.yVel

	if ball.x == 1 || ball.x == width-2 {
		ball.xVel = -ball.xVel
	}

	if ball.y == 1 || ball.y == height-2 {
		ball.yVel = -ball.yVel
	}

	if ball.y == paddle.position && ball.x >= 1 && ball.x <= 3 {
		ball.yVel = -ball.yVel
	}

	if ball.y == height-1 {
		gameOver()
	}

	if paddle.position < height-1 {
		paddle.position++
	}
}

func getInput(paddle *paddle) {
	var input string
	fmt.Scanln(&input)

	switch input {
	case "w":
		if paddle.position > 1 {
			paddle.position--
		}
	case "s":
		if paddle.position < height-2 {
			paddle.position++
		}
	case "q":
		os.Exit(0)
	}
}

func gameOver() {
	clearScreen()
	fmt.Println("Game Over!")
	os.Exit(0)
}

func main() {
	p := &paddle{position: height - 2}
	b := &ball{x: width / 2, y: height / 2, xVel: 1, yVel: 1}

	for {
		draw(p, b)
		getInput(p)
		update(p, b)
		time.Sleep(100 * time.Millisecond)
	}
}
