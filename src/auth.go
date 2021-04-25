package src

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

/**
获取签名
*/
func getSig(ts string) string {
	sig := AppId + ts

	/**
	md5
	*/
	md5Sig := fmt.Sprintf("%x", md5.Sum([]byte(sig)))

	/**
	hmac
	*/
	h := hmac.New(sha1.New, []byte(SecretKey))
	h.Write([]byte(md5Sig))
	sha1Sig := h.Sum(nil)

	/**
	base64
	 */
	signature :=  base64.StdEncoding.EncodeToString(sha1Sig)
	return signature
}
