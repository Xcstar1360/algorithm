package sort

/**
 * 经典排序算法
 * 分析角度:
 * 1.最好情况、最坏情况、平均情况时间复杂度
 * 2.时间复杂度的系数、常数、低阶(平常评估算法时不考虑, 但实际开发排序量一般较小, 需考虑进去)
 * 3.比较次数和交换(或移动)次数
 * 4.内存消耗 ( 空间复杂度为 O(1) 的排序算法称为 "原地排序算法" )
 * 5.稳定性
 * 以下实现不是重点,重点是各排序算法的思路,具体实现需要取决于场景与数据类型
 */

/**
 * 冒泡排序:
 * 原地排序,是稳定的排序算法
 * 时间复杂度: best-O(n)  worst-O(n^2)  avg-O(n^2)
 */
func bubbleSort(a []int) {
	var n = len(a)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		var flag = false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				flag = true
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
		if !flag {
			break
		}
	}
}

/**
 * 插入排序:
 * 原地排序,是稳定的排序算法
 * 时间复杂度: best-O(n)  worst-O(n^2)  avg-O(n^2)
 */
func insertionSort(a []int) {
	var n = len(a)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		var tmp = a[i]
		var j = i - 1
		for ; j >= 0; j-- {
			if a[j] > tmp {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = tmp
	}
}

/**
 * 选择排序
 * 原地排序,是不稳定的排序算法
 * 时间复杂度: best-O(n^2)  worst-O(n^2)  avg-O(n^2)
 */
func selectionSort(a []int) {
	var n = len(a)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		var min, pos = a[i], i
		for j := i; j < n; j++ {
			if a[j] < min {
				pos = j
				min = a[j]
			}
		}
		a[i], a[pos] = a[pos], a[i]
	}
}

/**
 * 归并排序
 * 稳定性取决于 merge 函数,所以是稳定的排序算法
 * 空间复杂度 O(n),所以不是原地排序算法
 * 时间复杂度: best-O(n*log(n))  worst-O(n*log(n))  avg-O(n*log(n))
 */
func mergeSort(a []int) {
	var n = len(a)
	if n <= 1 {
		return
	}

	mergeSortInternally(a, 0, n-1)
}

func mergeSortInternally(a []int, p int, r int) {
	if p >= r {
		return
	}

	var mid = (p + r) / 2
	mergeSortInternally(a, p, mid)
	mergeSortInternally(a, mid+1, r)
	merge(a, p, mid, r)
}

func merge(a []int, p int, q int, r int) {
	var tmp = make([]int, r-p+1)
	var i, j, k = p, q + 1, 0
	for i <= q && j <= r {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			k++
			i++
		} else {
			tmp[k] = a[j]
			k++
			j++
		}
	}

	var start, end = i, q
	if j <= r {
		start, end = j, r
	}
	for start <= end {
		tmp[k] = a[start]
		k++
		start++
	}

	for s := p; s <= r; s++ {
		a[s] = tmp[s-p]
	}
}

/**
 * 快速排序
 * 是不稳定的排序算法
 * 空间复杂度: 分区函数采用原地分区,不占用额外空间,所以为 O(1),即为原地排序算法
 * 时间复杂度: best-O(n*log(n))  worst-O(n^2)  avg-O(n*log(n))
 */
func quickSort(a []int) {
	var n = len(a)
	if n <= 1 {
		return
	}

	quickSortInternally(a, 0, n-1)
}

func quickSortInternally(a []int, p int, r int) {
	if p >= r {
		return
	}

	var q = partition(a, p, r)
	quickSortInternally(a, p, q-1)
	quickSortInternally(a, q+1, r)
}

// 原地分区,不占用额外空间
func partition(a []int, p int, r int) int {
	//var pivot = a[r] 换一个分界点,速度几何式提升
	var mid = (p + r) / 2
	var pivot = a[mid]
	a[r], a[mid] = a[mid], a[r]

	var i = p
	for j := p; j < r; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	return i
}

/**
 * 桶排序
 * 应用场景:
 * 适用于外部排序(即所排数据不能一次装入内存)
 * 待排数据需要很容易划分范围,每个数据范围就相当于一个桶
 * 每个数据范围内使用快排,各个数据范围内的数据量分布需要比较均匀,否则会退化为 O(n * log(n)) 算法
 */

/**
 * 计数排序(桶排的特殊情况)
 * 是稳定的排序算法,非原地排序
 * 时间复杂度: best-O(n)  worst-O(n)  avg-O(n)
 * 应用场景: 复杂度 O(n) 的排序算法对场景都有限制
 * 只能用在数据范围不大的场景,如果范围 k 比数据量 n 大很多,就不合适
 * 同时注意构造中间数组 c 的时候使用 a[i] 做 c 的下标,所以只能给非负整数排序
 */
func countingSort(a []int) {
	var n = len(a)
	if n <= 1 {
		return
	}

	var max = a[0]
	for i := 1; i < n; i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	var c = make([]int, max+1)
	for i := 0; i < n; i++ {
		c[a[i]]++
	}
	for i := 1; i <= max; i++ {
		c[i] = c[i] + c[i-1]
	}

	var r = make([]int, n)
	// 倒序遍历维持排序稳定性
	for i := n - 1; i >= 0; i-- {
		var pos = c[a[i]]
		r[pos-1] = a[i]
		c[a[i]]--
	}
	for i := 0; i < n; i++ {
		a[i] = r[i]
	}
}

/**
 * 基数排序
 * 对数据按位排,位数不同补一些不影响排序的值,对每位数据使用桶排序或计数排序
 * 假设分为 k 位,那么时间复杂度为 O(k * n),当 k 不大时(手机号 11 位),时间复杂度就近似 O(n)
 * 先排低位再排高位,所以需要使用稳定排序算法
 */
