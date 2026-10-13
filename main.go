
package main

import (
    "log"

    tea "github.com/charmbracelet/bubbletea"

    "mini-agent/ui"
)

func main() {
    program := tea.NewProgram(ui.NewModel())

    if _, err := program.Run(); err != nil {
        log.Fatal(err)
    }
}