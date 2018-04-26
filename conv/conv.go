package conv

import (
	"tcpserver/config"

	"github.com/djimenez/iconv-go"
)

// Convert req message ebcdic->utf8
func ConvertMsg(src []byte, l int) ([]byte, error) {
	cd, err := iconv.NewConverter(config.FromEncoding, config.ToEncoding)
	if err != nil {
		return nil, err
	}
	defer cd.Close()

	desc := make([]byte, l*3)
	bytesRead, bytesWritten, err := cd.Convert(src, desc)
	if err != nil || bytesRead != l {
		return nil, err
	}
	//log.Printf("<<< read:%d, writen:%d\n", bytesRead, bytesWritten)
	return desc[:bytesWritten], nil
}

func ConvertMsgFixLen(src []byte, l int) ([]byte, error) {
	cd, err := iconv.NewConverter(config.FromEncoding, config.ToEncoding)
	if err != nil {
		return nil, err
	}
	defer cd.Close()

	desc := make([]byte, l*3)
	bytesRead, _, err := cd.Convert(src, desc)
	if err != nil || bytesRead != l {
		return nil, err
	}
	//log.Printf("<<< read:%d, writen:%d\n", bytesRead, bytesWritten)
	return desc[:l], nil
}
