package maps

import "testing"

func TestGetOr(t *testing.T) {
	type args struct {
		m      map[string]any
		key    string
		defval int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get value of the key in the map if it exists",
			args: args{
				m:      map[string]any{"a": 1},
				key:    "a",
				defval: 0,
			},
			want: 1,
		},
		{
			name: "return the default value if the key does not exist in the map",
			args: args{
				m:      map[string]any{"a": 1},
				key:    "b",
				defval: 0,
			},
			want: 0,
		},
		{
			name: "return the default value if the value of the key is not of the type of the default value",
			args: args{
				m:      map[string]any{"a": "b"},
				key:    "a",
				defval: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOr[int](tt.args.m, tt.args.key, tt.args.defval); got != tt.want {
				t.Errorf("GetOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
