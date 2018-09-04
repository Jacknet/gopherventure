// GopherVenture!

package main

import (
	"bufio"
	"fmt"
	"os"
	//"time"
	//"math/rand"
)

func main() {

	// Game intro and reader is declared
	fmt.Println("-=GopherVenture!=-\nAn adventure game made in Go!")
	reader := bufio.NewReader(os.Stdin)

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
			obj++
			fmt.Println("\nYou pick up the key from the floor with your paw.")
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
			obj++
			fmt.Println("\nYou unlock the door and twist the knob. The door opens and you exit the room.")
		} else {
			fmt.Println("\nYou wonder what that means, but you can't seem to understand.")
		}
	}

	// Third objective is to pick up the sword
	fmt.Println("You hear something approach you. You notice a sword as you look down.\nOptions: [pick up sword]")
	for obj == 2 {
		byteOption, _, _ := reader.ReadLine()
		option := string(byteOption)
		if option == "pick up sword" {
			obj++
			fmt.Println("\nYou pick up the sword with your paw. The sword feels solid.")
		} else {
			fmt.Println("\nYou wonder what that means, but you can't seem to understand.")
		}
	}

	// Now you must fight!
	fmt.Println("A bug approaches you! You must fight using the sword!\nOptions: [slash enemy]")
	for obj == 3 {
		byteOption, _, _ := reader.ReadLine()
		option := string(byteOption)
		if option == "slash enemy" {
			obj++
			fmt.Println("\nYou slay the bug with your sword.")
		} else {
			fmt.Println("\nYou wonder what that means, but you can't seem to understand.")
		}
	}

	// TODO: RNG mechanic fight with monster
	//t := time.Now()
	//fmt.Println(rand.Intn(t.Nanosecond()))
	//fmt.Println(rand.Intn(t.Nanosecond()))
	//fmt.Println(rand.Intn(t.Nanosecond()))

	// Temporary ending message for end of game
	fmt.Println("\nYou vision somehow fades and your adventure temporarily halts...\nPress Enter/Return to exit.")
	reader.ReadString('\n')

}
