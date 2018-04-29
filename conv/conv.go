package conv

import (
	"log"

	"github.com/djimenez/iconv-go"
)

// Convert req message ebcdic->utf8
func ConvertMsg(cd *iconv.Converter, src []byte) ([]byte, error) {
	l := len(src)
	desc := make([]byte, l*3)
	bytesRead, bytesWritten, err := cd.Convert(src, desc)
	if err != nil || bytesRead != l {
		return nil, err
	}
	return desc[:bytesWritten], nil
}

// Convert request message with fixed length
func ConvertMsgFixLen(cd *iconv.Converter, src []byte) ([]byte, error) {
	l := len(src)
	desc := make([]byte, l*3)
	bytesRead, bytesWritten, err := cd.Convert(src, desc)
	if err != nil || bytesRead != l {
		log.Printf("<<< read:%d, writen:%d\n", bytesRead, bytesWritten)
		return nil, err
	}
	return desc[:l], nil
}
