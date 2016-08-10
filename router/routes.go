package router

import (
	"github.com/shandilyamanvesh/CrudAPIInGo/controllers"
)

var muxroutes = routes{
	route{
		"Index",
		"GET",
		"/",
		controllers.Index,
	},
	route{
		"BlogIndex",
		"GET",
		"/blogs",
		controllers.BlogIndex,
	},
	route{
		"BlogShow",
		"GET",
		"/blogs/{blogId}",
		controllers.Show,
	},
	route{
		"BlogSave",
		"POST",
		"/blogs",
		controllers.Save,
	},
}
