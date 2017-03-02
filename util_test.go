package main

import "testing"

func TestName(t *testing.T) {
	name := GetName()
	t.Log(name)
}
