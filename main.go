package main

import (
	"math"
)

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

type node struct {
	Label       int   //　nodeの値
	Parent      *node // 親nodeへのポインタ
	ChildLeft   *node // 左の子nodeへのポインタ
	ChildRight  *node // 右の子nodeへのポインタ
	heightLeft  int   // 左の部分木の高さ (部分木が空なら-1)
	heightRight int   // 右の部分木の高さ (部分木が空なら-1)
}

type AVL struct {
	root *node // rootへのポインタ
}

func (tree AVL) getCloseNode(x int) node {
	// xに最も近いnodeを探す

	searchNode := node(*tree.root)
	for {
		if (searchNode.Label < x) && (searchNode.ChildLeft != nil) {
			searchNode = *searchNode.ChildLeft
		} else if (searchNode.Label > x) && (searchNode.ChildRight != nil) {
			searchNode = *searchNode.ChildRight
		} else {
			break
		}
	}
	return searchNode
}

func (tree AVL) Search(x int) bool {
	// xに最も近いnodeがxなら存在
	return (tree.getCloseNode(x).Label == x)
}

func (tree AVL) Min() int {
	// intの最小値に最も近いnodeが最小値
	return tree.getCloseNode(math.MinInt).Label
}

func (tree AVL) Max() int {
	// intの最大値に最も近いnodeが最大値
	return tree.getCloseNode(math.MaxInt).Label
}

func (tree AVL) rotateRight(axis node) {
	a := axis
	b := axis.ChildLeft
	a.ChildLeft = b.ChildRight
	a.ChildLeft.Parent = &a
	b.ChildRight = &a
	b.Parent = a.Parent
	a.Parent = b

	a.heightLeft = max(a.ChildLeft.heightLeft, a.ChildLeft.heightRight) + 1
	b.heightRight = max(a.heightLeft, a.heightRight) + 1

}

func (tree AVL) rotateLeft(axis node) {
	a := axis
	b := axis.ChildRight
	a.ChildRight = b.ChildLeft
	a.ChildRight.Parent = &a
	b.ChildLeft = &a
	b.Parent = a.Parent
	a.Parent = b

	a.heightRight = max(a.ChildRight.heightLeft, a.ChildRight.heightRight) + 1
	b.heightLeft = max(a.heightLeft, a.heightRight) + 1

}

func (tree AVL) getFailsNode(startNode node) (failsNode node) {
	foucsNode := startNode
	var updateFlagL, updateFlagR bool
	var newHeight int

	for {
		if foucsNode.Parent == nil {
			break
		}

		// 左の部分木高さの更新
		updateFlagL = false
		if foucsNode.ChildLeft != nil {
			newHeight = max(foucsNode.ChildLeft.heightLeft, foucsNode.ChildLeft.heightRight) + 1
			updateFlagL = newHeight != foucsNode.heightLeft
			foucsNode.heightLeft = newHeight
		}

		// 右の部分木の高さの更新
		updateFlagR = false
		if foucsNode.ChildRight != nil {
			newHeight = max(foucsNode.ChildLeft.heightLeft, foucsNode.ChildLeft.heightRight) + 1
			updateFlagR = newHeight != foucsNode.heightLeft
			foucsNode.heightRight = newHeight
		}

		// どちらも更新されなければ終了
		if !(updateFlagL || updateFlagR) {
			break
		}

		// 条件を満たさない点を発見したら終了
		if abs(foucsNode.heightLeft-foucsNode.heightRight) > 1 {
			failsNode = failsNode
			break
		}
	}
	return
}

func (tree AVL) Insert(x int) bool {
	if tree.root == nil {
		// treeがからの場合、rootを作成
		tree.root = &node{Label: x, heightLeft: -1, heightRight: -1}
		return true
	}

	// xの親となるべきnode
	parentNode := tree.getCloseNode(x)

	// nodeの追加
	if parentNode.Label > x {
		parentNode.ChildLeft = &node{Label: x, Parent: &parentNode, heightLeft: -1, heightRight: -1}
		parentNode.heightLeft = 0
	} else if parentNode.Label < x {
		parentNode.ChildRight = &node{Label: x, Parent: &parentNode, heightLeft: -1, heightRight: -1}
		parentNode.heightRight = 0
	} else {
		// xが既に存在している
		return false
	}

}

func main() {

}
