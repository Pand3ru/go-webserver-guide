package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	userCount = 0
	store     = sessions.NewCookieStore([]byte("test"))
)

func configureStore() {
	store.Options = &sessions.Options{
		Path:     "/",      // Valid under all paths
		MaxAge:   3600 * 8, // Expiries in 8 hours
		HttpOnly: true,     // Enhances security by making cookie inaccessible to JavaScript
		Secure:   true,     // Ensures the cookie is sent over HTTPS
	}
}

func main() {
	configureStore()
	router()
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have reached the homepage")
}

func router() {
	http.HandleFunc("/", homepage)
	http.Handle("/middleware", middleware(homepage))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}

func middleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session")
		if err != nil {
			log.Println("Error retrieving session: ", err)
		}

		if session.Values["UserCount"] == nil {
			session.Values["UserCount"] = userCount
			userCount++
			err = session.Save(r, w)
			if err != nil {
				log.Println("Could not save session: ", err)
				return
			}
		}
		log.Printf("Function called by %d's User", session.Values["UserCount"])
		next.ServeHTTP(w, r)
	})
}
