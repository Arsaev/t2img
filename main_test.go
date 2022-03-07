package main

import (
	"testing"
)

func Test_generate(t *testing.T) {
	type args struct {
		input    string
		fontSize float64
		fontPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "default test",
			args: args{
				input:    "hello",
				fontSize: 24.0,
				fontPath: "SFNSMono.ttf",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generate(tt.args.input, tt.args.fontSize, tt.args.fontPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("generate() = %v, want not empty byte array", got)
			}
		})
	}
}
