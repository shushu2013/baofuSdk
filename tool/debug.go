package tool

import (
	"fmt"
	"runtime"
)

func DumpStacks() string {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	return fmt.Sprintf("=== BEGIN stack dump ===\n%s\n=== END stack dump ===", buf)
}
