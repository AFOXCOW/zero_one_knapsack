package zero_one_backpack

func exist(arr []int, num int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return true
		}
	}
	return false
}
func SA_BackTrack(w []int, v []int, capa int, num int) (things []thing, best int) {
	things, best = SA(w, v, capa, num)
	var arr []int
	var path []int
	var totalweight int
	for i := 0; i < len(things); i++ {
		arr = append(arr, things[i].id)
	}
	for i := 0; i < num/2; i++ {
		if exist(arr, i) {
			path = append(path, 1)
		} else {
			path = append(path, 0)
		}
	}
	for i := 0; i < len(path); i++ {
		if path[i] == 1 {
			totalweight += w[i]
		}
	}
	index := 0
	for i := 0; i < len(things); i++ {
		if things[index].id < len(path) {
			index++
		}
	}
	things = things[:index]
	capa -= totalweight
	num -= len(path)
	w = w[len(path):]
	v = v[len(path):]
	things_bk, _ := BackTracking(w, v, capa, num)
	for i := 0; i < len(things_bk); i++ {
		things_bk[i].id += int(num/2 + 1)
		things = append(things, things_bk[i])
	}
	best = 0
	for i := 0; i < len(things); i++ {
		best += things[i].value
	}
	return
}
