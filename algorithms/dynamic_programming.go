package algorithms

import "math"

//锯钢条问题
//长度	1	2	3	4	5	6	7	8	9	10
//价格	1	5	8	9	10	17	17	20	24	30
//BenchmarkCutRod-8         112447              9048 ns/op
//BenchmarkCutRodTopToBottom-8     1639650               689 ns/op
//BenchmarkCutRodBottomToTop-8     1957590               598 ns/op

//递归算法
func CutRod(p[]int,n int)int{
	if n==0{
		return 0
	}
	q:=-int(^uint32(0)>>1)
	for i:=1;i<=n;i++{
		q= int(math.Max(float64(CutRod(p,n-i)+p[i-1]),float64(q)))
	}
	return q
}

//自上而下的动态规划,需要带备忘
func CutRodTopToBottom(p[]int,n int)int{
	if n==0{
		return 0
	}
	r:=make([]int,n+1)
	for i:=0;i<n;i++{
		r[i]=-1
	}
	return memorizedTopToBottom(p,r,n)
}

func memorizedTopToBottom(p,r[]int,n int)int{
	//如果r中存了最大值，就不用走下面的递归了
	if r[n]>0{
		return r[n]
	}
	var q int
	if n==0{
		q=0
	}else {
		q=-int(^uint32(0)>>1)
		for i:=1;i<=n;i++{
			q= int(math.Max(float64(memorizedTopToBottom(p,r,n-i)+p[i-1]),float64(q)))
		}
	}
	r[n]=q
	return q
}

//自底向上
func CutRodBottomToTop(p[]int,n int)int{
	r:=make([]int,n+1)
	r[0]=0
	for i:=1;i<=n;i++{
		q:=-1
		for j:=1;j<=i;j++{
			q=int(math.Max(float64(q),float64(r[i-j]+p[j-1])))
		}
		r[i]=q
	}
	return r[n]
}