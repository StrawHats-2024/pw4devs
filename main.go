package main

import (
	"fmt"
	"log"
	"os"

	"github.com/StrawHats-2024/pm4devs/manager"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	log.SetOutput(f)

	m := manager.Model{Loading: true}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
