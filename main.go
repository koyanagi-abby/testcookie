package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/login", loginHandler)

	if err := http.ListenAndServe(":4000", mux); err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Header)

	beforeLogin := []byte(`
<html>
<body>
  <h1>Before login</h1>
  <a href="/login">Login</a>
</body>
</html>
`)

	afterLogin := []byte(`
	   <html>
	   <body>
	     <h1>After login</h1>
	   </body>
	   </html>
	   `)

	cookie, err := r.Cookie("session")

	if err != nil || cookie.Value == "" {
		w.Write(beforeLogin)
	} else {
		w.Write(afterLogin)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Header)

	w.Header().Set("Set-Cookie", "session=abcdef123456; HttpOnly; Secure; SameSite=Lax")
	w.Header().Set("Location", "https://github.com")
	w.WriteHeader(http.StatusSeeOther)
}
