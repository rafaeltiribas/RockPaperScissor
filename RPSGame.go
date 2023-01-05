// Classic Rock Paper and Scissors game.
package main

import (
	"fmt"
	"math/rand"
)

func DrawElement(x int) {
	paper :=
		`	______
	---'    ____)____
			   ______)
			  _______)
			 _______)
	---.__________)`

	scissor :=
		`_______
	---'   ____)____
			  ______)
		   __________)
		  (____)
	---.__(___)`

	rock :=
		` _______
	---'   ____)
		  (_____)
		  (_____)
		  (____)
	---.__(___)`

	switch x {
	case 0:
		fmt.Println(rock)
	case 1:
		fmt.Println(paper)
	case 2:
		fmt.Println(scissor)
	default:
		fmt.Println("ERROR!")
	}
}

func Game() {
	var player int // Player and Machine element choices.
	fmt.Println("<<Choose your element>>\n\n<<  Rock - 0  >>\n<<  Paper - 1  >>\n<<  Scissor - 2  >>\n\nElement: ")
	fmt.Scan(&player)
	for player < 0 || player > 2 {
		fmt.Println("<<Insert a valid option>>\nElement ==>")
		fmt.Scan(&player)
	}
	fmt.Println("\n\n<<You chose>>\n")
	DrawElement(player)
	machine := rand.Intn(3) // Computer picks an random number from 0 to 3.
	fmt.Println("\n\n<<Computer chose>>\n")
	DrawElement(machine)
	result := Win(player, machine)
	switch result {
	case 0:
		fmt.Println("\n\n<<YOU WIN!!!>>\n")
	case 1:
		fmt.Println("\n\n<<YOU LOST :(>>")
	case 2:
		fmt.Println("\n\n<<DRAW!>>")
	}
}

func Win(p int, m int) int {
	//	Verify each win possibility << 0 = Player Win | 1 = Machine Win || 2 - Draw>>
	//	<<  Rock - 0  >> <<  Paper - 1  >> <<  Scissor - 2  >>

	// Draw Situation
	if p == m {
		return 2
	}

	// Scissor
	if p == 2 {
		if m == 1 {
			return 0
		} else {
			return 1
		}
	}

	// Rock
	if p == 0 {
		if m == 2 {
			return 0
		} else {
			return 1
		}
	}

	// Paper

	if p == 1 {
		if m == 0 {
			return 0
		} else {
			return 1
		}
	}

	// Error
	return -1
}

func main() {
	Game()

}
