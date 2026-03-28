package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", HandleForm)

	port := "8080"

	fmt.Printf("server started on port: %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		fmt.Println(err)
	}
}

func HandleForm(res http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		fmt.Fprintf(res, "not the expected request method")
		return
	}

	content := req.FormValue("content")
	font := req.FormValue("font")

	// Run our Ascii Art function here and continue below.
	art, err := Art(content, font)
	if err != nil {
		fmt.Fprintf(res, "Could not process art!\n%s", err)
		fmt.Println(err)
	}

	fmt.Fprintf(res, "%s\n", art)
}