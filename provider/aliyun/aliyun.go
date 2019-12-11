package aliyun

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/messagebird/sachet"
)

type AliyunConfig struct {
	RegionId  string `yaml:"region_id"`
	APIKey    string `yaml:"api_key"`
	APISecret string `yaml:"api_secret"`
}

type Aliyun struct {
	client *dysmsapi.Client
}

func NewAliyun(config AliyunConfig) (*Aliyun, error) {
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.APIKey, config.APISecret)

	if err != nil {
		return nil, err
	}

	return &Aliyun{client: client}, nil
}

func (aliyun *Aliyun) Send(message sachet.Message) error {
	for _, recipent := range message.To {
		request := dysmsapi.CreateSendSmsRequest()

		request.Scheme = "https"
		request.SignName = "BlockFin"
		request.PhoneNumbers = recipent
		request.TemplateCode = "SMS_175485223"
		request.TemplateParam = "{code: 1234}"

		fmt.Println(request)

		if response, err := aliyun.client.SendSms(request); err != nil {
			fmt.Println(response)
			return err
		}
	}

	return nil
}
