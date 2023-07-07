package call_test_dep

import (
	"fmt"
	"github.com/hkhk366/test_dep"
)

func CallTestDep() {
	fmt.Println("Hello, world!")
	test_dep.PrintVersion()
}
