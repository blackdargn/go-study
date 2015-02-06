package study

type Sortable interface {
	// 大于比较，true：>; 否则<=
	Great(other interface{}) bool
}
/////////////////////////////////////////////////////////////////////////
func QuicksortT(data []Sortable, ascend bool) {
	quicksortT(data, 0, len(data)-1, ascend)
}

func quicksortT(data []Sortable, low int, high int, ascend bool) {
	if (low >= high) {
		return
	}
	pivot := data[low]
	pl, ph := low, high
	for pl < ph {
		for pl < ph && data[ph].Great(pivot) == ascend {
			ph = ph-1
		}
		data[pl] = data[ph]
		for pl < ph && !data[pl].Great(pivot) == ascend {
			pl = pl+1
		}
		data[ph] = data[pl]
	}
	data[pl] = pivot
	quicksortT(data, low, pl-1, ascend)
	quicksortT(data, pl+1, high, ascend)
}

func Quicksort(data []int, ascend bool) {
	quicksort(data, 0, len(data)-1, ascend)
}

func quicksort(data []int, low int, high int, ascend bool) {
	if (low >= high) {
		return
	}
	pivot := data[low]
	pl, ph := low, high
	for pl < ph {
		for pl < ph && data[ph]>pivot == ascend {
			ph = ph-1
		}
		data[pl] = data[ph]
		for pl < ph && data[pl]<=pivot == ascend {
			pl = pl+1
		}
		data[ph] = data[pl]
	}
	data[pl] = pivot
	quicksort(data, low, pl-1, ascend)
	quicksort(data, pl+1, high, ascend)
}
/////////////////////////////////////////////////////////////////////////
/**
  * 将节点i进行正确的下沉调整,左右子叶子为 2*i + 1, 2*i +2
  * @param data
  * @param i 调整的节点i
  * @param n 长度为n
  */
func minHeapFixDown(data []int, i, n int) {

	temp := data[i]
	// 左叶子
	j := 2 * i + 1
	for j < n {
		// 从左右叶子中找出最小的节点
		if j+1 < n && data[j+1] < data[j] {
			j++
		}
		// 无效调整
		if data[j] > temp {
			break
		}
		data[i] = data[j]
		i = j
		j = 2*i+1
	}
	data[i] = temp
}

/**
 * 删除最小堆的第一个节点
 * @param data
 * @param n
 */
func minHeapDel(data []int, n int) {
	temp := data[0]
	data[0] = data[n-1]
	data[n-1] = temp
	minHeapFixDown(data, 0, n-1)
}

/**
 * 将节点i进行正确的上升调整，父节点为 (i-1)/2
 * @param data
 * @param i
 */
func minHeapFixUp(data []int, i int) {
	temp := data[i]
	// 父节点
	j := (i - 1) / 2
	for j >= 0 && i != 0 {
		if data[j] <= temp {
			break
		}
		data[i] = data[j]
		i = j
		j = (i-1)/2
	}
	data[i] = temp
}

/**
 * 最小堆添加节点i
 * @param data
 * @param i
 * @param num
 */
func minHeapAdd(data []int, i, num int) {
	data[i] = num
	minHeapFixUp(data, i)
}

/**
 * 初始化最小堆
 * @param data
 */
func minHeapMake(data []int) {
	size := len(data)
	for i := size / 2 - 1 ; i >= 0 ; i-- {
		minHeapFixDown(data, i, size)
	}
}

/**
 * 最小堆排序
 * @param data
 */
func MinHeapSort(data []int) {
	minHeapMake(data)
	for n := len(data); n > 1 ; n-- {
		minHeapDel(data, n)
	}
}
/////////////////////////////////////////////////////////////////////////
/**
 * 归并排序，递归合并有序子序列
 * @param data
 */
func MergeSort(data []int) {
	mergeSort(data, 0, len(data)-1, make([]int, len(data)))
}

func mergeSort(data []int, first, last int, sorted []int) {
	if first < last {
		// 取中间位
		mid := (first + last) / 2
		// 排序前段
		mergeSort(data, first, mid, sorted)
		// 排序后段
		mergeSort(data, mid+1, last, sorted)
		// 合并前后段
		merge(data, first, mid, last, sorted)
	}
}

func merge(data []int, first, mid, last int, sorted []int) {
	// 前段指针
	i := first
	// 后端指针
	j := mid + 1
	// 有序指针
	k := 0

	// 找出前段 与 后段中有序元素，直到一个完
	for i <= mid && j <= last {
		if data[i] < data[j] {
			sorted[k] = data[i]
			k++
			i++
		}else {
			sorted[k] = data[j]
			k++
			j++
		}
	}
	// 前段续接
	for i <= mid {
		sorted[k] = data[i]
		k++
		i++
	}
	// 后段续接
	for j <= last {
		sorted[k] = data[j]
		k++
		j++
	}
	// 更新有序序列
	for n := 0; n < k; n++ {
		data[first+n] = sorted[n]
	}
}
/////////////////////////////////////////////////////////////////////////
