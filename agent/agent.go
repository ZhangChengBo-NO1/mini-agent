
package agent

import (
    "fmt"
    "strings"

    "mini-agent/llm"
    "mini-agent/tools"
)

type Agent struct {
    llm   *llm.Client
    tools []tools.Tool
}

func NewAgent() *Agent {
    return &Agent{
        llm:   llm.NewClient(),
        tools: tools.GetTools(),
    }
}

// Run 根据用户输入决定执行工具还是调用大模型。
func (a *Agent) Run(input string) (string, error) {
    input = strings.TrimSpace(input)
    if input == "" {
        return "", fmt.Errorf("请输入内容")
    }

    // V1：使用简单规则识别工具调用。
    if expression, ok := extractExpression(input); ok {
        return a.executeTool("calculator", expression)
    }

    if strings.Contains(input, "时间") ||
        strings.Contains(strings.ToLower(input), "time now") {
        return a.executeTool("time", "")
    }

    // 没有匹配到工具，就调用大模型回答问题。
    return a.llm.Chat(input)
}

// executeTool 根据工具名称找到对应的工具并执行。
func (a *Agent) executeTool(name, input string) (string, error) {
    for _, tool := range a.tools {
        if tool.Name == name {
            result, err := tool.Function(input)
            if err != nil {
                return "", fmt.Errorf("工具 %s 执行失败: %w", name, err)
            }

            return fmt.Sprintf(
                "[调用工具: %s]\n工具结果: %s",
                tool.Name,
                result,
            ), nil
        }
    }

    return "", fmt.Errorf("未找到工具: %s", name)
}

// extractExpression 识别输入中以“计算”开头的表达式。
// 例如：计算 (10+2)*3
func extractExpression(input string) (string, bool) {
    if !strings.HasPrefix(input, "计算") {
        return "", false
    }

    expression := strings.TrimSpace(strings.TrimPrefix(input, "计算"))
    return expression, expression != ""
}