package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lab2/utils/crack"
)

func main() {
	targetHash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"

	file, err := os.Open("nord_vpn.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false
	crackedPassword := ""

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
		if line == "" {
			continue
		}

		lineHash := crack.HashSHA1(line)

		if lineHash == targetHash {
			crackedPassword = line
			found = true
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	outputFile, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	if found {
		fmt.Printf("Password cracked successfully! The plaintext is: %s\n", crackedPassword)
		outputFile.WriteString(fmt.Sprintf("Password cracked successfully! The plaintext is: %s\n", crackedPassword))
	} else {
		fmt.Println("Password not found in the wordlist.")
		outputFile.WriteString("Password not found in the wordlist.\n")
	}
}