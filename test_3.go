package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)
func fetchBeefData() (string, error) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// CountBeef takes raw text and returns the count of each type of beef
func countBeef(text string) map[string]int {
	re := regexp.MustCompile(`\b[\w-]+\b`)
	words := re.FindAllString(strings.ToLower(text), -1)

	beefCount := make(map[string]int)
	for _, word := range words {
		beefCount[word]++
	}
	return beefCount
}


func summaryHandler(c *gin.Context) {
	text, err := fetchBeefData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	beefCount := countBeef(text)
	c.JSON(http.StatusOK, gin.H{"beef": beefCount})
}

func main() {
	router := gin.Default()
	router.GET("/beef/summary", summaryHandler)

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
