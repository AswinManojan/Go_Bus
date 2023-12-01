package main

import "gobus/di"

func main() {
	server := di.Init()
	server.StartServer()

	// db := db.ConnectDB()
	// st := entities.NewSeatLAyout(db)

	// st.Layout1()
	// st.Layout2()
	// st.Layout3()
}
