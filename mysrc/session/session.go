package main

import (
	"html/template"
	"jagre/session"
	"net/http"
)

func main6() {
	println("Begin")
	http.HandleFunc("/Counter", Counter)
	http.ListenAndServe(":2323", nil)
	print("End")
}

// Counter is counting accessed no.
func Counter(w http.ResponseWriter, r *http.Request) {

	// sManager, _ := session.NewManager("memory", "gSessionId", 3600)
	// session := sManager.SessionStart(w, r)
	mSession := session.MemorySessionManager.SessionStart(w, r)
	count := mSession.Get("accessCount")
	if count == nil {
		mSession.Set("accessCount", 1)
	} else {
		mSession.Set("accessCount", count.(int)+1)
	}
	t, _ := template.ParseFiles("counter.html")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, mSession.Get("accessCount"))
}
