package main

import "fmt"

/*
 * Write a program which allows the user to get information about a predefined
 * set of animals. Three animals are predefined, cow, bird, and snake. Each
 * animal can eat, move, and speak. The user can issue a request to find out
 * one of three things about an animal: 1) the food that it eats, 2) its
 * method of locomotion, and 3) the sound it makes when it speaks. The
 * following table contains the three animals and their associated data which
 * should be hard-coded into your program.
 *
 * | Animal | Food eaten | Locomotion method | Spoken sound |
 * |--------|------------|-------------------|--------------|
 * |  cow   |   grass    |      walk         |     moo      |
 * |  bird  |   worms    |      fly          |     peep     |
 * |  snake |   mice     |      slither      |     hsss     |
 *
 * Your program should present the user with a prompt, “>”, to indicate that
 * the user can type a request. Your program accepts one request at a time
 * from the user, prints out the answer to the request, and prints out a new
 * prompt. Your program should continue in this loop forever. Every request
 * from the user must be a single line containing 2 strings. The first string
 * is the name of an animal, either “cow”, “bird”, or “snake”. The second
 * string is the name of the information requested about the animal, either
 * “eat”, “move”, or “speak”. Your program should process each request by
 * printing out the requested data.
 *
 * You will need a data structure to hold the information about each animal.
 * Make a type called Animal which is a struct containing three fields:food,
 * locomotion, and noise, all of which are strings. Make three methods called
 * Eat(), Move(), and Speak(). The receiver type of all of your methods should
 * be your Animal type. The Eat() method should print the animal’s food, the
 * Move() method should print the animal’s locomotion, and the Speak() method
 * should print the animal’s spoken sound. Your program should call the
 * appropriate method when the user makes a request.
 */

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	// Data
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	// Infinite loop
	for {
		var animal_name, action_name string
		fmt.Print("> ")
		fmt.Scanf("%s", &animal_name)
		fmt.Scanf("%s", &action_name)
		var animal Animal
		switch {
		case animal_name == "cow":
			animal = cow
		case animal_name == "bird":
			animal = bird
		case animal_name == "snake":
			animal = snake
		}
		switch {
		case action_name == "eat":
			animal.Eat()
		case action_name == "move":
			animal.Move()
		case action_name == "speak":
			animal.Speak()
		}
	}
}
