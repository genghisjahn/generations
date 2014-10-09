package main

import (
	"math/rand"
	"time"
)

type Allele struct {
	Pos1 bool
	Pos2 bool
}

func (a *Allele) Select() bool {
	rand.Seed(time.Now().UnixNano())
	pick := rand.Float64()
	if pick < .5 {
		return a.Pos1
	}
	return a.Pos2
}

func GenerateAllele() Allele {
	rand.Seed(time.Now().UnixNano())
	allele := Allele{Pos1: false, Pos2: false}
	pick := rand.Float64()
	if pick > .5 {
		allele.Pos1 = true
	}
	pick = rand.Float64()
	if pick > .5 {
		allele.Pos2 = true
	}
	return allele
}
