package slice

import (
	"github.com/nnsgmsone/GaeaDB/curry"
	"github.com/nnsgmsone/GaeaDB/functor"
	"github.com/nnsgmsone/GaeaDB/typeclass"
)

func Nub(xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return xs
	default:
		return append([]typeclass.Ord{xs[0]}, Nub(Filter(curry.NotEq(xs[0]), xs[1:]))...)
	}
}

// 从数组xs中提取满足f要求的元素
func Filter(f func(typeclass.Ord) bool, xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return xs
	case f(xs[0]):
		return append([]typeclass.Ord{xs[0]}, Filter(f, xs[1:])...)
	default:
		return Filter(f, xs[1:])
	}
}

// 判断数组中是否存在某个元素
func Elem(x typeclass.Ord, xs []typeclass.Ord) typeclass.Ord {
	switch {
	case len(xs) == 0:
		return nil
	case xs[0].Eq(x):
		return xs[0]
	default:
		return Elem(x, xs[1:])
	}
}

// 判断数组中是否存在某个元素，存在返回下标，不存在返回-1
func ElemIndex(x typeclass.Ord, xs []typeclass.Ord) int {
	switch {
	case len(xs) == 0:
		return -1
	case xs[0].Eq(x):
		return 0
	default:
		switch n := ElemIndex(x, xs[1:]); n {
		case -1:
			return -1
		default:
			return n + 1
		}
	}
}

// 删除数组中的元素x，不重复删除
func Delete(x typeclass.Ord, xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return xs
	case xs[0].Eq(x):
		return xs[1:]
	default:
		return append([]typeclass.Ord{xs[0]}, Delete(x, xs[1:])...)
	}
}

// 删除数组xs中满足f条件的元素，不重复删除
func DeleteBy(f func(typeclass.Ord) bool, xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return xs
	case f(xs[0]):
		return xs[1:]
	default:
		return append([]typeclass.Ord{xs[0]}, DeleteBy(f, xs[1:])...)
	}
}

func Qsort(xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return xs
	default:
		return append(append(Qsort(Filter(curry.Lt(xs[0]), xs[1:])), xs[0]),
			Qsort(Filter(curry.Gt(xs[0]), xs[1:]))...)
	}
}

// 假设xs是个有序队列(从小到大)，然后insert(x)后满足x0 <= x <= 1
func Push(x typeclass.Ord, xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return []typeclass.Ord{x}
	case xs[0].Ge(x):
		return append([]typeclass.Ord{x}, xs...)
	default:
		return append(xs[:1], Push(x, xs[1:])...)
	}
}

// 将一个列表根据f映射成另一个元素
func Map(f functor.MapFunc, xs []typeclass.Ord) []typeclass.Ord {
	switch {
	case len(xs) == 0:
		return xs
	default:
		return append([]typeclass.Ord{f(xs[0])}, Map(f, xs[1:])...)
	}
}

// 将一个列表根据x和f从左到右折叠成一个新的元素
func Foldl(f functor.FoldFunc, x interface{}, xs []typeclass.Ord) interface{} {
	switch {
	case len(xs) == 0:
		return x
	default:
		return Foldl(f, f(xs[0], x), xs[1:])
	}
}

// 将一个列表根据x和f从右到左折叠成一个新的元素
func Foldr(f functor.FoldFunc, x interface{}, xs []typeclass.Ord) interface{} {
	switch {
	case len(xs) == 0:
		return x
	default:
		return f(xs[0], Foldr(f, x, xs[1:]))
	}
}
