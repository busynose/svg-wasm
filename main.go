package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"syscall/js"
	"time"

	svg "github.com/ajstarks/svgo"
)

func draw(this js.Value, args []js.Value) interface{} {
	rand.NewSource(time.Now().Unix())
	buf := bytes.NewBufferString("")
	canvas := svg.New(buf)
	x := args[1].Int()
	y := args[2].Int()
	r := rand.Intn(255)
	g := rand.Intn(255)
	b := rand.Intn(255)
	rgb := fmt.Sprintf("fill:rgb(%d,%d,%d)", r, g, b)
	canvas.Circle(x, y, args[0].Int(), rgb)
	return js.ValueOf(buf.String())
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("draw", js.FuncOf(draw))
	<-done
}
