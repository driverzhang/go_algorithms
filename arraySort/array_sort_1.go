package arraySort

import "go.zhuzi.me/go/log"

/* 冒泡排序 O(n²)
 * 冒泡排序只会操作相邻的两个数据。
 * 每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求。如果不满足就让它俩互换。
 * a是数组，n表示数组大小, 要求从小到大排序: 3,5,4,1,2,6
 */
func BubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ { // 外轮，末尾已经排序数
		// 提前退出标志
		flag := false
		for j := 0; j < n-i-1; j++ { // 内轮，可排序数
			if a[j] > a[j+1] { // 当内轮都循环完毕且都是 a[j] <= a[j+1] 则表示无需再交换数据 flag = false
				a[j], a[j+1] = a[j+1], a[j]
				//此次冒泡有数据交换
				flag = true
			}
		}
		// 如果没有交换数据，提前退出
		if !flag {
			break
		}
	}
}

/* 插入排序 O(n²)
 * 一个有序的数组，要往里面添加一个新数据后，继续保持数据是有序的
 * 将数组中的数据分为两个区间，【已排序区间】和【未排序区间】
 * 插入排序也包含两种操作，一种是【元素的比较】，一种是【元素移动】
 * 插入排序要比冒泡排序更受欢迎：冒泡排序的数据交换要比插入排序的数据移动要复杂，冒泡排序需要 3 个赋值操作，而插入排序只需要 1 个。
 * a是数组，n表示数组大小, 要求从小到大排序: 5,4,6,1,3,2
 */
func InsertionSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		//查找要插入的位置并移动数据
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
				log.Error(a) // 需要打印出来看，不然不好理解过程！
			} else {
				break
			}
		}
		a[j+1] = value
		log.Error(a)

	}
}

/* 选择排序 O(n²)
 * 选择排序算法的实现思路有点类似插入排序，也分【已排序区间】和【未排序区间】。
 * 但是选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾。
 */
func SelectionSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		mixIndex := i // mixIndex 不是表示最小的值，而是最小值的小标位置
		for j := i + 1; j < n; j++ {
			if a[j] < a[mixIndex] {
				mixIndex = j // 这里改变 最小值的下标
			}
		}

		a[i], a[mixIndex] = a[mixIndex], a[i] // 将后面最小值 与 开头的值 交换
	}

}

// 使用性能上从强到弱：【插入】 > 【冒泡】 > 【选择排序】
