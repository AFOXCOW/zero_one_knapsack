package zero_one_knapsack

func max(a float64, b float64) (c float64) {
	if a > b {
		c = a
		return
	}
	c = b
	return
}

func Dyn_program(Items ItemsInterface, capa int) (bestItems []Item, value float64) {
	var tmp [][]float64
	num := Items.Len()
	//ceate the space for tmp array
	for i := 0; i < num; i++ {
		x := make([]float64, capa+1)
		tmp = append(tmp, x)
	}
	//initialize the tmp array
	for j := 1; j < capa+1; j++ {
		if Items.Weight(0) <= capa {
			tmp[0][j] = Items.Value(0)
		}
	}
	//fill the tmp array with the recursive expression
	for i := 1; i < num; i++ {
		for j := 1; j < capa+1; j++ {
			if j-Items.Weight(i) >= 0 {
				tmp[i][j] = max(tmp[i-1][j], Items.Value(i)+tmp[i-1][j-Items.Weight(i)])
			} else {
				tmp[i][j] = tmp[i-1][j]
			}

		}
	}
	//from the best value position ,back to the start position to find one of the best value solution.
	i := num - 1
	j := capa
	value = tmp[i][j]
	for i > 0 && j > 0 {
		if tmp[i][j] != tmp[i-1][j] {
			bestItems = append(bestItems, Item{i, Items.Weight(i), Items.Value(i)})
			j = j - Items.Weight(i)
			i = i - 1
		} else {
			i = i - 1
		}
	}
	if i == 0 && j != 0 {
		bestItems = append(bestItems, Item{i, Items.Weight(i), Items.Value(i)})
	}
	return
}
