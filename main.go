package main

import (
	//"log"
	//"net"
	"fmt"
)

type Item struct {
	title string
	body  string
}

var database []Item

func GetByName(Title string) Item {

	var getItem Item

	for _, val := range database {

		if val.title == Title {

			getItem = val

		}

	}
	return getItem

}

func CreateItem(item Item) Item {

	database = append(database, item)

	return item

}

func EditItem(Title string, edit Item) Item {

	var change Item

	for idx, val := range database {

		if val.title == Title {

			database[idx] = edit
			change = edit

		}
	}
	return change
}

func DeleteItem(item Item) Item {

	var del Item

	for idx, val := range database {

		if val.title == item.title && val.body == item.body {

			database = append(database[:idx], database[idx+1:]...)
			break
		}

	}

	return del
}

func main() {

	fmt.Println("Initial datbase:  ", database)

	a := Item{"first", "Divyam Gupta"}
	b := Item{"Second", "Divyansh Gupta"}
	c := Item{"Third", " Ramesh Gupta"}
	d := Item{"Fourth", "Sumit Gupta"}

	CreateItem(a)
	CreateItem(b)
	CreateItem(c)
	fmt.Println("Second Database", database)

	DeleteItem(d)
	fmt.Println("Fourth Database", database)

	EditItem("Third", Item{"third", "Prakash Gupta"})
	fmt.Println("Third Database", database)

	x := GetByName("first")
	y := GetByName("Second")

	fmt.Println(x, y)

}
