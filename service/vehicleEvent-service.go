package service

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/droatl2000/demo-ledger-rest/model"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type VehicleEventService interface {
	Save(model.VehicleEvent) model.VehicleEvent
	FindAll() []model.VehicleEvent
	FindByVin(string) []model.VehicleEvent
}

type vehicleEventService struct {
	vehicleEvents []model.VehicleEvent
}

func NewEvent() VehicleEventService {
	return &vehicleEventService{
		vehicleEvents: []model.VehicleEvent{},
	}
}

func (service *vehicleEventService) Save(vehicleEvent model.VehicleEvent) model.VehicleEvent {
	var client = GetHederaClient()

	myTopicId, err := hedera.TopicIDFromString(os.Getenv("VEHICLE_EVENT_TOPIC_ID"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("The topic ID = %v\n", myTopicId)

	ma, err := json.Marshal(vehicleEvent)
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

	return vehicleEvent
}

func (service *vehicleEventService) FindByVin(searchVin string) []model.VehicleEvent {
	var client = GetHederaClient()

	myTopicId, err := hedera.TopicIDFromString(os.Getenv("VEHICLE_EVENT_TOPIC_ID"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("The topic ID = %v\n", myTopicId)

	var results []model.VehicleEvent

	sub, err := hedera.NewTopicMessageQuery().
		SetTopicID(myTopicId).
		SetStartTime(time.Unix(0, 0)).
		Subscribe(client, func(message hedera.TopicMessage) {
			var ma model.VehicleEvent
			err := json.Unmarshal(message.Contents, &ma)
			if err != nil {
				println(err.Error(), ": error Unmarshalling")
			}
			fmt.Println(ma.Vin, "-", ma.Servicer, "-", ma.SelectedFileName, "-", ma.Technician)
			if (ma.Vin == searchVin) || (searchVin == "") {
				results = append(results, ma)
			}
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

func (service *vehicleEventService) FindAll() []model.VehicleEvent {
	return service.FindByVin("")
}
