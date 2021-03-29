package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var (
	B             BST
	RootNode      *Node
	InsertCounter int = 0
)

// for json unmarshalling
type result struct {
	Appt []Appointment `json:"results"`
}

type Appointment struct {
	Id       int
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
	var r result

	content, err := os.ReadFile("handlers/data/appt.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal(content, &r)
	// fmt.Println(r.Appt[0])
	// func(t **Node) {
	// 	for _, appointment := range r.Appt {
	// 		B.Insert(t, &appointment) // being &Root being 'captured'
	// 	}
	// }(&RootNode)

	for i := 0; i < len(r.Appt); i++ {
		B.Insert(&RootNode, &r.Appt[i])
	}

	// B.Display(RootNode)
	// var dAppt []Appointment
	// B.List(RootNode, &dAppt)

	// fmt.Println(dAppt)
}

func (b *BST) Insert(t **Node, a *Appointment) error {
	// fmt.Println("=====================================")
	// fmt.Printf("Insert: %d\n", InsertCounter)
	// InsertCounter++

	if *t == nil {
		*t = &Node{
			Item: a,
		}

		b.size++
		return nil
	}

	if a.Customer < (*t).Item.Customer {
		b.Insert(&((*t).Left), a)
	} else if a.Customer > (*t).Item.Customer {
		b.Insert(&((*t).Right), a)
	}

	return nil
}

func (b *BST) Search(r **Node, name string) (*Appointment, error) {
	if *r == nil {
		return nil, errors.New("no appointment been made yet")
	}

	if name == (*r).Item.Customer {
		return (*r).Item, nil
	} else if name < (*r).Item.Customer {
		b.Search(&((*r).Left), name)
	} else {
		b.Search(&((*r).Right), name)
	}

	return nil, errors.New("appointment not found")

}

func (b *BST) List(t *Node, l *[]Appointment) {
	if t != nil {
		*l = append(*l, *t.Item)
		b.List(t.Left, l)
		b.List(t.Right, l)
	}
}

func (b *BST) Display(t *Node) {
	if t != nil {
		fmt.Println(t.Item)
		b.Display(t.Left)
		b.Display(t.Right)
	}
}
