package controllers

import (
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
)

func ScrapesController(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Scrapes, %s!", c.URLParams["site_password"])
}
