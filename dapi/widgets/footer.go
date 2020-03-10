package widgets

import "github.com/rivo/tview"

// Footer is a widget displayed at the bottom of the app
type Footer struct {
	root *tview.Flex
	Help *tview.TextView
}

func newFooter() *Footer {
	footer := Footer{}
	footer.root = tview.NewFlex()

	footer.Help = tview.NewTextView()
	footer.Help.SetText("Help")
	footer.Help.SetBackgroundColor(ColorBlack)

	footer.root.AddItem(footer.Help, 0, 1, false)

	return &footer
}

// Root returns the root element
func (f *Footer) Root() tview.Primitive {
	return f.root
}
