package main

import (
	"context"
	"fmt"
	"plugin"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/benji-vesterby/atomizer"
)

func main() {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Create an instance of the atomizer to test the add conductor with
	mizer := atomizer.Atomize(ctx)

	if p, err := plugin.Open("plugintest.so"); err == nil {

		var sym plugin.Symbol
		if sym, err = p.Lookup("MYATOM"); err == nil {
			spew.Dump(sym)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

	mizer.Exec()
}
