package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type ApiResponse struct {
	Result string             `json:"result"`
	Rates  map[string]float64 `json:"conversion_rates"`
}

func main() {
	// Check if correct number of arguments are provided
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go [from_currency] [to_currency] [amount]")
		fmt.Println("Example: go run main.go USD INR 10.5")
		return
	}

	apiKey := "fcbdb75d8c5cde8f95a76c5d" // Note: In production, use environment variables
	from := os.Args[1]
	to := os.Args[2]

	// Convert amount argument to float64
	amount, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println(" Invalid amount. Please enter a valid number!!")
		return
	}

	// Validate currency codes (should be 3 letters)
	if len(from) != 3 || len(to) != 3 {
		fmt.Println(" Currency codes must be 3 letters (e.g., USD, EUR, INR)!!")
		return
	}

	// Get latest rates for base currency
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/%s", apiKey, from)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(" Request error xx:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("API returned status xx:", resp.Status)
		return
	}

	var result ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("xx Decode error:", err)
		return
	}

	if result.Result != "success" {
		fmt.Println("xx API response was unsuccessful.")
		return
	}

	rate, exists := result.Rates[to]
	if !exists {
		fmt.Printf("xx Could not find exchange rate for %s\n", to)
		return
	}

	convertedAmount := amount * rate
	fmt.Printf("\n💱 Currency Conversion Result:\n")
	fmt.Printf("➡️ %.2f %s = %.2f %s\n", amount, from, convertedAmount, to)
	fmt.Printf("🔁 Exchange Rate: 1 %s = %.4f %s\n", from, rate, to)
}
