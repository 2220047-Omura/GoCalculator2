package main

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/zserge/lorca"
)

var ui lorca.UI
var l, N, v2 int
var str string
var contFib bool

type fib struct {
	number      int
	level       int
	N           int
	calc        bool
	append      bool
	left, right *fib
}

type lv struct {
	level  int
	number []int
	next   *lv
}

func (f *fib) fibTree(n int) *fib {
	N += 1
	if f == nil {
		if n != 1 && n != 0 {
			f = &fib{number: n, level: l, N: N}
		} else {
			f = &fib{number: 1, level: l, N: N, calc: true}
		}
		if n > 1 {
			l = f.level + 1
			f.left = f.left.fibTree(n - 1)
			l = f.level + 1
			f.right = f.right.fibTree(n - 2)
		}
	}
	return f
}

func (f *fib) nodeCalc() *fib {
	contFib = false
	if f.left != nil && f.right != nil {
		if f.left.calc == true && f.right.calc == true {
			contFib = true
			f.number = f.left.number + f.right.number
			f.calc = true
		}
	}
	if f.left != nil {
		if f.left.calc == false {
			f.left = f.left.nodeCalc()
		}
	}
	if f.right != nil {
		if f.right.calc == false {
			f.right = f.right.nodeCalc()
		}
	}
	return f
}

func (f *fib) printTree() {
	fmt.Println(f)
	if f.left != nil {
		f.left.printTree()
	}
	if f.right != nil {
		f.right.printTree()
	}
}

func (lv1 *lv) setLv(v int) *lv {
	if lv1 == nil {
		lv1 = &lv{level: v2}
		v2 += 1
		if v2 < v {
			lv1.next = lv1.next.setLv(v)
		}
	}
	return lv1
}

func (f *fib) appendLv(lv *lv) {
	if f.level == lv.level && f.append == false {
		lv.number = append(lv.number, f.number)
		f.append = true
	} else {
		if lv.next != nil {
			f.appendLv(lv.next)
		}
	}
	if f.left != nil {
		f.left.appendLv(lv)
	}
	if f.right != nil {
		f.right.appendLv(lv)
	}
}

func (lv *lv) printLv() {
	fmt.Println(lv)
	if lv.next != nil {
		lv.next.printLv()
	}
}

func (lv *lv) lvToStr() {
	str += "<center>"
	for _, v := range lv.number {
		a := strconv.Itoa(v)
		str += (a + " ")
	}
	str += ("</center>" + "<br/>")
	if lv.next != nil {
		lv.next.lvToStr()
	} else {
		//V = str
		return
	}
}

func (f *fib) search(n int, c1 chan int) {
	/*
		if l == f.level && n == f.number {
			c1 <- f.number
		}
	*/
	if n == f.N {
		c1 <- f.number
	}

	if f.left != nil {
		f.left.search(n, c1)
	}
	if f.right != nil {
		f.right.search(n, c1)
	}
}

func (f *fib) searchGo(n int, c2 chan int) {
	/*
		if l == f.level && n == f.number {
			c2 <- f.number
		}
	*/
	if n == f.N {
		c2 <- f.number
	}
	if f.left != nil {
		go f.left.searchGo(n, c2)
	}
	if f.right != nil {
		go f.right.searchGo(n, c2)
	}
}

func main() {
	ui, _ = lorca.New("", "", 1200, 900, "--remote-allow-origins=*")
	defer ui.Close()

	// GoのfuncをJavaScriptで呼び出せるようBind
	ui.Bind("fibTree", fibTree)
	ui.Bind("searchTree", searchTree)

	// 今回はurl.PathEscapeでそのままHTMLを配置
	// ※下記はHTTPサーバをたててそこから取得する例
	//   https://github.com/zserge/lorca/tree/master/examples/counter
	//<input type="button" value="buttun1" onclick="clickBtn3()" />
	//<button onclick="fibTree()">make</button>
	ui.Load("data:text/html," + url.PathEscape(`
	<!doctype html>
    <html lang="ja">
    <head>

	<script>
		function clickBtn3() {
    		const number2 = document.getElementById("number2");
			document.getElementById("span2").textContent = number2.value;
			return number2.value
		}
		function clickSearch() {
			const number3 = document.getElementById("number3");
			return number3.value
		}
    </script>
	</head>
    <body>
	<p>terms : <span id="span2"></span></p>
	<input type="number" id="number2" value="4" max="2000" />
	<input type="button" value="make" onclick="fibTree()" />
	<br/>
	number : <input type="number" id="number3" value="0" max="2000" />
	<input type="button" value="search" onclick="searchTree()" />
        <div id="content3"></div>
		<br/>
		<div id="content2"></div>
    </body>
    </html>
	`))

	<-ui.Done()
}

//class="fit-picture"
/*
crossorigin=anonymous
		referrerpolicy=no-referrer
		<img
		src="/Users/omura/OC/img/mc2.png"
		alt="mc-pic"
		width="1169"
		height="826"
		 />
*/

var f *fib
var Lv *lv

func fibTree() {

	//var str2 string
	// JavaScriptのfunction呼び出し
	str1 := ui.Eval(`clickBtn3();`).String()
	v, _ := strconv.Atoi(str1)
	// HTMLに反映
	f = f.fibTree(v)
	for {
		f = f.nodeCalc()
		if contFib == false {
			break
		}
	}
	f.printTree()
	Lv = Lv.setLv(v)

	f.appendLv(Lv)
	Lv.printLv()

	Lv.lvToStr()
	//fmt.Println(str2)
	ui.Eval(`document.getElementById('content2').innerHTML="` + str + `";`)

	l, v2 = 0, 0
	str=""
}

func searchTree() {
	str2 := ui.Eval(`clickSearch();`).String()
	/*
		Split := strings.Split(str2, ",")
		level, _ := strconv.Atoi(Split[0])
		number, _ := strconv.Atoi(Split[1])
	*/
	number, _ := strconv.Atoi(str2)
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)

	t1 := time.Now()
	//f.search(level, number, c1)
	f.search(number, c1)
	Ans1 := <-c1
	t2 := time.Now().Sub(t1)
	fmt.Println(Ans1, t2)

	t3 := time.Now()
	//go f.searchGo(level, number, c2)
	go f.searchGo(number, c2)
	Ans2 := <-c2
	t4 := time.Now().Sub(t3)
	fmt.Println(Ans2, t4)

	ui.Eval(`document.getElementById('content3').innerHTML="` + "Seq : " + t2.String() + "   Go : " + t4.String() + `";`)
}
