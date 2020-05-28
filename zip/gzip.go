package zip

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"github.com/sirupsen/logrus"
)

/**
  gzip 压缩，且 16进制转换
 */
func Gzip(json string) string {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err := zw.Write([] byte(json))
	if err != nil {
		logrus.Fatal(err)
	}

	if err := zw.Close(); err != nil {
		logrus.Fatal(err)
	}

	encoded := hex.EncodeToString(buf.Bytes())
	return encoded
}
