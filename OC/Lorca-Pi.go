package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand/v2"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/zserge/lorca"
)

var ui lorca.UI
var str string

// "--remote-allow-origins=*"
func main() {
	ui, _ = lorca.New("", "", 1200, 900, "--remote-allow-origins=*")
	defer ui.Close()

	// GoのfuncをJavaScriptで呼び出せるようBind
	ui.Bind("test", test)
	ui.Bind("reset", reset)

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
		function clickBtn() {
    		const plot = document.getElementById("plot");
			const thread = document.getElementById("thread");
			return plot.value+","+thread.value
		}
    </script>
	</head>
    <body>
	Plot : <input type="number" id="plot" value="1000000" max="1000000000" /><br/>
	Thread : <input type="number" id="thread" value="1" max="2000" /><br/>
	<input type="button" value="Test" onclick="test()" />
	<input type="button" value="Reset" onclick="reset()" />
    <div id="content"></div>
    </body>
    </html>
	`))

	<-ui.Done()
}

/*
<input type="button" value="Seq Float" onclick="SeqFloat()" />
	<input type="button" value="Go Float" onclick="GoFloat()" />
	<input type="button" value="Seq Big" onclick="SeqBig()" />
	<input type="button" value="Go Big" onclick="GoBig()" />
*/

func test() {
	var Pi1, Pi1Sum, Pi2, Pi2Sum, Pi3, Pi3Sum, Pi4, Pi4Sum int
	var str2 string
	str += "<tr>"

	input := ui.Eval(`clickBtn();`).String()
	Split := strings.Split(input, ",")
	Tstr := Split[1]
	plot, _ := strconv.Atoi(Split[0])
	thread, _ := strconv.Atoi(Split[1])
	n := plot / thread

	//fmt.Println("MCFloat")
	c1 := make(chan int, thread)
	t1 := time.Now()
	for i := 0; i < thread; i++ {
		MCFloat(n, c1)
	}
	for i := 0; i < thread; i++ {
		Pi1 = <-c1
		Pi1Sum += Pi1
	}
	t2 := time.Now()
	var Ans1 float64 = (float64(Pi1Sum) / float64(plot)) * 4
	//fmt.Println(Ans1)
	//str += ("<td> Seq Float : " + t2.Sub(t1).String() + "</td>")
	str += ("<td>" + t2.Sub(t1).String() + "</td>")
	fmt.Println("MCFloat Ans: ", Ans1, " Time: ", t2.Sub(t1), "\n")

	//fmt.Println("MCgoFloat")
	c2 := make(chan int, thread)
	t3 := time.Now()
	for i := 0; i < thread; i++ {
		go MCgoFloat(n, c2)
	}
	for i := 0; i < thread; i++ {
		Pi2 = <-c2
		Pi2Sum += Pi2
	}
	t4 := time.Now()
	var Ans2 float64 = (float64(Pi2Sum) / float64(plot)) * 4

	t24 := strconv.FormatFloat(float64(t2.Sub(t1))/float64(t4.Sub(t3)), 'f', 2, 64)
	//fmt.Print(Ans3)
	//str += ("<td>&ensp; Go Float : " + t4.Sub(t3).String() + "</td>")
	str += ("<td>" + t4.Sub(t3).String() + " (th=" + Tstr + ", x" + string(t24) + ")</td>")
	fmt.Println("MCgoFloat Ans: ", Ans2, " Time: ", t4.Sub(t3), "\n")

	//fmt.Println("MCBig")
	c3 := make(chan int, thread)
	t5 := time.Now()
	for i := 0; i < thread; i++ {
		MCBig(n, c3)
	}
	for i := 0; i < thread; i++ {
		Pi3 = <-c3
		Pi3Sum += Pi3
	}
	t6 := time.Now()
	var Ans3 float64 = (float64(Pi3Sum) / float64(plot)) * 4
	//fmt.Println(Ans2)
	//str += ("<td>&ensp; Seq Big : " + t6.Sub(t5).String() + "</td>")
	str += ("<td>" + t6.Sub(t5).String() + "</td>")
	fmt.Println("MCBig Ans: ", Ans3, " Time: ", t5.Sub(t6), "\n")

	//fmt.Println("MCgoBig")
	c4 := make(chan int, thread)
	t7 := time.Now()
	for i := 0; i < thread; i++ {
		go MCgoBig(n, c4)
	}
	for i := 0; i < thread; i++ {
		Pi4 = <-c4
		Pi4Sum += Pi4
	}
	t8 := time.Now()
	var Ans4 float64 = (float64(Pi4Sum) / float64(plot)) * 4

	t68 := strconv.FormatFloat(float64(t6.Sub(t5))/float64(t8.Sub(t7)), 'f', 2, 64)
	//fmt.Println(Ans2)a
	//str += ("<td>&ensp; Go Big : " + t8.Sub(t7).String() + "</td>")
	str += ("<td>" + t8.Sub(t7).String() + " (th=" + Tstr + ", x" + string(t68) + ")</td>")
	fmt.Println("MCgoBig Ans: ", Ans4, " Time: ", t8.Sub(t7), "\n")

	str += "</tr>"
	str2 = "<table>" + "<tr><td>|------- Float (Seq) -------|</td><td>|------- Float (Goroutine) -------|</td><td>|------- Big (Seq) -------|</td><td>|------- Big (Goroutine) -------|</td></tr>" + str + "</table>"
	ui.Eval(`document.getElementById('content').innerHTML="` + str2 + `";`)
}

func reset() {
	str = ""
	ui.Eval(`document.getElementById('content').innerHTML="` + str + `";`)
}

func MCFloat(n int, c chan int) {
	local_Pi := 0
	for i := 0; i < n; i++ {
		x := rand.Float64()
		X := math.Pow(x, 2)
		y := rand.Float64()
		Y := math.Pow(y, 2)
		if X+Y <= 1 {
			local_Pi += 1
		}
	}
	c <- local_Pi
}

func MCBig(n int, c chan int) {
	local_Pi := 0
	X := new(big.Float).SetPrec(1024)
	Y := new(big.Float).SetPrec(1024)
	Z := new(big.Float).SetPrec(1024)
	for i := 0; i < n; i++ {
		x := rand.Float64()
		xBig := new(big.Float).SetFloat64(x)
		X.Mul(xBig, xBig)
		y := rand.Float64()
		yBig := new(big.Float).SetFloat64(y)
		Y.Mul(yBig, yBig)
		Z.Add(X, Y)
		cmp := Z.Cmp(big.NewFloat(1))
		if cmp != 1 {
			local_Pi += 1
		}
	}
	c <- local_Pi
}

func MCgoFloat(n int, c chan int) {
	local_Pi := 0
	for i := 0; i < n; i++ {
		x := rand.Float64()
		X := math.Pow(x, 2)
		y := rand.Float64()
		Y := math.Pow(y, 2)
		if X+Y <= 1 {
			local_Pi += 1
		}
	}
	c <- local_Pi
	//wg.Done()
}

func MCgoBig(n int, c chan int) {
	local_Pi := 0
	X := new(big.Float).SetPrec(1024)
	Y := new(big.Float).SetPrec(1024)
	Z := new(big.Float).SetPrec(1024)
	for i := 0; i < n; i++ {
		x := rand.Float64()
		xBig := new(big.Float).SetFloat64(x)
		X.Mul(xBig, xBig)
		y := rand.Float64()
		yBig := new(big.Float).SetFloat64(y)
		Y.Mul(yBig, yBig)
		Z.Add(X, Y)
		cmp := Z.Cmp(big.NewFloat(1))
		if cmp != 1 {
			local_Pi += 1
		}
	}
	c <- local_Pi
	//wg.Done()
}
