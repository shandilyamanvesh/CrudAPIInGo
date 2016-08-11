package router

import (
	"github.com/shandilyamanvesh/CrudAPIInGo/controllers"
)

// All routes with their corresponding name,verb,url and http handler methods.
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
		"BlogCreate",
		"POST",
		"/blogs",
		controllers.Create,
	},
	route{
		"BlogUpdate",
		"PUT",
		"/blogs/{blogId}",
		controllers.Update,
	},
	route{
		"BlogDelete",
		"DELETE",
		"/blogs/{blogId}",
		controllers.Delete,
	},
}
