package main

import (
	"fmt"
	"net/http"

	"labra-backend/internal/api/routes"
)

const PORT = "8080"

func main() {
	mux := http.NewServeMux()

	routes.Oauth(mux)
	// TODO: probably switch this to TLS
	fmt.Println("Server starting on :", PORT)
	http.ListenAndServe("localhost:"+PORT, mux)
}
