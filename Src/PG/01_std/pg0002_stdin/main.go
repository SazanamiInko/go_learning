// ///////////////////
// /
// / 標準入力
// /
// ////////////////////
package main

import "fmt"

//メイン関数
func main() {

	var ans string
	fmt.Println("あなたは誰")
	fmt.Scan(&ans)
	fmt.Println("僕も")
	fmt.Println(ans)
	fmt.Println("です")
}
