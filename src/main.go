package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	store := new(Store)
	if err := store.Init(); err != nil {
		log.Fatalf("unable to initialise store: %v", err)
	}

	p := tea.NewProgram(initialModel(store))
	if _, err := p.Run(); err != nil {
		fmt.Printf("unable to run tui: %v", err)
		os.Exit(1)
	}
}

