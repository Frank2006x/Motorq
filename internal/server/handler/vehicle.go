package handler

import (
	"strconv"

	"Frank2006xmotorq/internal/service"

	"github.com/gofiber/fiber/v3"
)

type VehicleHandler struct {
	svc *service.VehicleService
}

func NewVehicleHandler(svc *service.VehicleService) *VehicleHandler {
	return &VehicleHandler{svc: svc}
}

func (h *VehicleHandler) Create(c fiber.Ctx) error {
	var req service.CreateVehicleRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if req.Model == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "model is required",
		})
	}

	if req.Status == "" {
		req.Status = "active"
	}

	vehicle, err := h.svc.Create(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(vehicle)
}

func (h *VehicleHandler) Get(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	vehicle, err := h.svc.Get(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(vehicle)
}

func (h *VehicleHandler) List(c fiber.Ctx) error {
	vehicles, err := h.svc.List(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(vehicles)
}

func (h *VehicleHandler) ListByFleet(c fiber.Ctx) error {
	fleetID, err := strconv.ParseInt(c.Query("fleet_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid fleet_id",
		})
	}

	vehicles, err := h.svc.ListByFleet(c.Context(), fleetID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(vehicles)
}

func (h *VehicleHandler) Update(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var req service.UpdateVehicleRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	req.ID = id
	vehicle, err := h.svc.Update(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(vehicle)
}

func (h *VehicleHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	_, err = h.svc.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}