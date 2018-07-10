package zero_one_knapsack

type ItemsInterface interface {
	// Id is the unique label of item.
	Id(i int) int
	// Len is the number of elements in the collection.
	Len() int
	/*
		Weight is the weight of the i(th) item.
		In most condition, the weight is not the type(int),such as 1.5kg.
		But to solve the problem,we must define the smallest step of weight.
		So the smallest step is 1.
		If you want to use 1.5kg. You just change the measure to 1500g or 15 hundred g.
	*/

	Weight(i int) int
	//Value is the value of the i(th) item. of cause it is type(float64).
	Value(i int) float64
}

//here is the standard output item.
//whatever you input.
//The output Items is a slice of struct Item once your input satisfy the interface.
type Item struct {
	id int

	weight int

	value float64
}

//
type NumInterface interface {
	Less(i, j int) bool
}
