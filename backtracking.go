package zero_one_knapsack

import (
	"math"
	"sort"
)

/*
	the defination of the solution tree are as follows.
	use the Len  Less Swap  to satisfy the sort interface.
	use the SearchTree to find the nodeid'x' in TreeNodeSlice.
*/
type TreeNode struct {
	nodeid  uint64
	visited bool
}

type TreeNodeSlice []TreeNode

func (t TreeNodeSlice) Len() int {
	return len(t)
}

func (t TreeNodeSlice) Less(i, j int) bool {
	return t[i].nodeid < t[j].nodeid
}
func (t TreeNodeSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func SearchTree(a TreeNodeSlice, x uint64) int {
	return sort.Search(len(a), func(i int) bool { return a[i].nodeid >= x })
}

/*
	the node2path`s function is that it can transfer the nodeid to the path from Treeroot to current_node.
	such as 3 will be transfer to 11.and the 11 means the 0th item and 1th item are choosed.
*/
func node2path(current_node uint64) (path []int) {
	var arr []int
	for current_node != 0 {
		arr = append(arr, int(current_node%2))
		current_node = (current_node - 1) / 2
	}
	for i := 0; i < len(arr); i++ {
		path = append(path, arr[len(arr)-i-1])
	}
	return
}

/*
	apparently a path got a value so far
*/
func path2value(path []int, items ItemsInterface) (sum float64) {
	sum = 0
	for i := 0; i < len(path); i++ {
		if path[i] == 1 {
			sum += items.Value(i)
		}
	}
	return
}

/*
	judge the current node.
	if it`s overweighted then  it`s a dead node.
*/
func current_dead(Items ItemsInterface, capa int, curr uint64) bool {

	path := node2path(curr)
	totalweight := 0
	for i := 0; i < len(path); i++ {
		if path[i] == 1 {
			totalweight += Items.Weight(i)
		}
	}
	if totalweight > capa {
		return true
	}
	return false
}

/*
	if current node is the leaf node.
	then it`s a optional solution but maybe it`s not the best solution.
*/
func end(curr uint64, node_num uint64) bool {
	if 2*curr+1 >= node_num {
		return true
	}
	return false
}

/*
	if current node is visited,then return true
	else return false.
*/
func visited(Tree TreeNodeSlice, nodeid uint64) bool {
	i := SearchTree(Tree, nodeid)
	if i < len(Tree) && Tree[i].nodeid == nodeid && Tree[i].visited == true {
		return true
	} else {
		return false
	}
}

/*
	BackTracking function,select the best solution(one of the best) from all optional solution.
*/
func BackTracking(Items ItemsInterface, capa int) (bestItems []Item, best_value float64) {
	var Tree TreeNodeSlice
	num := Items.Len()
	node_num := uint64(math.Pow(2, float64(num+1)) - 1)
	var curr, pre uint64
	curr, pre = 0, 0
	var paths [][]int
	var arr []int

	sort.Sort(Tree)

	for curr != 0 || pre != 2 {
		Tree = append(Tree, TreeNode{curr, true})
		if current_dead(Items, capa, curr) {
			tmp := curr
			curr = pre
			pre = tmp
		} else if end(curr, node_num) {
			path := node2path(curr)
			for i := 0; i < len(path); i++ {
				arr = append(arr, path[len(path)-i-1])
			}
			paths = append(paths, arr)
			arr = nil
			tmp := curr
			curr = pre
			pre = tmp
		} else if visited(Tree, 2*curr+1) == false {
			pre = curr
			curr = 2*curr + 1
		} else if visited(Tree, 2*curr+2) == false {
			pre = curr
			curr = 2*curr + 2
		} else {
			pre = curr
			curr = (curr - 1) / 2
		}
	}
	max := float64(0)
	index := 0
	for i := 0; i < len(paths); i++ {
		value := path2value(paths[i], Items)
		if value > max {
			index = i
			max = value
		}
	}
	for i := 0; i < len(paths[index]); i++ {
		if paths[index][i] != 0 {
			bestItems = append(bestItems, Item{i, Items.Weight(i), Items.Value(i)})
		}
	}
	best_value = max
	return
}
