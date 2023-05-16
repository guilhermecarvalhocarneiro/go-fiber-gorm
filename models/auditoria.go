package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	DjangoUserName string `json:"django_username"`
}

type Permissions struct {
	gorm.Model
	Name      string `json:"name"`
	AppName   string `json:"app_name"`
	ModelName string `json:"model_name"`
}

type Data struct {
	gorm.Model
	Data string `json:"data"`
}

type Auditoria struct {
	gorm.Model
	CreatedAt        time.Time
	UsuarioRefer     int         `json:"usuario_id"`
	Usuario          Usuario     `json:"usuario" gorm:"foreignKey:UsuarioRefer"`
	IpAddress        string      `json:"ip_address"`
	PermissionsRefer int         `json:"permissions_id"`
	Permissions      Permissions `json:"permissions" gorm:"foreignKey:PermissionsRefer"`
	DataRefer        int         `json:"data_id"`
	Data             Data        `json:"data" gorm:"foreignKey:DataRefer"`
	// DataAuditoria    pkg.JSONMap `json:"data_auditoria" gorm:"type:jsonb"`
	DataAuditoria postgres.Jsonb `json:"data_auditoria"`
}
