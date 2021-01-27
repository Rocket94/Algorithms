package leetcode

//func lengthOfLongestSubstring(s string) int {
//
//}

func Convert(s string, numRows int) string {
	if numRows==1{
		return s
	}
	var data []byte=[]byte(s)
	var strArr [][]byte=make([][]byte,numRows)
	for i:=0;i<numRows;i++{
		strArr[i]=make([]byte, len(s))
	}
	var cnt int
	var col int
	strArr[0][0]=data[0]
	cnt++
	for cnt< len(data){
		for i:=1;i<numRows&&cnt<len(data);i++{
			strArr[i][col]=data[cnt]
			cnt++
		}
		for j:=numRows-2;j>=0&&cnt<len(data);j--{
			col++
			strArr[j][col]=data[cnt]
			cnt++
		}
	}
	var cnt2 int
	for i:=0;i<len(strArr);i++{
		for j:=0;j<len(strArr[0]);j++{
			if	strArr[i][j]!=0{
				data[cnt2]=strArr[i][j]
				cnt2++
			}
		}
	}
	return string(data)
}
