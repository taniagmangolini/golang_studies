package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Estrutura para representar um livro
type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Year   string `json:"year"`
}

// Lista de livros como banco de dados em memória
var books = []Book{
    {ID: "1", Title: "The Catcher in the Rye", Author: "J.D. Salinger", Year: "1951"},
    {ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Year: "1960"},
    {ID: "3", Title: "1984", Author: "George Orwell", Year: "1949"},
}

// Função para obter todos os livros
func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

// Função para obter um livro por ID
func getBookByID(c *gin.Context) {
    id := c.Param("id")
    for _, book := range books {
        if book.ID == id {
            c.IndentedJSON(http.StatusOK, book)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// Função para adicionar um novo livro
func createBook(c *gin.Context) {
    var newBook Book

    // BindJSON converte o JSON recebido em um objeto Go
    if err := c.BindJSON(&newBook); err != nil {
        return
    }

    // Adiciona o novo livro à lista
    books = append(books, newBook)
    c.IndentedJSON(http.StatusCreated, newBook)
}

// Função para deletar um livro por ID
func deleteBookByID(c *gin.Context) {
    id := c.Param("id")
    for index, book := range books {
        if book.ID == id {
            books = append(books[:index], books[index+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// Função principal para iniciar o servidor
func main() {
    router := gin.Default()

    // Rotas da API
    router.GET("/books", getBooks)             // Rota para obter todos os livros
    router.GET("/books/:id", getBookByID)      // Rota para obter um livro por ID
    router.POST("/books", createBook)          // Rota para criar um novo livro
    router.DELETE("/books/:id", deleteBookByID) // Rota para deletar um livro por ID

    // Inicia o servidor na porta 8080
    router.Run("localhost:8080")
}
