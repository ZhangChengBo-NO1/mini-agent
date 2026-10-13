package tools

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculate(input string) (string, error) {
	nums := strings.Split(input, "+")

	if len(nums) != 2 {
		// 真正地返回 error
		return "", fmt.Errorf("参数错误：目前只支持两个数相加，您输入的是 %s", input)
	}

	// 处理解析错误
	a, err := strconv.Atoi(strings.TrimSpace(nums[0]))
	if err != nil {
		return "", fmt.Errorf("第一个数 %s 不是有效的数字", nums[0])
	}

	b, err := strconv.Atoi(strings.TrimSpace(nums[1]))
	if err != nil {
		return "", fmt.Errorf("第二个数 %s 不是有效的数字", nums[1])
	}

	// 成功则返回结果
	return strconv.Itoa(a + b), nil
}