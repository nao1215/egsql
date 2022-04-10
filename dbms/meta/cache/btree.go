package cache

import (
	"encoding/json"
	"sort"
)

const (
	// maxDegree is max number of sub trees of a node.
	maxDegree = 3
)

// BTree is a structure for managing nodes in the B-tree
type BTree struct {
	// Top is the top node of the Btree.
	Top *node `json:"top"`
	// Length is the number of data inserted.
	Length int `json:"length"`
}

// node is a Btree node. It has data and child node(s).
type node struct {
	// Items is the data that the node has.
	Items items `json:"items"`
	// Children are the child nodes of the node.
	Children []*node `json:"children"`
}

// items is the data that the node has.
type items []Item

// Int32Item is item of int32
type Int32Item int32

// Int64Item is item of int64
type Int64Item int64

// Item must be comarable for b-tree implementation.
type Item interface {
	// Less returns true if Item less than the argument.
	Less(than Item) bool
}

// NewBTree returns an initialized BTree structure.
func NewBTree() *BTree {
	return &BTree{
		Top:    nil,
		Length: 0,
	}
}

// SerializeBTree serializes a BTree structure to json format.
func SerializeBTree(tree *BTree) ([]byte, error) {
	return json.Marshal(tree)
}

// DeserializeBTree deserializes a BTree structure from json format.
func DeserializeBTree(bytes []byte) (*BTree, error) {
	var tree BTree
	err := json.Unmarshal(bytes, &tree)
	return &tree, err
}

// Insert inserts an item into the BTree.
func (b *BTree) Insert(item Item) {
	b.Length++

	if b.Top == nil {
		b.Top = new(node)
		b.Top.Items.insertAt(0, item)
		return
	}

	b.Top.insert(item)
}

// Find searches for items in the BTree structure.
// If the item is found, it returns true and the index at which the item exists.
func (b *BTree) Find(item Item) (bool, int) {
	if b.Top == nil {
		return false, -1
	}

	return b.Top.find(item)
}

// Get gets an Item from a BTree structure.
func (b *BTree) Get(key Item) Item {
	if b.Top == nil {
		return nil
	}

	return b.Top.get(key)
}

// Len returns the length of the BTree structure.
func (b *BTree) Len() int {
	return b.Length
}

// Less returns true if Int32Item i is less than the Item than passed in the argument.
func (i Int32Item) Less(than Item) bool {
	v, ok := than.(Int32Item)
	if !ok {
		return false
	}
	return i < v
}

// Less returns true if Int64Item i is less than the Item than passed in the argument.
func (i Int64Item) Less(than Item) bool {
	v, ok := than.(Int64Item)
	if !ok {
		return false
	}
	return i < v
}

func (i items) MarshalJSON() ([]byte, error) {
	var intItems []Int32Item
	for _, item := range i {
		intItems = append(intItems, item.(Int32Item))
	}
	return json.Marshal(intItems)
}

func (i *items) UnmarshalJSON(b []byte) error {
	var intItems []Int32Item
	err := json.Unmarshal([]byte(b), &intItems)

	for index, item := range intItems {
		i.insertAt(index, item)
	}

	return err
}

// find returns a bool type indicating whether the item is
// found or not and the index where the item is located.
func (i *items) find(item Item) (bool, int) {
	for index, itm := range *i {
		if !itm.Less(item) {
			if !item.Less(itm) {
				return true, index
			}
			return false, index
		}
	}
	return false, len(*i)
}

// insertAt inserts an item at the specified index.
func (i *items) insertAt(index int, item Item) {
	*i = append(*i, nil)
	if index < len(*i) {
		copy((*i)[index+1:], (*i)[index:])
	}
	(*i)[index] = item
}

// insert inserts an item.
func (n *node) insert(item Item) {
	found, index := n.Items.find(item)
	if found {
		return
	}

	if len(n.Children) == 0 {
		n.Items.insertAt(index, item)

		if len(n.Items) == maxDegree {
			n.splitMe()
		}

		return
	}

	if len(n.Children[index].Items) == maxDegree-1 {
		n.splitChild(index, item)

		if len(n.Items) == maxDegree {
			n.splitMe()
		}

		return
	}

	n.Children[index].insert(item)
}

// deleteChildAt deletes the item at the specified index
func (n *node) deleteChildAt(index int) {
	first := n.Children[:index]
	second := n.Children[index+1:]
	n.Children = append(first, second...)
}

// splitChild splits child nodes.
func (n *node) splitChild(index int, item Item) {
	_, innerIndex := n.Children[index].Items.find(item)
	n.Children[index].Items.insertAt(innerIndex, item)

	leftItem := n.Children[index].Items[maxDegree/2-1]
	midItem := n.Children[index].Items[maxDegree/2]
	rightItem := n.Children[index].Items[maxDegree/2+1]

	n.deleteChildAt(index)

	_, midIndex := n.Items.find(midItem)
	n.Items.insertAt(midIndex, midItem)

	left := new(node)
	left.insert(leftItem)

	right := new(node)
	right.insert(rightItem)

	n.Children = append(n.Children, left)
	n.Children = append(n.Children, right)

	sort.Slice(n.Children, func(i, j int) bool {
		return n.Children[i].Items[0].Less(n.Children[j].Items[0])
	})
}

// splitMe splits the node itself.
func (n *node) splitMe() {
	left := new(node)
	left.Items.insertAt(0, n.Items[maxDegree/2-1])

	right := new(node)
	right.Items.insertAt(0, n.Items[maxDegree/2+1])

	mid := n.Items[maxDegree/2]
	n.Items = append([]Item{}, mid)

	if len(n.Children) == maxDegree+1 {
		var nodes []*node

		left.Children = append(left.Children, n.Children[0])
		left.Children = append(left.Children, n.Children[1])
		nodes = append(nodes, left)

		right.Children = append(right.Children, n.Children[2])
		right.Children = append(right.Children, n.Children[3])
		nodes = append(nodes, right)
		n.Children = nodes
	} else {
		n.Children = append(n.Children, left)
		n.Children = append(n.Children, right)
	}
}

// get returns the item corresponding to the specified key.
func (n *node) get(key Item) Item {
	found, i := n.Items.find(key)
	if found {
		return n.Items[i]
	} else if len(n.Children) > 0 {
		return n.Children[i].get(key)
	}
	return nil
}

// find returns whether the node corresponding to the argument
// item exists or not with bool type. If the node exists, its
// index is returned; if there are no child nodes, -1 is returned.
func (n *node) find(item Item) (bool, int) {
	found, index := n.Items.find(item)
	if found {
		return found, index
	}

	if len(n.Children) == 0 {
		return false, -1
	}

	return n.Children[index].find(item)
}
