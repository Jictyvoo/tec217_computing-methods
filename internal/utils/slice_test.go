package utils

import (
	"testing"
)

func TestReduceSlice(t *testing.T) {
	type args struct {
		input    []int
		callback func(old, a int) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty Input",
			args: args{
				input:    []int{},
				callback: func(old, a int) int { return old + a },
			},
			want: 0,
		},
		{
			name: "Single Element",
			args: args{
				input:    []int{10},
				callback: func(old, a int) int { return old + a },
			},
			want: 10,
		},
		{
			name: "Multiple Elements",
			args: args{
				input:    []int{1, 2, 3, 4, 5},
				callback: func(old, a int) int { return old + a },
			},
			want: 15,
		},
		{
			name: "Negative Numbers",
			args: args{
				input:    []int{-1, -2, -3, -4, -5},
				callback: func(old, a int) int { return old + a },
			},
			want: -15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReduceSlice(tt.args.input, tt.args.callback); got != tt.want {
				t.Errorf("ReduceSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
