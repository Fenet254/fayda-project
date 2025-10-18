package main

import (
	"context"
	"log"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var (
	OIDCProvider *oidc.Provider
	OAuth2Config oauth2.Config
)

func InitOIDC() {
	issuerURL := os.Getenv("OIDC_ISSUER_URL")
	if issuerURL == "" {
		issuerURL = "https://esignet.ida.fayda.et"
	}

	ctx := context.Background()

	var err error
	OIDCProvider, err = oidc.NewProvider(ctx, issuerURL)
	if err != nil {
		log.Fatal("❌ Failed to initialize OIDC provider:", err)
	}

	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		clientID = "dummy_client_id"
	}

	redirectURI := os.Getenv("REDIRECT_URI")
	if redirectURI == "" {
		redirectURI = "http://localhost:3000/callback"
	}

	OAuth2Config = oauth2.Config{
		ClientID:    clientID,
		RedirectURL: redirectURI,
		Scopes:      []string{oidc.ScopeOpenID, "profile", "email"},
		Endpoint:    OIDCProvider.Endpoint(),
	}

	log.Println("✅ OIDC Client is ready")
}
