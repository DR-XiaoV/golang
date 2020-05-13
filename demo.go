package main

func main() {
	//print("hello world")
	//const name string = "韩梅梅"
	//print(name)
	//const a,b,c,d = 1,2,3,4
	//print(a,b,c,d)
	//const width,height = 100,100
	//var area int
	//area = width * height
	//println(area)

	const (
		a = iota
		b
		c
		d = 19
		e
		f = iota
		g
		h = "123"
		i = iota
	)
	const (
		v = iota
	)
	println(a,b,c,d,e,f,g,h,i,v)
}
