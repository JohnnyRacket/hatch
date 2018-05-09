package controllers

import (
	"fmt"
	"hatch/app-service/middleware"
	"hatch/app-service/session-manager"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// AuthenticationController type that holds a sessionmanager
type AuthenticationController struct {
	middleware     *middleware.Manager
	sessionManager *session.Manager
}

// NewAuthenticationController returns new instance of authentication controller.
func NewAuthenticationController(sessionManager *session.Manager) *AuthenticationController {
	return &AuthenticationController{middleware: middleware.NewManager(sessionManager), sessionManager: sessionManager}
}

// RegisterRoutes registers the routes for the controller
func (a *AuthenticationController) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/logout", a.logout).Methods("GET")
	r.HandleFunc("/login", a.initiatelogin).Methods("POST")
	r.HandleFunc("/authenticate", a.authenticate).Methods("GET").Queries("code", "{code}")
	r.HandleFunc("/secretroute", a.middleware.ValidateAuthentication(a.secretroute)).Methods("GET")
}

func (a *AuthenticationController) secretroute(w http.ResponseWriter, r *http.Request) {
	session, err := a.sessionManager.FetchUserIdentityCookieFromRequest(r)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Could not retrieve session", http.StatusInternalServerError)
	}

	fmt.Fprintln(w, fmt.Sprintf("your super secret secret is %s", session.Values[a.sessionManager.UserKey]))
}

func (a *AuthenticationController) initiatelogin(w http.ResponseWriter, r *http.Request) {
	// TODO: Get email from body, send email on to auth API to handle email auth flow
}

func (a *AuthenticationController) authenticate(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if code == "" {
		http.Error(w, "No code passed", http.StatusBadRequest)
	}

	// TODO: trade code with auth server for user access token
	token := uuid.New().String()

	cookie, err := a.sessionManager.FetchUserIdentityCookieFromRequest(r)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Could not retrieve session", http.StatusInternalServerError)
		return
	}

	cookie.Values[a.sessionManager.UserKey] = token

	cookie.Save(r, w)

	// Redirect to app which will now be securely authenticated
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *AuthenticationController) logout(w http.ResponseWriter, r *http.Request) {
	success := handleLogout(a, w, r)
	if !success {
		http.Error(w, "Could not log out", http.StatusInternalServerError)
	}
	// Redirect to app which will now be securely authenticated
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func handleLogout(a *AuthenticationController, w http.ResponseWriter, r *http.Request) bool {
	session, err := a.sessionManager.FetchUserIdentityCookieFromRequest(r)
	if err != nil {
		fmt.Print(err)
		return false
	}

	session.Values[a.sessionManager.UserKey] = ""

	session.Save(r, w)
	return true
}
