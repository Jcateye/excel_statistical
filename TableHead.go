package main

import 	"github.com/deckarep/golang-set"


type TableHead struct {
	Index    int    `json:"index"`
	Name     string `json:"name"`
	ValueSet mapset.Set
}
