// ///////////////////
// /
// / 小数点桁数指定表示
// /
// ////////////////////
package main

import (
	"fmt"
	"math"
)

// メイン関数
func main() {
	var lastYearNum float64
	var nownum float64

	lastYearNum = 356
	nownum = 121
	var rate = (lastYearNum - nownum) / lastYearNum * 100 * 1000

	rate = math.Round(rate) / 1000
	fmt.Print("前年に比べて")
	fmt.Printf("%.3f", rate)
	fmt.Print("%減です")
}
