package handlers

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Authenticate gh here

	// Todo: replace redirect host
	ouathConfig := &oauth2.Config{
		ClientID:     "GH_CLIENT_ID",
		ClientSecret: "GH_CLIENT_SECRET",
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
