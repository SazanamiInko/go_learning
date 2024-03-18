// ///////////////////
// /
// / 文字列型変数
// /
// ////////////////////
package main

import "fmt"

//メイン関数
func main() {

	var ans string
	var res string = "オウム"
	fmt.Println("あなたは誰")
	fmt.Scan(&ans)
	fmt.Print("僕は")
	fmt.Print(res)
	fmt.Print("です")
}
