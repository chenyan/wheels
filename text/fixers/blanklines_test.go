package fixers

import (
	"strings"
	"testing"
)

func TestRemoveExtraBlankLines(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "空字符串",
			input:   "",
			want:    "",
			wantErr: false,
		},
		{
			name:    "单行文本",
			input:   "hello",
			want:    "hello",
			wantErr: false,
		},
		{
			name:    "无空行",
			input:   "line1\nline2\nline3",
			want:    "line1\nline2\nline3",
			wantErr: false,
		},
		{
			name:    "单个空行",
			input:   "line1\n\nline2",
			want:    "line1\n\nline2",
			wantErr: false,
		},
		{
			name:    "多个连续空行",
			input:   "line1\n\n\n\nline2",
			want:    "line1\n\nline2",
			wantErr: false,
		},
		{
			name:    "开头多个空行",
			input:   "\n\n\nline1",
			want:    "\nline1",
			wantErr: false,
		},
		{
			name:    "结尾多个空行",
			input:   "line1\n\n\n",
			want:    "line1",
			wantErr: false,
		},
		{
			name:    "只有空行",
			input:   "\n\n\n",
			want:    "",
			wantErr: false,
		},
		{
			name:    "包含空格的空行",
			input:   "line1\n  \n\t\n  \nline2",
			want:    "line1\n  \nline2",
			wantErr: false,
		},
		{
			name:    "超大输入",
			input:   strings.Repeat("a\n", MaxInputSize),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveExtraBlankLines(tt.input, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveExtraBlankLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("RemoveExtraBlankLines() = %q, want %q", got, tt.want)
			}
		})
	}
}

// 性能测试
func BenchmarkRemoveExtraBlankLines(b *testing.B) {
	// 准备测试数据
	smallInput := strings.Repeat("line\n\n\n", 10)   // 小输入
	largeInput := strings.Repeat("line\n\n\n", 1000) // 大输入

	b.Run("SmallInput", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = RemoveExtraBlankLines(smallInput, 1)
		}
	})

	b.Run("LargeInput", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = RemoveExtraBlankLines(largeInput, 1)
		}
	})
}
