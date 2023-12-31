package secretmanager

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"twitterGo/awsgo"
	"twitterGo/models"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var datosSecret models.Secret

	fmt.Println("> Requesting Secret " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal([]byte(*key.SecretString), &datosSecret)
	if err != nil {
		return models.Secret{}, err
	}

	fmt.Println("> Reading Secret OK " + secretName)

	return datosSecret, nil
}
