package main

import (
	"fmt"
	"inno/hw3/types"
)

func main() {
	b := types.Bold{}
	i := types.Italics{}
	c := types.Code{}
	p := types.Plain{}

	fmt.Println(b.Format("bold"))
	fmt.Println(i.Format("italic"))
	fmt.Println(c.Format("code"))
	fmt.Println(p.Format(i.Format("plain from italic")))

	ch := types.ChainFormatter{}
	fmt.Println(ch.AddFormatter(b).AddFormatter(i).AddFormatter(c).Format("chain format"))
}
