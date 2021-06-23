package service

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/droatl2000/demo-ledger-rest/model"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type VerifiedServicerService interface {
	Save(model.VerifiedServicer) model.VerifiedServicer
	FindAll() []model.VerifiedServicer
}

type verifiedServicerService struct {
	verifiedServicers []model.VerifiedServicer
}

func NewServicer() VerifiedServicerService {
	return &verifiedServicerService{
		verifiedServicers: []model.VerifiedServicer{},
	}
}

func (service *verifiedServicerService) Save(verifiedServicer model.VerifiedServicer) model.VerifiedServicer {
	var client = GetHederaClient()

	myTopicId, err := hedera.TopicIDFromString(os.Getenv("VEHICLE_SERVICER_TOPIC_ID"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("The topic ID = %v\n", myTopicId)

	ma, err := json.Marshal(verifiedServicer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ma))

	//Create the transaction
	transaction := hedera.NewTopicMessageSubmitTransaction().
		SetTopicID(myTopicId).
		SetMessage([]byte(string(ma)))

	//Sign with the client operator private key and submit the transaction to a Hedera network
	txResponse, err := transaction.Execute(client)
	if err != nil {
		panic(err)
	}

	//Request the receipt of the transaction
	transactionReceipt, err := txResponse.GetReceipt(client)
	if err != nil {
		panic(err)
	}

	//Get the transaction consensus status
	transactionStatus := transactionReceipt.Status

	fmt.Printf("The transaction consensus status is %v\n", transactionStatus)
	//v2.0.0

	return verifiedServicer
}

func (service *verifiedServicerService) FindAll() []model.VerifiedServicer {
	var client = GetHederaClient()

	myTopicId, err := hedera.TopicIDFromString(os.Getenv("VEHICLE_SERVICER_TOPIC_ID"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("The topic ID = %v\n", myTopicId)

	var results []model.VerifiedServicer

	sub, err := hedera.NewTopicMessageQuery().
		SetTopicID(myTopicId).
		SetStartTime(time.Unix(0, 0)).
		Subscribe(client, func(message hedera.TopicMessage) {
			var ma model.VerifiedServicer
			err := json.Unmarshal(message.Contents, &ma)
			if err != nil {
				println(err.Error(), ": error Unmarshalling")
			}
			fmt.Println(ma.Name, "-", ma.StreetAddress, "-", ma.Services, "-", ma.Technicians)
			results = append(results, ma)
		})

	if err != nil {
		println(err.Error(), ": error subscribing to the topic")
		return results
	}

	time.Sleep(3 * time.Second)
	sub.Unsubscribe()

	if err != nil {
		panic(err)
	}

	return results
}
