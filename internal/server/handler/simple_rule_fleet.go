package handler

import (
	"strconv"

	"Frank2006xmotorq/internal/service"

	"github.com/gofiber/fiber/v3"
)

type SimpleRuleFleetHandler struct {
	svc *service.SimpleRuleFleetService
}

func NewSimpleRuleFleetHandler(svc *service.SimpleRuleFleetService) *SimpleRuleFleetHandler {
	return &SimpleRuleFleetHandler{svc: svc}
}

func (h *SimpleRuleFleetHandler) Create(c fiber.Ctx) error {
	var req service.CreateSimpleRuleFleetRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if req.FleetID == 0 || req.TargetField == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "fleet_id and target_field are required",
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

func (h *SimpleRuleFleetHandler) Get(c fiber.Ctx) error {
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

func (h *SimpleRuleFleetHandler) ListByFleet(c fiber.Ctx) error {
	fleetID, err := strconv.ParseInt(c.Query("fleet_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "fleet_id is required",
		})
	}

	rules, err := h.svc.ListByFleet(c.Context(), fleetID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(rules)
}

func (h *SimpleRuleFleetHandler) Delete(c fiber.Ctx) error {
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