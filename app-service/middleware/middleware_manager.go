package middleware

import (
	"fmt"
	"hatch/app-service/session-manager"
	"net/http"
)

// Manager holds the session and provides middleware functions
type Manager struct {
	sessionmanager *sessionmanager.SessionManager
}

// NewManager creates a new manager instance with a sessionmanager internally
func NewManager(sessionmanager *sessionmanager.SessionManager) *Manager {
	return &Manager{sessionmanager: sessionmanager}
}

// ValidateAuthentication of the request
func (m *Manager) ValidateAuthentication(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := m.sessionmanager.FetchUserIdentityCookieFromRequest(r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Could not retrieve session", http.StatusInternalServerError)
			return
		}

		cookieID, ok := cookie.Values[m.sessionmanager.UserCookieKey].(string)
		if !ok || cookieID == "" {
			fmt.Println("no cookie found, sad cookie monster")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		fmt.Println(cookieID)

		f(w, r)
	}
}

// ValidateAuthenticationForHandler wraps authentication for routes requiring handlers
func (m *Manager) ValidateAuthenticationForHandler(h http.Handler) http.Handler {
	return m.ValidateAuthentication(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
