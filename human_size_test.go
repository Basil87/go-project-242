package code

import "testing"

func TestFormatSize(t *testing.T) {
	tests := []struct {
		size int64
		want string
	}{
		{0, "0B"},
		{999, "999B"},
		{1000, "1.0KB"},
		{1500, "1.5KB"},
		{1000000, "1.0MB"},
		{1000000000, "1.0GB"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := FormatSize(tt.size, true)
			if got != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, got)
			}
		})
	}
}
