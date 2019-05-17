//Package opts defines a struct-tag based API for
//rapidly building command-line interfaces. For example:
//
//  package main
//
//  import (
//  	"log"
//  	"github.com/jpillora/opts"
//  )
//
//  func main() {
//  	type config struct {
//  		File  string `opts:"help=file to load"`
//  		Lines int    `opts:"help=number of lines to show"`
//  	}
//  	c := config{}
//  	opts.Parse(&c)
//  	log.Printf("%+v", c)
//  }
//
//Build and run:
//
//  $ go build -o my-prog
//  $ ./my-prog --help
//
//    Usage: my-prog [options]
//
//    Options:
//    --file, -f   file to load
//    --lines, -l  number of lines to show
//    --help, -h
//
//  $ ./my-prog -f foo -l 12
//  {File:foo Lines:42}
//
//See https://github.com/jpillora/opts for more information.
package opts
