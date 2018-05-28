package main

import (
	"context"
	"fmt"
	"io"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	fmt.Println("<html style='font-family: arial,sans-serif; background: blue; color: white;'><body><h1>Version 1</h1></body></html>")
}
