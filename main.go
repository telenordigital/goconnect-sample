package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/telenordigital/goconnect"
)

var templates = template.Must(template.ParseFiles("html/main.html", "html/start.html"))

type templateData struct {
}

// The login page. This is the default start page.
func startPageHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "start.html", templateData{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "main.html", templateData{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Launch the demo service
func main() {
	db, err := sql.Open("sqlite3", "connect-sessions.db?cache=share&mode=twc&_foreign_keys=1&_journal_mode=wal")
	if err != nil {
		panic(err)
	}
	config := goconnect.NewDefaultConfig(goconnect.ClientConfig{
		Host:                      goconnect.StagingHost,
		ClientID:                  "telenordigital-connectexample-web",
		Password:                  "",
		LoginCompleteRedirectURI:  "/main.html",
		LogoutCompleteRedirectURI: "/",
	})
	// Note: Ignoring errors since the store might already exist
	goconnect.SQLStorePrepare(db)

	dbstore, err := goconnect.NewSQLStorage(db)
	if err != nil {
		panic(err)
	}
	connect := goconnect.NewConnectIDWithStorage(config, dbstore)

	// The /css and /images endpoints
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("html/css"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("html/images"))))

	// This the default start page
	http.HandleFunc("/", startPageHandler)

	// Main page - requires authentication
	http.HandleFunc("/main.html", connect.NewAuthHandlerFunc(mainPageHandler))

	// A protected resource - requires authentication
	http.Handle("/extra/", connect.NewAuthHandler(
		http.StripPrefix("/extra/", http.FileServer(http.Dir("html/extra")))))

	// API endpoint - requires authentication
	http.HandleFunc("/api/oneliner", connect.NewAuthHandlerFunc(OneLinerHandlerFunc))

	// This isn't required since the /auth/info endpoint provides the same information but
	// this is just to show an example of how session information can be used in resources.
	http.HandleFunc("/api/userinfo", connect.NewAuthHandlerFunc(UserInfoHandlerFunc))

	// The ConnectID endpoints is added to /auth/. The actual endpoint depends on your
	// client configuration in ConnectID; ie what redirects are required.
	http.Handle("/connect/", connect.Handler())

	// Show the logged in user's properties.
	http.HandleFunc("/connect/profile", connect.SessionProfile)

	fmt.Println("Serving data on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
