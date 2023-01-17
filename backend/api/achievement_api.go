package api

import (
	"fmt"
	"main/db"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getAchievementsRequest struct {
	DateStamp int32 `uri:"DateStamp" binding:"required"`
}

func (server *Server) getAchievements(ctx *gin.Context) {
	var req getAchievementsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	event := db.ReadAchievements(server.store, req.DateStamp)
	ctx.JSON(http.StatusOK, event)
}

func (server *Server) postAchievements(ctx *gin.Context) {
	var req models.Achievement
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req)

	event := db.CreateAchievement(server.store, req)
	ctx.JSON(http.StatusOK, event)
}

type deleteAchievementsRequest struct {
	ID int64 `uri:"ID" binding:"required"`
}

func (server *Server) deleteAchievements(ctx *gin.Context) {
	var req deleteAchievementsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req)
	db.DeleteAchievement(server.store, req.ID)
	ctx.JSON(http.StatusOK, req.ID)
}

func (server *Server) putAchievements(ctx *gin.Context) {
	var req models.Achievement
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req)

	event := db.UpdateAchievement(server.store, req)
	ctx.JSON(http.StatusOK, event)
}
