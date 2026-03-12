package main

import (
	"fmt"
	"net/http"

	"labra-backend/internal/api/routes"

	"github.com/lpernett/godotenv"
)

const PORT = "8080"

func main() {
	err := godotenv.Load("./../.env")
	if err != nil {
		fmt.Println(err)
	}

	mux := http.NewServeMux()

	routes.Oauth(mux)
	// TODO: probably switch this to TLS
	fmt.Println("Server starting on :", PORT)
	http.ListenAndServe("localhost:"+PORT, mux)
}
