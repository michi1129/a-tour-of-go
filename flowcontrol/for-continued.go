package main

import "fmt"

func main() {
	sum := 1
	// フォーマッタがセミコロンを削除してしまうため
	// 実質的に for-is-gos-white.go と同じ
	// for ; sum < 1000;  {
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
