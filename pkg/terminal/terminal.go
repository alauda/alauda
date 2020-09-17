package terminal

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// Terminal defines the main terminal window for the Kubernetes cluster
type Terminal struct {
	application *tview.Application
	root        *tview.Grid
}

// NewTerminal returns a new terminal window
func NewTerminal(profileFile string) *Terminal {
	return &Terminal{
		application: tview.NewApplication(),
	}
}

// Start starts the terminal application
func (term *Terminal) Start() error {
	term.initializeLayout()

	err := term.application.Run()
	if err != nil {
		return err
	}

	return nil
}

// Stop stops the terminal application
func (term *Terminal) Stop() {
	term.application.Stop()
}

func (term *Terminal) initializeLayout() {
	grid := tview.NewGrid()
	term.application.SetRoot(grid, true).SetInputCapture(term.keyCapture)
	term.root = grid
}

func (term *Terminal) keyCapture(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyRune {
		rune := event.Rune()
		if rune == 'q' {
			term.Stop()
			return nil
		}
	}

	return event
}
