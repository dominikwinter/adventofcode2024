package lib

import (
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	type args struct {
		slice []int
		i     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Remove first element",
			args: args{slice: []int{1, 2, 3, 4}, i: 0},
			want: []int{2, 3, 4},
		},
		{
			name: "Remove middle element",
			args: args{slice: []int{1, 2, 3, 4}, i: 2},
			want: []int{1, 2, 4},
		},
		{
			name: "Remove last element",
			args: args{slice: []int{1, 2, 3, 4}, i: 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Remove from single element slice",
			args: args{slice: []int{1}, i: 0},
			want: []int{},
		},
		{
			name: "Remove with index out of range",
			args: args{slice: []int{1, 2, 3, 4}, i: 4},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.slice, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

