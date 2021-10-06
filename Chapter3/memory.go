package memory

import "fmt"

func main() {
	five := [5]string{"Annie", "Betty", "Charley", "Doug", "Bill"}

	for i, v := range five {
		fmt.Printf("Value[%s]\tAddress[%p]\tIndexAddr[%p]\n", v, &v, &five[i])
	}
}
