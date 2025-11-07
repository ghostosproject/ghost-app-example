package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ghostosproject/ghost-app/model"
)

type GhostAppModel struct {
	height int
	width  int
}

func (m GhostAppModel) Init() tea.Cmd {
	return nil
}

func (m GhostAppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			cmds = append(cmds, nil)
		}
	}
	return m, tea.Batch(cmds...)
}

func (m GhostAppModel) View() string {
	//return m.list.View()
	return "This is the Ghost App Model"
}

func MakeApp(width, height int) (model.GhostApp, error) {
	return GhostAppModel{
		height: height,
		width:  width,
	}, nil
}
