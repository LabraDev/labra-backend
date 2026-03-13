package handlers

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Authenticate gh here

	gh_client := os.Getenv("GH_CLIENT_ID")
	gh_secret := os.Getenv("GH_CLIENT_SECRET")

	fmt.Println(gh_client, "\n---\n", gh_secret)
	// Todo: replace redirect host
	ouathConfig := &oauth2.Config{
		ClientID:     gh_client,
		ClientSecret: gh_secret,
		Scopes:       []string{"repo", "user"},
		Endpoint:     github.Endpoint,
		RedirectURL:  "http://localhost:8080/callback",
	}

	// This protects against CSRF attacks
	verifier := oauth2.GenerateVerifier()

	url := ouathConfig.AuthCodeURL("state", oauth2.S256ChallengeOption(verifier))
	fmt.Printf("Please visit %v for the auth\n", url)

	// http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
