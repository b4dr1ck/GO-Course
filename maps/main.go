package main

import "fmt"

// verwende Type-Alias
type urlsMap map[string]string

func main() {
	/*
		 	// ohne Alias
			urls := map[string]string{
				"google":    "https://www.google.com/",
				"homepage":  "https://badricks-world.at/",
				"localhost": "http://127.0.0.1/",
			}
	*/
	urls := urlsMap{
		"google":    "https://www.google.com/",
		"homepage":  "https://badricks-world.at/",
		"localhost": "http://127.0.0.1/",
	}

	fmt.Println(urls)

	urls["deviantart"] = "https://www.deviantart.com/"

	delete(urls, "google")

	fmt.Println(urls)
}
