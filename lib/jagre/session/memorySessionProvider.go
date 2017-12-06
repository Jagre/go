package session

import (
	"container/list"
	"sync"
	"time"
)

// SessionStore struct
type SessionStore struct {
	sid      string
	accessed time.Time
	value    map[interface{}]interface{}
}

var MemorySessionManager *Manager

// Set is the method of SessionStore
func (store *SessionStore) Set(key, value interface{}) error {
	store.value[key] = value
	mProvider.SessionUpdate(store.sid)
	return nil
}

// Get session
func (store *SessionStore) Get(key interface{}) interface{} {
	mProvider.SessionUpdate(store.sid)
	if v, ok := store.value[key]; ok {
		return v
	}
	return nil
}

func (store *SessionStore) Del(key interface{}) error {
	delete(store.value, key)
	mProvider.SessionUpdate(store.sid)
	return nil
}

func (store *SessionStore) SessionID() string {
	return store.sid
}

// MemoryProvider struct
type MemoryProvider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element //Save sessionstore in memory
	list     *list.List               //Save as GC
}

var mProvider = &MemoryProvider{list: list.New()}

func (mProvider *MemoryProvider) SessionInit(sid string) (Session, error) {
	mProvider.lock.Lock()
	defer mProvider.lock.Unlock()

	v := make(map[interface{}]interface{}, 0)
	store := &SessionStore{sid: sid, accessed: time.Now(), value: v}
	elem := mProvider.list.PushBack(store)
	mProvider.sessions[sid] = elem

	return store, nil
}

func (mProvider *MemoryProvider) SessionRead(sid string) (Session, error) {
	if elem, ok := mProvider.sessions[sid]; ok {
		return elem.Value.(*SessionStore), nil
	}
	return mProvider.SessionInit(sid)
}

func (mProvider *MemoryProvider) SessionDestroy(sid string) error {
	if elem, ok := mProvider.sessions[sid]; ok {
		delete(mProvider.sessions, sid)
		mProvider.list.Remove(elem)
	}
	return nil
}

func (mProvider *MemoryProvider) SessionGC(maxLifeTime int64) {
	mProvider.lock.Lock()
	defer mProvider.lock.Unlock()

	for {
		if elem := mProvider.list.Back(); elem == nil {
			break
		} else if elem.Value.(*SessionStore).accessed.Unix()+maxLifeTime < time.Now().Unix() {
			delete(mProvider.sessions, elem.Value.(*SessionStore).sid)
			mProvider.list.Remove(elem)
		} else {
			break
		}
	}
}

// SessionUpdate is Update the session accessed time & move the fornt of sessionStore
func (mProvider *MemoryProvider) SessionUpdate(sid string) error {
	mProvider.lock.Lock()
	defer mProvider.lock.Unlock()
	if elem, ok := mProvider.sessions[sid]; ok {
		//update access time
		elem.Value.(*SessionStore).accessed = time.Now()
		//Move the element to the first
		mProvider.list.MoveToFront(elem)
	}
	return nil
}

func init() {
	//Register provider
	mProvider.sessions = make(map[string]*list.Element, 0)
	Register("memory", mProvider)

	//Init Session Manager
	MemorySessionManager, _ = NewManager("memory", "goSessionId", 3600)
	go MemorySessionManager.GC()
}
