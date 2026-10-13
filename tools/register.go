
package tools

func GetTools() []Tool {
    return []Tool{
        {
            Name:        "calculator",
            Description: "计算数学表达式，支持加减乘除、括号和小数",
            Function:    Calculate,
        },
        {
            Name:        "time",
            Description: "获取当前系统时间",
            Function:    GetTime,
        },
    }
}
