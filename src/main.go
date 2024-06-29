package main

import (
	"dir_to_json/pkg"
	"fmt"
	"os"
)

func main() {

	//check if there are enough arguments
	if len(os.Args) < 3 {
		fmt.Println("\n\033[31mNot enough arguments\033[0m")
		fmt.Println("\nUsage: dir_to_json <path> <level> [show]")
		fmt.Print("Example: dir_to_json /tmp 2 show\n\n")
		return
	}

	//get arguments from command line
	args := os.Args[1:]

	// First argument is the path
	path := args[0]

	// Second argument is the level
	level := 0
	levelString := args[1]
	fmt.Sscanf(levelString, "%d", &level)

	// Third argument is the show hide
	showHide := false
	if len(os.Args) > 3 {
		showHide = args[3] == "show"
	}

	// Print arguments
	fmt.Println("Path:", path)
	fmt.Println("Level:", level)
	fmt.Println("Show/Hide:", showHide)

	out, err := pkg.DirToJson(path, level, showHide)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(out)

}
