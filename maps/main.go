package main

import "fmt"

func main() {
	var colorsEmpty map[string]string
	colorsMade := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}

	fmt.Println(colorsEmpty)
	fmt.Println(colorsMade)

	fmt.Println(colors)

	colors["white"] = "#ffffff"
	fmt.Println(colors)
	delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// use map when you have a collection of closely related props: like color name to hex
//   - may add more in the future
// use struct for entities that make more sense as object
