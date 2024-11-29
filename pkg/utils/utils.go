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
		handleError(fmt.Errorf("Error creating request: %v", err))
	}

	sessionCookie := getSessionCookie()

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	resp, err := client.Do(req)
	if err != nil {
		handleError(fmt.Errorf("Error fetching input for day %d: %v", day, err))
	}
	defer resp.Body.Close()

	dirName := fmt.Sprintf("day%02d", day)
	filePath := filepath.Join(dirName, "input.txt")

	file, err := os.Create(filePath)
	if err != nil {
		handleError(fmt.Errorf("Error creating file %s: %v\n", filePath, err))
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		handleError(fmt.Errorf("Error writing input to file %s: %v\n", filePath, err))
	}

	fmt.Printf("Successfully fetched input for day %d and saved to %s\n", day, filePath)
}

func getSessionCookie() string {
	err := godotenv.Load()
	if err != nil {
		handleError(fmt.Errorf("Error loading .env"))
	}

	sessionCookie := os.Getenv("SESSION_COOKIE")
	if sessionCookie == "" {
		handleError(fmt.Errorf("Error: SESSION_COOKIE not set"))
	}
	return sessionCookie
}

func handleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
