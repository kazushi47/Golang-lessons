package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	/***** 変数 *****/
	var num1 int = 123
	// 型が明白の時は型名を省略可
	var num2 = 123
	// := を用いると var も省略可
	num3 := 123
	// まとめて定義
	var (
		num4_1 = 123
		num4_2 = 456
	)
	name, age := "Yamada", 26
	// 上記たちを確認
	fmt.Println("***** 変数 *****")
	fmt.Println("num1:", num1, ", num2:", num2, ", num3:", num3, ", num4_1:", num4_1, ", num4_2:", num4_2, ", name:", name, ", age", age)

	/***** 定数 ******/
	// 型名は省略可で多くの場合指定しない
	const foo1 = 100
	const (
		foo2 = 100
		foo3 = 200
	)
	// 上記たちを確認
	fmt.Println("***** 定数 *****")
	fmt.Println("foo1:", foo1, ", foo2:", foo2, ", foo3:", foo3)

	/***** 型 *****/
	// 別名を付ける
	type UtcTime string
	type JstTime string
	var t1 UtcTime = "00:00:00"
	var t2 JstTime = "00:00:00"
	// 上記は型が異なるので t1 = t2 のような代入はエラーになる
	// 型名()で型変換
	var a1 uint16 = 1234
	var a2 uint32 = uint32(a1)
	// 上記たちを確認
	fmt.Println("***** 型 *****")
	fmt.Println("t1:", t1, ", t2:", t2, ", a1:", a1, ", a2:", a2)
	// NULL値
	fmt.Println("NULL値:", nil)

	/***** 演算子 *****/
	x := 7
	y := 4
	sum := x + y
	sub := x - y
	mul := x * y
	div := x / y
	mod := x % y
	x++
	y--
	x += y
	// 上記たちを確認
	fmt.Println("***** 演算子 *****")
	fmt.Println("x:", x, ", y:", y, ", sum:", sum, " ,sub:", sub, ", mul:", mul, ", div:", div, ", mod:", mod)

	/***** 配列 *****/
	array1 := [3]string{}
	array1[0] = "Red"
	array1[1] = "Green"
	array1[2] = "Blue"
	// 初期化によって個数が決まる場合は...と省略可
	array2 := [...]string{"Red", "Green", "Blue"}
	// 上記たちを確認
	fmt.Println("***** 配列 *****")
	fmt.Println(array1)
	fmt.Println(array2)

	/***** スライス *****/
	// 個数を変更可能なものをスライスと呼ぶ
	slice1 := []string{}
	// 要素を追加
	slice1 = append(slice1, "Red")
	slice1 = append(slice1, "Green")
	slice1 = append(slice1, "Blue")
	// 上記たちを確認
	fmt.Println("***** スライス *****")
	fmt.Println(slice1)
	//len()で長さを取得できる
	fmt.Println("slice1の長さ=", len(slice1))

	/***** マップ *****/
	map1 := map[string]int{
		"x": 100,
		"y": 200,
	}
	// 要素を追加
	map1["z"] = 300
	map1["p"] = 123
	// 要素を削除
	delete(map1, "x")
	// 上記たちを確認
	fmt.Println("***** マップ *****")
	fmt.Println(map1)

	/***** if文 *****/
	fmt.Println("***** if文 *****")
	if 3 > 2 {
		fmt.Println("OK")
	} else if 2 > 2 {
		fmt.Println("Good")
	} else {
		fmt.Println("NG")
	}

	/***** switch文 *****/
	fmt.Println("***** switch文 *****")
	switch 3 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Any")
	}

	/***** for文 *****/
	// Goにはwhile文が無く、繰り返しは全てforを用いる
	fmt.Println("***** for文 *****")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 配列やスライスなどにはrangeを用いる
	colors := [...]string{"Red", "Green", "Blue"}
	for i, color := range colors {
		fmt.Println(i+1, "番目は", color, "です。")
	}

	/***** 関数(func) *****/
	fmt.Println("***** 関数(func) *****")
	// add関数を下部で作成しておく
	fmt.Println(add(5, 3))
	// 複数の値を返却することもできる
	rtn1, rtn2 := addMinus(5, 3)
	fmt.Println("rtn1:", rtn1, ", rtn2:", rtn2)
	// 不要な返却にはブランク変数 _ を使用する
	_, rtn3 := addMinus(4, 5)
	fmt.Println("rtn3:", rtn3)
	// 可変引数のある関数
	fmt.Println(getSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	/***** 構造体(struct) *****/
	// クラスの代わりに用いられる
	fmt.Println("***** 構造体 *****")
	var p1 Person
	p1.SetPerson("Yamada", 26, "Y")
	p1Name, p1Age := p1.GetPerson()
	fmt.Println(p1Name, p1Age)
	// 構造体を使用した初期化
	p2 := Person{"Yamada", 26, "Y"}
	fmt.Println(p2)

	/***** 遅延実行 *****/
	/*
		関数から戻る直前に処理を遅延実行する。
		そのため、リソースを忘れずに解放できる
	*/
	fmt.Println("***** 遅延実行 *****")
	f, err := os.Open("D:\\GolangProjects\\Lesson01\\sample.txt")
	if err != nil {
		fmt.Println("i/o error")
	}
	// 遅延実行(defer)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	fmt.Println(string(b))

	/***** ゴルーチン *****/
	fmt.Println("***** ゴルーチン *****")
	// 高速な並行処理を実現
	go funcA()
	for i := 0; i < 10; i++ {
		fmt.Println("M")
		time.Sleep(20 * time.Millisecond)
	}
}

/***** 関数(func) *****/
func add(x int, y int) int {
	return x + y
}

/***** 関数(func) *****/
func addMinus(x int, y int) (int, int) {
	return x + y, x - y
}

/***** 関数(func) *****/
func getSum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

/***** 構造体(struct) *****/

type Person struct {
	// 構造体にはメンバ変数のみ定義
	name string
	age  int
	// 大文字で始まるものはパッケージ外からアクセス可能
	Nickname string
}

/***** 構造体(struct) *****/
func (p *Person) SetPerson(name string, age int, Nickname string) {
	// p = thisに相当する変数
	p.name = name
	p.age = age
	p.Nickname = Nickname
}

/***** 構造体(struct) *****/
func (p *Person) GetPerson() (string, int) {
	return p.name, p.age
}

/***** ゴルーチン *****/
func funcA() {
	for i := 0; i < 10; i++ {
		fmt.Println("A")
		time.Sleep(10 * time.Millisecond)
	}
}
