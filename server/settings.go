package server

import (
	"github.com/anovafawzi/socialmedia/db"
	"github.com/anovafawzi/socialmedia/domain/relation"
	"github.com/gin-gonic/gin"
)

// Cors : allow CORS
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// OptionsUser : set header
func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

// LoadRoute : return all possible route
func LoadRoute() *gin.Engine {
	r := gin.Default()

	// create controller
	var (
		sqliteRepo         = db.NewSQLiteRepository("./dbsocmed.db")
		relationController = relation.NewController(sqliteRepo)
	)

	r.Use(Cors())
	v1 := r.Group("api/v1")
	{
		v1.POST("/friendconnection", relationController.PostFriendConnection)
		v1.POST("/friendlist", relationController.PostFriendList)
		v1.POST("/friendcommonlist", relationController.PostFriendCommonList)
		v1.POST("/friendsubscribe", relationController.PostFriendSubscribe)
		v1.POST("/friendblock", relationController.PostFriendBlock)
		v1.POST("/friendupdates", relationController.PostFriendUpdates)
	}

	return r
}
