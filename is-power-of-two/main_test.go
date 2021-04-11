package main

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{i: 0},
			want: false,
		},
		{
			name: "test 2",
			args: args{i: 15},
			want: false,
		},
		{
			name: "test 3",
			args: args{i: 16},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.i); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}