package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Comic contains fields from XKCD comics
type Comic struct {
	Alt         string `json:"alt"`
	Day         string `json:"day"`
	Img         string `json:"img"`
	Link        string `json:"link"`
	Month       string `json:"month"`
	News        string `json:"news"`
	Num         int    `json:"num"`
	SafeTitle   string `json:"safe_title"`
	Title       string `json:"title"`
	Transcript  string `json:"transcript"`
	Year        string `json:"year"`
}

func main () {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s numXKCD\n", os.Args[0])
		os.Exit(1)
	}
	
	i, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Usage: %s numXKCD\n", os.Args[0])
		os.Exit(1)
	}

	
	comic := getXKCD(fmt.Sprintf("https://xkcd.com/%d/info.0.json", i))

	fmt.Printf("Alt: %s\n", comic.Alt)
	fmt.Printf("Day: %s\n", comic.Day)
	fmt.Printf("Img: %s\n", comic.Img)
	fmt.Printf("Link: %s\n", comic.Link)
	fmt.Printf("Month: %s\n", comic.Month)
	fmt.Printf("News: %s\n", comic.News)
	fmt.Printf("Num: %d\n", comic.Num)
	fmt.Printf("Safe title: %s\n", comic.SafeTitle)
	fmt.Printf("Title: %s\n", comic.Title)
	fmt.Printf("Transcript: %s\n", comic.Transcript)
	fmt.Printf("Year: %s\n", comic.Year)
}

func getXKCD(url string) Comic {

	responseJSON, err := http.Get(url)

	if err != nil {
		fmt.Print(err)
	}
	defer responseJSON.Body.Close()

	byteValue, _ := ioutil.ReadAll(responseJSON.Body)

	var comic Comic

	json.Unmarshal(byteValue, &comic)

	return comic
}