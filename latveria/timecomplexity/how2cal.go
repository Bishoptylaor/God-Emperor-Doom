package timecomplexity

import (
	"fmt"
	"math"
)

/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2024/7/29 -- 17:34
 @Author  : bishop ❤️ MONEY
 @Description: how to calculate time complexity
*/

// O(1)<O(logN)<O(N)<O(NlogN)<O(N2)<O(2N)<O(N!)

var N = math.MaxInt64

func do() {
	fmt.Println("doing")
}

func O1() {
	// O(1)
	do()

	// O(100) = O(1)
	for i := 0; i < 100; i++ {
		do()
	}
}

func OLogN() {
	// N = 2^m m 为具体高度，每次操作都是二分
	// O(logN)
	i := N
	for i > 0 {
		do()
		i = i / 2
	}
}

func ON() {
	// O(N)
	for i := 0; i < N; i++ {
		do()
	}

	// O(100N) = O(N)
	for i := 0; i < N; i++ {
		for j := 0; j < 100; j++ {
			do()
		}
	}
}

func ONLogN() {
	// O(NlogN)
	i := N
	for i > 0 {
		i = i / 2
		for j := 0; j < N; j++ {
			do()
		}
	}
}

func ON2() {
	// O(N^2)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			do()
		}
	}
}

func O2N() {
	// O(2^N)
	// 每轮执行会派生或者新增两个新的调用：1,2,4,8,16...
	// 这里仅仅是个例子，如果递归调用只调用一次，那么其时间复杂度应为 O(N)
	var Fibonacci func(int) int
	Fibonacci = func(n int) int {
		if n <= 1 {
			return n
		}
		return Fibonacci(n-2) + Fibonacci(n-1)
	}
}

func ON_() {
	// N! 阶乘，
	// N,N-1,N-2,N-3...1
	// 下一次处理的数量是当前处理数量 - 1，本身有乘数关系
	var oo func(n int) int
	oo = func(n int) int {
		if n <= 1 {
			return 1
		}
		return n * oo(n-1)
	}
}
