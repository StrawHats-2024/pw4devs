package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	authmodel "github.com/StrawHats-2024/pw4devs/auth_model"
	listmodel "github.com/StrawHats-2024/pw4devs/list_model"
	tea "github.com/charmbracelet/bubbletea"
)

var currentUserId = -1

func main() {
	isLoggedIn := false

	if token, err := readFileContent("./token.txt"); err == nil {
		if res, err := authmodel.VerifyToken(token); err == nil && res.Valid {
			isLoggedIn = true
			currentUserId = res.UserID
		} else if err != nil {
			log.Fatal(err)
		}
	} else if !errors.Is(err, ErrFileNotFound) {
		log.Fatal(err)
	}

	if _, err := tea.NewProgram(model{
		list:  listmodel.NewModel(),
		auth:  authmodel.InitialModel(),
		login: isLoggedIn,
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
