package handler

import (
	"strconv"

	"Frank2006xmotorq/internal/service"

	"github.com/gofiber/fiber/v3"
)

type SimpleRuleVehicleHandler struct {
	svc *service.SimpleRuleVehicleService
}

func NewSimpleRuleVehicleHandler(svc *service.SimpleRuleVehicleService) *SimpleRuleVehicleHandler {
	return &SimpleRuleVehicleHandler{svc: svc}
}

func (h *SimpleRuleVehicleHandler) Create(c fiber.Ctx) error {
	var req service.CreateSimpleRuleVehicleRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if req.VehicleID == 0 || req.TargetField == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "vehicle_id and target_field are required",
		})
	}

	if req.Priority == "" {
		req.Priority = "low"
	}

	rule, err := h.svc.Create(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(rule)
}

func (h *SimpleRuleVehicleHandler) Get(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	rule, err := h.svc.Get(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(rule)
}

func (h *SimpleRuleVehicleHandler) ListByVehicle(c fiber.Ctx) error {
	vehicleID, err := strconv.ParseInt(c.Query("vehicle_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "vehicle_id is required",
		})
	}

	rules, err := h.svc.ListByVehicle(c.Context(), vehicleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(rules)
}

func (h *SimpleRuleVehicleHandler) Delete(c fiber.Ctx) error {
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