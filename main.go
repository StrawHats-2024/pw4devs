package main

import (
	"fmt"
	"os"

	authmodel "github.com/StrawHats-2024/pw4devs/auth_model"
	listmodel "github.com/StrawHats-2024/pw4devs/list_model"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if _, err := tea.NewProgram(model{
		list:  listmodel.NewModel(),
		auth:  authmodel.InitialModel(),
		login: false,
	},
		tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("err :", err)
		os.Exit(1)
	}
	// err := authmodel.Login("palegar.parikshith@gmail.com", "parikshith44")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
