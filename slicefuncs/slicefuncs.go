package slicefuncs

// FindAll 在slice中搜索所有符合条件（condition函数返回true）的元素.
func FindAll[T interface{}](slice []T, contdition func(T) bool) []T {
	result := []T{}
	for _, item := range slice {
		if contdition(item) {
			result = append(result, item)
		}
	}
	return result
}

// FindFirst 在slice中搜索第一个符合条件（condition函数返回true）的元素，返回元素的索引和元素值，若未找到，返回-1和该型变量初始值
func FindFirst[T interface{}](slice []T, contdition func(T) bool) (int, T) {
	for i := 0; i < len(slice); i++ {
		if contdition(slice[i]) {
			return i, slice[i]
		}
	}
	var t T
	return -1, t
}

// FindLast 在slice中搜索最后一个符合条件（condition函数返回true）的元素，返回元素的索引和元素值，若未找到，返回-1和该型变量初始值
func FindLast[T interface{}](slice []T, contdition func(T) bool) (int, T) {
	for i := len(slice) - 1; i >= 0; i-- {
		if contdition(slice[i]) {
			return i, slice[i]
		}
	}
	var t T
	return -1, t
}

// ForEach 对slice中的每个元素执行action函数，适用于读取slice场景。
func ForEach[T interface{}](slice []T, action func(e *T)) {
	for i := 0; i < len(slice); i++ {
		action(&(slice[i]))
	}
}

// ForEachParal 对slice中的每个元素多线程并行执行action函数，适用于读取slice场景。
// 参数n为并行协程数
func ForEachParal[T interface{}](slice []T, n int, action func(e *T)) {
	if len(slice) <= 0 {
		return
	}
	if n <= 0 {
		n = 1 //若n设置错误，默认单线程执行
	} else {
		if n > len(slice) {
			n = len(slice) //若n小于slice的长度，则取n为slice的长度
		}
	}

	// 初始化token池
	token := make(chan bool, n)
	defer close(token)
	// 执行action，每次执行时先占用token，执行完后释放
	for i := 0; i < len(slice); i++ {
		token <- true
		go func(index int) {
			action(&slice[index])
			<-token
		}(i)
	}
	// 收回token，确保所有协程都已完成
	for i := 0; i < n; i++ {
		token <- true
	}
}

// Map 使用映射函数将 []T1 转换成 []T2。
// 映射函数 f 接受两个类型类型 T1 和 T2。
func Map[T1 interface{}, T2 interface{}](slice []T1, mapper func(T1) T2) []T2 {
	mapped := make([]T2, len(slice))
	for i := 0; i < len(slice); i++ {
		mapped[i] = mapper(slice[i])
	}
	return mapped
}

// MapParal 多线程并行使用映射函数将 []T1 转换成 []T2。
// 映射函数 f 接受两个类型类型 T1 和 T2。
func MapParal[T1 interface{}, T2 interface{}](slice []T1, n int, mapper func(T1) T2) []T2 {
	if n <= 0 {
		n = 1 //若n设置错误，默认单线程执行
	} else {
		if n > len(slice) {
			n = len(slice) //若n小于slice的长度，则取n为slice的长度
		}
	}
	mapped := make([]T2, len(slice))

	// 初始化token池
	token := make(chan bool, n)
	defer close(token)
	// 执行action，每次执行时先占用token，执行完后释放
	for i := 0; i < len(slice); i++ {
		token <- true
		go func(index int) {
			mapped[index] = mapper(slice[index])
			<-token
		}(i)
	}
	// 收回token，确保所有协程都已完成
	for i := 0; i < n; i++ {
		token <- true
	}
	return mapped
}

// Reduce 使用汇总函数将 []T1 切片汇总成一个结果。
func Reduce[T1, T2 interface{}](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for i := 0; i < len(s); i++ {
		r = f(r, s[i])
	}
	return r
}

// Filter 使用过滤函数过滤切片中的数据。
// 该函数返回新的切片，只会保留调用 f 返回 true 的元素。
func Filter[T interface{}](s []T, f func(T) bool) []T {
	var r []T
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			r = append(r, s[i])
		}
	}
	return r
}

// Reverse 反转Slice中元素的顺序。
func Reverse[T interface{}](s []T) []T {
	var r []T
	l := len(s)
	for i := 0; i < l; i++ {
		r = append(r, s[l-1-i])
	}
	return r
}

// GroupCount 统计数组中各元素出现的频次，并将结果记入map
func GroupCount[T interface{}](s []T) map[interface{}]int {
	gmap := make(map[interface{}]int)
	for i := 0; i < len(s); i++ {
		gmap[s[i]]++
	}
	return gmap
}

// Max 取slice中最大的元素，返回元素的索引和元素值，若slice为空，则返回索引为-1；smaller(a,b)表示a是否小于b，小于则返回true。
func Max[T interface{}](slice []T, smaller func(T, T) bool) (index int, max T) {
	if len(slice) == 0 {
		return -1, max
	}
	max = slice[0]
	for i := 1; i < len(slice); i++ {
		if smaller(max, slice[i]) {
			max = slice[i]
			index = i
		}
	}
	return index, max
}

// Min 取slice中最小的元素，返回元素的索引和元素值，若slice为空，则返回索引为-1；smaller(a,b)表示a是否小于b，小于则返回true。
func Min[T interface{}](slice []T, smaller func(T, T) bool) (index int, min T) {
	if len(slice) == 0 {
		return -1, min
	}
	min = slice[0]
	for i := 1; i < len(slice); i++ {
		if smaller(slice[i], min) {
			min = slice[i]
			index = i
		}
	}
	return index, min
}
