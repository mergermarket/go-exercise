package misterPointy

func add(a, b int) int{
	return a + b
}

func addPointers(a, b *int) int {
	return *a + *b
}

func addReturnPointer(a, b int) *int{
	sum := a + b
	return &sum
}

func addPointersReturnPointer(a, b *int) *int{
	sum:= *a + *b
	return &sum
}

func addModifyArgument(a int, b *int) {
	sum:=a + *b
	*b = sum
}

func addPointerToPointer(a, b **int) int{
	sum := **a + **b
	return sum
}
