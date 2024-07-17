package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
)

type ExpandRequest struct {
	ShortURL string `json:"url"`
}

func (m *Manager) Expand(c echo.Context) error {
	var req ExpandRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	_, err := url.ParseRequestURI(req.ShortURL)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid URL")
	}

	s, err := m.usecase.Expand(req.ShortURL)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, s)
}
