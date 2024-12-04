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
func TestTranspose(t *testing.T) {
	type args struct {
		matrix [][]int
		steps  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Transpose 1 step",
			args: args{matrix: [][]int{{1, 2}, {3, 4}}, steps: 1},
			want: [][]int{{3, 1}, {4, 2}},
		},
		{
			name: "Transpose 2 steps",
			args: args{matrix: [][]int{{1, 2}, {3, 4}}, steps: 2},
			want: [][]int{{4, 3}, {2, 1}},
		},
		{
			name: "Transpose 3 steps",
			args: args{matrix: [][]int{{1, 2}, {3, 4}}, steps: 3},
			want: [][]int{{2, 4}, {1, 3}},
		},
		{
			name: "Transpose 4 steps (full rotation)",
			args: args{matrix: [][]int{{1, 2}, {3, 4}}, steps: 4},
			want: [][]int{{1, 2}, {3, 4}},
		},
		{
			name: "Transpose 0 steps (no rotation)",
			args: args{matrix: [][]int{{1, 2}, {3, 4}}, steps: 0},
			want: [][]int{{1, 2}, {3, 4}},
		},
		{
			name: "Transpose 5 steps (1 step effectively)",
			args: args{matrix: [][]int{{1, 2}, {3, 4}}, steps: 5},
			want: [][]int{{3, 1}, {4, 2}},
		},
		{
			name: "Transpose 10x10 matrix 1 step",
			args: args{matrix: [][]int{
				{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
				{31, 32, 33, 34, 35, 36, 37, 38, 39, 40},
				{41, 42, 43, 44, 45, 46, 47, 48, 49, 50},
				{51, 52, 53, 54, 55, 56, 57, 58, 59, 60},
				{61, 62, 63, 64, 65, 66, 67, 68, 69, 70},
				{71, 72, 73, 74, 75, 76, 77, 78, 79, 80},
				{81, 82, 83, 84, 85, 86, 87, 88, 89, 90},
				{91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
				{101, 102, 103, 104, 105, 106, 107, 108, 109, 110},
			}, steps: 1},
			want: [][]int{
				{101, 91, 81, 71, 61, 51, 41, 31, 21, 11},
				{102, 92, 82, 72, 62, 52, 42, 32, 22, 12},
				{103, 93, 83, 73, 63, 53, 43, 33, 23, 13},
				{104, 94, 84, 74, 64, 54, 44, 34, 24, 14},
				{105, 95, 85, 75, 65, 55, 45, 35, 25, 15},
				{106, 96, 86, 76, 66, 56, 46, 36, 26, 16},
				{107, 97, 87, 77, 67, 57, 47, 37, 27, 17},
				{108, 98, 88, 78, 68, 58, 48, 38, 28, 18},
				{109, 99, 89, 79, 69, 59, 49, 39, 29, 19},
				{110, 100, 90, 80, 70, 60, 50, 40, 30, 20},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Transpose(tt.args.matrix, tt.args.steps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTransposeStrings(t *testing.T) {
	type args struct {
		matrix []string
		steps  int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Transpose 1 step",
			args: args{matrix: []string{"ab", "cd"}, steps: 1},
			want: []string{"ca", "db"},
		},
		{
			name: "Transpose 2 steps",
			args: args{matrix: []string{"ab", "cd"}, steps: 2},
			want: []string{"dc", "ba"},
		},
		{
			name: "Transpose 3 steps",
			args: args{matrix: []string{"ab", "cd"}, steps: 3},
			want: []string{"bd", "ac"},
		},
		{
			name: "Transpose 4 steps (full rotation)",
			args: args{matrix: []string{"ab", "cd"}, steps: 4},
			want: []string{"ab", "cd"},
		},
		{
			name: "Transpose 0 steps (no rotation)",
			args: args{matrix: []string{"ab", "cd"}, steps: 0},
			want: []string{"ab", "cd"},
		},
		{
			name: "Transpose 5 steps (1 step effectively)",
			args: args{matrix: []string{"ab", "cd"}, steps: 5},
			want: []string{"ca", "db"},
		},
		{
			name: "Transpose 3x3 matrix 1 step",
			args: args{matrix: []string{"abc", "def", "ghi"}, steps: 1},
			want: []string{"gda", "heb", "ifc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransposeStrings(tt.args.matrix, tt.args.steps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransposeStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTransposeDiagonalStrings(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "3x3 matrix",
			args: args{input: []string{"abc", "def", "ghi"}},
			want: []string{"a", "bd", "ceg", "fh", "i"},
		},
		{
			name: "2x2 matrix",
			args: args{input: []string{"ab", "cd"}},
			want: []string{"a", "bc", "d"},
		},
		{
			name: "4x4 matrix",
			args: args{input: []string{"abcd", "efgh", "ijkl", "mnop"}},
			want: []string{"a", "be", "cfi", "dgjm", "hkn", "lo", "p"},
		},
		{
			name: "Single row",
			args: args{input: []string{"abcd"}},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "Single column",
			args: args{input: []string{"a", "b", "c", "d"}},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "Empty input",
			args: args{input: []string{}},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransposeDiagonalStrings(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransposeDiagonalStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestSliceToInts(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Valid integers",
			args: args{list: []string{"1", "2", "3"}},
			want: []int{1, 2, 3},
		},
		{
			name: "Empty list",
			args: args{list: []string{}},
			want: []int{},
		},
		{
			name: "List with negative integers",
			args: args{list: []string{"-1", "-2", "-3"}},
			want: []int{-1, -2, -3},
		},
		{
			name: "List with mixed valid and invalid integers",
			args: args{list: []string{"1", "a", "3"}},
			want: []int{1, 0, 3},
		},
		{
			name: "List with all invalid integers",
			args: args{list: []string{"a", "b", "c"}},
			want: []int{0, 0, 0},
		},
		{
			name: "List with empty strings",
			args: args{list: []string{"", "", ""}},
			want: []int{0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToInts(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
