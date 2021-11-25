package main

import "fmt"

type notifier interface {
	notify()
}

func sendNotifyMessage(n notifier) {
	n.notify()
}

type Company struct {
	admin
	u user
	level int
}

type admin struct {
	name string
}

func (a admin) notify()  {
	fmt.Printf("admin %s notify\n", a.name)
}

type user struct {
	name string
}

func (u user) notify()  {
	fmt.Printf("user %s notify\n", u.name)
}
func (u *user) changeName(newName string) {
	u.name = newName
}
func main() {
	u := user{name: "li"}
	a := admin{name: "admin"}
	c := Company{
		level: 1,
		u: user{
			"liliya",
		},
		admin:admin{
			name: "aaa",
		},
	}
	sendNotifyMessage(u)
	sendNotifyMessage(a)
	sendNotifyMessage(c.u)
	sendNotifyMessage(c.admin)
}