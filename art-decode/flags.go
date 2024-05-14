package artdec

import "flag"

var (
	MultilineDecode  = flag.Bool("m", false, "Decode multiple lines from input file")
	SinglelineEncode = flag.Bool("e", false, "Encode a single line")
	MultilineEncode  = flag.Bool("me", false, "Encode multiple lines from input file")
	Help             = flag.Bool("h", false, "Print how to use program")
)

func init() {
	flag.Parse()
}
