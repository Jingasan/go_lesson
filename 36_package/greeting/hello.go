// パッケージ名はディレクトリと同じものにする必要がある
package greeting

import "fmt"

// 他ディレクトリ（パッケージ）から呼び出される関数は大文字から始まる必要がある
func Hello() {
	fmt.Println("Hello")
}
