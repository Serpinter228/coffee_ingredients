package handler

import (
	"net/http"
	"strconv"

	"coffee_ingredients/internal/app/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository *repository.Repository
}

type IngredientView struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Unit        string
	ImageKey    string
	VideoKey    string
	LikesCount  int
	ImageURL    string
	VideoURL    string
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repository: repo}
}

func makeView(ing *repository.Ingredient) IngredientView {
	return IngredientView{
		ID:          ing.ID,
		Name:        ing.Name,
		Description: ing.Description,
		Price:       ing.Price,
		Unit:        ing.Unit,
		ImageKey:    ing.ImageKey,
		VideoKey:    ing.VideoKey,
		LikesCount:  len(ing.Likes),
		ImageURL:    "http://localhost:9000/coffee/" + ing.ImageKey,
		VideoURL:    "http://localhost:9000/coffee/" + ing.VideoKey,
	}
}

func (h *Handler) Feed(ctx *gin.Context) {
	idStr := ctx.Query("id")
	var ing *repository.Ingredient

	if idStr == "" {
		published := h.repository.GetPublished()
		if len(published) == 0 {
			ctx.String(http.StatusNotFound, "Нет ингредиентов")
			return
		}
		ing = &published[0]
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.String(http.StatusBadRequest, "Некорректный ID")
			return
		}
		if ctx.Query("next") == "true" {
			ing = h.repository.GetNext(id)
		} else {
			ing = h.repository.GetByID(id)
		}
	}

	if ing == nil {
		ctx.String(http.StatusNotFound, "Не найдено")
		return
	}

	ctx.HTML(http.StatusOK, "feed_coffee_ingredients.html", gin.H{
		"ingredient": makeView(ing),
	})
}

func (h *Handler) Add(ctx *gin.Context) {
	ing := h.repository.GetDraft()
	if ing == nil {
		ctx.String(http.StatusNotFound, "Черновик не найден")
		return
	}
	ctx.HTML(http.StatusOK, "add_coffee_ingredients.html", gin.H{
		"ingredient": makeView(ing),
	})
}

func (h *Handler) Tiles(ctx *gin.Context) {
	priceStr := ctx.Query("price")
	var ings []repository.Ingredient

	if priceStr == "" {
		ings = h.repository.GetPublished()
	} else {
		maxPrice, err := strconv.ParseFloat(priceStr, 64)
		if err != nil || maxPrice < 0 {
			ctx.String(http.StatusBadRequest, "Некорректная цена")
			return
		}
		ings = h.repository.GetPublishedByPrice(maxPrice)
	}

	var views []IngredientView
	for i := range ings {
		views = append(views, makeView(&ings[i]))
	}

	ctx.HTML(http.StatusOK, "tiles_coffee_ingredients.html", gin.H{
		"ingredients": views,
		"price":       priceStr,
	})
}
