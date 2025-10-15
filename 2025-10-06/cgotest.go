package main

//// #cgo CFLAGS: -I/opt/homebrew/include
////#cgo LDFLAGS: -L/opt/homebrew/lib -lmpfi -lmpfr -lgmp
//
//#include <stdio.h>
//int f(int x) {
//    printf("Hello, World! (%d)\n", x);
//    return x + 1;
//}
//int g(int y) {
//    printf("Hello, World! (%d)\n", y);
//    return y + 1;
//}
//void myPrint3(const char* format, int a, int b) {
//    printf(format, a, b);
//}
//
import "C"
import "sync"

var wg sync.WaitGroup

func func1(x int, c *chan int) {
	defer wg.Done()
	t := C.f(C.int(x))
	*c <- int(t)
}
func func2(y int, c *chan int) {
	defer wg.Done()
	t := C.g(C.int(y))
	*c <- int(t)
	// C.printf("Hello, World! (%d)\n", t)
	C.myPrint3(C.CString("Result: (%d, %d)\n"), C.int(y), C.int(t))
}

func main() {
	wg.Add(2)
	r1 := make(chan int, 1)
	r2 := make(chan int, 1)
	go func1(3, &r1)
	go func2(4, &r2)
	wg.Wait()
	x := <-r1
	y := <-r2
	println(x + y)
}
