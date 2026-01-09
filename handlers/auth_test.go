package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/invertedbit/gms/auth"
)

func TestHandleLoginView_SessionCookieSet(t *testing.T) {
	// Initialize session store
	auth.SessionStore = *session.New()

	// Create a test app
	app := New()

	// First request - should set session cookie
	req := httptest.NewRequest("GET", "/auth/login", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check if session cookie is set
	cookies := resp.Cookies()
	var sessionCookie bool
	for _, cookie := range cookies {
		if cookie.Name == "session_id" {
			sessionCookie = true
			if cookie.Value == "" {
				t.Error("Session cookie value is empty")
			}
			break
		}
	}

	if !sessionCookie {
		t.Error("Expected session cookie to be set on first visit, but it was not")
	}

	// Second request with the session cookie - should recognize returning visitor
	req2 := httptest.NewRequest("GET", "/auth/login", nil)
	for _, cookie := range cookies {
		req2.AddCookie(cookie)
	}
	resp2, err := app.Test(req2, -1)
	if err != nil {
		t.Fatalf("Failed to send second request: %v", err)
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", resp2.StatusCode)
	}
}
