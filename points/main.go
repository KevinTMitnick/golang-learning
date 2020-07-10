package main

import "fmt"

// 指针就是地址
// &value	对变量去地址
// *ptr		对地址取值

func main1()  {
	a := 100

	var p *int = &a

	a = 101
	fmt.Println("a=",a )
	fmt.Println("p=",p)

	*p = 32000
	fmt.Println(a)
	fmt.Println("*p=", *p)

	a = 44000
	fmt.Println(a)
	fmt.Println("*p=", *p)
}

func main2()  {
	 var a int = 123

	 fmt.Printf("a type is %T\n", a)
	 fmt.Printf("a value is %v\n", a)
	 fmt.Printf("a addres is %p\n", &a)

	 aPointer := &a

	 fmt.Printf("aPointer type is: %T\n ", aPointer)
	 fmt.Printf("aPointer value is: %v\n", aPointer)
	 fmt.Printf("aPointer address is: %p\n", &aPointer)

	 fmt.Println(*aPointer)		// 取aPointer地址所存的值


}

func main()  {
	// 定义指针，指向内存空间
	var p *int		// nil 指向内存编号为 0x0空间 , 不允许用户进行读写操作的

	//*p = 123		// 直接写数据，报错

	p = new(int)	// new(数据类型)
	*p = 123		// 写数据
	fmt.Println(*p)

}
