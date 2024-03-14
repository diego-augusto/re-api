package packets

import (
	"math"
	"sort"
)

// PackSrv is the interface for the pack service
type PackSrv interface {
	GetPacks([]int, int) map[int]int
}

type packSrv struct{}

// New returns a new instance of the pack service
func New() packSrv {
	return packSrv{}
}

// GetPacks returns the number of packs needed to fulfill the order
func (packSrv) GetPacks(packs []int, order int) map[int]int {

	if order <= 0 {
		return map[int]int{}
	}

	if len(packs) == 0 {
		return map[int]int{}
	}

	sort.Slice(packs, func(i, j int) bool {
		return packs[i] > packs[j]
	})

	combination := make([]int, 0)
	combinations := make(map[int][][]int)
	memo := make(map[int]bool)
	find(order, packs, combination, combinations, memo)

	if len(combinations) == 0 {
		return map[int]int{}
	}

	// find combination with the least amount of items to fulfill the order
	max := math.MinInt
	for k := range combinations {
		if k > max {
			max = k
		}
	}

	// find the combination with the least amount of items
	// if there are multiple combinations with the same amount of items
	minLegth := math.MaxInt
	minArr := make([]int, 0)
	for _, v := range combinations[max] {
		if len(v) < minLegth {
			minLegth = len(v)
			minArr = v
		}
	}

	// count the amount of packs needed
	mapArr := make(map[int]int)
	for _, n := range minArr {
		mapArr[n]++
	}

	return mapArr
}

// Find is a recursive function that finds all the combinations of packs.
// Can be used to fulfill the order with memoization.
func find(orderAmount int, packs []int, currentComb []int, combMap map[int][][]int, memo map[int]bool) {
	if _, exists := memo[orderAmount]; exists {
		return
	}
	if orderAmount <= 0 {
		combMap[orderAmount] = append(combMap[orderAmount], append([]int(nil), currentComb...))
		memo[orderAmount] = true
		return
	}
	for i, v := range packs {
		find(orderAmount-v, append(packs[i:], packs[:i]...), append(currentComb, v), combMap, memo)
	}
	memo[orderAmount] = true
}
