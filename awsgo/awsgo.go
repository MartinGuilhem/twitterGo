package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	Ctx context.Context
	Cfg aws.Config
	err error
)

func InicializeAWS() {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		panic("Error while loading config .aws/config " + err.Error())
	}
}
