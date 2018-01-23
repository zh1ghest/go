package static

import (
	"github.com/v2pro/wombat/generic"
)

func init() {
	encodeAnything.ImportFunc(encodeBinary)
}

var encodeBinary = generic.DefineFunc(
	"EncodeBinary(dst DT, src ST)").
	Param("DT", "the dst type to copy into").
	Param("ST", "the src type to copy from").
	Source(`
dst.WriteBinary(src)
	`)