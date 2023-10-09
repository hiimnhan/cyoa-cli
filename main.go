package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	storyMap, err := os.ReadFile("gopher.json")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

	story := parseStory(storyMap)

	p := tea.NewProgram(&Model{
		chapter: story["intro"],
	},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("Could not run program", err)
		os.Exit(1)
	}
}
