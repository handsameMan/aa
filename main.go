package main

import (
	"archie/connection"
	"archie/models"
	"archie/routes"
	"archie/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createDataTable(db *gorm.DB, model interface{}) (hasTable bool) {
	hasTable = db.HasTable(model)

	if !hasTable {
		db.CreateTable(model)
	}

	return
}

func InitTable() {
	db, err := connection.GetDB()

	utils.Check(err)
	defer db.Close()

	// init table
	createDataTable(db, models.UserOrganization{})
	createDataTable(db, models.User{})
	createDataTable(db, models.Organization{})
}

func main() {
	InitTable()

	organization := models.Organization{
		OrganizeName: "选课精灵",
		Description:  "That's fucking awesome!",
	}

	organization.NewOrganization()

	routes.Serve()
}
