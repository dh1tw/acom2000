package acom2000

import (
	"testing"
)

func Test_dec2Ascii(t *testing.T) {
	type args struct {
		dec int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"31923", args{dec: 31923}, "7<;3", false},
		{"negative", args{dec: -31923}, "", true},
		{"large", args{dec: 329812331923}, "4<<:58<=93", false},
		{"zero", args{dec: 0}, "0", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dec2Ascii(tt.args.dec)
			if (err != nil) != tt.wantErr {
				t.Errorf("dec2Ascii() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("dec2Ascii() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ascii2Dec(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"7<;3", args{s: "7<;3"}, 31923, false},
		{"illegal chars", args{s: "7test3"}, 0, true},
		{"zero", args{s: "0"}, 0, false},
		{"large", args{s: "4<<:58<=93"}, 329812331923, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ascii2Dec(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ascii2Dec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ascii2Dec() = %v, want %v", got, tt.want)
			}
		})
	}
}
