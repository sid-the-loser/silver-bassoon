package main

import (
	"flag"
	"fmt"
	"log"
)

const version = "pr1.0.0"
const commonErrorSuffix = "Use \"-help\" for more information about silver bassoon."

func main() {
	versionCheck := flag.Bool("version", false, "Shows the version of silver-bassoon")

	//winWidth := flag.Int("width", 800, "The width of the Pygame window")
	//winHeight := flag.Int("height", 600, "The height of the Pygame window")
	//
	//useScaling := flag.Bool("scalable", false, "Makes Pygame window scalable")
	//useResize := flag.Bool("resizable", false, "Makes Pygame window resizable")

	flag.Parse()

	// checks if the version flag is enabled, if so, show version and stop program
	// todo: try adding GitHub API based version checkups
	if *versionCheck {
		fmt.Println("silver-bassoon", version)
		return
	}

	filename := flag.Arg(0)

	if len(filename) == 0 {
		log.Fatalln("No filename provided to generate the Pygame script! " + commonErrorSuffix)
	}
}
