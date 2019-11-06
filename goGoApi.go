package main

import (
	"fmt"
	"github.com/blhack/goGoApi/config"
	"github.com/blhack/goGoApi/database"
	"github.com/blhack/goGoApi/routes"
	"github.com/blhack/goGoApi/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"net/http"
	"net/url"
	"os"
)

func setCacheHeader(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=86400")
		h.ServeHTTP(w, r)
	}
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+config.Domain+req.RequestURI, 302)
}

func defaultHeaders(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		Origin, err := url.Parse(r.Header.Get("Origin"))
		utils.CheckErr(err)

		if config.WhiteListOrigins[Origin.Host] {
			w.Header().Add("Access-Control-Allow-Headers", "origin, content-type, credentials")
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		}
		h(w, r)
	}
}

func main() {
	var err error

	utils.SayHi()

	f, err := os.OpenFile("./logfile.log", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	database.Init()

	if err != nil {
		panic(err)
	}

	var port = os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.Handle("/", setCacheHeader(handlers.CombinedLoggingHandler(f, http.FileServer(http.Dir("./static")))))

	http.HandleFunc("/auth/login", defaultHeaders(routes.Auth))
	http.HandleFunc("/auth/logout", defaultHeaders(routes.DeAuth))
	http.HandleFunc("/auth/whoami", defaultHeaders(routes.WhoAmI))
	http.HandleFunc("/auth/register", defaultHeaders(routes.Register))
	http.HandleFunc("/things/addThing", defaultHeaders(routes.AddThing))
	http.HandleFunc("/things/getThings", defaultHeaders(routes.GetThings))

	go http.ListenAndServe(fmt.Sprintf(":%s", config.HttpPort), http.HandlerFunc(redirect))
	err = http.ListenAndServeTLS(":443", config.CertFile, config.KeyFile, nil)
	utils.CheckErr(err)
}
