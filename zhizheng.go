package main

func main(){
	var a = 20  // 实际变量
	var ip *int  // 指针变量

	ip = &a     // 指针地址
	println(&a,ip,*ip)
	foo()
}

func fun(){
	const all int = 9
}

func foo (){
	const name string = "chujian"
	print(name)
}
