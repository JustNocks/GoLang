package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type MediaContainer struct {
	Size   int     `json:"size"`
	Videos []Video `xml:"Video"`
}

type Video struct {
	Title     string `xml:"title,attr"`
	RatingKey string `xml:"ratingKey,attr"`
}

func main() {

	// Get the absolute path to the directory where the Go source file is located
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	// Construct the path to the urls.txt file relative to the .go file
	urlsPath := filepath.Join(dir, "data", "urls.txt")

	// Open the urls.txt file
	urlsFile, err := os.Open(urlsPath)
	if err != nil {
		fmt.Println("Error opening urls.txt:", err)
		return
	}
	defer urlsFile.Close()

	scanner := bufio.NewScanner(urlsFile)
	var urls []string
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if len(urls) != 2 {
		panic("Expected exactly 2 URLs in the file")
	}

	// Send API requests to both servers and decode the XML response
	res1, err1 := http.Get(urls[0])
	fmt.Println(res1)
	if err1 != nil {
		panic(err1)
	}
	defer res1.Body.Close()

	res2, err2 := http.Get(urls[1])
	fmt.Println(res2)
	if err2 != nil {
		panic(err2)
	}
	defer res2.Body.Close()

	var container1 MediaContainer
	err3 := xml.NewDecoder(res1.Body).Decode(&container1)
	if err3 != nil {
		panic(err3)
	}

	var container2 MediaContainer
	err4 := xml.NewDecoder(res2.Body).Decode(&container2)
	if err4 != nil {
		panic(err4)
	}

	// Find duplicates by comparing the rating keys
	var duplicates []string
	for _, video1 := range container1.Videos {
		for _, video2 := range container2.Videos {
			if video1.RatingKey == video2.RatingKey {
				duplicates = append(duplicates, video1.Title)
				break
			}
		}
	}

	// Print the duplicates
	if len(duplicates) > 0 {
		fmt.Println("Duplicates found:")
		for _, title := range duplicates {
			fmt.Println(title)
		}
	} else {
		fmt.Println("No duplicates found")
	}
	// Write the duplicates to a file
	if len(duplicates) > 0 {
		file, err := os.Create("duplicates.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush()

		for _, title := range duplicates {
			fmt.Fprintln(writer, title)
		}

		fmt.Println("Duplicates written to duplicates.txt")
	} else {
		fmt.Println("No duplicates found")
	}
}
