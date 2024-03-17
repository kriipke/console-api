package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kriipke/console-api/pkg/models"
	"gorm.io/gorm"
)

type ClusterController struct {
	DB *gorm.DB
}

func NewClusterController(DB *gorm.DB) ClusterController {
	return ClusterController{DB}
}

// [...] Create Cluster Handler
func (cc *ClusterController) CreateCluster(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateClusterRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newCluster := models.Cluster{
		ApiServerHost:  payload.ApiServerHost,
		ApiServerPort:  payload.ApiServerPort,
		Name:   		payload.Name,
		Image:     		payload.Image,
		AddedBy:      	currentUser.ID,
		CreatedAt: 		now,
		UpdatedAt: 		now,
	}

	result := cc.DB.Create(&newCluster)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Cluster with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newCluster})
}

// [...] Update Cluster Handler
func (cc *ClusterController) UpdateCluster(ctx *gin.Context) {
	clusterId := ctx.Param("clusterId")

	// TO-DO: Add LastUpdatedBy field of Cluster and use this def below to populate it
	//currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateCluster
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedCluster models.Cluster
	result := cc.DB.First(&updatedCluster, "id = ?", clusterId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No cluster with that title exists"})
		return
	}
	now := time.Now()
	clusterToUpdate := models.Cluster{
		ApiServerHost:  payload.ApiServerHost,
		ApiServerPort:  payload.ApiServerPort,
		Name:   		payload.Name,
		Image:     		payload.Image,
		AddedBy:      	updatedCluster.AddedBy,
		CreatedAt: updatedCluster.CreatedAt,
		UpdatedAt: now,
	}

	cc.DB.Model(&updatedCluster).Updates(clusterToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedCluster})
}

// [...] Get Single Cluster Handler
func (cc *ClusterController) FindClusterById(ctx *gin.Context) {
	clusterId := ctx.Param("clusterId")

	var cluster models.Cluster
	result := cc.DB.First(&cluster, "id = ?", clusterId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No cluster with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": cluster})
}

// [...] Get All Clusters Handler
func (cc *ClusterController) FindClusters(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var clusters []models.Cluster
	results := cc.DB.Limit(intLimit).Offset(offset).Find(&clusters)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(clusters), "data": clusters})
}

// [...] Delete Cluster Handler
func (cc *ClusterController) DeleteCluster(ctx *gin.Context) {
	clusterId := ctx.Param("clusterId")

	result := cc.DB.Delete(&models.Cluster{}, "id = ?", clusterId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No cluster with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
