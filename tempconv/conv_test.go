package tempconv

import (
	"testing"
)

func TestCtoF(t *testing.T) {
	type args struct {
		c Celsius
	}
	tests := []struct {
		name string
		args args
		want Fahrenheit
	}{
		{
			name: "Should_ConvertTo32F_When0C",
			args: args{
				c: Celsius(0),
			},
			want: Fahrenheit(32),
		},
		{
			name: "Should_ConvertTo120F_When50C",
			args: args{
				c: Celsius(50),
			},
			want: Fahrenheit(122),
		},
		{
			name: "Should_ConvertTo14F_When-10C",
			args: args{
				c: Celsius(-10),
			},
			want: Fahrenheit(14),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtoF(tt.args.c); got != tt.want {
				t.Errorf("CtoF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFtoC(t *testing.T) {
	type args struct {
		f Fahrenheit
	}
	tests := []struct {
		name string
		args args
		want Celsius
	}{
		{
			name: "Should_ConvertTo0C_When32F",
			args: args{
				f: Fahrenheit(32),
			},
			want: Celsius(0),
		},
		{
			name: "Should_ConvertTo-20C_When-4F",
			args: args{
				f: Fahrenheit(-4),
			},
			want: Celsius(-20),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FtoC(tt.args.f); got != tt.want {
				t.Errorf("FtoC() = %v, want %v", got, tt.want)
			}
		})
	}
}
