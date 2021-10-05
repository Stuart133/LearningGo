package main

import "fmt"

func main() {
	tripleNotation()
}

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}

func tripleNotation() {
	slice1 := []string{"A", "B", "C", "D", "E"}
	slice2 := slice1[2:4:4] // The third value controls the slice capacity. As it ends up being 2, a new backing array is created when we try and append

	slice2 = append(slice2, "CHANGED")

	inspectSlice(slice1)
	inspectSlice(slice2)
}

func slicingTest() {
	slice1 := []string{"A", "B", "C", "D", "E"}
	slice2 := slice1[2:4]

	slice2[0] = "CHANGED"

	inspectSlice(slice1)
	inspectSlice(slice2)
}

func overwriteSliceTest() {
	slice := make([]string, 5, 8)
	slice[0] = "Apple"
	slice[1] = "Orange"
	slice[2] = "Banana"
	slice[3] = "Grape"
	slice[4] = "Plum"

	inspectSlice(slice)

	slice2 := slice[0:6]
	slice2[5] = "Test"

	inspectSlice(slice2)

	slice = append(slice, "NewValue")

	inspectSlice(slice)
	inspectSlice(slice2)
}
