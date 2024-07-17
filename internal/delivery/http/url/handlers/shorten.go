package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

func (m *Manager) Shorten(c echo.Context) error {
	var req ShortenRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	_, err := url.ParseRequestURI(req.URL)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid URL")
	}

	s, err := m.usecase.Shorten(req.URL)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, s)
}
