package zero_one_backpack

import (
	"strconv"
)

type thing struct {
	id     int
	weight int
	value  int
}

func (th thing) String() string {
	return "ID:" + strconv.Itoa(th.id) + "  weight:" + strconv.Itoa(th.weight) + "  value:" + strconv.Itoa(th.value)
}

func max(a int, b int) (c int) {
	if a > b {
		c = a
		return
	}
	c = b
	return
}
func Dyn_program(w []int, v []int, capa int, num int) (things []thing, value int) {
	var tmp [][]int
	for i := 0; i < num; i++ {
		x := make([]int, capa+1)
		tmp = append(tmp, x)
	}
	for j := 1; j < capa+1; j++ {
		if w[0] <= capa {
			tmp[0][j] = v[0]
		}
	}
	for i := 1; i < num; i++ {
		for j := 1; j < capa+1; j++ {
			if j-w[i] >= 0 {
				tmp[i][j] = max(tmp[i-1][j], v[i]+tmp[i-1][j-w[i]])
			} else {
				tmp[i][j] = tmp[i-1][j]
			}

		}
	}
	i := num - 1
	j := capa
	value = tmp[i][j]
	for i > 0 && j > 0 {
		if tmp[i][j] != tmp[i-1][j] {
			things = append(things, thing{i, w[i], v[i]})
			j = j - w[i]
			i = i - 1
		} else {
			i = i - 1
		}
	}
	if i == 0 && j != 0 {
		things = append(things, thing{i, w[i], v[i]})
	}
	return
}
