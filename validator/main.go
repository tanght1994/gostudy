package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	required()
	omitempty()
	structonly()
	dive()
}

func required() {
	// required 表示, 字段不能是golang"零值"
	type T1 struct {
		A int            `validate:"required"` // 不能是 0
		B string         `validate:"required"` // 不能是 ""
		C []int          `validate:"required"` // 不能是 nil
		D map[string]int `validate:"required"` // 不能是 nil
	}

	t1 := T1{A: 0, B: "", C: nil, D: nil}
	t2 := T1{A: 0, B: "", C: []int{}, D: map[string]int{}}
	t3 := T1{A: 1, B: "a", C: []int{}, D: map[string]int{}}
	validate := validator.New()

	err := validate.Struct(t1)
	fmt.Println(err) // A B C D 四个字段都不合法

	err = validate.Struct(t2)
	fmt.Println(err) // A B 不合法

	err = validate.Struct(t3)
	fmt.Println(err) // 无err
}

func omitempty() {
	// omitempty 如果字段为零值, 则不进行此字段的后续校验(如果用了required, 则omitempty就没有意义了)
	type T1 struct {
		A int `validate:"omitempty,gt=5"`
	}

	t1 := T1{A: 0} // A为零值, 且有omitempty, 所以不进行gt=5校验
	t2 := T1{A: 1} // A不是零值, 进行gt=5校验, 所以校验失败

	validate := validator.New()
	err := validate.Struct(t1)
	fmt.Println(err)

	err = validate.Struct(t2)
	fmt.Println(err)
}

func structonly() {
	// structonly修饰的嵌套字段, 不会进入到字段内部执行校验
	type T1 struct {
		A string `validate:"required"`
		B int    `validate:"required,gt=5"`
	}

	type T2 struct {
		C string `validate:"required"`
		D T1     `validate:"required"`
	}

	type T3 struct {
		C string `validate:"required"`
		D T1     `validate:"required,structonly"` // structonly表示, 只要D存在即可, 不要进入到T1这个结构体的内部进行校验
	}

	type T4 struct {
		C  string                           `validate:"required"`
		T1 `validate:"required,structonly"` // structonly对匿名字段同样生效
	}

	t1 := T1{A: "a", B: 3}
	t2 := T2{C: "c", D: t1}
	t3 := T3{C: "c", D: t1}
	t4 := T4{C: "c", T1: t1}

	validate := validator.New()

	err := validate.Struct(t2)
	fmt.Println(err) // err非空, 原因是 t2.D.B 字段不满足 gt=5

	err = validate.Struct(t3)
	fmt.Println(err) // 无err, 原因是 t3.D 字段为 structonly, 不会进一步对 t3.D 字段进行深入的校验, 即便 t3.D.B 字段不满足 gt=5

	err = validate.Struct(t4)
	fmt.Println(err) // 无err, 原因同上
}

func dive() {
	// dive 下潜一层 只用于数组

	// [][]string `validate:"gte=1"`  于这个二维数组来说, 我们的validate只对这个二维数组的最外维生效
	// 也就是说我们只限制了这个二维数组的第1维是 gte=1 (对于数组来书就是指长度>=1)
	// data := [][]string{[]string{}, []string{}} 这个data是合法的, 尽管它内部的各个元素都是[]string{}(长度为0的数组)
	// 我们如何限制这个二维数组的内部元素呢? 用dive下潜一层就行了
	// [][]string `validate:"gte=1,dive,gte=1"`  第一个gte=1限制第一维, 第二个gte=1限制第二维
	// data := [][]string{[]string{}, []string{}} 这个data就不合法了, 因为[]string{}被第二个 gte=1 限制了, 长度为0, 不满足gte=1
	// 但是 data := [][]string{[]string{"", "", ""}} 这个data还是合法的, 尽管它的第3层元素为空字符串("零值")
	// 我们可以继续下潜(可以随意下潜, 下潜多深都OK)
	// [][]string `validate:"gte=1,dive,gte=1,dive,required"`
	// 我们的 required 限制了第三层的元素不能为 "零值"
	// 所以 data := [][]string{[]string{"", "", ""}} 这个data就不合法了, 因为第三层元素为""
	// data := [][]string{[]string{"abc", "def"}} 这个data是合法的

	// 例一: [][]string `validate:"gt=0,dive,len=1,dive,required"`
	// gt=0 用于 []
	// len=1 用于 []string
	// required 用于 string

	// 例二: [][]string `validate:"gt=0,dive,dive,required"`
	// gt=0 用于 []
	// []string 没有规则限制
	// required 用于 string

	type T1 struct {
		A [][]string `validate:"gte=1,dive,gte=2,dive,required"`
	}

	t1 := T1{A: [][]string{}}
	t2 := T1{A: [][]string{[]string{"abc"}}}
	t3 := T1{A: [][]string{[]string{"a", "", "a"}}}
	t4 := T1{A: [][]string{[]string{"a", "a"}}}

	validate := validator.New()
	err := validate.Struct(t1)
	fmt.Println(err) // [][]string 的长度为0, 不满足gte=1

	err = validate.Struct(t2)
	fmt.Println(err) // [][]string 的长度为1, 满足gte=1, []string 长度为1, 不满足gte=2

	err = validate.Struct(t3)
	fmt.Println(err) // [][]string 的长度为1, 满足gte=1, []string 长度为3, 满足gte=2, ["a", "", "a"] 存在零值, 不满足required

	err = validate.Struct(t4)
	fmt.Println(err) // OK
}
