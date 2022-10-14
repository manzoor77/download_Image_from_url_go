package main

import (
	"net/http"
	"os"
)

func main() {
	img, _ := os.Create("profile.png")

	defer img.Close()

	resp, _ := http.Get("https://cdn.pixabay.com/photo/2014/12/22/00/07/tree-576847__480.png")
	defer resp.Body.Close()

	// b, _ := io.Copy(img, resp.Body)
	// fmt.Println("File size: ", b)
}
