// ///////////////////
// /
// / 標準入力
// /
// ////////////////////
package main

import (
	"fmt"
)

// メイン関数
func main() {
	var leftOn uint8 = 1
	var leftOff uint8 = 0
	var rightOn uint8 = 1
	var rightOff uint8 = 0

	fmt.Printf("{%t}||{%t}=>{%t}\r\n", leftOn, rightOn, leftOn || rightOn)
	fmt.Printf("{%t}||{%t}=>{%t}\r\n", leftOn, rightOff, leftOn || rightOff)
	fmt.Printf("{%t}||{%t}=>{%t}\r\n", leftOff, rightOn, leftOff || rightOn)
	fmt.Printf("{%t}||{%t}=>{%t}\r\n", leftOff, rightOff, leftOff || rightOff)

}
