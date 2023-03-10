package heap

type Heap2 struct {
	Size int
	// 使用内部的数组来模拟树
	// 一个节点下标为 i，那么父亲节点的下标为 (i-1)/2
	// 一个节点下标为 i，那么左儿子的下标为 2i+1，右儿子下标为 2i+2
	Array []int
}

func NewHeap2(array []int) *Heap2 {
	return &Heap2{Array: array}
}

// Push 上浮操作
func (h *Heap2) Push(x int) {
	// x为堆中插入的第一个元素
	if h.Size == 0 {
		h.Array[0] = x
		h.Size++
		return
	}

	// i 是要插入节点的下标
	i := h.Size
	// 如果下标存在
	// 将小的值 x 一直上浮
	for i > 0 {
		parent := (i - 1) / 2
		// 如果插入的值小于等于父亲节点，那么可以直接退出循环，因为父亲仍然是最大的
		if x <= h.Array[parent] {
			break
		}

		// 否则交换
		h.Array[i] = h.Array[parent]
		i = parent
	}

	// 将该值 x 放在不会再翻转的位置
	h.Array[i] = x

	// 堆数量加一
	h.Size++
}

// Pop 下沉操作
func (h *Heap2) Pop() int {
	if h.Size == 0 {
		return -1
	}

	// 取出根节点
	ret := h.Array[0]

	// 取出最后一个节点，将根节点放最后
	h.Size--
	x := h.Array[h.Size]
	h.Array[h.Size] = ret

	i := 0
	for {
		lc := 2*i + 1 // 左孩子
		rc := 2*i + 2 // 右孩子

		// 判断是否有左子树, 如果没有左子树 说明没有右子树
		if lc >= h.Size {
			break
		}

		// 有右子树, 比较左右子树的值大小 拿到最大值下标
		if rc < h.Size && h.Array[rc] > h.Array[lc] {
			lc = rc
		}

		// 父亲节点的值都大于或等于两个儿子较大的那个，不需要向下继续翻转了，返回
		if x >= h.Array[lc] {
			break
		}

		// 将较大的儿子与父亲交换，维持这个最大堆的特征
		h.Array[i] = h.Array[lc]

		// 继续往下操作
		i = lc
	}

	// 将最后一个元素值放在不再翻转的位置
	h.Array[i] = x

	return ret
}
