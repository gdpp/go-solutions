package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Milestone Project: Video Stats (Simplified)
	// Create a program that analyzes the statistics of a video and determines if it was successful.

	channelName := "Go Academy"

	var videoTitle string = "Go Programming Tutorial"
	var views int = 1250
	var likes int = 180
	var comments int = 45
	// var isMonetized bool = true

	engagement := (float64(likes+comments) / float64(views)) * 100

	fmt.Println("--- VIDEO STATS ANALYZER ---")
	fmt.Println(strings.ToUpper("channel:" + channelName))
	fmt.Println(strings.ToUpper("video:" + videoTitle))
	fmt.Printf("Engagement %.2f%%\n", engagement)

	fmt.Println(math.Sqrt(float64(views)))

	if views > 1000 && engagement > 10 {
		fmt.Println("This video is a success!")
	} else {
		fmt.Println("This video needs improvement.")
	}

	if totalInteractions := likes + comments; totalInteractions > 200 {
		fmt.Printf("Total interactions: %d - Great engagement!\n", totalInteractions)
	} else {
		fmt.Printf("Total interactions: %d - Could be better.\n", totalInteractions)
	}

	const daysInWeek = 5
	for i := 1; i <= daysInWeek; i++ {
		views += 100
	}

	fmt.Printf("Projected views after %d days: %d\n", daysInWeek, views)

	switch {
	case engagement >= 20:
		fmt.Println("Viral")
	case engagement >= 10:
		fmt.Println("Popular")
	case engagement >= 5:
		fmt.Println("Growing")
	default:
		fmt.Println("new")
	}
}
