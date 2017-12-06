package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Session interface
// Defined all actions for session
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Del(key interface{}) error
	SessionID() string
}

// Provider interface
// Defined all operation for sesion
type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

var providers = make(map[string]Provider)

// Register makes a session provide available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(sessionName string, provider Provider) {
	if provider == nil {
		panic("Session provider is nil")
	}

	if _, existed := providers[sessionName]; existed {
		panic("Session provider has been registed")
	}
	providers[sessionName] = provider
}

// Manager struct
type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

// NewManager is a global manager of the session
func NewManager(providerName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}
	// provider := providers[providerName]
	// if provider == nil {
	// 	return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	// }
	return &Manager{cookieName: cookieName, provider: provider, maxLifeTime: maxLifeTime}, nil
}

// SessionID create a new id for session
func (manager *Manager) SessionID() string {
	b := make([]byte, 32)
	_, e := io.ReadFull(rand.Reader, b)
	if e != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// SessionStart is Manager.SessionStart
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	//manager.lock.Lock()
	//defer manager.lock.Unlock()

	cookie, e := r.Cookie(manager.cookieName)
	if e != nil || cookie.Value == "" {
		//Have error or no cookie your specified
		sid := manager.SessionID()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

// SessionDestroy is Manager.SessionDestroy
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, e := r.Cookie(manager.cookieName)
	if e != nil || cookie.Value == "" {
		return
	}
	manager.lock.Lock()
	defer manager.lock.Unlock()
	sid, _ := url.QueryUnescape(cookie.Value)
	manager.provider.SessionDestroy(sid)
	expiration := time.Now()
	cookie2 := http.Cookie{
		Name:     manager.cookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  expiration,
		MaxAge:   int(manager.maxLifeTime)}
	http.SetCookie(w, &cookie2)
}

// GC is Manager.GC
func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Lock()

	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime)*time.Second, func() { manager.GC() })
}
