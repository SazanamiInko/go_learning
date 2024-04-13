// ///////////////////
// /
// / Or演算
// /
// ////////////////////
package main

import "fmt"

//メイン関数
func main() {
	var leftOn bool = true
	var leftOff bool = false
	var rightOn bool = true
	var rightOff bool = false

	fmt.Printf("{%t}||{%t}=>{%t}\r\n", leftOn, rightOn, leftOn || rightOn)
	fmt.Printf("{%t}||{%t}=>{%t}\r\n", leftOn, rightOff, leftOn || rightOff)
	fmt.Printf("{%t}||{%t}=>{%t}\r\n", leftOff, rightOn, leftOff || rightOn)
	fmt.Printf("{%t}||{%t}=>{%t}\r\n", leftOff, rightOff, leftOff || rightOff)

}
