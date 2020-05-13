package main

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			print(j)
			print("*")
			print(i)
			print("=")
			print(i*j, "   ")
		}
		println("")
	}
}
