package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func FetchInput(day int) {
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		os.Exit(1)
	}

	sessionCookie := getSessionCookie()
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching input for day %d: %v\n", day, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	dirName := fmt.Sprintf("day%02d", day)
	filePath := filepath.Join(dirName, "input.txt")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Error writing input to file %s: %v\n", filePath, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully fetched input for day %d and saved to %s\n", day, filePath)
}

func getSessionCookie() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env")
		os.Exit(1)
	}

	sessionCookie := os.Getenv("SESSION_COOKIE")
	if sessionCookie == "" {
		fmt.Println("Error: SESSION_COOKIE not set")
		os.Exit(1)
	}
	return sessionCookie
}
