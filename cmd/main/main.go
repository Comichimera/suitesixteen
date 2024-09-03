package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"suitesixteen/internal/file"
	"suitesixteen/internal/github"
	"suitesixteen/internal/input"
)

func main() {
	currentDateTime := time.Now().Format("20060102150405")
	filename := currentDateTime + ".csv"

	if file.Exists(filename) {
		fmt.Printf("The file %s already exists. Please delete the file to proceed or wait for the next second to log more data.\n", filename)
		return
	}

	csvFile, err := file.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer csvFile.Close()

	writer := bufio.NewWriter(csvFile)
	defer writer.Flush()

	gameMode := input.SelectGameMode()
	if gameMode == "" {
		fmt.Println("Invalid input for game mode, please enter 1, 2, or 3.")
		return
	}

	track := input.SelectTrack()

	var committedRecords []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nStart entering leaderboard data. Type 'commit' to save, 'rollback' to delete the last committed record, 'track' to change the track, 'mode' to change the game mode, or 'push' to write all records to the file and exit.")

	var currentRecord []string

	for {
		if scanner.Scan() {
			userInput := scanner.Text()
			userInput = strings.TrimSpace(userInput)

			switch strings.ToLower(userInput) {
			case "commit":
				if len(currentRecord) == 0 {
					fmt.Println("No data to commit.")
					continue
				}
				record := fmt.Sprintf("%s,%s,%s\n", gameMode, track, strings.Join(currentRecord, ","))
				committedRecords = append(committedRecords, record)
				currentRecord = nil
				fmt.Println("Record committed.")

			case "rollback":
				if len(committedRecords) > 0 {
					committedRecords = committedRecords[:len(committedRecords)-1]
					fmt.Println("Last committed record has been rolled back.")
				} else {
					fmt.Println("No records to rollback.")
				}
				currentRecord = nil

			case "push":
				for _, record := range committedRecords {
					_, err := writer.WriteString(record)
					if err != nil {
						fmt.Println("Error writing to file:", err)
						return
					}
				}
				writer.Flush()
				fmt.Println("All records saved.")

				fmt.Print("Enter a commit message for GitHub: ")
				if scanner.Scan() {
					commitMessage := scanner.Text()
					commitMessage = strings.TrimSpace(commitMessage)

					owner := "Comichimera"
					repo := "datatesting"
					path := "data/uploads/" + filename
					token := os.Getenv("GITHUB_TOKEN")

					fmt.Println("Token: ", token)

					err := github.UploadFileToGitHub(owner, repo, path, token, commitMessage, filename)
					if err != nil {
						fmt.Printf("Error uploading file: %v\n", err)
					}
				}

				return

			case "track":
				track = input.SelectTrack()
				fmt.Println("Track changed. Continue entering data.")

			case "mode":
				gameMode = input.SelectGameMode()
				fmt.Println("Game mode changed. Continue entering data.")

			default:
				if len(userInput) > 0 {
					currentRecord = append(currentRecord, userInput)
				}
			}
		}
	}
}
