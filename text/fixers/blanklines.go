package fixers

import (
	"bufio"
	"errors"
	"strings"
)

// 错误定义
var (
	ErrInputTooLarge = errors.New("input size exceeds maximum allowed size")
	// maxInputSize 定义允许处理的最大输入大小（10MB）
	MaxInputSize = 10 << 20
)

// RemoveExtraBlankLines 移除文本中的多余空行，保留最多一个连续的空行
// 如果输入超过最大大小限制，将返回错误
func RemoveExtraBlankLines(content string) (string, error) {
	// 处理边界情况
	if len(content) == 0 {
		return "", nil
	}

	if len(content) > MaxInputSize {
		return "", ErrInputTooLarge
	}

	// 预分配 Builder 容量，假设结果长度不会超过原始长度
	var result strings.Builder
	result.Grow(len(content))

	scanner := bufio.NewScanner(strings.NewReader(content))
	blankLineCount := 0

	// 大文件使用 Scanner
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			blankLineCount++
			if blankLineCount <= 1 {
				result.WriteString(line + "\n")
			}
		} else {
			blankLineCount = 0
			result.WriteString(line + "\n")
		}
	}

	// 检查扫描错误
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.TrimRight(result.String(), "\n"), nil
}
