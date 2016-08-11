package dal

import (
	"github.com/shandilyamanvesh/CrudAPIInGo/models"
)

//Db error msg
var errMsg = "No blog available at this Id."

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
	return nil, &models.Error{Msg: errMsg}
}

func UpdateBlogById(blog models.Blog) (*models.Blog, error) {
	for i, b := range blogs {
		if b.Id == blog.Id {
			blogs[i] = blog
			return &blogs[i], nil
		}
	}
	// return error if not found
	return nil, &models.Error{Msg: errMsg}
}

func DeleteBlogById(id int) error {
	for i, b := range blogs {
		if b.Id == id {
			blogs = append(blogs[:i], blogs[i+1:]...)
			return nil
		}
	}
	// return error if not found
	return &models.Error{Msg: errMsg}
}
