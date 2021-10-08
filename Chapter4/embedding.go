package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}

type admin struct {
	person user // Not embedding
	level  string
}

type adminEmbed struct {
	user
	level string
}

type adminPointerEmbed struct {
	*user
	level string
}

func main() {
	ad := adminPointerEmbed{
		user: &user{
			name:  "John Smith",
			email: "john@gmail.com",
		},
		level: "super",
	}

	sendNotification(ad)
	ad.user.notify()
	ad.notify()
}
