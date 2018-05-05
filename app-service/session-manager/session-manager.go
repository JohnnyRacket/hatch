package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// Manager type holds a session store
type Manager struct {
	cookieName string
	UserKey    string
	store      *sessions.CookieStore
}

// NewManager creates a new session store
func NewManager(authenticationKey []byte, encryptionkey []byte) *Manager {
	return &Manager{
		cookieName: "user-identity",
		UserKey:    "user-values-id",
		store:      sessions.NewCookieStore(authenticationKey, encryptionkey)}
}

// FetchUserIdentityCookieFromRequest (request) returns cookie by name
func (s *Manager) FetchUserIdentityCookieFromRequest(r *http.Request) (*sessions.Session, error) {
	session, err := s.store.Get(r, s.cookieName)

	return session, err
}
