package widgets

import "github.com/rivo/tview"

// Content holds the main view of the app
type Content struct {
	root *tview.Flex
}

func newContent() *Content {
	c := Content{}
	c.root = tview.NewFlex()
	return &c
}

// Root returns the root element
func (c *Content) Root() tview.Primitive {
	return c.root
}
