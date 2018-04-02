package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	user
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending email to admin: %s<%s>\n", a.name, a.email)
}

func (u *user) notify() {
	fmt.Printf("Sending email to user: %s<%s>\n", u.name, u.email)
}

func main() {
	u := user{name: "ahmad tabbakha", email: "tabba.ahmad@gmail.com"}
	ad := admin{
		user:  user{name: "ahmad leet", email: "ahmad1337@admin.com"},
		level: "super",
	}
	sendNotification(&u)
	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
