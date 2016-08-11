package dal

import (
	"github.com/shandilyamanvesh/CrudAPIInGo/models"
)

//keeps track of primary key:blog id
var currentId int

//list of blogs
var blogs models.Blogs

//Initialize repo
func init() {
	blogs = models.Blogs{}
}

func CreateBlog(b models.Blog) models.Blog {
	currentId += 1
	b.Id = currentId
	blogs = append(blogs, b)
	return b
}

func GetAllBlogs() models.Blogs {
	return blogs
}

func GetBlogById(id int) (*models.Blog, error) {
	for _, b := range blogs {
		if b.Id == id {
			return &b, nil
		}
	}
	// return error if not found
	return nil, &models.Error{Msg: "No blog available at this Id."}
}
