package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type LightningComputer interface {
	LightningConnection()
}

type USBComputer interface {
	USBConnection()
}

type MacComputer struct {
}

func (c MacComputer) LightningConnection() {
	fmt.Println("Mac lightning connected")
}

type Adapter struct {
	C LightningComputer
}

func (a *Adapter) USBConnection() {
	fmt.Println("USB connection to lightning")
	a.C.LightningConnection()
}

func XmlToJSON(x []byte) []byte {
	v := B{}
	xml.Unmarshal(x, &v)
	d, _ := json.Marshal(v.X)
	return d
}

type B struct {
	XMLName xml.Name `xml:"Bstruct"`
	X       []string `xml:"a"`
}

func main() {
	a := MacComputer{}
	b := Adapter{C: a}
	b.USBConnection()
}
