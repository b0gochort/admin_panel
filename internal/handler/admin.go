package handler

import (
	"encoding/json"
	"fmt"
	"github.com/b0gochort/admin_panel/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/valyala/fasthttp"
	"resenje.org/logging"
)

func (h *Handler) GetCompetents(ctx *fasthttp.RequestCtx) {
	if !ctx.IsGet() {
		logging.Info("invalid method")
		ctx.Error("handler NewChat check method: %v", fasthttp.StatusMethodNotAllowed)
		return
	}

	cToken := ctx.Request.Header.Cookie("token")

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(string(cToken), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("bot"), nil
	})

	if err != nil {
		logging.Info(fmt.Sprintf("handler.GetCompetents.ParseWithClaims: %v", err))
		ctx.Error("ParseWithClaims", fasthttp.StatusInternalServerError)
		return
	}

	sub, err := claims.GetSubject()
	if err != nil {
		logging.Info(fmt.Sprintf("handler.GetCompetents.GetSubject: %v", err))
		ctx.Error("GetSubject", fasthttp.StatusInternalServerError)
		return
	}

	if sub != "admins" {
		logging.Info("only admins have rules get cmpt")
		ctx.Error("not allowed", fasthttp.StatusForbidden)
		return
	}

	competents, err := h.services.AdminService.GetCompetent()
	if err != nil {
		if err != nil {
			logging.Info(fmt.Sprintf("handler.GetCompetents.AdminService.GetCompetent.%v", err))
			ctx.Error("GetCompetent", fasthttp.StatusInternalServerError)

			return
		}
	}

	res, err := json.Marshal(competents)
	if err != nil {
		logging.Info(fmt.Sprintf("handler.GetCompetents.Marshal: %v", err))
		ctx.Error("something went wrong ", fasthttp.StatusInternalServerError)

		return
	}

	ctx.Write(res)
	return
}

func (h *Handler) ChangeCompetent(ctx *fasthttp.RequestCtx) {
	var req model.CompetentRes

	if !ctx.IsPut() {
		logging.Info("invalid method")
		ctx.Error("handler NewChat check method: %v", fasthttp.StatusMethodNotAllowed)
		return
	}

	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		logging.Info(fmt.Sprintf("handler.GetCompetents.Unmarshal: %v", err))
		ctx.Error("unprocessable entity", fasthttp.StatusUnprocessableEntity)
		return
	}

	cToken := ctx.Request.Header.Cookie("token")

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(string(cToken), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("bot"), nil
	})

	if err != nil {
		logging.Info(fmt.Sprintf("handler.GetCompetents.ParseWithClaims: %v", err))
		ctx.Error("ParseWithClaims", fasthttp.StatusInternalServerError)
		return
	}

	sub, err := claims.GetSubject()
	if err != nil {
		logging.Info(fmt.Sprintf("handler.GetCompetents.GetSubject: %v", err))
		ctx.Error("GetSubject", fasthttp.StatusInternalServerError)
		return
	}

	if sub != "users" {
		logging.Info("only admins have rules")
		ctx.Error("not allowed", fasthttp.StatusForbidden)
		return
	}

	competent, err := h.services.AdminService.ChangeCompetent(req)
	if err != nil {
		if err != nil {
			logging.Info(fmt.Sprintf("handler.GetCompetents.AdminService.GetCompetent.%v", err))
			ctx.Error("GetCompetent", fasthttp.StatusInternalServerError)

			return
		}
	}

	res, err := json.Marshal(competent)
	if err != nil {
		logging.Info(fmt.Sprintf("handler.GetCompetents.Marshal: %v", err))
		ctx.Error("something went wrong ", fasthttp.StatusInternalServerError)

		return
	}

	ctx.Write(res)
	return
}

func (h *Handler) AddCompetent(ctx *fasthttp.RequestCtx) {
	var req model.AddCompetentRes

	if !ctx.IsPut() {
		logging.Info("invalid method")
		ctx.Error("handler NewChat check method: %v", fasthttp.StatusMethodNotAllowed)
		return
	}

	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		logging.Info(fmt.Sprintf("handler.AddCompetent.Unmarshal: %v", err))
		ctx.Error("unprocessable entity", fasthttp.StatusUnprocessableEntity)
		return
	}

	cToken := ctx.Request.Header.Cookie("token")

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(string(cToken), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("bot"), nil
	})

	if err != nil {
		logging.Info(fmt.Sprintf("handler.AddCompetent.ParseWithClaims: %v", err))
		ctx.Error("ParseWithClaims", fasthttp.StatusInternalServerError)
		return
	}

	sub, err := claims.GetSubject()
	if err != nil {
		logging.Info(fmt.Sprintf("handler.AddCompetent.GetSubject: %v", err))
		ctx.Error("GetSubject", fasthttp.StatusInternalServerError)
		return
	}

	if sub != "users" {
		logging.Info("only admins have rules")
		ctx.Error("not allowed", fasthttp.StatusForbidden)
		return
	}

	err = h.services.AdminService.AddCompetent(req.Uid)
	if err != nil {
		logging.Info(fmt.Sprintf("handler.AddCompetent.AdminService.AddCompetent.%v", err))
		ctx.Error("GetCompetent", fasthttp.StatusInternalServerError)

		return
	}

	result := model.AddCompetentRes{
		Uid:    req.Uid,
		Status: "Successful add competent",
	}

	res, err := json.Marshal(result)
	if err != nil {
		logging.Info(fmt.Sprintf("handler.AddCompetent.Marshal: %v", err))
		ctx.Error("something went wrong ", fasthttp.StatusInternalServerError)
		return
	}

	ctx.Write(res)
	return
}
