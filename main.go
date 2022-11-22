package main

import (
	"fmt"
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

func (tree *AVL) getCloseNode(x int) *node {
	// xに最も近いnodeを探す

	// rootからの探索
	searchNode := (*tree).root

	for {

		// 着目してるnodeのLabelよりxが大きい(小さい)かつ右(左)部分木が存在するなら探索を続ける
		if (searchNode.Label < x) && (searchNode.ChildRight != nil) {
			searchNode = searchNode.ChildRight
		} else if (searchNode.Label > x) && (searchNode.ChildLeft != nil) {
			searchNode = searchNode.ChildLeft
		} else {
			// labelとxが一致するか葉に到達したら終了
			break
		}
	}
	return searchNode
}

func (tree *AVL) Search(x int) bool {
	// xに最も近いnodeがxなら存在
	return (tree.getCloseNode(x).Label == x)
}

func (tree *AVL) Min() int {
	// intの最小値に最も近いnodeが最小値
	return tree.getCloseNode(math.MinInt).Label
}

func (tree *AVL) Max() int {
	// intの最大値に最も近いnodeが最大値
	return tree.getCloseNode(math.MaxInt).Label
}

func (tree *AVL) rotateRight(axis *node) {
	// aを起点に時計回りに回転
	a := axis
	b := axis.ChildLeft
	a.ChildLeft = b.ChildRight
	if a.ChildLeft != nil {
		a.ChildLeft.Parent = a
	}
	b.ChildRight = a
	b.Parent = a.Parent
	a.Parent = b

	if a.ChildLeft == nil {
		a.heightLeft = -1
	} else {
		a.heightLeft = max(a.ChildLeft.heightLeft, a.ChildLeft.heightRight) + 1
	}
	b.heightRight = max(a.heightLeft, a.heightRight) + 1

	if b.Parent != nil {
		if b.Parent.ChildLeft == a {
			b.Parent.ChildLeft = b
			b.Parent.heightLeft = max(b.heightLeft, b.heightRight) + 1
		} else {
			b.Parent.ChildRight = b
			b.Parent.heightRight = max(b.heightLeft, b.heightRight) + 1
		}
	}
}

func (tree *AVL) rotateLeft(axis *node) {
	// aを起点に反時計回りに回転
	a := axis
	b := axis.ChildRight
	a.ChildRight = b.ChildLeft
	if a.ChildRight != nil {
		a.ChildRight.Parent = a
	}
	b.ChildLeft = a
	b.Parent = a.Parent
	a.Parent = b

	if a.ChildRight == nil {
		a.heightRight = -1
	} else {
		a.heightRight = max(a.ChildRight.heightLeft, a.ChildRight.heightRight) + 1
	}
	b.heightLeft = max(a.heightLeft, a.heightRight) + 1

	if b.Parent != nil {
		if b.Parent.ChildLeft == a {
			b.Parent.ChildLeft = b
			b.Parent.heightLeft = max(b.heightLeft, b.heightRight) + 1
		} else {
			b.Parent.ChildRight = b
			b.Parent.heightRight = max(b.heightLeft, b.heightRight) + 1
		}
	}
}

func (tree *AVL) getFailsNode(startNode *node) (failsNode *node, isFound bool) {
	// 部分木の高さを修正しながらAVL木の条件に満たさなくなったnodeを取得する
	foucsNode := startNode
	var updateFlagL, updateFlagR bool // 部分木の高さが更新されたかどうかのフラグ
	var newHeight int                 // 部分木の新しい高さ

	isFound = false
	for {
		// 左の部分木高さの更新
		updateFlagL = false
		if foucsNode.ChildLeft != nil {
			newHeight = max(foucsNode.ChildLeft.heightLeft, foucsNode.ChildLeft.heightRight) + 1
			updateFlagL = newHeight != foucsNode.heightLeft
			foucsNode.heightLeft = newHeight
		} else {
			updateFlagL = foucsNode.heightLeft != -1
			foucsNode.heightLeft = -1
		}

		// 右の部分木の高さの更新
		updateFlagR = false
		if foucsNode.ChildRight != nil {
			newHeight = max(foucsNode.ChildRight.heightLeft, foucsNode.ChildRight.heightRight) + 1
			updateFlagR = newHeight != foucsNode.heightLeft
			foucsNode.heightRight = newHeight
		} else {
			updateFlagR = foucsNode.heightRight != -1
			foucsNode.heightRight = -1
		}

		// どちらも更新されなければ終了
		if !(updateFlagL || updateFlagR) {
			break
		}

		// 条件を満たさない点を発見したら終了
		if abs(foucsNode.heightLeft-foucsNode.heightRight) > 1 {
			failsNode = foucsNode
			isFound = true
			break
		}

		// rootにたどり着いたら終了
		if foucsNode.Parent == nil {
			break
		}

		// 親をたどっていき条件を満たさない点を探索する
		foucsNode = foucsNode.Parent
	}
	return
}

func (tree *AVL) solveTree(failsNode *node) {

	// 条件に合わなくなった点がrootであるかどうか
	rootFlag := failsNode.Parent == nil

	var failsChild *node // failsNodeのうち高い方の部分木
	if failsNode.heightLeft > failsNode.heightRight {
		failsChild = failsNode.ChildLeft
		if failsChild.heightLeft >= failsChild.heightRight {
			// 一重回転
			tree.rotateRight(failsNode)
		} else {
			// 二重回転
			tree.rotateLeft(failsChild)
			tree.rotateRight(failsNode)
		}
	} else {
		failsChild = failsNode.ChildRight
		if failsChild.heightRight >= failsChild.heightLeft {
			// 一重回転
			tree.rotateLeft(failsNode)
		} else {
			// 二重回転
			tree.rotateRight(failsChild)
			tree.rotateLeft(failsNode)
		}
	}

	if rootFlag {
		// rootの付け替え
		tree.root = failsNode.Parent
	}
}

func (tree *AVL) Insert(x int) bool {
	if tree.root == nil {
		// treeがからの場合、rootを作成
		tree.root = &node{Label: x, heightLeft: -1, heightRight: -1}
		return true
	}

	// xの親となるべきnode
	parentNode := tree.getCloseNode(x)

	// nodeの追加
	if parentNode.Label > x {
		parentNode.ChildLeft = &node{Label: x, Parent: parentNode, heightLeft: -1, heightRight: -1}
	} else if parentNode.Label < x {
		parentNode.ChildRight = &node{Label: x, Parent: parentNode, heightLeft: -1, heightRight: -1}
	} else {
		// xが既に存在している
		return false
	}

	// 追加したnodeの親を起点にAVL木の条件を解決する
	// 条件を満たさなくなった点の取得
	failsNode, isFound := tree.getFailsNode(parentNode)
	if !isFound {
		return true
	}

	// 条件の解決
	tree.solveTree(failsNode)
	return true
}

func (tree *AVL) Delete(x int) bool {

	// 削除対象のnode
	deleteNode := tree.getCloseNode(x)

	// 削除対象が存在しない場合false
	if deleteNode.Label != x {
		return false
	}

	// タスクをnodeの削除から葉の削除に変える
	var removeNode *node
	if deleteNode.ChildLeft != nil {
		// 削除対象が左部分木を持っている時
		// 削除位置のnodeのラベルを部分木の最大値に置き換えてその最大値(葉)を削除対象にする
		leftTree := &AVL{root: deleteNode.ChildLeft}
		removeNode = leftTree.getCloseNode(x)
		deleteNode.Label = removeNode.Label

	} else if deleteNode.ChildRight != nil {
		// 削除対象が右部分木を持っている時
		// 削除対象のnodeのラベルを部分木の最小値に置き換えてその最小値(葉)を削除対象にする
		rightTree := &AVL{root: deleteNode.ChildRight}
		removeNode = rightTree.getCloseNode(x)
		deleteNode.Label = removeNode.Label
	} else {
		// 削除対象が部分木を持たない時(= 削除対象が葉)
		removeNode = deleteNode
		if removeNode.Parent == nil {
			tree.root = nil
			return true
		}
	}

	// 削除対象の親のnodeからの参照を捨てる
	parentNode := removeNode.Parent
	if parentNode.ChildLeft == removeNode {
		parentNode.ChildLeft = nil
	} else {
		parentNode.ChildRight = nil
	}

	// 削除対象の親を起点にAVL木の条件を解決する
	// 条件を満たさない点の取得
	failsNode, isFound := tree.getFailsNode(parentNode)
	if !isFound {
		return true
	}

	// 条件の解決
	tree.solveTree(failsNode)
	return true
}

func InitTree(x ...int) *AVL {
	tree := &AVL{}
	for _, xi := range x {
		tree.Insert(xi)
	}
	return tree
}

func main() {
	tree := InitTree(2, 1, 5, 4, 6, 3)
	tree.Delete(4)
	fmt.Println(tree.Max())
}
