package main

import (
    "myapp/internal/adapters/api"
    "myapp/internal/adapters/db"
    "myapp/internal/app"
    "github.com/gin-gonic/gin"
)

func main() {
    // Inicializa o repositório (fake in-memory) e o serviço
    bookRepo := db.NewBookRepository()
    bookService := app.NewBookService(bookRepo)
    
    // Inicializa o handler e as rotas
    bookHandler := api.NewBookHandler(bookService)
    
    router := gin.Default()
    router.GET("/books", bookHandler.GetBooks)
    router.GET("/books/:id", bookHandler.GetBookByID)
    router.POST("/books", bookHandler.CreateBook)
    router.DELETE("/books/:id", bookHandler.DeleteBook)

    // Inicia o servidor
    router.Run(":8080")
}



package db

import (
    "errors"
    "myapp/internal/domain"
)

// bookRepository é uma implementação fake da interface BookRepository.
type bookRepository struct {
    books []domain.Book
}

// NewBookRepository cria uma nova instância de bookRepository.
func NewBookRepository() *bookRepository {
    return &bookRepository{
        books: []domain.Book{
            {ID: "1", Title: "Go Programming", Author: "John Doe", Year: "2020"},
        },
    }
}

func (r *bookRepository) GetAll() ([]domain.Book, error) {
    return r.books, nil
}

func (r *bookRepository) GetByID(id string) (domain.Book, error) {
    for _, book := range r.books {
        if book.ID == id {
            return book, nil
        }
    }
    return domain.Book{}, errors.New("book not found")
}

func (r *bookRepository) Create(book domain.Book) error {
    r.books = append(r.books, book)
    return nil
}

func (r *bookRepository) Delete(id string) error {
    for index, book := range r.books {
        if book.ID == id {
            r.books = append(r.books[:index], r.books[index+1:]...)
            return nil
        }
    }
    return errors.New("book not found")
}


package api

import (
    "myapp/internal/domain"
    "myapp/internal/ports"
    "net/http"
    "github.com/gin-gonic/gin"
)

type BookHandler struct {
    service ports.BookService
}

// NewBookHandler cria uma nova instância de BookHandler.
func NewBookHandler(s ports.BookService) *BookHandler {
    return &BookHandler{service: s}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
    books, err := h.service.GetAllBooks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
    id := c.Param("id")
    book, err := h.service.GetBookByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
    var book domain.Book
    if err := c.BindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.service.CreateBook(book); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
    id := c.Param("id")
    if err := h.service.DeleteBook(id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
        return
    }
    c.Status(http.StatusNoContent)
}


package app

import (
    "myapp/internal/domain"
    "myapp/internal/ports"
)

// bookService é a implementação da interface BookService.
type bookService struct {
    repository ports.BookRepository
}

// NewBookService cria uma nova instância de bookService.
func NewBookService(repo ports.BookRepository) ports.BookService {
    return &bookService{
        repository: repo,
    }
}

func (s *bookService) GetAllBooks() ([]domain.Book, error) {
    return s.repository.GetAll()
}

func (s *bookService) GetBookByID(id string) (domain.Book, error) {
    return s.repository.GetByID(id)
}

func (s *bookService) CreateBook(book domain.Book) error {
    return s.repository.Create(book)
}

func (s *bookService) DeleteBook(id string) error {
    return s.repository.Delete(id)
}


package ports

import "myapp/internal/domain"

// BookService define os métodos do serviço para interagir com o domínio.
type BookService interface {
    GetAllBooks() ([]domain.Book, error)
    GetBookByID(id string) (domain.Book, error)
    CreateBook(book domain.Book) error
    DeleteBook(id string) error
}


package ports

import "myapp/internal/domain"

// BookRepository é a interface que define as operações de persistência para o domínio Book.
type BookRepository interface {
    GetAll() ([]domain.Book, error)
    GetByID(id string) (domain.Book, error)
    Create(book domain.Book) error
    Delete(id string) error
}
44package domain

// Book é a entidade de domínio
type Book struct {
    ID     string
    Title  string
    Author string
    Year   string
}

