package controllers

import (
	"reflect"
	"testing"
)

// test for the helper function that check if rooms in slice b are found in any rooms that are already in the slice of paths
func Test_contains(t *testing.T) {
	type args struct {
		a [][]string
		b []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Slice exists in 2D slice",
			args: args{
				a: [][]string{
					{"a", "b", "c"},
					{"d", "e", "f"},
					{"g", "h", "i"},
				},
				b: []string{"d", "e", "f"},
			},
			want: true,
		},
		{
			name: "Slice does not exist in 2D slice",
			args: args{
				a: [][]string{
					{"a", "b", "c"},
					{"d", "e", "f"},
					{"g", "h", "i"},
				},
				b: []string{"x", "y", "z"},
			},
			want: false,
		},
		{
			name: "Empty 2D slice",
			args: args{
				a: [][]string{},
				b: []string{"a", "b", "c"},
			},
			want: false,
		},
		{
			name: "Empty slice to search for",
			args: args{
				a: [][]string{
					{"a", "b", "c"},
					{"d", "e", "f"},
					{"g", "h", "i"},
				},
				b: []string{},
			},
			want: false,
		},
		{
			name: "Both slices are empty",
			args: args{
				a: [][]string{},
				b: []string{},
			},
			want: false,
		},
		{
			name: "Partial match but not exact slice",
			args: args{
				a: [][]string{
					{"a", "b", "c"},
					{"d", "e", "f"},
					{"g", "h", "i"},
				},
				b: []string{"d", "e"},
			},
			want: true,
		},
		{
			name: "Duplicate slices in 2D slice",
			args: args{
				a: [][]string{
					{"a", "b", "c"},
					{"d", "e", "f"},
					{"d", "e", "f"},
				},
				b: []string{"d", "e", "f"},
			},
			want: true,
		},
		{
			name: "Different lengths of paths",
			args: args{
				a: [][]string{
					{"a", "b"},
					{"c", "d", "e"},
					{"a", "b"},
				},
				b: []string{"a", "b"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

// test for optimizer function2
// optimizing paths 2
// returns unique paths only
func Test_optimize2(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Unique paths only",
			args: args{
				paths: [][]string{
					{"a", "b", "c"},
					{"d", "e", "f"},
					{"g", "h", "i"},
				},
			},
			want: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
				{"g", "h", "i"},
			},
		},
		{
			name: "Duplicate paths present",
			args: args{
				paths: [][]string{
					{"a", "b", "c"},
					{"d", "e", "f"},
					{"a", "b", "c"},
				},
			},
			want: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
			},
		},
		{
			name: "One path",
			args: args{
				paths: [][]string{
					{"a", "b", "c"},
				},
			},
			want: [][]string{{"a", "b", "c"}},
		},
		{
			name: "Multiple duplicate paths",
			args: args{
				paths: [][]string{
					{"a", "b", "c"},
					{"a", "b", "c"},
					{"a", "b", "c"},
				},
			},
			want: [][]string{
				{"a", "b", "c"},
			},
		},
		{
			name: "Mixed duplicate and unique paths",
			args: args{
				paths: [][]string{
					{"a", "b", "c"},
					{"d", "e", "f"},
					{"g", "h", "i"},
					{"d", "e", "f"},
					{"g", "h", "i"},
					{"a", "b", "c"},
				},
			},
			want: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
				{"g", "h", "i"},
			},
		},
		{
			name: "Single path",
			args: args{
				paths: [][]string{
					{"a", "b", "c"},
				},
			},
			want: [][]string{
				{"a", "b", "c"},
			},
		},
		{
			name: "Different lengths of paths",
			args: args{
				paths: [][]string{
					{"a", "b"},
					{"c", "d", "e"},
					{"a", "b"},
				},
			},
			want: [][]string{
				{"a", "b"},
				{"c", "d", "e"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimize2(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("optimize2() = %v, want %v", got, tt.want)
			}
		})
	}
}
