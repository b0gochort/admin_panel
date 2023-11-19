package handler

import (
	"fmt"
	"github.com/b0gochort/admin_panel/internal/service"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}

}

func (h *Handler) InitRoutes(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	switch string(ctx.Path()) {
	case "/get-competents":
		h.GetCompetents(ctx)
	case "/change-competents":
		h.ChangeCompetent(ctx)
	case "/add-competents":
		h.AddCompetent(ctx)
	}
}

func ping(ctx *fasthttp.RequestCtx) {
	fmt.Println("pong")
}
