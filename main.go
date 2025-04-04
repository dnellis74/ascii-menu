package main

import (
	"ascii-menu/config"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the application
	app := tview.NewApplication()

	// Create the main menu
	menu := tview.NewList().
		ShowSecondaryText(false).
		SetMainTextColor(tcell.ColorWhite).
		SetSelectedTextColor(tcell.ColorBlack).
		SetSelectedBackgroundColor(tcell.ColorGreen)

	// Create description text view
	description := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetTextColor(tcell.ColorYellow)

	// Create instructions text view
	instructions := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetTextColor(tcell.ColorBlue)

	// Set instructions text
	instructions.SetText(strings.Join([]string{
		"Controls:",
		"↑/↓ - Navigate games",
		"Enter/A Button - Select game",
		"Space - Return to menu",
		"q - Quit",
	}, "\n"))

	// Add games to the menu
	for _, game := range cfg.Games {
		game := game // Create a new variable to avoid closure issues
		menu.AddItem(game.Name, "", 0, func() {
			// Suspend the application
			app.Suspend(func() {
				// Launch the game
				cmd := exec.Command(game.Path)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					log.Printf("Game exited with error: %v", err)
				}
			})
		})
	}

	// Add a quit option
	menu.AddItem("Quit", "", 'q', func() {
		app.Stop()
	})

	// Update description when selection changes
	menu.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if index < len(cfg.Games) {
			description.SetText(cfg.Games[index].Description)
		} else {
			description.SetText("")
		}
	})

	// Create the main layout
	mainFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(menu, 0, 1, true).
		AddItem(description, 0, 2, false)

	// Create the full layout with instructions at the bottom
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(mainFlex, 0, 1, true).
		AddItem(instructions, 5, 0, false)

	// Set the root and run the application
	if err := app.SetRoot(flex, true).Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
