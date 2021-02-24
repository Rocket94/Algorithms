package algorithms

import (
	"Algorithms/util"
	"math"
)

//动态规划，如果一个问题包含其子问题的最优解，我们称它具有最优子结构性质
//两个最优解子问题必须是无关的才行，资源不能共享
//分治每一步都产生新问题 ，而动态规划是反复求解相同问题，把每次重叠问题存入内存即可

//锯钢条问题
//长度	1	2	3	4	5	6	7	8	9	10
//价格	1	5	8	9	10	17	17	20	24	30
//BenchmarkCutRod-8         112447              9048 ns/op
//BenchmarkCutRodTopToBottom-8     1639650               689 ns/op
//BenchmarkCutRodBottomToTop-8     1957590               598 ns/op

//递归算法
func CutRod(p []int, n int) int {
	if n == 0 {
		return 0
	}
	q := util.INT_MIN
	for i := 1; i <= n; i++ {
		//i代表前一段长度，p[i-1]即该长度的价格，递归子式代表后一段价格
		q = util.CompareMax(CutRod(p, n-i)+p[i-1], q)
	}
	return q
}

//自上而下的动态规划,需要带备忘
func CutRodTopToBottom(p []int, n int) int {
	if n == 0 {
		return 0
	}
	//创建备忘录，索引代表该长度下最大价格
	r := make([]int, n+1)
	for i := 0; i < n; i++ {
		r[i] = -1	//给r初始化，都是负数
	}
	return memorizedTopToBottom(p, r, n)
}

func memorizedTopToBottom(p, r []int, n int) int {
	//如果r中存了最大值，就不用走下面的递归了
	if r[n] > 0 {
		return r[n]		//如果r值是正数，说明已经求过该条件下最优解了，就不用继续递归直接返回就好
	}
	var q int
	if n == 0 {
		q = 0		//终止条件n=0无法在进行切割
	} else {
		q = util.INT_MIN
		for i := 1; i <= n; i++ {
			q = util.CompareMax(memorizedTopToBottom(p, r, n-i)+p[i-1], q)
		}
	}
	//记录该长度最优解
	r[n] = q
	return q
}

//自底向上
func CutRodBottomToTop(p []int, n int) int {
	//自底向上的话更为简单，直接从小到大记录，最优解为前一个最优解加上这个
	r := make([]int, n+1)
	r[0] = 0
	//代表长度从1~n的钢条
	for i := 1; i <= n; i++ {
		q := -1
		//分离出一截和另一个子问题，记录所有里面最大值为q
		for j := 1; j <= i; j++ {
			q = util.CompareMax(q, r[i-j]+p[j-1])
		}
		r[i] = q
	}
	return r[n]
}


func LCSLength(x, y []string) int {
	var c = make([][]int, len(x))
	for i := 0; i < len(x); i++ {
		c[i] = make([]int, len(y))
	}

	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y); j++ {
			if x[i] == y[j] {
				if i == 0 || j == 0 {
					c[i][j] = 0 + 1
				} else {
					c[i][j] = c[i-1][j-1] + 1
				}
			} else {
				m := i - 1
				n := j - 1
				if m < 0 {
					m = 0
				}
				if n < 0 {
					n = 0
				}
				c[i][j] = int(math.Max(float64(c[m][j]), float64(c[i][n])))
			}
		}
	}
	return c[len(x)-1][len(y)-1]
}
