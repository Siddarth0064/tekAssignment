package sumavg

import "fmt"

func SumAvg(arr []int) {
	var sum int // this contains sum of the elements from slice
	var avg int // this contains avg of the elements from slice
	len := len(arr)
	for _, v := range arr {
		sum = sum + v
	}
	avg = sum / len
	fmt.Println("the sum of the numbers of slice is :", sum)
	fmt.Println("the average of the numbers of slice  is :", avg)

}
