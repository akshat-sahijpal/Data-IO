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

func CreateData(name, linked string, next *Data, age, dob int, time2 time.Time) Data {
	return Data{
		Name: name,
		Age:  age,
		Dob:  dob,
		Linker: Link{
			Linked: linked,
			Hash:   HashedIn{Next: next},
		},
		CreatedAt: &CreateTime{
			T: time2,
		},
	}
}
func main() {
	a := CreateData("aks", "as", nil, 20, 2001, time.Now())
	b := CreateData("aks32", "as3", &a, 23, 2004, time.Now())
	c := CreateData("de2", "as3d", &b, 23, 2024, time.Now())
	d := CreateData("rks32", "asef3", &c, 323, 2204, time.Now())
	e := CreateData("afrfks32", "aqs3", &d, 233, 2204, time.Now())
	f := CreateData("aksww32", "avs3", &e, 2323, 2404, time.Now())
	g := CreateData("aks", "as", &f, 20, 2001, time.Now())
	h := CreateData("aks32", "as3", &g, 23, 2004, time.Now())
	i := CreateData("de2", "as3d", &h, 23, 2024, time.Now())
	j := CreateData("rks32", "asef3", &i, 323, 2204, time.Now())
	l := CreateData("afrfks32", "aqs3", &j, 233, 2204, time.Now())
	m := CreateData("aksww32", "avs3", &l, 2323, 2404, time.Now())
	marshal, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}
