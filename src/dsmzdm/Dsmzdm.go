package main

import (
	"fetchhtml"
	"flag"
	"logo"
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
	logo.Log.Info(page, commentcnt, vote)
	return page, commentcnt, vote

}

func main() {
	logo.Log.Info("Start")
	page, commentcnt, votecnt := ParamsParse()
	fetchhtml.HandelAllUrl(page, commentcnt, votecnt)
}
