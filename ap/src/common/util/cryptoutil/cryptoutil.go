package cryptoutil

import (
	"alma-server/ap/src/common/error/chk"
	"crypto/aes"
	"encoding/hex"
)

const configSeacretKey = "iK!a|HwY$9QsT_$Ru/MD(bNFkYZ3GLcW"

// EncPassword .
func EncPassword(src string) string {
	return Enc(src, configSeacretKey)
}

// DecPassword .
func DecPassword(hexCipherText string) string {
	return Dec(hexCipherText, configSeacretKey)
}

// Enc .
func Enc(src string, seacretKey string) string {
	srcTxtBytes := []byte(src)

	block, err := aes.NewCipher([]byte(seacretKey))
	chk.SE(err)

	encTxtBytes := make([]byte, len(srcTxtBytes))

	// Encrypt
	block.Encrypt(encTxtBytes, srcTxtBytes)

	return hex.EncodeToString(encTxtBytes)
}

// Dec .
func Dec(hexCiperText string, seacretKey string) string {

	bCipherText, _ := hex.DecodeString(hexCiperText)
	block, err := aes.NewCipher([]byte(seacretKey))
	chk.SE(err)

	decryptedText := make([]byte, len(bCipherText))
	block.Decrypt(decryptedText, bCipherText)

	return string(decryptedText)
}
