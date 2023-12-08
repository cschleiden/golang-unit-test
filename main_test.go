package main

import "testing"

func TestFoo(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1 + 2 = 3",
			args{1, 2},
			3,
		},
		{
			"2 + 2 = 4",
			args{2, 2},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Foo(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Foo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBar(t *testing.T) {
	// do nothing
}
