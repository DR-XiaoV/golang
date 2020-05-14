package main

func main() {

	var arr [10]int
	var a,b int

	for a = 0; a < 10; a++{
		arr[a] = 100+a
	}

	for b = 0; b < 10; b++{
		println(arr[b])
	}
}