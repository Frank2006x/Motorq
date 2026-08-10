package handler

import (
	"strconv"
	"time"

	"Frank2006xmotorq/internal/service"

	"github.com/gofiber/fiber/v3"
)

type TelemetryHandler struct {
	svc        *service.TelemetryService
	evaluator  *service.RuleEvaluator
}

func NewTelemetryHandler(svc *service.TelemetryService, evaluator *service.RuleEvaluator) *TelemetryHandler {
	return &TelemetryHandler{
		svc:       svc,
		evaluator: evaluator,
	}
}

func (h *TelemetryHandler) Create(c fiber.Ctx) error {
	var req service.CreateTelemetryRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if req.VehicleID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "vehicle_id is required",
		})
	}

	if req.Timestamp.IsZero() {
		req.Timestamp = time.Now()
	}

	telemetry, err := h.svc.Create(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Evaluate rules for the new telemetry
	// Run in background to not block the response
	go func() {
		if h.evaluator != nil {
			_ = h.evaluator.CheckRulesAndUpdateStatus(c.Context(), telemetry.ID)
		}
	}()

	return c.Status(fiber.StatusCreated).JSON(telemetry)
}

func (h *TelemetryHandler) Get(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	telemetry, err := h.svc.Get(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(telemetry)
}

func (h *TelemetryHandler) ListByVehicle(c fiber.Ctx) error {
	vehicleID, err := strconv.ParseInt(c.Query("vehicle_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "vehicle_id is required",
		})
	}

	telemetry, err := h.svc.ListByVehicle(c.Context(), vehicleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(telemetry)
}

func (h *TelemetryHandler) ListLatestByVehicle(c fiber.Ctx) error {
	vehicleID, err := strconv.ParseInt(c.Query("vehicle_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "vehicle_id is required",
		})
	}

	limit := int32(10)
	if limitStr := c.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = int32(l)
		}
	}

	telemetry, err := h.svc.ListLatestByVehicle(c.Context(), vehicleID, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(telemetry)
}