package handler

import (
	"car-service/internal/broker"
	"car-service/internal/model"
	"car-service/internal/service"
	"car-service/internal/validation"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type VehicleHandler struct {
	service *service.VehicleService
}

func NewVehicleHandler(s *service.VehicleService) *VehicleHandler {
	return &VehicleHandler{service: s}
}

func (h *VehicleHandler) Register(r *gin.Engine) {
	r.POST("/vehicles", h.Create)
	r.GET("/vehicles/:id", h.GetByID)
	r.GET("/vehicles", h.ListAll)
}

// @Summary Создать автомобиль
// @Tags vehicles
// @Accept json
// @Produce json
// @Param vehicle body model.Vehicle true "Автомобиль"
// @Success 201 {object} model.Vehicle
// @Failure 400 {object} map[string]string
// @Router /vehicles [post]
func (h *VehicleHandler) Create(c *gin.Context) {
	var v model.Vehicle
	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.service.Create(&v); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save"})
		return
	}

	if err := validation.ValidateVehicle(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := broker.PublishVehicleCreated(v); err != nil {
		logrus.Errorf("failed to publish Kafka event: %v", err)
	}

	c.JSON(http.StatusCreated, v)
}

// @Summary Получить автомобиль по id
// @Tags vehicles
// @Accept json
// @Produce json
// @Param id path uint true "ID автомобиля"
// @Success 200 {object} model.Vehicle
// @Failure 404 {object} map[string]string
// @Router /vehicles/{id} [get]
func (h *VehicleHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	v, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, v)
}

// @Summary Получить список всех автомобилей
// @Tags vehicles
// @Accept json
// @Produce json
// @Success 200 {array} model.Vehicle
// @Failure 404 {object} map[string]string
// @Router /vehicles [get]
func (h *VehicleHandler) ListAll(c *gin.Context) {
	v, err := h.service.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list vehicles"})
		return
	}

	c.JSON(http.StatusOK, v)
}
