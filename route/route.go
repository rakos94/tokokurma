package route

import (
	"tokokurma/config"
	"tokokurma/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Run(v *gin.Engine, addr string) error {
	config.GetConfig()
	r = v
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))
	RunRoute()
	err := r.Run(addr)
	if err != nil {
		return err
	}

	return nil
}

func GetR() *gin.Engine {
	return r
}

func RunRoute() {

	r.POST("/login", controller.Login.Login)
	authorized := r.Group("/")
	authorized.Use(controller.Login.Auth())
	authorized.GET("/authenticated", controller.Login.Authenticated)
	NewCRUDRoute(authorized, "/item", controller.Item)
	authorized.GET("/item/:id/history", controller.Item.History)
	NewCRUDRoute(authorized, "/customer", controller.Customer)
	NewCRUDRoute(authorized, "/order", controller.Order)
	NewCRUDRoute(authorized, "/user", controller.User)
}

func NewCRUDRoute(r *gin.RouterGroup, group string, ctr CRUDRouteInterface) {
	g := r.Group(group)
	{
		g.GET("", ctr.Get)
		g.POST("/add", ctr.Add)
		g.GET("/:id", ctr.First)
		g.POST("/:id/update", ctr.Update)
		g.POST("/:id/delete", ctr.Delete)
	}
}

type CRUDRouteInterface interface {
	Get(c *gin.Context)
	First(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
