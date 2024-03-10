package packets

import (
	"reflect"
	"testing"
)

func Test_GetPacks(t *testing.T) {
	testCases := []struct {
		desc  string
		items int
		sizes []int
		want  []packet
	}{
		{
			desc:  "items 0",
			items: 0,
			sizes: []int{1000, 500},
			want:  []packet{},
		},
		{
			desc:  "itemsles than 0",
			items: -1,
			sizes: []int{1000, 500},
			want:  []packet{},
		},
		{
			desc:  "no sizes",
			items: 501,
			sizes: []int{},
			want:  []packet{},
		},
		{
			desc:  "has lower number",
			items: 501,
			sizes: []int{1000, 500},
			want: []packet{
				{Size: 500, Quantity: 2},
			},
		},
		{
			desc:  "less than lower size",
			items: 1,
			sizes: []int{500, 250},
			want: []packet{
				{Size: 250, Quantity: 1},
			},
		},
		{
			desc:  "fit in two diferent packs",
			items: 501,
			sizes: []int{1000, 500, 250},
			want: []packet{
				{Size: 500, Quantity: 1},
				{Size: 250, Quantity: 1},
			},
		},
		{
			desc:  "fit in three diferent packs",
			items: 12001,
			sizes: []int{5000, 2000, 1000, 500, 250},
			want: []packet{
				{Size: 5000, Quantity: 2},
				{Size: 2000, Quantity: 1},
				{Size: 250, Quantity: 1},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := New().GetPacks(tc.items, tc.sizes)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
