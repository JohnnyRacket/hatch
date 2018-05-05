package sessionmanager

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// SessionManager type holds a session store
type SessionManager struct {
	sessionID     string
	UserCookieKey string
	store         *sessions.CookieStore
}

// NewSessionManager creates a new session store
func NewSessionManager(authenticationKey []byte, encryptionkey []byte) *SessionManager {
	return &SessionManager{
		sessionID:     "user-identity",
		UserCookieKey: "cookie-id",
		store:         sessions.NewCookieStore(authenticationKey, encryptionkey)}
}

// FetchUserIdentityCookieFromRequest (request) returns cookie by name
func (s *SessionManager) FetchUserIdentityCookieFromRequest(r *http.Request) (*sessions.Session, error) {
	session, err := s.store.Get(r, s.sessionID)

	return session, err
}
