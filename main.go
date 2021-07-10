package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type CreateTime struct {
	T time.Time `json:"t"`
}
type Data struct {
	Name      string      `json:"name"`
	Age       int         `json:"age"`
	Dob       int         `json:"dob"`
	Linker    Link        `json:"linker"`
	CreatedAt *CreateTime `json:"created_at"`
}
type Link struct {
	Linked string   `json:"-"`
	Hash   HashedIn `json:"hash"`
}
type HashedIn struct {
	Next *Data `json:"next"`
}

func main() {
	f := Data{
		Name: "ME",
		Age:  233,
		Dob:  03,
		Linker: Link{
			Linked: "CA",
		},
		CreatedAt: &CreateTime{
			T: time.Now(),
		},
	}

	e := Data{
		Name: "SE",
		Age:  23,
		Dob:  03,
		Linker: Link{
			Linked: "CA",
			Hash: HashedIn{
				Next: &f,
			},
		},
	}
	d := Data{
		Name: "Akshat",
		Age:  20,
		Dob:  2001,
		Linker: Link{
			Linked: "JSON",
			Hash: HashedIn{
				Next: &e,
			},
		},

		CreatedAt: &CreateTime{
			T: time.Now(),
		},
	}
	marshal, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}
