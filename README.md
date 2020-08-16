# minimum-time-visiting-all-points

## 題目解讀：

### 題目來源:
[minimum-time-visiting-all-points](https://leetcode.com/problems/minimum-time-visiting-all-points/)

### 原文:
On a plane there are n points with integer coordinates points[i] = [xi, yi]. Your task is to find the minimum time in seconds to visit all points.

You can move according to the next rules:

In one second always you can either move vertically, horizontally by one unit or diagonally (it means to move one unit vertically and one unit horizontally in one second).
You have to visit the points in the same order as they appear in the array.


### 解讀：

給定一組座標點points 每個座標點point[i] = [xi, yi]

xi, yi 代表知道在x軸的位置 在y軸的位置

求出照順序連結每個point的最短路徑秒數

連接規則如下：

1 必須找次序連 也就是 point[i] 只能連到 point[i+1]

2 可以連接對角 或是直線


## 初步解法:
### 初步觀察:

首先因為照訊序連接 因此可以

計算 point[i+1] - point[i] 的具例 需要offset 多少 x 多少 y

因為是距離 所以其實可以只看絕對值 x, y

從觀察結果上 可以知道 if x > y

則需要移動 [y, y] + [ x-y, 0]  =  y + (x - y) 秒

如果是 y > x 

則需要移動 [x, x] + [0 , y -x] = x + (y - x) 秒

因此可以知道 需要花的秒數就是 max(x,y)秒

### 初步設計:

Given a set of points p

Step 0: let idx = 0, result = 0 

Step 1: if idx >= length of p go to step 4

Step 2:  set result + = max(|x[idx]-x[idx+1]|,|y[idx] - y[idx+1]|)

Step 3:  go to step 1

Step 4: return result
## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package min_time_to_visit_all_points

import "math"

func minTimeToVisitsAllPoints(points [][]int) int {
	result := 0
	upperbound := len(points)
	for idx := 1; idx < upperbound; idx++ {
		result += calculateDistance(points[idx-1], points[idx])
	}
	return result
}
func calculateDistance(p1 []int, p2 []int) int {
	xdiff := int(math.Abs(float64(p1[0] - p2[0])))
	ydiff := int(math.Abs(float64(p1[1] - p2[1])))
	return max(xdiff, ydiff)
}
func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

```
## 測資的撰寫
```golang
package min_time_to_visit_all_points

import "testing"

func Test_minTimeToVisitsAllPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				points: [][]int{{1, 1}, {3, 4}, {-1, 0}},
			},
			want: 7,
		},
		{
			name: "Example2",
			args: args{
				points: [][]int{{3, 2}, {-2, 2}},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToVisitsAllPoints(tt.args.points); got != tt.want {
				t.Errorf("minTimeToVisitsAllPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record
[golang leetcode 30day 19thday](https://hackmd.io/IN7vkzFDTVCa4ReiA4frTw?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)