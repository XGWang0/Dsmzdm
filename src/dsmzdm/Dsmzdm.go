package main

import (
	"commlib"
	"fetchhtml"
	"flag"
	"os"
)

func ParamsParse() (int, int, int) {
	var cmdflagset = flag.NewFlagSet("cmdflag", flag.ExitOnError)
	var (
		page       int
		commentcnt int
		vote       int
	)
	cmdflagset.IntVar(&page, "p", 10, "How many pages need to be clawer")
	cmdflagset.IntVar(&commentcnt, "c", 2, "Filter product whose comment count is greater then your setting")
	cmdflagset.IntVar(&vote, "v", 3, "Filter products whose vote count is greater then your setting")
	cmdflagset.Parse(os.Args[1:])
	commlib.Mtrloggger.Println(page, commentcnt, vote)
	return page, commentcnt, vote

}

func main() {
	commlib.Mtrloggger, _ = commlib.InitLogger()
	commlib.Mtrloggger.Println("Start")
	page, commentcnt, votecnt := ParamsParse()
	fetchhtml.HandelAllUrl(page, commentcnt, votecnt)
}
