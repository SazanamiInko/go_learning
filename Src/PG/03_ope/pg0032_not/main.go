// ///////////////////
// /
// / Not演算
// /
// ////////////////////
package main

import "fmt"

//メイン関数
func main() {
	var On bool = true
	var Off bool = false

	fmt.Printf("{%t}=>{%t}\r\n", On, !On)
	fmt.Printf("{%t}=>{%t}\r\n", Off, !Off)
}
