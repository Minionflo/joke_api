package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var jokes []string = nil

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	var joke string = jokes[rand.Intn(len(jokes))]
	jokee := `{"body":"` + joke + `"}`
	fmt.Fprintf(w, jokee)
}

func main() {
	var lines int = 0
	file, err := os.Open("./jokes.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			jokes = append(jokes, line)
			lines++
		}
	}
	file.Close()
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
