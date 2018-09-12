// GopherVenture!

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	// Game intro and reader is declared
	fmt.Println("-=GopherVenture!=-\nAn adventure game made in Go!")
	reader := bufio.NewReader(os.Stdin)

	// Set random number seed based on time
	rand.Seed(time.Now().UTC().UnixNano())

	// Game asks for name.
	fmt.Println("\nWhat is your name?")
	byteName, _, _ := reader.ReadLine()
	name := string(byteName)

	// Checks to see if you typed your name and tells name.
	// Otherwise, use fallback name "player."
	if name == "" {
		name = "player"
	}
	fmt.Printf("\nAlrighty, %s.\n", name)

	// Declare variable used to track objective
	obj := 0

	// Starting scenario is given along with the available commands
	fmt.Println("You are a gopher inside a room with the door locked. You see a key on the floor. What do you want to do?\nOptions: [open door] [pick up key]")

	// First objective is to pick up the key
	for obj == 0 {
		byteOption, _, _ := reader.ReadLine()
		option := string(byteOption)
		if option == "pick up key" {
			fmt.Println("\nYou pick up the key from the floor with your paw.")
			obj++
		} else if option == "open door" {
			fmt.Println("\nYour try to twist the knob with your paw, but it won't budge. The door is locked.")
		} else {
			fmt.Println("\nYou wonder what that means, but you can't seem to understand.")
		}
	}

	// Second objective is to open the door
	for obj == 1 {
		byteOption, _, _ := reader.ReadLine()
		option := string(byteOption)
		if option == "pick up key" {
			fmt.Println("\nYou already have the key on your paw.")
		} else if option == "open door" {
			fmt.Println("\nYou unlock the door and twist the knob. The door opens and you exit the room.")
			obj++
		} else {
			fmt.Println("\nYou wonder what that means, but you can't seem to understand.")
		}
	}

	// Third objective is to pick up the sword
	fmt.Println("You hear something approaching. You notice a sword as you look down.\nOptions: [pick up sword]")
	for obj == 2 {
		byteOption, _, _ := reader.ReadLine()
		option := string(byteOption)
		if option == "pick up sword" {
			fmt.Println("\nYou pick up the sword with your paw. The sword feels solid.")
			obj++
		} else {
			fmt.Println("\nYou wonder what that means, but you can't seem to understand.")
		}
	}

	// Now you must fight!
	fmt.Println("A bug approaches you! You must fight using the sword!\nOptions: [slash enemy]")

	// Sets enemy health
	enemy := 10

	for obj == 3 {
		byteOption, _, _ := reader.ReadLine()
		option := string(byteOption)
		if option == "slash enemy" {
			fmt.Println("\nYou hit the bug with your sword.")

			// Sword varies in damage from 0 to 5
			sword := rand.Intn(6)
			enemy -= sword

			// Checks if enemy is defeated before moving
			if enemy > 0 {
				// Checks if sword did no damage to say appropriate message
				if sword == 0 {
					fmt.Println("Your sword has inflicted no damage.")
				} else {
					fmt.Println("Your sword has inflicted", sword, "damage.")
				}

				//Prints remaining enemy health
				fmt.Println("The enemy has", enemy, "HP left.")
				fmt.Println("Keep slashing!")
			} else {
				fmt.Println("Your sword has inflicted", sword, "damage.")
				fmt.Println("The enemy has been slain and you move onward!")
				obj++
			}
		} else {
			fmt.Println("\nYou wonder what that means, but you can't seem to understand.")
		}
	}

	// Temporary ending message for end of game
	fmt.Println("\nYou vision somehow fades and your adventure temporarily halts...\nPress Enter/Return to exit.")
	reader.ReadString('\n')

}
