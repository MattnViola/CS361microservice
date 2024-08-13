package userdata

import (
	"net/http"

	"01-Login/platform/storage"

	"github.com/gin-gonic/gin"
)

func Handler(ctx *gin.Context) {
    var input struct {
        Key string `json:"key"`
    }

    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    profile, err := storage.DownloadAndParseJSON("cs361microservicedata", input.Key)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to download and parse JSON"})
        return
    }

    ctx.JSON(http.StatusOK, profile)
}