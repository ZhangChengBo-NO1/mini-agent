
package ui

import (
    "fmt"
    "strings"

    tea "github.com/charmbracelet/bubbletea"

    "mini-agent/agent"
)

type responseMsg struct {
    text string
    err  error
}

type Model struct {
    input   string
    output  string
    busy    bool
    agent   *agent.Agent
}

func NewModel() Model {
    return Model{
        agent: agent.NewAgent(),
        output: "你好！我是 Mini Agent。\n" +
            "试试：计算 (10+2)*3，或者问我 Go 相关知识。\n",
    }
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "ctrl+d":
            return m, tea.Quit

        case "enter":
            if m.busy {
                return m, nil
            }

            input := strings.TrimSpace(m.input)
            if input == "" {
                return m, nil
            }

            m.output += "\n你: " + input + "\n"
            m.input = ""
            m.busy = true

            a := m.agent
            return m, func() tea.Msg {
                result, err := a.Run(input)
                return responseMsg{text: result, err: err}
            }

        case "backspace":
            if len(m.input) > 0 {
                // 删除最后一个 UTF-8 字符，避免中文输入被截断。
                _, size := lastRune(m.input)
                m.input = m.input[:len(m.input)-size]
            }

        default:
            // 只追加可打印字符，避免把按键名称写入输入框。
            if len(msg.Runes) > 0 {
                m.input += string(msg.Runes)
            }
        }

    case responseMsg:
        m.busy = false
        if msg.err != nil {
            m.output += "错误: " + msg.err.Error() + "\n"
        } else {
            m.output += "Agent: " + msg.text + "\n"
        }
    }

    return m, nil
}

func lastRune(s string) (rune, int) {
    for i := len(s) - 1; i >= 0; i-- {
        if i == 0 || (s[i]&0xC0) != 0x80 {
            r := []rune(s[i:])[0]
            return r, len(string(r))
        }
    }
    return 0, 0
}

func (m Model) View() string {
    status := ""
    if m.busy {
        status = "\n正在处理，请稍候..."
    }

    return fmt.Sprintf(
        "========== Mini Agent ==========\n\n%s\n%s\n\n> %s\n\nEnter 发送 | Ctrl+C 退出\n",
        m.output,
        status,
        m.input,
    )
}