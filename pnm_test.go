package asciiart

import (
	"reflect"
	"testing"
)

func Test_parsePNM(t *testing.T) {
	type args struct {
		pnm string
	}
	tests := []struct {
		name string
		args args
		want PNM
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePNM(tt.args.pnm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePNM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfoPnm(t *testing.T) {
	type args struct {
		pnm string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InfoPnm(tt.args.pnm)
		})
	}
}

func TestRenderPNMBW(t *testing.T) {
	type args struct {
		pnm string
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
			if got := RenderPNMBW(tt.args.pnm); got != tt.want {
				t.Errorf("RenderPNMBW() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_renderP2BW(t *testing.T) {
	type args struct {
		img PNM
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
			if got := renderP2BW(tt.args.img); got != tt.want {
				t.Errorf("renderP2BW() = %v, want %v", got, tt.want)
			}
		})
	}
}
