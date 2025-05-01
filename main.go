package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	var language string
	fmt.Print("Enter programming language: ")
	fmt.Scanln(&language)

	searchRepos(language)
}

func searchRepos(language string) {
	baseURL := "https://api.github.com/search/repositories"

	query := fmt.Sprintf("language:%s good-first-issues:>0", language)
	params := url.Values{}
	params.Add("q", query)
	params.Add("sort", "stars")
	params.Add("order", "desc")
	params.Add("per_page", "5")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error fetching from GitHub:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var result struct {
		Items []struct {
			FullName    string `json:"full_name"`
			Description string `json:"description"`
			Stars       int    `json:"stargazers_count"`
			URL         string `json:"html_url"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		os.Exit(1)
	}

	fmt.Println("\nTop repositories:")
	for _, repo := range result.Items {
		fmt.Printf("\nName: %s\nDescription: %s\nStars: %d\nURL: %s\n", repo.FullName, repo.Description, repo.Stars, repo.URL)
	}
}
