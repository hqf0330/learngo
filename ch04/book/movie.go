package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year  int  `json:"released"`
	Color bool `json:"color,omitempty"`
	Actor []string
}

func main() {
	movies := []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actor: []string{"Humphrey Bogart", "Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actor: []string{"Paul Newman"}},
		{Title: "Bullet", Year: 1968, Color: false, Actor: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	fmt.Println("---------------------------------------------------------------------")

	data2, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data2)

	fmt.Println("---------------------------------------------------------------------")

	var titles []struct{ Title string }
	if err = json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
