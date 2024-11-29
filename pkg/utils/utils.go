package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func FetchInput(day, year int, outFile, outDir string) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		handleError(fmt.Errorf("error creating request: %v", err))
	}

	sessionCookie := getSessionCookie()

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	resp, err := client.Do(req)
	if err != nil {
		handleError(fmt.Errorf("error fetching input for day %d: %v", day, err))
	}
	defer resp.Body.Close()

	filePath := filepath.Join(outDir, outFile)

	file, err := os.Create(filePath)
	if err != nil {
		handleError(fmt.Errorf("error creating file %s: %v", filePath, err))
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		handleError(fmt.Errorf("error writing input to file %s: %v", filePath, err))
	}

	fmt.Printf("Successfully fetched input for day %d and saved to %s\n", day, filePath)
}

func getSessionCookie() string {
	godotenv.Load()
	sessionCookie := os.Getenv("SESSION_COOKIE")
	if sessionCookie == "" {
		handleError(fmt.Errorf("error: SESSION_COOKIE not set"))
	}
	return sessionCookie
}

func handleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
