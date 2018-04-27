package conv

import (
	"github.com/djimenez/iconv-go"
)

// Convert req message ebcdic->utf8
func ConvertMsg(cd *iconv.Converter, src []byte, l int) ([]byte, error) {
	desc := make([]byte, l*3)
	bytesRead, bytesWritten, err := cd.Convert(src, desc)
	if err != nil || bytesRead != l {
		return nil, err
	}
	//log.Printf("<<< read:%d, writen:%d\n", bytesRead, bytesWritten)
	return desc[:bytesWritten], nil
}

// Convert request message with fixed length
func ConvertMsgFixLen(cd *iconv.Converter, src []byte, l int) ([]byte, error) {
	desc := make([]byte, l*3)
	bytesRead, _, err := cd.Convert(src, desc)
	if err != nil || bytesRead != l {
		return nil, err
	}
	//log.Printf("<<< read:%d, writen:%d\n", bytesRead, bytesWritten)
	return desc[:l], nil
}
