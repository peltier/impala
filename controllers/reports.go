package controllers

import (
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
)

func ReportsController(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Reports, %s!", c.URLParams["site_password"])
}
