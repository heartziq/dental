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

// for json unmarshalling (for Customers)
// type result struct {
// 	Appt []Appointment `json:"results"`
// }

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

// bst
type BST struct {
	Root *Node
	size int
}

func init() {
	// var r result

	content, err := os.ReadFile("handlers/data/appt.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	// per user basis - only Admin will use BST algo
	// json.Unmarshal(content, &r)
	json.Unmarshal(content, &adminResult)
	// fmt.Println(adminResult["Edward"])
	// func(t **Node) {
	// 	for _, appointment := range r.Appt {
	// 		B.Insert(t, &appointment) // being &Root being 'captured'
	// 	}
	// }(&RootNode)
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

	// fmt.Printf("newName: %s\n", name)
	// fmt.Printf("Item.Customer: %s\n", (*r).Item.Customer)

	// fmt.Printf("newName == Item.Customer: %v\n", name == (*r).Item.Customer)
	// fmt.Printf("newName == Item.Customer: %v\n", strings.Compare(name, (*r).Item.Customer))

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

func (b *BST) Display(t *Node) {
	if t != nil {
		fmt.Println(t.Item)
		b.Display(t.Left)
		b.Display(t.Right)
	}
}
