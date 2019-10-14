package asciiart

import "testing"

func TestRenderASC(t *testing.T) {
	type args struct {
		asc string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RenderASC(tt.args.asc); got != tt.want {
				t.Errorf("RenderASC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetASCII(t *testing.T) {
	tests := []struct {
		name string
		nr   byte
		want rune
	}{
		{"!", 33, '!'},
		{"\n", 10, '\n'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetASCII(tt.nr); got != tt.want {
				t.Errorf("GetASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}
