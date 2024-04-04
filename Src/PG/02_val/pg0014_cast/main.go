// ///////////////////
// /
// / 型変換
// /
// ////////////////////
package main

//メイン関数
func main() {
	const tax_rate = 0.10
	var price = 1000.00
	var in_tax = int(price*(1.00+tax_rate) + 0.5)
	println("税込み価格は", in_tax)
}
