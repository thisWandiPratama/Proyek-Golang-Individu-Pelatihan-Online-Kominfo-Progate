package main

import (
	"log"
	"path/filepath"
	"todos/todos"
	webHandler "todos/web/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// konek database
	dsn := "root:12345678@tcp(127.0.0.1:3306)/todos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	todoRepository := todos.NewRepository(db)
	todoService := todos.NewService(todoRepository)

	// web
	todoWebHandler := webHandler.NewTodoHandler(todoService)

	router := gin.Default()
	router.Use(cors.Default())

	router.HTMLRender = loadTemplates("./web/templates")

	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")

	// categori
	router.GET("/", todoWebHandler.Index)
	router.GET("/todos/new", todoWebHandler.New)
	router.POST("/todos", todoWebHandler.Create)
	router.GET("/todos/edit/:id", todoWebHandler.Edit)
	router.POST("/todos/update/:id", todoWebHandler.Update)
	router.GET("/todos/done/:id", todoWebHandler.Done)
	router.GET("/todos/delete/:id", todoWebHandler.Delete)

	router.Run()
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
