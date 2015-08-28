// Copyright 2015, David Howden
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import (
	"reflect"
	"strings"
	"testing"
)

func stringToPath(s string) Path {
	return stringSliceToPath(strings.Split(s, PathSeparator))
}

func stringSliceToPath(s []string) Path {
	p := make(Path, len(s))
	for i, x := range s {
		p[i] = Key(x)
	}
	return p
}

func TestPathEqual(t *testing.T) {
	tests := []struct {
		p, q  Path
		equal bool
	}{
		{
			Path(nil), Path(nil),
			true,
		},
		{
			Path([]Key{}), Path([]Key{}),
			true,
		},
		{
			stringToPath("a"), stringToPath("a"),
			true,
		},
		{
			stringToPath("a:b"), stringToPath("a"),
			false,
		},
		{
			Path(nil), stringToPath("a"),
			false,
		},
	}

	for ii, tt := range tests {
		if tt.p.Equal(tt.q) != tt.equal {
			t.Errorf("[%d] (%#v).Equal(%#v) = %v, expected %v", ii, tt.p, tt.q, !tt.equal, tt.equal)
		}
	}
}

func TestOrderedIntersection(t *testing.T) {
	tests := []struct {
		in  [][]Path
		out []Path
	}{
		{
			in:  nil,
			out: []Path{},
		},

		{
			in: [][]Path{
				{stringToPath("A")},
			},
			out: []Path{stringToPath("A")},
		},

		{
			in: [][]Path{
				{stringToPath("A")},
				{stringToPath("B")},
			},
			out: []Path{},
		},

		{
			in: [][]Path{
				{stringToPath("A")},
				{stringToPath("B"), stringToPath("A")},
			},
			out: []Path{stringToPath("A")},
		},

		{
			in: [][]Path{
				{stringToPath("A"), stringToPath("B")},
				{stringToPath("B"), stringToPath("A")},
				{stringToPath("A"), stringToPath("B"), stringToPath("C")},
				{stringToPath("C"), stringToPath("A"), stringToPath("B"), stringToPath("B")},
			},
			out: []Path{stringToPath("B"), stringToPath("A")},
		},
	}

	for ii, tt := range tests {
		got := OrderedIntersection(tt.in...)
		if !reflect.DeepEqual(got, tt.out) {
			t.Errorf("[%d] got %#v, expected: %#v", ii, got, tt.out)
		}
	}
}
