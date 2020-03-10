package widgets

import "github.com/rivo/tview"

type Layout struct {
	root    *tview.Flex
	header  *Header
	content *Content
	footer  *Footer
}

func NewLayout() *Layout {
	l := Layout{}

	l.root = tview.NewFlex()
	l.root.SetDirection(tview.FlexRow)

	l.header = newHeader()
	l.root.AddItem(l.header.Root(), 1, 0, false)

	l.content = newContent()
	l.root.AddItem(l.content.Root(), 0, 1, false)

	l.footer = newFooter()
	l.root.AddItem(l.footer.Root(), 1, 0, false)

	return &l
}

func (l *Layout) Root() tview.Primitive {
	return l.root
}

func (l *Layout) SetTitle(text string) {
	l.header.Title.SetText("[::b]" + text)
}

func (l *Layout) SetHelpText(text string) {
	l.footer.Help.SetText(text)
}
