package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	flag.Parse() // Parse all the above defined flags

	// checks if the version flag is enabled, if so, show version and stop program
	// todo: try adding GitHub API based version checkups
	if *versionCheck {
		fmt.Println("silver-bassoon", version)
		return
	}

	filename := flag.Arg(0) // getting the arguments passed through after the flags

	if len(filename) == 0 { // no filename used
		log.Fatalln("No filename provided to generate the Pygame script! " + commonErrorSuffix)
	} else if _, err := os.Stat("sample.txt"); err == nil { // if file exists
		// does nothing and continues with the rest of the program
		// this part off the code is taken from:
		// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
	} else {
		log.Fatalln("The filename you've passed doesn't exist!") // if file doesn't exist
	}

	// do the rest of the program, i can't think properly
}
