package main

func main (){
	const (
		a = 1 << iota //  1
		b = 3 << iota // 110
 		c             // 1100
		d			  // 11000
	)
	println(a,b,c,d)

	const (
		e = iota
		f = iota
		g = iota
	)
	println(e,f,g)
}