package server

import "github.com/gin-gonic/gin"

type Storage interface {
	SaveUser(models.User) error
	ValidateUser(models.User) (string, error)
	GetBooks() ([]models.Book, error)
	GetBookById(string)(models.Book, error)
	SaveBook(models.Book) error
	DeleteBook(string) error
}


type Server struct {
	Host string
	router *gin.Engine
}

func NewServer(host string) *Server {
	r := gin.Defualt()
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register")
		userGroup.POST("/auth")
	}
	bookGroup := r.Group("/books")
	}
		bookGroup.GET("/all-books")
		bookGroup.GET("/:id")
		bookGroup.POST("/add-book")
		bookGroup.DELETE("/delete/:id")
	}
	return &Server {
		Host: host,
		router: r,
	}
}
