package main

type user struct {
	name  string
	email string
}

func main() {
	stayOnStack()
	escapeToHeap()
}

func stayOnStack() user {
	u := user{
		name:  "Stuart",
		email: "nicetry@nice.com",
	}

	return u
}

func escapeToHeap() *user {
	u := user{
		name:  "Stuart",
		email: "nicetry@nice.com",
	}

	return &u
}
