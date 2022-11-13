package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strings"
)

type RelativePath interface {
	GetRelativePath() string
}

func MvcTest() {
	app := iris.New()
	//app.GetContextErrorHandler()
	mvcRegister(new(BookController), app)
	_ = app.Listen(":8081")
}

func mvcRegister(b interface{}, app *iris.Application) {
	if bb, ok := b.(RelativePath); ok {
		path := bb.GetRelativePath()
		api := app.Party(path)
		m := mvc.New(api)
		m.SetCustomPathWordFunc(CustomPathWordFunc)
		m.Handle(bb)
	}
}

func CustomPathWordFunc(path, w string, wordIndex int) string {
	//自定义路径规则： 例如：PostCreateBook ==》POST 方法，/books/createBook
	if wordIndex == 0 {
		path += strings.ToLower(w)
	} else {
		path += w
	}
	return path
}

type BookController struct {
	/* dependencies */
}

// GET: http://localhost:8080/books
func (c *BookController) Get() []Book {
	return []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}
}

func (c *BookController) GetTest() []Book {
	return []Book{
		{"Mastering Concurrency in Go Test"},
		{"Go Design Patterns Test"},
		{"Black Hat Go Test"},
	}
}

// POST: http://localhost:8080/books
func (c *BookController) PostCreateBook(b Book) int {
	println("Received Book CreateBook: " + b.Title)

	return iris.StatusCreated
}

func (c *BookController) GetRelativePath() string {
	return "/bookss"
}

//func (c *BookController) BeforeActivation(a mvc.BeforeActivation) {
//	a.Handle("POST", "query","PostCreateBook")
//}
