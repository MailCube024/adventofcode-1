package main

import (
	"log"
	"strings"
	"strconv"
)

const instr string = "R4, R4, L1, R3, L5, R2, R5, R1, L4, R3, L5, R2, L3, L4, L3, R1, R5, R1, L3, L1, R3, L1, R2, R2, L2, R5, L3, L4, R4, R4, R2, L4, L1, R5, L1, L4, R4, L1, R1, L2, R5, L2, L3, R2, R1, L194, R2, L4, R49, R1, R3, L5, L4, L1, R4, R2, R1, L5, R3, L5, L4, R4, R4, L2, L3, R78, L5, R4, R191, R4, R3, R1, L2, R1, R3, L1, R3, R4, R2, L2, R1, R4, L5, R2, L2, L4, L2, R1, R2, L3, R5, R2, L3, L3, R3, L1, L1, R5, L4, L4, L2, R5, R1, R4, L3, L5, L4, R5, L4, R5, R4, L3, L2, L5, R4, R3, L3, R1, L5, R5, R1, L3, R2, L5, R5, L3, R1, R4, L5, R4, R2, R3, L4, L5, R3, R4, L5, L5, R4, L4, L4, R1, R5, R3, L1, L4, L3, L4, R1, L5, L1, R2, R2, R4, R4, L5, R4, R1, L1, L1, L3, L5, L2, R4, L3, L5, L4, L1, R3"
const test string = "R5, L5, R5, R3"
const test2 string = "R2, R2, R2"
const (
	left = iota
	up = iota
	right = iota
	down = iota
)

type Movement struct {
	x, y                          int
	currentDirection, totalBlocks int
}

func main() {
	var instructions []string = strings.Split(instr, ", ")
	mov := Movement{currentDirection:up, totalBlocks:0, x:0, y:0}
	for _, element := range instructions {
		direction := string(element[0])
		amount, err := strconv.Atoi(string(element[1]))
		if (err != nil) {
			log.Fatal(err)
		}
		
		log.Printf("New command: direction(%s)-Amount(%d)\n", direction, amount)
		
		mov.move(direction, amount)
		mov.print()
	}
	
	log.Printf("Distance from initial point to HQ - %d\n", mov.TotalBlocks())
}

func (m Movement) TotalBlocks() int {
	var x, y int = 0, 0
	if x = m.x; m.x < 0 {
		x = -m.x
	}
	if y = m.y; m.y < 0 {
		y = -m.y
	}
	
	return x + y
}

func (m *Movement) move(dir string, dist int) {
	m.changeDirection(dir)
	m.moveBlocks(dist)
}

func (m *Movement) moveBlocks(amount int) {
	switch m.currentDirection{
	case up: m.y += amount
	case down: m.y -= amount
	case left: m.x -= amount
	case right: m.x += amount
	}
}

func (m *Movement) changeDirection(command string) {
	if (command == "R") {
		switch m.currentDirection {
		case left: m.currentDirection = up
		case up: m.currentDirection = right
		case right: m.currentDirection = down
		case down: m.currentDirection = left
		}
	} else if (command == "L") {
		switch m.currentDirection {
		case left: m.currentDirection = down
		case up: m.currentDirection = left
		case right: m.currentDirection = up
		case down: m.currentDirection = right
		}
	}
}

func (m Movement) getFormattedDirection() string {
	switch  m.currentDirection{
	case up: return "UP"
	case down: return "DOWN"
	case left: return "LEFT"
	case right: return "RIGHT"
	default  : return ""
	}
}

func (m *Movement) print() {
	log.Printf("Current position (%d,%d): Direction (%s) - Total Blocks (%d)\n", m.x, m.y, m.getFormattedDirection(), m.TotalBlocks())
	
}