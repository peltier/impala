package controllers

import (
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
)

func UpdatesController(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updates, %s!", c.URLParams["site_password"])
}
