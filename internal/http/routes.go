package http

import (
	"github.com/HishamAlkahtani/yt-auto-archiver/internal/db"
	"github.com/gin-gonic/gin"
)

type AddChannelRequest struct {
	ChannelId string `json:"channel_id" binding:"required"`
}

func RegisterRoutes(r *gin.Engine) {
	var request AddChannelRequest

	r.POST("/add-channel", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}

		addChannel(request)
	})
}

func addChannel(request AddChannelRequest) {
	_, err := db.NewDB("internal/db/mydb.sqlite")

	if err != nil {
		panic("what")
	}
}
