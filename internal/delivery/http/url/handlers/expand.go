package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ExpandRequest struct {
	Hash string `param:"url"`
}

func (m *Manager) Expand(c echo.Context) error {
	var req ExpandRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	s, err := m.usecase.Expand(req.Hash)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, s)
}
