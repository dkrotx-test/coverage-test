package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dkrotx/coverage-test/pkg"
)

func main() {
	verbose := flag.Bool("v", false, "be vebose")
	flag.Parse()

	raw_expr := strings.Join(flag.Args(), "")
	if raw_expr == "" {
		flag.Usage()
		os.Exit(64)
	}

	tokens, err := pkg.ParseString(raw_expr)
	if err != nil {
		panic(fmt.Sprintf(`failed to parse "%v": %s`, raw_expr, err))
	}

	if *verbose {
		fmt.Printf("Tokens: %q\n", tokens)
	}

	rpn, err := pkg.BuildRPN(tokens)
	if err != nil {
		panic(fmt.Sprintf(`failed to build RPN from "%v" (%v)`, raw_expr, tokens))
	}

	if *verbose {
		fmt.Printf("RPN:    %q\n", rpn)
	}

	fmt.Println(pkg.EvalRPN(rpn))
}
