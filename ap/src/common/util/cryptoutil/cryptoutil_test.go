package cryptoutil_test

import (
	"alma-server/ap/src/common/util/cryptoutil"
	"log"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/common/util/cryptoutil

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	// AESの鍵長さは16byte，24byte，32byteのいずれかである必要がある（それぞれAES-128，AES-192，AES-256と呼ばれる）
	const seacretKey = "12345678901234567890123456789012"

	// 暗号化前文字列
	const srcText = "Almadestela110  "

	// 暗号化後文字列
	const encHexStr = "b9ea496bd9e9f399f891fb20fd3a5e69"

	g.Describe("文字列暗号化（AES）用のテスト", func() {

		g.It("32bytesの暗号キーで暗号化した場合", func() {
			ecnStr := cryptoutil.Enc(srcText, seacretKey)
			log.Println("asc str", ecnStr)
			g.Assert(ecnStr).Eql(encHexStr)
		})

		g.It("復元", func() {
			decStr := cryptoutil.Dec(encHexStr, seacretKey)
			log.Println("desc str", decStr)
			g.Assert(decStr).Eql(srcText)
		})

	})

}
