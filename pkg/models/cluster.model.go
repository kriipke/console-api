package models

import (
	"time"

	"github.com/google/uuid"
)

type Cluster struct {
	ID        				uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	ApiServerHost     string    `gorm:"uniqueIndex;not null" json:"api_server_host,omitempty"`
	ApiServerPort     string    `gorm:"not null" json:"api_server_port,omitempty"`
	Name   						string    `gorm:"not null" json:"name,omitempty"`
	Image     				string    `gorm:"not null" json:"image,omitempty"`
	AddedBy      			uuid.UUID `gorm:"not null" json:"added_by,omitempty"`
	CreatedAt 				time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt 				time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateClusterRequest struct {
	ApiServerHost     string    `json:"api_server_host"  binding:"required"`
	ApiServerPort     string    `json:"api_server_port"  binding:"required"`
	Name   						string    `json:"name" binding:"required"`
	Image     				string    `json:"image,omitempty"`
	AddedBy      			string    `json:"added_by,omitempty"`
	CreatedAt					time.Time `json:"created_at,omitempty"`
	UpdatedAt 				time.Time `json:"updated_at,omitempty"`
}

type UpdateCluster struct {
	ApiServerHost    	string    `json:"api_server_host,omitempty"`
	ApiServerPort 		string    `json:"api_server_port,omitempty"`
	Name   						string    `json:"name,omitempty"`
	Image     				string    `json:"image,omitempty"`
	AddedBy      			string    `json:"added_by,omitempty"`
	CreateAt  				time.Time `json:"created_at,omitempty"`
	UpdatedAt 				time.Time `json:"updated_at,omitempty"`
}

