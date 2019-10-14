package asciiart

import (
	"reflect"
	"testing"
)

func TestRenderFile(t *testing.T) {
	type args struct {
		fn string
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
			if got := RenderFile(tt.args.fn); got != tt.want {
				t.Errorf("RenderFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpenFile(t *testing.T) {
	type args struct {
		fn string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OpenFile(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetectFType(t *testing.T) {
	type args struct {
		fn string
	}
	tests := []struct {
		name   string
		args   args
		wantD  []byte
		wantFt string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, gotFt := DetectFType(tt.args.fn)
			if !reflect.DeepEqual(gotD, tt.wantD) {
				t.Errorf("DetectFType() gotD = %v, want %v", gotD, tt.wantD)
			}
			if gotFt != tt.wantFt {
				t.Errorf("DetectFType() gotFt = %v, want %v", gotFt, tt.wantFt)
			}
		})
	}
}

func TestGetFileType(t *testing.T) {
	type args struct {
		fn string
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
			if got := GetFileType(tt.args.fn); got != tt.want {
				t.Errorf("GetFileType() = %v, want %v", got, tt.want)
			}
		})
	}
}
