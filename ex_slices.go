package main
import (
	"fmt"
)
// Create a slice of integers with values [10, 20, 30, 40, 50].

// Print the slice using fmt.Println.

// Access and print the first and last elements of the slice.

func getSlice() {
	mySlice := []int{10, 20, 30, 40, 50}
	fmt.Println(mySlice)
	fmt.Println(mySlice[0])
	fmt.Println(mySlice[len(mySlice)-1])
}

// Start with a slice of integers: []int{1, 2, 3}.

// Use the append function to add the numbers 4, 5, and 6.

// Print the updated slice.

func appendToSlice() {
	mySlice := []int{1, 2, 3}
	mySlice = append(mySlice, 4, 5, 6)
	fmt.Println(mySlice)
}

// Create an array of strings: [5]string{"a", "b", "c", "d", "e"}.

// Create a slice from the array that includes elements from index 1 to 3.

// Print the slice and its length.

func sliceFromArray() []string {
	myArray := [5]string{"a", "b", "c", "d", "e"}
	mySlice := myArray[1:3]
	fmt.Println(len(mySlice))
	fmt.Println(mySlice)
	return mySlice

}

// Define a slice of integers: []int{10, 20, 30, 40, 50}.

// Create a new slice that includes elements from index 2 to 4.

// Modify the new slice and observe how the original slice changes.

func sliceFromSlice() []int {
	mySlice := []int{10, 20, 30, 40, 50}
	newSlice := mySlice[2:4]
	fmt.Println(newSlice)
	newSlice = append(newSlice, 3)
	fmt.Println(mySlice)
	fmt.Println(newSlice)

	return newSlice
}

// Use make to create a slice of float64 with length 3 and capacity 5.

// Assign values to the first three elements.

// Print the slice, its length, and its capacity.

func makeSlice() []float64 {
	mySlice := make([]float64, 3, 5)
	mySlice = append(mySlice, 4.7, 7, 9.0)
	fmt.Println(mySlice)
	return mySlice
}

// Create a slice of integers: []int{1, 2, 3, 4, 5}.

// Create another slice of the same length using make.

// Use the copy function to copy elements from the first slice to the second slice.

// Print both slices to verify.

func copySlice() []int {
	mySlice := []int{1, 2, 3, 4, 5}
	sliceTwo := make([]int, 5, 5)
	copy(sliceTwo, mySlice)
	fmt.Println(mySlice)
	fmt.Println(sliceTwo)
	return sliceTwo
}

// Create a slice of strings: []string{"Go", "Python", "Java", "C++"}.

// Use a for loop with range to iterate over the slice and print each element.

// Print the index and value for each iteration.

func iterateSlice() []string {
	mySlice := []string{"Go", "Python", "Java", "C++"}
	for _, lang := range(mySlice) {
		fmt.Println(lang)
	}
	return mySlice
}

// Start with an empty slice of integers: []int{}.

// Append numbers 1 to 10 using a loop.

// Print the resulting slice.

func loopSlice() []int {
	mySlice := []int{}
	for num := 1; num <= 10; num++ {
		mySlice = append(mySlice, num)
	}
	fmt.Println(mySlice)
	return mySlice
}

// Check Capacity Growth

// Create a slice of integers with an initial capacity of 2.

// Use a loop to append numbers 1 to 10.

// Print the slice, its length, and capacity after each append operation.

func checkCapacity() {
	mySlice := make([]int, 0, 2)
	for num := 1; num <= 10; num++ {
		mySlice = append(mySlice, num)
		fmt.Printf("Slice %v, capacity %v, len %v \n", mySlice, cap(mySlice), len(mySlice))
	}
}

// Create a multi-dimensional slice (a slice of slices):
// Iterate over the multi-dimensional slice and print each inner slice. Print the individual elements within the inner slices.

func iterateMultidimensional() {
	matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
	for _, row := range matrix {
		fmt.Println(row)
		for _, elem := range row {
			fmt.Println(elem)
		}
		
	}
}

func ex_slices() {
	fmt.Println("hey")
	getSlice()
	appendToSlice()
	sliceFromArray()
	sliceFromSlice()
	makeSlice()
	copySlice()
	iterateSlice()
	loopSlice()
	checkCapacity()
	iterateMultidimensional()

}