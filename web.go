package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/grsakea/kappastat/backend"
	"html/template"
	"log"
	"net/http"
)

var Backend *backend.Controller
var templates = template.Must(template.ParseFiles("templates/following.html",
	"templates/viewer.html",
	"templates/stat.html",
	"templates/index.html",
	"templates/head.inc",
	"templates/header.inc"))

func launchFrontend(c *backend.Controller) {
	m := martini.Classic()
	Backend = c
	m.Use(martini.Static("static"))
	m.Get("/", indexHandler)
	m.Get("/following", followHandler)
	m.Get("/stat", statHandler)
	m.Get("/viewer", viewerHandler)
	m.Get("/add/:streamer", addHandler)
	m.Get("/api/viewer/:streamer", apiViewer)
	m.Get("/api/stat/:streamer", apiStat)
	m.Get("/api/following", apiFollowing)

	log.Print("Started Web Server")
	m.Run()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func followHandler(w http.ResponseWriter, r *http.Request) {
	liste := Backend.ListStreams()
	templates.ExecuteTemplate(w, "following.html", liste)
}

func viewerHandler(w http.ResponseWriter, r *http.Request) {
	views := []backend.ViewerCount{}
	templates.ExecuteTemplate(w, "viewer.html", views)
}

func statHandler(w http.ResponseWriter, r *http.Request) {
	views := []backend.ViewerCount{}
	templates.ExecuteTemplate(w, "stat.html", views)
}

func addHandler(w http.ResponseWriter, r *http.Request, params martini.Params) {
	Backend.AddStream(params["streamer"])
	fmt.Fprintf(w, "Added %s", params["streamer"])
}
