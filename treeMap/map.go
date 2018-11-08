package treeMap

import "github.com/nnsgmsone/GaeaDB/rbtree"

func New() TreeMap {
	return &treeMap{rbtree.New()}
}
