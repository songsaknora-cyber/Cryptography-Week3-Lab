package main

import (
	"bufio"
	"fmt"
	"os"

	"lab1/utils/crack"
)

func main() {
	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"

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
		line := scanner.Text()
		
		lineHash := crack.HashMD5(line)

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

	fmt.Println("Verbose output saved to verbose.txt")
}