package main

import (
	"fmt"
)

func main() {
	var car = make(map[string]string)
	car["name"] = "BWM"
	car["color"] = "Black"

	message := createMessage(car)
	printMessage(message)
}

func createMessage(car map[string]string) string {
	name := car["name"]
	color := car["color"]

	message := fmt.Sprintf("Mobil %s berwarna %s", name, color)
	return message
}

func printMessage(message string) {
	fmt.Println(message)
}
