package main

func main(){
	var a = 20  // 实际变量
	var ip *int  // 指针变量

	ip = &a     // 指针地址

	println(&a,ip,*ip)
}