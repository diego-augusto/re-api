package packets

import (
	"sort"
)

// PackSrv is the interface for the pack service
type PackSrv interface {
	GetPacks(int, []int) []packet
}

type packSrv struct{}

// New returns a new instance of the pack service
func New() packSrv {
	return packSrv{}
}

type packet struct {
	Size     int `json:"size"`
	Quantity int `json:"quantity"`
}

// GetPacks returns the number of packs needed to fulfill the order
func (s packSrv) GetPacks(items int, sizes []int) []packet {

	if len(sizes) == 0 {
		return []packet{}
	}

	if items < 1 {
		return []packet{}
	}

	// sort sizes in descending order
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	// get the smallest size
	var smallest int
	if len(sizes) > 0 {
		smallest = sizes[len(sizes)-1]
	}

	packs := make([]packet, 0)

	for _, size := range sizes {
		if items >= size {
			quantity := items / size
			items %= size
			packs = append(packs, packet{Size: size, Quantity: quantity})
		}
	}

	// if there are items left, add a pack with the lower size
	if items > 0 {
		// if the last pack has smallest size, increment the quantity
		if len(packs) > 0 && packs[len(packs)-1].Size == smallest {
			packs[len(packs)-1].Quantity++
			return packs
		}
		// add a new smallest size pack, otherwise
		packs = append(packs, packet{Size: smallest, Quantity: 1})
	}

	return packs
}
