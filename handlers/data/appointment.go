package data

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var (
	B             BST
	RootNode      *Node
	InsertCounter int                      = 0
	adminResult   map[string][]Appointment // for Admin
)

type Appointment struct {
	Id       string
	Customer string
	Doctor   string
	Time     string
	Location string
}

// A Node represents any single element
type Node struct {
	Item  *Appointment
	Right *Node
	Left  *Node
}

// BST keeps track of the whole Binary Search Tree
type BST struct {
	Root *Node
	size int
}

func init() {

	content, err := os.ReadFile("handlers/data/appt.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal(content, &adminResult)
	for _, appointments := range adminResult {
		for i := 0; i < len(appointments); i++ {
			B.Insert(&RootNode, &appointments[i], "")
		}
	}
}

func (b *BST) Insert(t **Node, a *Appointment, name string) error {

	if *t == nil {
		*t = &Node{
			Item: a,
		}

		if name != "" {

			adminResult[name] = append(adminResult[name], *a)

		}

		b.size++
		return nil
	}
	if a.Customer <= (*t).Item.Customer {
		b.Insert(&((*t).Left), a, name)
	} else if a.Customer > (*t).Item.Customer {
		b.Insert(&((*t).Right), a, name)
	}

	return nil
}

func (b *BST) Search(r **Node, name string, a **Appointment) {

	if *r == nil {
		return
	}

	if strings.Compare(name, (*r).Item.Customer) == 0 {
		*a = (*r).Item
	} else if name < (*r).Item.Customer {
		b.Search(&((*r).Left), name, a)
	} else {
		b.Search(&((*r).Right), name, a)
	}
}

// ListAll list all appointments made by all users (Admin Only)
func (b *BST) ListAll(t *Node, l *[]Appointment) {
	if t != nil {
		*l = append(*l, *t.Item)
		b.ListAll(t.Left, l)
		b.ListAll(t.Right, l)
	}
}

// ListOne list only 1 user's appointments
func ListOne(cName string) []Appointment {
	return adminResult[cName]
}

func (b *BST) Display(t *Node, a *[]Appointment) {
	if t != nil {
		*a = append(*a, *t.Item)
		b.Display(t.Left, a)
		b.Display(t.Right, a)
	}
}
