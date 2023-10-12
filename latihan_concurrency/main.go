package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Player struct {
	Name string
	Hit  int
}

func main() {
	var player1, player2 string

	fmt.Print("Input nama player 1: ")
	fmt.Scanln(&player1)
	fmt.Print("Input nama player 2: ")
	fmt.Scanln(&player2)
	ball := make(chan int)
	done := make(chan bool)

	var a Player = Player{Name: player1, Hit: 0}
	var b Player = Player{Name: player2, Hit: 0}

	go play(ball, done, a)
	go play(ball, done, b)
	ball <- 1

	if <-done {
		log.Println("Selesai...")
	}
}

func play(ball chan int, done chan bool, player Player) {
	for {

		time.Sleep(1 * time.Second)
		// player 1 hit ball to player 2
		player.Hit++
		counter := <-ball
		log.Println("Player", player.Name, "= Hit", player.Hit, "| counter", counter)
		if counter%11 == 0 {
			log.Println("Player", player.Name, "kalah, total hit=", player.Hit, "kalah di nomor", counter)
			done <- true
			break
		}

		counter = rand.Intn(100-1) + 1
		// ball will be delivered to player 2
		ball <- counter
	}
}
