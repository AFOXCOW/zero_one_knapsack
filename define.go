package zero_one_knapsack

type ItemsInterface interface {
	Id(i int) int
	// Len is the number of elements in the collection.
	Len() int

	Weight(i int) int

	Value(i int) float64
}
type Item struct {
	id int

	weight int

	value float64
}
