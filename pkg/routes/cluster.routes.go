package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kriipke/console-api/pkg/controllers"
	"github.com/kriipke/console-api/pkg/middleware"
)

type ClusterRouteController struct {
	clusterController controllers.ClusterController
}

func NewRouteClusterController(clusterController controllers.ClusterController) ClusterRouteController {
	return ClusterRouteController{clusterController}
}

func (cc *ClusterRouteController) ClusterRoute(rg *gin.RouterGroup) {

	router := rg.Group("clusters")
	router.Use(middleware.DeserializeUser())
	router.POST("/", cc.clusterController.CreateCluster)
	router.GET("/", cc.clusterController.FindClusters)
	router.PUT("/:clusterId", cc.clusterController.UpdateCluster)
	router.GET("/:clusterId", cc.clusterController.FindClusterById)
	router.DELETE("/:clusterId", cc.clusterController.DeleteCluster)
}
