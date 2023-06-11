package handlers

import (
	"monolith/security"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	secret             = os.Getenv("SECRET")
	accessTokenExpiry  = time.Hour * 1
	refreshTokenExpiry = time.Hour * 24 * 7
)

func (h Handler) Refresh(c *gin.Context) {
	var refreshToken string
	for _, cookie := range c.Request.Cookies() {
		if cookie.Name == "refresh_token" {
			refreshToken = cookie.Value
			break
		}
	}

	if refreshToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	claims, err := security.ValidateToken(refreshToken, secret, security.RefreshToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	newAccessToken, err := security.CreateToken(claims.Email, secret, accessTokenExpiry, security.AccessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newRefreshToken, err := security.CreateToken("", secret, refreshTokenExpiry, security.RefreshToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    newRefreshToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}
