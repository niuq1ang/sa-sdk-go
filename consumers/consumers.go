package consumers

import (
	"time"

	"github.com/NiuQiang00/sa-sdk-go/structs"
	"github.com/NiuQiang00/sa-sdk-go/utils"
)

type Consumer interface {
	Send(data structs.EventData) error
	Flush() error
	Close() error
}

func send(url string, data string, to time.Duration, list bool) error {
	pdata := ""

	if list {
		rdata, err := utils.GeneratePostDataList(data)
		if err != nil {
			return err
		}
		pdata = rdata
	} else {
		rdata, err := utils.GeneratePostData(data)
		if err != nil {
			return err
		}
		pdata = rdata
	}

	err := utils.DoRequest(url, pdata, to)
	if err != nil {
		return err
	}

	return nil
}
