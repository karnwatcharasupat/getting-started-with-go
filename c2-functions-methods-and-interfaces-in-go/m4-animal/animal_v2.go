package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var animalMap = make(map[string]Animal)

type Animal interface {
	Eat()
	Speak()
	Move()
}

type Cow struct{}

func (Cow) Eat() {
	fmt.Println("grass")
}

func (Cow) Move() {
	fmt.Println("walk")
}

func (Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (Bird) Eat() {
	fmt.Println("worms")
}

func (Bird) Move() {
	fmt.Println("fly")
}

func (Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (Snake) Eat() {
	fmt.Println("mice")
}

func (Snake) Move() {
	fmt.Println("slither")
}

func (Snake) Speak() {
	fmt.Println("hsss")
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for true {
		loop(scanner)
	}
}

func loop(scanner *bufio.Scanner) {
	fmt.Print("> ")

	if scanner.Scan() {
		in := scanner.Text()
		s := strings.Split(in, " ")

		switch strings.ToLower(s[0]) {
		case "newanimal":
			newAnimal(s[1], strings.ToLower(s[2]))
		case "query":
			query(s[1], strings.ToLower(s[2]))
		}
	}

}

func query(name string, cmd string) {
	animal := animalMap[name]

	switch cmd {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func newAnimal(name string, species string) {
	switch species {
	case "cow":
		animalMap[name] = Cow{}
	case "bird":
		animalMap[name] = Bird{}
	case "snake":
		animalMap[name] = Snake{}
	}

	fmt.Println("Created it!")
}
