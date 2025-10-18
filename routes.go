package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/test-db", testDB)
	}

	r.GET("/", homeHandler)
	r.GET("/login", loginHandler)
	r.GET("/callback", callbackHandler)
	r.GET("/profile", profileHandler)
	r.GET("/logout", logoutHandler)
}

func testDB(c *gin.Context) {
	var now time.Time
	err := DB.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"now": now})
}

func homeHandler(c *gin.Context) {
	session := sessions.Default(c)
	tokenSet := session.Get("tokenSet")

	if tokenSet != nil {
		c.HTML(http.StatusOK, "home_logged_in.html", nil)
	} else {
		c.HTML(http.StatusOK, "home.html", nil)
	}
}

func loginHandler(c *gin.Context) {
	session := sessions.Default(c)

	state, err := generateState()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate state"})
		return
	}

	nonce, err := generateNonce()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate nonce"})
		return
	}

	session.Set("state", state)
	session.Set("nonce", nonce)
	session.Save()

	authURL := OAuth2Config.AuthCodeURL(state, oidc.Nonce(nonce))
	c.Redirect(http.StatusFound, authURL)
}

func callbackHandler(c *gin.Context) {
	session := sessions.Default(c)

	state := session.Get("state")
	if state == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State not found"})
		return
	}

	if c.Query("state") != state.(string) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State mismatch"})
		return
	}

	oauth2Token, err := OAuth2Config.Exchange(c, c.Query("code"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code"})
		return
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No id_token field in oauth2 token"})
		return
	}

	verifier := OIDCProvider.Verifier(&oidc.Config{ClientID: OAuth2Config.ClientID})
	idToken, err := verifier.Verify(c, rawIDToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify ID Token"})
		return
	}

	nonce := session.Get("nonce")
	if nonce == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nonce not found"})
		return
	}

	if idToken.Nonce != nonce.(string) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nonce mismatch"})
		return
	}

	var claims struct {
		Email    string `json:"email"`
		Profile  string `json:"profile"`
		Username string `json:"preferred_username"`
	}
	if err := idToken.Claims(&claims); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse claims"})
		return
	}

	session.Set("tokenSet", oauth2Token)
	session.Set("userinfo", claims)
	session.Save()

	c.Redirect(http.StatusFound, "/")
}

func profileHandler(c *gin.Context) {
	session := sessions.Default(c)
	tokenSet := session.Get("tokenSet")

	if tokenSet == nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	userinfo := session.Get("userinfo")
	c.HTML(http.StatusOK, "profile.html", gin.H{"userinfo": userinfo})
}

func logoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/")
}

// Helper functions for state and nonce generation
func generateState() (string, error) {
	return "random_state", nil // In production, generate a secure random state
}

func generateNonce() (string, error) {
	return "random_nonce", nil // In production, generate a secure random nonce
}
