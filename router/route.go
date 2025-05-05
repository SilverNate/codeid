package router

import (
	"codeid/murid"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	muridRepo := murid.NewMuridRepository(db)
	muridService := murid.NewMuridService(muridRepo)
	muridHandler := murid.NewMuridHandler(muridService)

	api := r.Group("/api")
	murid := api.Group("/murid")

	{
		murid.POST("", muridHandler.CreateMurid)
		murid.GET("", muridHandler.GetMurid)
		//murid.GET("/:id", muridHandler.GetAuthorByID)
		//murid.PUT("/:id", muridHandler.UpdateAuthor)
		//murid.DELETE("/:id", muridHandler.DeleteAuthor)
	}
	return r
}
