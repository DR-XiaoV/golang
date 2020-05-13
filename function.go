package main

//func main(){
//	var a = 123
//	var b = 456
//	var num = max(a,b)
//	print(num)
//}
//func max(num1,num2 int) int{
//	if(num1 > num2){
//		return num1
//	}else{
//		return num2
//	}
//}

func main() {
	a, b := demo("google", "microsoft")
	print(a, b)
}
func demo(x, y string) (string, string) {
	return y, x
}
