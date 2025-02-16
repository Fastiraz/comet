package menu

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Item struct {
	TitleStr, Desc string
}

func (i Item) Title() string       { return i.TitleStr }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.TitleStr }

type Model struct {
	list        list.Model
	Selected    Item
	ItemChosen  bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		} else if msg.String() == "enter" {
			if selectedItem, ok := m.list.SelectedItem().(Item); ok {
				m.Selected = selectedItem
				m.ItemChosen = true
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return docStyle.Render(m.list.View())
}

func NewMenu(items []list.Item, title string) Model {
	m := Model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = title
	return m
}
