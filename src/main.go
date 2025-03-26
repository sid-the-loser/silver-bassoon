package main

import (
	"flag"
	"fmt"
)

var version = "1.0.0"

func main() {
	winWidth := flag.Int("width", 800, "The width of the Pygame window")
	winHeight := flag.Int("height", 600, "The height of the Pygame window")

	useScaling := flag.Bool("scalable", false, "Makes Pygame window scalable")
	useResize := flag.Bool("resizable", false, "Makes Pygame window resizable")

	flag.Parse()

	filename := flag.Arg(0)

	formattedString := fmt.Sprintf("%s %d %d", filename, *winWidth, *winHeight)

	fmt.Printf(formattedString) // just printing the formatted text for now,
	// will work on the rest of the logic later
}
