package packets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetPacks(t *testing.T) {
	testCases := []struct {
		desc  string
		items int
		sizes []int
		want  map[int]int
	}{
		{
			items: 1000,
			sizes: []int{2000, 3000},
			want:  map[int]int{2000: 1},
		},
		{
			items: 500000,
			sizes: []int{23, 31, 53},
			want:  map[int]int{23: 2, 31: 7, 53: 9429},
		},
		{
			items: 14,
			sizes: []int{5, 12},
			want:  map[int]int{5: 3},
		},
		{
			items: -1,
			sizes: []int{1000, 500},
			want:  map[int]int{},
		},
		{
			items: 501,
			sizes: []int{},
			want:  map[int]int{},
		},
		{
			items: 500,
			sizes: []int{1000, 500},
			want:  map[int]int{500: 1},
		},
		{
			items: 501,
			sizes: []int{1000, 500},
			want:  map[int]int{1000: 1},
		},
		{
			items: 1,
			sizes: []int{1000, 500},
			want:  map[int]int{500: 1},
		},
		{
			items: 501,
			sizes: []int{1000, 500, 250},
			want:  map[int]int{500: 1, 250: 1},
		},
		{
			items: 12001,
			sizes: []int{5000, 2000, 1000, 500, 250},
			want: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := New().GetPacks(tc.sizes, tc.items)
			assert.Equal(t, tc.want, got)
		})
	}
}
