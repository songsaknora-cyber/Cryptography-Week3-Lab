package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func hashString(input string, algorithm string) string {
	var hashBytes []byte
	data := []byte(input)

	switch algorithm {
	case "MD5":
		hash := md5.Sum(data)
		hashBytes = hash[:]
	case "SHA1":
		hash := sha1.Sum(data)
		hashBytes = hash[:]
	case "SHA256":
		hash := sha256.Sum256(data)
		hashBytes = hash[:]
	case "SHA512":
		hash := sha512.Sum512(data)
		hashBytes = hash[:]
	}
	return hex.EncodeToString(hashBytes)
}

func proofMe(input1, input2 string) {
	fmt.Println("\n======== Name + Hashing Program ========")

	algorithms := []string{"MD5", "SHA1", "SHA256", "SHA512"}

	for _, algo := range algorithms {

		outputA := hashString(input1, algo)
		outputB := hashString(input2, algo)

		matchStatus := "No Match!"
		if outputA == outputB {
			matchStatus = "Match!"
		}

		fmt.Printf("Hash (%s): Output A / output B => %s\n", algo, matchStatus)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please input value 1: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	fmt.Print("Please input value 2: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)


	proofMe(input1, input2)
}