package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yutakahashi114/clean-architecture/usecase"
)

// RestaurantHandler .
type RestaurantHandler interface {
	GetByID(ctx *gin.Context)
}

// NewRestaurantHandler .
func NewRestaurantHandler(restaurantUseCase usecase.RestaurantUseCase) RestaurantHandler {
	return &restaurantHandler{
		restaurantUseCase: restaurantUseCase,
	}
}

// restaurantHandler .
type restaurantHandler struct {
	restaurantUseCase usecase.RestaurantUseCase
}

// GetByID .
func (h *restaurantHandler) GetByID(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		renderError(ctx, err)
		return
	}

	out, err := h.restaurantUseCase.GetByID(ctx, usecase.GetByIDInput{
		ID: id,
	})
	if err != nil {
		renderError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"restaurant": map[string]interface{}{
			"id":        out.ID,
			"name":      out.Name,
			"tags":      out.Tags,
			"clientUid": out.ClientUID,
		},
	})
}

func renderError(ctx *gin.Context, err error) {
	log.Println(err)
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": fmt.Sprintf("internal server error: %v", err),
	})
}

func test() {

	// db := DBInstance{
	// 	DB: "",
	// }
	IRepository := Repository{}

	h := Handler{

		UseCase: UseCase{

			Repository: IRepository,
		},
	}

	_ = h
}

type Handler struct {
	UseCase UseCase
}

type UseCase struct {
	Repository Repository
}

type Repository struct {
	DBInstance DBInstance
}

type DBInstance struct {
	DB string
}
