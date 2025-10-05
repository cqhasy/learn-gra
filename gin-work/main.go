package main

import "github.com/gin-gonic/gin"

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  string `json:"stock"`
	Place  string `json:"place"`
}

var books = make(map[string]Book)

func AddBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBind(&newBook); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if _, exists := books[newBook.ID]; exists {
		c.JSON(409, gin.H{"error": "该图书 ID 已存在"})
		return
	}

	books[newBook.ID] = newBook
	c.JSON(200, newBook)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	delete(books, id)
	c.JSON(200, gin.H{"message": "删除成功"})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var changedBook Book
	if err := c.ShouldBind(&changedBook); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if id != changedBook.ID {
		c.JSON(400, gin.H{"error": "路径 ID 与请求体 ID 不一致"})
		return
	}

	books[id] = changedBook
	c.JSON(200, changedBook)
}

func SearchAllBook(c *gin.Context) {
	c.JSON(200, books)
}

func main() {
	r := gin.Default()

	r.GET("/book/search", SearchAllBook)
	r.POST("/book/add", AddBook)
	r.DELETE("/book/delete/:id", DeleteBook)
	r.PUT("/book/update/:id", UpdateBook)

	r.Run()
}
