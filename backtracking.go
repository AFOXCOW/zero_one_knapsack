package zero_one_backpack

import (
	"fmt"
	"math"
)

func Num2path(num int) (path []int) {
	for num != 0 {
		path = append(path, num%2)
		num = (num - 1) / 2
	}
	return
}
func path2value(path []int, value []int) (sum int) {
	sum = 0
	for i := 0; i < len(path); i++ {
		if path[i] == 1 {
			sum += value[i]
		}
	}
	return
}

func dead(w []int, capa int, curr int) bool {
	arr := Num2path(curr)
	var path []int
	for i := 0; i < len(arr); i++ {
		path = append(path, arr[len(arr)-i-1])
	}
	totalweight := 0
	for i := 0; i < len(path); i++ {
		if path[i] == 1 {
			totalweight += w[i]
		}
	}
	if totalweight > capa {
		return true
	}
	return false
}
func end(curr int, node_num int) bool {
	if 2*curr+1 >= node_num {
		return true
	}
	return false
}
func BackTracking(w []int, v []int, capa int, num int) (things []thing, best int) {
	node_num := int(math.Pow(2, float64(num+1)) - 1)
	Tree := make([]int, node_num)
	curr, pre := 0, 0
	var paths [][]int
	var arr []int
	for curr != 0 || pre != 2 {
		Tree[curr] = 1
		if dead(w, capa, curr) {
			tmp := curr
			curr = pre
			pre = tmp
		} else if end(curr, node_num) {
			path := Num2path(curr)
			for i := 0; i < len(path); i++ {
				arr = append(arr, path[len(path)-i-1])
			}
			paths = append(paths, arr)
			arr = nil
			fmt.Println()
			tmp := curr
			curr = pre
			pre = tmp
		} else if Tree[2*curr+1] == 0 {
			pre = curr
			curr = 2*curr + 1
		} else if Tree[2*curr+2] == 0 {
			pre = curr
			curr = 2*curr + 2
		} else {
			pre = curr
			curr = (curr - 1) / 2
		}
	}
	max := 0
	index := 0
	for i := 0; i < len(paths); i++ {
		value := path2value(paths[i], v)
		if value > max {
			index = i
			max = value
		}
	}
	for i := 0; i < len(paths[index]); i++ {
		if paths[index][i] != 0 {
			things = append(things, thing{i, w[i], v[i]})
		}
	}
	best = max
	return
}
