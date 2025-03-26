package main

import (
	"flag"
	"log"
)

var version = "1.0.0"

func main() {
	// winWidth := flag.Int("width", 800, "The width of the Pygame window")
	// winHeight := flag.Int("height", 600, "The height of the Pygame window")

	// useScaling := flag.Bool("scalable", false, "Makes Pygame window scalable")
	// useResize := flag.Bool("resizable", false, "Makes Pygame window resizable")

	flag.Parse()

	filename := flag.Arg(0)

	if len(filename) == 0 {
		log.Fatalln("No filename provided to generate the Pygame script!")
	}
}
