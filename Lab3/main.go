package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lab3/utils/crack"
)

func main() {
	targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

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

		lineHash := crack.HashSHA512(line)

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