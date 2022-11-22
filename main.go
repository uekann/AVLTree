package main

import (
	"fmt"
)

type node struct {
	Label       int
	ChildLeft   *node
	ChildRight  *node
	heightLeft  int
	heightRight int
}

type AVL node

func (tree AVL) Search(x int) bool {
	searchNode := node(tree)
	for{
		if searchNode.Label == x{
			return true
		} else if (searchNode.Label < x) && (searchNode.ChildLeft != nil){
			searchNode = *searchNode.ChildLeft
		} else if (searchNode.Label > x) && (searchNode.ChildRight != nil){
			searchNode = *searchNode.ChildRight
		} else {
			return false
		}
	}
}

func (tree AVL) Min() int {
	searchNode := node(tree)
	for{
		if searchNode.ChildLeft == nil{
			return searchNode.Label
		} else {
			searchNode = *searchNode.ChildLeft
		}
	}
}

func (tree AVL) Max() int {
	searchNode := node(tree)
	for {
		if searchNode.ChildRight == nil{
			return searchNode.Label
		} else {
			searchNode = *searchNode.ChildRight
		}
	}
}

func (tree AVL) Insert(x int) bool {
	if tree.Search(x) {
		return false
	}

	
}