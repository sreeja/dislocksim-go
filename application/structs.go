package main

// import (
// 	"encoding/json"
// )

type Exectime struct {
	Name string `json:"name"`
	Time int    `json:"time"`
}

type OpLock struct {
	Op    string `json:"op"`
	Locks []Tok  `json:"locks"`
}

type Tok struct {
	Name string `json:"name"`
	Mode string `json:"mode"`
}

type LockType struct {
	Name      string `json:"name"`
	Param     string `json:"param"`
	Placement string `json:"placement"`
}

type Lock struct {
	Name string
	Mode string
	Type LockType
}
