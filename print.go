package zero_one_knapsack

import (
	"strconv"
)

func (i Item) String() string {
	return "ID:" + strconv.Itoa(i.id) + "  weight:" + strconv.Itoa(i.weight) + "  value:" + strconv.FormatFloat(i.value, 'E', -1, 64)
}
