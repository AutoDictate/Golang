package pointers

func SimplePointer(a *int, b *int)  (add int, mul int){

	*a = *a + 10
	*b = *b * *a

	add = *a
	mul = *b

	return

}