// This file is part of random-json
// Copyright (C) 2022 Nandakumar Edamana
// Under the terms of the GNU General Public License version 3.

package generator

import (
	"encoding/json"
	"log"
	"math/rand"
	
	"github.com/google/uuid"
)

const (
	Object = iota
	Array
	String
	Number
	Null
	TYPE_GUARD
)

type RandomJsonGenerator struct {
	maxDepth int
	maxElems int
}

func NewRandomJsonGenerator() RandomJsonGenerator {
	return RandomJsonGenerator{
		4,
		8,
	}
}

func (this *RandomJsonGenerator) Generate() interface{} {
	return this.GenerateWithMaxDepth(this.maxDepth)
}

func (this *RandomJsonGenerator) GenerateWithMaxDepth(maxDepth int) interface{} {
	// Make sure objects have the highest chance and arrays have the next

	rnd := rand.Intn(100)
	
	typ := Object
	if rnd < 25 {
		typ = rand.Intn(TYPE_GUARD)
	} else if rnd < 50 {
		typ = rand.Intn(TYPE_GUARD)
	}
	
	// ------------------------------------------------------------------

	switch typ {
	case Object:
		return this.GenerateObject(maxDepth)
	case Array:
		return this.GenerateArray(maxDepth)
	case String:
		return this.GenerateString()
	case Number:
		return rand.Float64()
	case Null:
		return nil
	default:
		panic("unhandled type")
	}
}

func (this *RandomJsonGenerator) GenerateAsString() string {
	jobj := this.Generate()

	str, err := json.Marshal(jobj)
	if err != nil {
		log.Fatalf("error marshalling json: %s", err.Error())
	}
	
	return string(str)
}

func (this *RandomJsonGenerator) GenerateArray(maxDepth int) interface{} {
	arr := make([]interface{}, 0)

	if maxDepth > 0 {
		for i := 0; i < rand.Intn(this.maxElems); i++ {
			arr = append(arr, this.GenerateWithMaxDepth(maxDepth - 1))
		}
	}

	return arr
}

func (this *RandomJsonGenerator) GenerateObject(maxDepth int) interface{} {
	obj := make(map[string]interface{})

	if maxDepth > 0 {
		for i := 0; i < rand.Intn(this.maxElems); i++ {
			obj[uuid.NewString()] = this.GenerateWithMaxDepth(maxDepth - 1)
		}
	}

	return obj
}

func (this *RandomJsonGenerator) GenerateString() string {
	s := uuid.NewString()

	if rand.Intn(100) < 50 {
		s = s + " " + uuid.NewString() + " ☺"
	}

	return s
}
