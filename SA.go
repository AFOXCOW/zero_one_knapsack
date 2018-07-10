package zero_one_knapsack

import (
	"math"
	"math/rand"
	"time"
)

func policy(choose []bool, indexx int, indexy int, Items ItemsInterface) (deltaV float64, deltaW int, Case int) {
	if indexx == indexy {
		if choose[indexx] == true {
			deltaV = -Items.Value(indexx)
			deltaW = -Items.Weight(indexx)
			Case = 4
			return
		} else {
			deltaV = Items.Value(indexx)
			deltaW = Items.Weight(indexx)
			Case = 5
			return
		}
	} else if choose[indexx] == false && choose[indexy] == true {

		deltaV = Items.Value(indexx) - Items.Value(indexy)
		deltaW = Items.Weight(indexx) - Items.Weight(indexy)
		Case = 0
		return
	} else if choose[indexx] == true && choose[indexy] == false {

		deltaV = Items.Value(indexy) - Items.Value(indexx)
		deltaW = Items.Weight(indexy) - Items.Weight(indexx)
		Case = 1
		return
	} else if choose[indexx] == false && choose[indexy] == false {

		deltaV = Items.Value(indexx) + Items.Value(indexy)
		deltaW = Items.Weight(indexx) + Items.Weight(indexy)
		Case = 2
		return
	} else {

		deltaV = -Items.Value(indexy)
		deltaW = -Items.Weight(indexy)
		Case = 3
		return
	}
}
func doitbyCase(choose []bool, Case int, x int, y int) {
	switch Case {
	case 0:
		choose[x] = true
		choose[y] = false
	case 1:
		choose[x] = false
		choose[y] = true
	case 2:
		choose[x] = true
		choose[y] = true
	case 3:
		choose[y] = false
	case 4:
		choose[x] = false
	case 5:
		choose[x] = true
	}
}

func Possibility(curr_weight int, deltaV float64, deltaW int, max_w int, T float64) (po float64) {
	if curr_weight+deltaW > max_w {
		po = 0
	} else if deltaV > 0 {
		po = 1
	} else {
		po = math.Exp(float64(deltaV) / T)
	}
	return
}
func randX(n int) (x int) {
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(n)
	return
}
func SA(Items ItemsInterface, capa int) (best_items []Item, best_value float64) {
	num := Items.Len()
	choose := make([]bool, num)
	curr_value := float64(0)
	curr_weight := 0
	T := 100.0     //初始温度
	t_min := 0.1   //终止温度
	ratio := 0.999 //温度下降率
	var (
		deltaV float64
		deltaW int
		Case   int
	)
	for T > t_min {
		x := randX(num)
		y := randX(num)
		deltaV, deltaW, Case = policy(choose, x, y, Items)
		po := Possibility(curr_weight, deltaV, deltaW, capa, T)
		test := rand.Float64()
		if test < po {
			doitbyCase(choose, Case, x, y)
			curr_value += deltaV
			curr_weight += deltaW
		}
		T = ratio * T
	}
	for i := 0; i < len(choose); i++ {
		if choose[i] == true {
			best_items = append(best_items, Item{i, Items.Weight(i), Items.Value(i)})
		}
	}
	best_value = curr_value
	return
}
