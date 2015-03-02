package main

import (
	. "github.com/jmervine/GoT"
	"testing"
)

func TestConditions_AreMet(T *testing.T) {
	// from: pushevent_test.go
	push := UnmarshalJson()

	good := Conditions{}
	Go(T).Assert(good.AreMet(push))

	good = Conditions{Branches: []string{"gh-pages"}}
	Go(T).Assert(good.AreMet(push))

	g1 := good
	g1.Owner = true
	Go(T).Assert(g1.AreMet(push))

	g2 := good
	g2.Admin = true
	Go(T).Assert(g2.AreMet(push))

	// bad
	b1 := good
	b1.Branches = []string{"master"}
	Go(T).Refute(b1.AreMet(push))

	b2 := good
	b2.Master = true
	Go(T).Refute(b2.AreMet(push))
}
