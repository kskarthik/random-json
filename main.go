// This file is part of random-json
// Copyright (C) 2022 Nandakumar Edamana
// Under Apache License, Version 2.0

package main

import (
	"fmt"
	"math/rand"
	"random-json/generator"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	gen := generator.NewRandomJsonGenerator()

	fmt.Println(gen.GenerateAsString())
}
