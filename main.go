package main

import (
	"gobus/di"
	"gobus/utils"
)

func main() {
	utils.LoadEnvironmentVariables()
	server := di.Init()
	server.StartServer()

	// db := db.ConnectDB()
	// st := entities.NewSeatLayout(db)

	// st.Layout1()
	// st.Layout2()
	// st.Layout3()
}
