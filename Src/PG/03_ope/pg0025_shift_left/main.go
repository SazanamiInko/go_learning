// ///////////////////
// /
// / 左シフト演算
// /
// ////////////////////
package main

import "fmt"

//メイン関数
func main() {
	var target_mini uint8 = 1
	var target_max uint8 = 255
	var calc_target_mini = target_mini >> 1
	var calc_target_max = target_max >> 1

	fmt.Printf("{%d}を左シフトすると{%d}になります\r\n", target_mini, calc_target_mini)
	fmt.Printf("バイナリ{%b}を左シフトすると{%b}になります\r\n", target_mini, calc_target_mini)
	fmt.Printf("{%d}を左シフトすると{%d}になります\r\n", target_max, calc_target_max)
	fmt.Printf("バイナリ{%b}を左シフトすると{%b}になります\r\n", target_max, calc_target_max)
}
