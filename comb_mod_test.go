package combmod

import "testing"

func Test_combmod(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"6-3", args{6, 3}, 20},
		{"(100 choose 50) mod (1e9+7)", args{100, 50}, 538992043},
		{"(100000 choose 4999) mod (1e9+7)", args{100000, 4999}, 533759731},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combmod(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("combmod() = %v, want %v", got, tt.want)
			}
		})
	}
}
