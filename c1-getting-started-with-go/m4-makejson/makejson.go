/*
Write a program which prompts the user to first enter a name,
and then enter an address.

Your program should create a map
and add the name and address to the map using the keys “name” and “address”, respectively.

Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.

Submit your source code for the program,
“makejson.go”.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	data := make(map[string]string)

	fmt.Print("Enter your name: ")
	if scanner.Scan() {
		name := scanner.Text()
		data["name"] = name
	}

	fmt.Print("Enter your address: ")
	if scanner.Scan() {
		addr := scanner.Text()
		data["address"] = addr
	}

	json, _ := json.Marshal(data)
	fmt.Println(string(json))
}
