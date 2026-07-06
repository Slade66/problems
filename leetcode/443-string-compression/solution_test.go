package main

import "testing"

func TestCompress(t *testing.T) {
	testCases := []struct {
		name       string
		chars      []byte
		wantLength int
		wantPrefix []byte
	}{
		{
			name:       "official example 1",
			chars:      []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			wantLength: 6,
			wantPrefix: []byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			name:       "official example 2",
			chars:      []byte{'a'},
			wantLength: 1,
			wantPrefix: []byte{'a'},
		},
		{
			name:       "official example 3",
			chars:      []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			wantLength: 4,
			wantPrefix: []byte{'a', 'b', '1', '2'},
		},
		{
			name:       "group length exactly ten",
			chars:      []byte{'z', 'z', 'z', 'z', 'z', 'z', 'z', 'z', 'z', 'z'},
			wantLength: 3,
			wantPrefix: []byte{'z', '1', '0'},
		},
		{
			name:       "all unique characters",
			chars:      []byte{'a', 'b', 'C', '1', '!'},
			wantLength: 5,
			wantPrefix: []byte{'a', 'b', 'C', '1', '!'},
		},
		{
			name:       "multiple digit and single groups",
			chars:      []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'b', 'c', 'c'},
			wantLength: 6,
			wantPrefix: []byte{'a', '1', '2', 'b', 'c', '2'},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotLength := compress(tc.chars)
			if gotLength != tc.wantLength {
				t.Errorf("compress(%v) length = %d, want %d", tc.chars, gotLength, tc.wantLength)
			}
			for i := 0; i < tc.wantLength; i++ {
				if tc.chars[i] != tc.wantPrefix[i] {
					t.Errorf("compress prefix at index %d = %q, want %q; got prefix %q", i, tc.chars[i], tc.wantPrefix[i], tc.chars[:tc.wantLength])
				}
			}
		})
	}
}
