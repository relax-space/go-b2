package b2

import (
	"fmt"

	"github.com/relax-space/go-a1/a1"
	"github.com/relax-space/go-a2/a2"
)

func Hello() string {
	return fmt.Sprintf("%v %v", a1.Hello(), a2.Hello())
}
