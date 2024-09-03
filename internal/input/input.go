package input

import (
	"fmt"
)

func SelectGameMode() string {
	var gameMode string

	fmt.Println("Which game mode are you documenting?")
	fmt.Println("[1] Time Trial")
	fmt.Println("[2] Pro Race")
	fmt.Println("[3] Battle Race")
	fmt.Print("Enter your choice (1, 2, or 3): ")
	fmt.Scanln(&gameMode)

	switch gameMode {
	case "1":
		return "Time Trial"
	case "2":
		return "Pro Race"
	case "3":
		return "Battle Race"
	default:
		return ""
	}
}

func SelectTrack() string {
	fmt.Println("\nWhich track are you documenting?")
	fmt.Println("[1] Rocky Road Speedway")
	fmt.Println("[2] Lighthouse Loop")
	fmt.Println("[3] Great Dragon Road")
	fmt.Println("[4] Aviator Ranch")
	fmt.Println("[5] Boathouse Bay")
	fmt.Println("[6] Cannon Crossover")
	fmt.Println("[7] Forbidden Fortress")
	fmt.Println("[8] Gold Rush Road")
	fmt.Println("[9] Badlands Motorplex")
	fmt.Println("[0] Twisted Canyon")
	fmt.Print("Enter your choice (0 to 9): ")

	var track string
	fmt.Scanln(&track)

	switch track {
	case "1":
		return "Rocky Road Speedway"
	case "2":
		return "Lighthouse Loop"
	case "3":
		return "Great Dragon Road"
	case "4":
		return "Aviator Ranch"
	case "5":
		return "Boathouse Bay"
	case "6":
		return "Cannon Crossover"
	case "7":
		return "Forbidden Fortress"
	case "8":
		return "Gold Rush Road"
	case "9":
		return "Badlands Motorplex"
	case "0":
		return "Twisted Canyon"
	default:
		fmt.Println("Invalid input for track, please enter a number from 0 to 9.")
		return SelectTrack()
	}
}
