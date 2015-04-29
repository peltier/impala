package controllers

import (
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
)

func ErrorsController(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing to see here. Move along.")
}
