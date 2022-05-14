package handlers

import (
	"net/http"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/repository"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) URLHandler {
	return URLHandler{
		repo: repo,
	}
}

func (h *URLHandler) Get(c *gin.Context) {
	// TODO: answer here
	shortURL := c.Param("shortURL")
	entityURL, err := h.repo.Get(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.Writer.Header().Set("Location", entityURL.LongURL)
	c.JSON(http.StatusFound, gin.H{
		"message": "Redirecting...",
	})
}

func (h *URLHandler) Create(c *gin.Context) {
	// TODO: answer here
	var entityURL entity.URL
	errReq := c.ShouldBindJSON(&entityURL)
	if errReq != nil {
		c.JSON(http.StatusBadRequest, entity.ErrBadRequest)
	}

	entityURLResp, err := h.repo.Create(entityURL.LongURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"Data": entityURLResp,
	})

}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	// TODO: answer here
	var entityURL entity.URL
	errReq := c.ShouldBindJSON(&entityURL)
	if errReq != nil {
		c.JSON(http.StatusBadRequest, entity.ErrBadRequest)
	}

	entityURLResp, err := h.repo.CreateCustom(entityURL.LongURL, entityURL.ShortURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"Data": entityURLResp,
	})
}
