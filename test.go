package main
func main(){
	test()
}

func test(){
	println("chujian")
	foo1(12)
}
func foo1(num int){
	if(num > 10){
		print("你输入的数字是:",num)
	}else{
		print("小于10 自动忽略")
	}
}