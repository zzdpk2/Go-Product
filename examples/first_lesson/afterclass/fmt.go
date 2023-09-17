package main

import "fmt"

func main() {
	// rune本质是int32，不是字符
	// var r1 rune = 'A'
	// var r2 rune = '😃'
	byteData := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}
	fmt.Println(printBytes(byteData))
	floatNum := 3.7890
	fmt.Println(printNumWith2(floatNum))

}

// 输出两位小数
func printNumWith2(float642 float64) string {
	return fmt.Sprintf("%.2f", float642)
}

func printBytes(data []byte) string {
	return fmt.Sprintf("%x", data)
}
