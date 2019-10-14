package asciiart

import "testing"

func TestBlockelems(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{" ", "▀ ▁ ▂ ▃ ▄ ▅ ▆ ▇ █ ▉ ▊ ▋ ▌ ▍ ▎ ▏ ▐ ░ ▒ ▓ ▔ ▕ ▖ ▗ ▘ ▙ ▚ ▛ ▜ ▝ ▞ ▟"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Blockelems(); got != tt.want {
				t.Errorf("Blockelems() = %v, want %v", got, tt.want)
			}
		})
	}
}
