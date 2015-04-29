package controllers

import (
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
)

func AnnouncementsController(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Announcements, %s!", c.URLParams["site_password"])
}
