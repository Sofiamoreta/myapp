package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Tillåt att ange port som flagga: -port=8081
	portFlag := flag.String("port", "", "port to serve on")
	flag.Parse()

	// Om ingen flagga används, kolla PORT miljövariabeln
	port := *portFlag
	if port == "" {
		port = os.Getenv("PORT")
	}
	if port == "" {
		port = "8080" // standardport
	}

	fmt.Printf("Server is running on port %s\n", port)

	// Enkel HTTP-handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hej från Go-servern! 🚀")
	})

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
