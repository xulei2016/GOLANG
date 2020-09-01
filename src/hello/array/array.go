package array

import "fmt"



//声明初始化1
func f1(){
	var a [3]int8			//零值初始化
	var b = [4]int8{1,2,3}	//赋值初始化，未赋值的为类型零值
	var c = [4]string{"北京","上海","广州","深圳"}	//完成初始化

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}


//声明初始化2 执行推断数组长度
func f2(){
	var a = [...]int8{1,2}
	var b = [...]string{"北京", "上海", "深圳"}
	fmt.Println(a)
	fmt.Println(b)
}

//使用指定索引值赋值初始化
func f3(){
	var a [3]int8
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)
}

//同上
func f4(){
	var a = [3]int8{0:2,1:2,2:4}
	var b = [...]string{0:"北京",1:"上海",2:"天津"}
	fmt.Println(a)
	fmt.Println(b)
}

//遍历
func loop(){
	var a = [...]string{"北京", "上海", "深圳"}

	//for
	for i := 0; i < len(a); i++{
		fmt.Printf("数组a,下标%d 的值是%s\n",i,a[i])
	}

	//range
	for i,v := range a{
		fmt.Printf("下标%d 的值为%s\n",i,v)
	}
}

//多维数组
//多维数组遍历


func Pratice (){
	//f1()
	//f2()
	//f3()
	//f4()
	//loop()

	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]  // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4]  // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}
