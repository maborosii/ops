// package main

// import (
// 	"fmt"
// 	"os"

// 	"ops/pkg/setting"

// 	"cuelang.org/go/cue"
// 	"cuelang.org/go/cue/cuecontext"
// )

// func main() {
// 	var (
// 		c *cue.Context
// 		v cue.Value
// 	)

// 	// create a context
// 	c = cuecontext.New()

// 	// read and compile value
// 	d, _ := os.ReadFile("allconfig.cue")
// 	v = c.CompileBytes(d)
// 	iter, _ := v.LookupPath(cue.ParsePath("minstone.app.frontend")).List()

// 	// print the value
// 	// fmt.Println(demo)
// 	for iter.Next() {
// 		var a setting.BaseAppCfg
// 		// fmt.Println(iter.Value().LookupPath(cue.ParsePath("package_type")).String())
// 		tmp := iter.Value()
// 		fmt.Println(tmp)
// 		_ = tmp.Decode(&a)
// 		fmt.Println(a)

// 	}
// }
package main

import (
	"fmt"

	"cuelang.org/go/cue"
)

type ab struct{ A, B int }

var r cue.Runtime

var x ab

func main() {
	i, _ := r.Compile("test", `{A: 2, B: 4}`)
	_ = i.Value().Decode(&x)
	fmt.Println(x)
	// i, _ := r.Compile("test", `{B: "foo"}`)
	// _ = i.Value().Decode(&x)
	// fmt.Println(x)
}
