package api

import (
	"fmt"
	"main/db"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getPeriodsRequest struct {
	DateStamp int32 `uri:"DateStamp" binding:"required"`
}

func (server *Server) getPeriods(ctx *gin.Context) {
	var req getPeriodsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	event := db.ReadPeriods(server.store, req.DateStamp)
	ctx.JSON(http.StatusOK, event)
}

func (server *Server) postPeriods(ctx *gin.Context) {
	var req models.Period
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req)

	event := db.CreatePeriod(server.store, req)
	ctx.JSON(http.StatusOK, event)
}

type deletePeriodsRequest struct {
	ID int64 `uri:"ID" binding:"required"`
}

func (server *Server) deletePeriods(ctx *gin.Context) {
	var req deletePeriodsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req)
	db.DeletePeriod(server.store, req.ID)
	ctx.JSON(http.StatusOK, req.ID)
}

func (server *Server) putPeriods(ctx *gin.Context) {
	var req models.Period
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req)

	event := db.UpdatePeriod(server.store, req)
	ctx.JSON(http.StatusOK, event)
}
