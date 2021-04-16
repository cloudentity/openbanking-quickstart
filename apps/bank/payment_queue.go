package main

import (
	"time"

	paymentModels "github.com/cloudentity/openbanking-quickstart/openbanking/paymentinitiation/models"
	logrus "github.com/sirupsen/logrus"
)

type PaymentQueue struct {
	repo  UserRepo
	queue chan paymentModels.OBWriteDomesticResponse5
}

func NewPaymentQueue(repo UserRepo) (PaymentQueue, error) {
	return PaymentQueue{
		repo:  repo,
		queue: make(chan paymentModels.OBWriteDomesticResponse5),
	}, nil
}

func (pq *PaymentQueue) Start() {
	for payment := range pq.queue {
		payment := payment
		time.AfterFunc(10*time.Second, func() {
			if err := pq.repo.SetDomesticPaymentStatus(*payment.Data.DomesticPaymentID, AcceptedSettlementCompleted); err != nil {
				logrus.WithError(err).Errorf("failed to update domestic payment status for id %s (%s)", *payment.Data.DomesticPaymentID, err)
			}
		})
	}
}
