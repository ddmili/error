package error

import (
	"testing"
)

func TestNewError(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test 1", args{"error message"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewError(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("NewError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewFail(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test 1", args{"error message"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewFail(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("NewFail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewInfo(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test 1", args{"error message"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewInfo(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("NewInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewPanic(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test 1", args{"error message"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewPanic(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("NewPanic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_print(t *testing.T) {
	type args struct {
		level string
		str   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			print(tt.args.level, tt.args.str)
		})
	}
}
