package main
import (
    "fmt"
    "math/rand"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	var j int
	var array []int
	array = append(array, this.nums...)
	for i := 0; i < len(this.nums); i++ {
		j = rand.Intn(len(this.nums)-i) + i
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    obj := Constructor(nums)
    param_3 := obj.Shuffle()
    param_1 := obj.Reset()
    param_2 := obj.Shuffle()
    fmt.Println(param_3, param_1, param_2)
}
