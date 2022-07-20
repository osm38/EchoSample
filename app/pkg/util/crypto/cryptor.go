package crypto

import(
	"crypto/sha256"
	"echosample/pkg/util/convert"
	"time"
	"fmt"
)

func GetSHA256() string {
	t := time.Now()
	s := convert.TimeToStr(t)
	data := sha256.Sum256([]byte(s))
	hash := fmt.Sprintf("%x", data)

	return hash
}