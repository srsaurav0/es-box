package routers

import (
	"es-box/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.SearchController{}, "get:Get")                  // Get the search page
	web.Router("/search", &controllers.SearchController{}, "post:SearchResults") // Handles POST request for the search form
}
