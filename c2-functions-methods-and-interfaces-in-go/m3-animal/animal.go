package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

var animalMaps = map[string]Animal{
	"cow": Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	},
	"bird": Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	},
	"snake": Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	},
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for true {
		loop(scanner)
	}
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func loop(scanner *bufio.Scanner) {
	fmt.Print(">")
	if scanner.Scan() {
		in := scanner.Text()
		s := strings.Split(in, " ")
		name := strings.ToLower(s[0])
		info := strings.ToLower(s[1])

		animal := animalMaps[name]
		switch info {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		}

	}
}
