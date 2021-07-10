package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	marshal, err := json.Marshal(bulkData())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	fmt.Println(GetJsonReads())
}

// GetJsonReads Reading Data
func GetJsonReads() Data {
	file, err := os.OpenFile("out.json", os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		panic("Error Reading Json!!!!")
	}
	// Convert *File to bytes using ioutil
	all, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		panic("Error &")
	}
	v := Data{}
	err1 := json.Unmarshal(all, &v)
	if err1 != nil {
		panic("Error")
	}
	return v
}
func bulkData() Data {
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
	a1 := CreateData("aks", "as", &m, 20, 2001, time.Now())
	b1 := CreateData("aks32", "as3", &a1, 23, 2004, time.Now())
	c1 := CreateData("de2", "as3d", &b1, 23, 2024, time.Now())
	d1 := CreateData("rks32", "asef3", &c1, 323, 2204, time.Now())
	e1 := CreateData("afrfks32", "aqs3", &d1, 233, 2204, time.Now())
	f1 := CreateData("aksww32", "avs3", &e1, 2323, 2404, time.Now())
	g1 := CreateData("aks", "as", &f1, 20, 2001, time.Now())
	h1 := CreateData("aks32", "as3", &g1, 23, 2004, time.Now())
	i1 := CreateData("de2", "as3d", &h1, 23, 2024, time.Now())
	j1 := CreateData("rks32", "asef3", &i1, 323, 2204, time.Now())
	l1 := CreateData("afrfks32", "aqs3", &j1, 233, 2204, time.Now())
	m1 := CreateData("aksww32", "avs3", &l1, 2323, 2404, time.Now())
	a11 := CreateData("aks", "as", &m1, 20, 2001, time.Now())
	b11 := CreateData("aks32", "as3", &a11, 23, 2004, time.Now())
	c11 := CreateData("de2", "as3d", &b11, 23, 2024, time.Now())
	d11 := CreateData("rks32", "asef3", &c11, 323, 2204, time.Now())
	e11 := CreateData("afrfks32", "aqs3", &d11, 233, 2204, time.Now())
	f11 := CreateData("aksww32", "avs3", &e11, 2323, 2404, time.Now())
	g11 := CreateData("aks", "as", &f11, 20, 2001, time.Now())
	h11 := CreateData("aks32", "as3", &g11, 23, 2004, time.Now())
	i11 := CreateData("de2", "as3d", &h11, 23, 2024, time.Now())
	j11 := CreateData("rks32", "asef3", &i11, 323, 2204, time.Now())
	l11 := CreateData("afrfks32", "aqs3", &j11, 233, 2204, time.Now())
	m11 := CreateData("aksww32", "avs3", &l11, 2323, 2404, time.Now())
	return m11
}
