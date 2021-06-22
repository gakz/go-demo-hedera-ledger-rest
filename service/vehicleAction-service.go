package service

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/droatl2000/demo-ledger-rest/model"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

type VehicleActionService interface {
	Save(model.VehicleAction) model.VehicleAction
	FindAll() []model.VehicleAction
	FindByVin(string) []model.VehicleAction
}

type vehicleActionService struct {
	vehicleActions []model.VehicleAction
}

func New() VehicleActionService {
	return &vehicleActionService{
		vehicleActions: []model.VehicleAction{},
	}
}

func GetHederaClient() *hedera.Client {
	//Loads the .env file and throws an error if it cannot load the variables from that file correctly
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("Unable to load environment variables from .env file. Error:\n%v\n", err))
	}

	//Grab your testnet account ID and private key from the .env file
	myAccountId, err := hedera.AccountIDFromString(os.Getenv("ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}

	myPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	//Print your testnet account ID and private key to the console to make sure there was no error
	fmt.Printf("The account ID is = %v\n", myAccountId)

	//Create your testnet client
	client := hedera.ClientForTestnet()
	client.SetOperator(myAccountId, myPrivateKey)

	return client
}

func (service *vehicleActionService) Save(vehicleAction model.VehicleAction) model.VehicleAction {
	var client = GetHederaClient()

	myTopicId, err := hedera.TopicIDFromString(os.Getenv("TOPIC_ID"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("The topic ID = %v\n", myTopicId)

	ma, err := json.Marshal(vehicleAction)
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

	return vehicleAction
}

func (service *vehicleActionService) FindByVin(searchVin string) []model.VehicleAction {
	var client = GetHederaClient()

	myTopicId, err := hedera.TopicIDFromString(os.Getenv("TOPIC_ID"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("The topic ID = %v\n", myTopicId)

	var results []model.VehicleAction

	sub, err := hedera.NewTopicMessageQuery().
		SetTopicID(myTopicId).
		SetStartTime(time.Unix(0, 0)).
		Subscribe(client, func(message hedera.TopicMessage) {
			var ma model.VehicleAction
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

func (service *vehicleActionService) FindAll() []model.VehicleAction {
	return service.FindByVin("")
}
