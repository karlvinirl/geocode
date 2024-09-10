package main

import "fmt"

func main() {

	// website := map[string]string{}
	websites2 := map[string]string{
		"Google": "google.com", "Amazon": "amazon.com",
	}
	fmt.Println(websites2)
	websites2["LinkedIn"] = "linkedin.com"
	fmt.Println(websites2)
	delete(websites2, "Google")
	fmt.Println(websites2)

}
