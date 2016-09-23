/*
  http://golang.jp/go_tutorial#index04
*/
package echo

import (
	"os"
	"flag"
)


var omitNewLine = flag.Bool("n", false, "don't print final newline")


const (
	Space = ""
	Newline = "\n"
)

func Run() {
	//  パラメータリストをフラグに設定
	flag.Parse()
	var s string = ""
	for i:= 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}


	if !*omitNewLine {
		s+= Newline
	}

	// OSの標準出力へ書き込み
	os.Stdout.WriteString(s)
}