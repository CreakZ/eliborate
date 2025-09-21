package responses

import (
	"net/http"
)

func NewJwtCookieResponse(jwt string) *http.Cookie {
	return &http.Cookie{
		Name:     "jwt",
		Value:    jwt,
		Path:     "/",
		Domain:   "",    // fill this later
		Secure:   false, // temporary false
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}
