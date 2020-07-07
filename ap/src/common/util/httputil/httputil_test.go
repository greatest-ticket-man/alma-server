package httputil_test

import (
	"reflect"
	"testing"

	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/util/httputil"

	"context"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/common/util/httputil

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("httputil", func() {

		g.It("test", func() {

			rt := reflect.TypeOf("")

			result, err := httputil.GetJSON(context.Background(), "https://example.com/", nil, rt)
			chk.SE(err)

			g.Assert(result.Result.(string) == "").IsFalse()

		})

	})

}
