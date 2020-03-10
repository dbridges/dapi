package widgets

import "github.com/rivo/tview"

// Widget represents a general widget or view of the app
type Widget interface {
	Root() tview.Primitive
}
