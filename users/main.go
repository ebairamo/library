package main

import (
	"user_manager/ui"
	"user_manager/web"
)

func main() {
	users := ui.InitDefaultUsers()
	web.SetUsers(users)
	// ui.RunConsoleUI()
	web.StartServer()
}
