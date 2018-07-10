package zero_one_knapsack

import (
	"container/heap"
	"fmt"
	"sort"
)

type HeapNode struct {
	upper  float64 //node`s upper board:the priority attribution
	value  float64 //node`s value
	weight float64 //node`s weight
	level  int     //node`s level
	lkid   bool
	index  int //node`s index in max-heap
}

type PriorityQueue []*HeapNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].value > pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	//exchange the node  but keep the index the original value
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	heapNode := x.(*HeapNode)
	heapNode.index = n
	*pq = append(*pq, heapNode)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	heapnode := old[n-1]
	heapnode.index = -1 // for safety
	*pq = old[0 : n-1]
	return heapnode
}

func (pq *PriorityQueue) update(heapnode *HeapNode, level int, value float64, weight float64, upper float64, lkid bool) {
	heapnode.level = level
	heapnode.value = value
	heapnode.weight = weight
	heapnode.upper = upper
	heapnode.lkid = lkid
	heap.Fix(pq, heapnode.index)
}

func max_bound(t int, curr_weight int, curr_value float64, capa int, Items ItemsInterface) (upper float64) {
	n := Items.Len()
	left := capa - curr_weight
	upper = float64(curr_value)
	for t < n && Items.Weight(t) <= left {
		left -= Items.Weight(t)
		upper += float64(Items.Value(t))
		t++
	}
	if t < n {
		upper += float64(Items.Value(t)) / float64(Items.Weight(t)) * float64(left)
	}
	return
}
func addLiveNode(pq *PriorityQueue, upper float64, curr_value float64, curr_weight float64, level int, lkid bool, num int) {
	heapnode := HeapNode{upper, curr_value, curr_weight, level, lkid, pq.Len()}
	if level <= num {
		heap.Push(pq, &heapnode)
	}
}
func Max_queue(Items ItemsInterface, capa int) (bestvalue float64, arr []HeapNode) {
	num := Items.Len()
	curr_weight := 0
	curr_value := float64(0)
	var (
		Lflag bool
		Rflag bool
	)
	i := 0
	pq := make(PriorityQueue, 0)
	upper := max_bound(i, curr_weight, curr_value, capa, Items)
	for true {
		Lflag = false
		Rflag = false
		if i < num && curr_weight+Items.Weight(i) <= capa {
			if float64(upper) > bestvalue {
				Lflag = true
				if float64(curr_value+Items.Value(i)) > bestvalue {
					bestvalue = float64(curr_value + Items.Value(i))
				}
				addLiveNode(&pq, upper, float64(curr_value+Items.Value(i)), float64(curr_weight+Items.Weight(i)), i+1, true, num)
			}
		}
		upper = max_bound(i+1, curr_weight, curr_value, capa, Items)
		if i < num && upper >= bestvalue {
			Rflag = true
			addLiveNode(&pq, upper, float64(curr_value), float64(curr_weight), i+1, false, num)
		}
		if pq.Len() == 0 && !Lflag && !Rflag {
			return
		}
		if pq.Len() != 0 {
			heapnode := heap.Pop(&pq).(*HeapNode)
			curr_weight = int(heapnode.weight)
			curr_value = heapnode.value
			upper = heapnode.upper
			i = heapnode.level
			arr = append(arr, *heapnode)
		}
	}
	return
}

func Nodes2Path(arr []HeapNode, num int) (path []bool) {
	for i := 0; i < len(arr); i++ {
		if arr[len(arr)-1].level != num {
			arr = append(arr[:len(arr)-1])
		}
	}
	for i := 0; i < len(arr); i++ {
		path = append(path, arr[i].lkid)
	}
	return
}

type Pair struct {
	key   int
	Value float64
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func SortByV_W(w []int, v []int) (weight []int, value []int) {
	var v_w PairList
	for i := 0; i < len(w); i++ {
		v_w = append(v_w, Pair{i, float64(v[i]) / float64(w[i])})
	}
	sort.Sort(v_w)
	for i := 0; i < len(v_w); i++ {
		j := v_w[i].key
		value = append(value, v[j])
		weight = append(weight, w[j])
	}
	return
}

func PQpathPrint(arr []bool, Items ItemsInterface) {
	num := Items.Len()
	for i := len(arr) - num; i < len(arr); i++ {
		if arr[i] == true {
			fmt.Printf("ID : %d  weight : %d value : %d\n", i-len(arr)+num, Items.Weight(i-len(arr)+num), Items.Value(i-len(arr)+num))
		}
	}
}
