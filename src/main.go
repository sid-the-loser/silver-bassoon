package main

import (
	"flag"
	"fmt"
)

func main() {
	winWidth := flag.Int("width", 800, "The width of the Pygame window")
	winHeight := flag.Int("height", 600, "The height of the Pygame window")

	flag.Parse()

	filename := flag.Arg(0)

	formattedString := fmt.Sprintf("%s %d %d", filename, *winWidth, *winHeight)

	fmt.Printf(formattedString)
}
