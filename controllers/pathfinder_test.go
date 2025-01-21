package controllers

import "testing"

//test for the helper function that check if rooms in slice b are found in any rooms that are already in the slice of paths

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


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
