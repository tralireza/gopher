package gopher

// 307m Range Sum Query - Mutable
type NumArray struct {
	Tree *SNode307
}

type SNode307 struct { // Segment Tree :: Node
	Start, End  int
	Left, Right *SNode307
	Sum         int
}

func Constructor307(nums []int) NumArray {
	var Build func(l, r int) *SNode307
	Build = func(l, r int) *SNode307 {
		if l > r {
			return nil
		}

		n := &SNode307{Start: l, End: r}
		if l == r {
			n.Sum = nums[l]
		} else {
			m := l + (r-l)>>1
			n.Left = Build(l, m)
			n.Right = Build(m+1, r)
			n.Sum = n.Left.Sum + n.Right.Sum
		}
		return n
	}

	return NumArray{Tree: Build(0, len(nums)-1)}
}

func (o *NumArray) Update(index, v int) {
	var Adjust func(*SNode307)
	Adjust = func(n *SNode307) {
		if n.Start == n.End {
			n.Sum = v
			return
		}

		m := n.Start + (n.End-n.Start)>>1
		if index <= m {
			Adjust(n.Left)
		} else {
			Adjust(n.Right)
		}
		n.Sum = n.Left.Sum + n.Right.Sum
	}

	Adjust(o.Tree)
}

func (o *NumArray) SumRange(left, right int) int {
	var Sum func(n *SNode307, l, r int) int
	Sum = func(n *SNode307, l, r int) int {
		if n.Start == l && n.End == r {
			return n.Sum
		}

		m := n.Start + (n.End-n.Start)>>1
		if l >= m+1 {
			return Sum(n.Right, l, r)
		} else if r <= m {
			return Sum(n.Left, l, r)
		}
		return Sum(n.Left, l, m) + Sum(n.Right, m+1, r)
	}

	return Sum(o.Tree, left, right)
}
