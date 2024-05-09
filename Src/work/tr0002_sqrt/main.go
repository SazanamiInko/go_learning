// ///////////////////
// /
// / 平方根の近似値
// /
// ////////////////////
package main

import (
	"fmt"
	"math"
)

// メイン関数
func main() {
	var x = 37.0
	var sqrt_ans = math.Sqrt(x)
	var near = nearSqrt(x)
	fmt.Printf("%.0fの平方根は%.2fです", x, sqrt_ans)
	fmt.Println("")
	fmt.Printf("近似値は%.2fです", near)
}

// 平方根の近似値を求める関数
func nearSqrt(root float64) float64 {
	var para = math.Floor(math.Sqrt(root))
	return para + (root-para*para)/(2*para)
}
