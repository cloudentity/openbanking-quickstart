package main

import "github.com/cloudentity/openbanking-quickstart/openbanking/obbr/consents/models"

func OBBRMapError(err error) interface{} {
	return models.OpenbankingBrasilResponseError{}
}
