// ///////////////////
// /
// / 改行なし標準出力
// /
// ////////////////////
package main

import "fmt"

//メイン関数
func main() {

	var ans string
	fmt.Println("あなたは誰")
	fmt.Scan(&ans)
	fmt.Print("僕も")
	fmt.Print(ans)
	fmt.Print("です")
}
