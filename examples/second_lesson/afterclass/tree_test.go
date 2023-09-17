package main

import (
	"testing"
)

func TestInsert(t *testing.T) {
	mt := &MultiChildTree{}

	// 插入序列
	values := []int{10, 5, 15, 20, 7}
	for _, v := range values {
		mt.Insert(v)
	}

	// 验证树的结构
	if mt.Root.Value != 10 {
		t.Errorf("Root value should be 10 but got %d", mt.Root.Value)
	}
	if len(mt.Root.Children) != 2 || mt.Root.Children[0].Value != 5 || mt.Root.Children[1].Value != 15 {
		t.Errorf("Root children values should be [5, 15]")
	}
	if len(mt.Root.Children[0].Children) != 1 || mt.Root.Children[0].Children[0].Value != 7 {
		t.Errorf("Child of 5 should have one child with value 7")
	}
	if len(mt.Root.Children[1].Children) != 1 || mt.Root.Children[1].Children[0].Value != 20 {
		t.Errorf("Child of 15 should have one child with value 20")
	}
}
