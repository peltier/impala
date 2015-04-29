package main

import (
	"regexp"

	"github.com/film42/impala/controllers"
	"github.com/zenazn/goji"
)

func SetupRoutesAndStartServer() {
	announcementsExpression := regexp.MustCompile("^/(?P<site_password>[a-zA-Z0-9]{32})/announce")
	goji.Get(announcementsExpression, controllers.AnnouncementsController)

	reportsExpression := regexp.MustCompile("^/(?P<site_password>[a-zA-Z0-9]{32})/report")
	goji.Get(reportsExpression, controllers.ReportsController)

	scrapesExpression := regexp.MustCompile("^/(?P<site_password>[a-zA-Z0-9]{32})/scrape")
	goji.Get(scrapesExpression, controllers.ScrapesController)

	updatesExpression := regexp.MustCompile("^/(?P<site_password>[a-zA-Z0-9]{32})/update")
	goji.Get(updatesExpression, controllers.UpdatesController)

	goji.NotFound(controllers.ErrorsController)

	goji.Serve()
}
