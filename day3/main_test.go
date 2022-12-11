package main

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "intersect",
			args: args{
				a: []byte("abc"),
				b: []byte("bcd"),
			},
			want: []byte("bc"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asciiToPrio(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A",
			args: args{
				c: 'A',
			},
			want: 27,
		},
		{
			name: "a",
			args: args{
				c: 'a',
			},
			want: 1,
		},
		{
			name: "p",
			args: args{
				c: 'p',
			},
			want: 16,
		},
		{
			name: "P",
			args: args{
				c: 'P',
			},
			want: 42,
		},
		{
			name: "Z",
			args: args{
				c: 'Z',
			},
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asciiToPrio(tt.args.c); got != tt.want {
				t.Errorf("asciiToPrio() = %v, want %v", got, tt.want)
			}
		})
	}
}
