# go-demo-ledger-rest

This sample app is a REST API that reads and writes messages to the Hedera distrubuted ledger.  It is written in Go and utilizes the [Hedera Go SDK](https://github.com/hashgraph/hedera-sdk-go). The function of this demo app is to log vehicle maintenance/repairs onto a public, immutable ledger, in this case the Hedera network.  App users could use it to prove that they performed regualar maintenance on their vehicle, which would be usefull for resale evalution purposes.  Orgazizations, like police departments and repair shops could also use it to record incidents like accident damage.  Think of it as a decentralized [Carfax](https://www.carfax.com/vehicle-history-reports/).  The service providers that work on the vehicles come from a Hedera topic conforming to the [Hedera Hashgraph DID Method & Verifiable Credentials](https://hedera.com/blog/decentralized-identity-on-hedera-consensus-service) models.

## Setup

This sample app assumes you have already installed the GO distribution.  If not, you can find instructions [here](https://golang.org/doc/install)

Adiitionally, you will need a Hedera Portal profile. To create your Hedera Portal profile register [here](https://portal.hedera.com/register).  Once registered, you'll need to note your Account ID and your Private Key.  These credential will be used by the the app to access any Hedera network services uned in the demo.

Before starting the project, create an .env file in the project root directory.  This file will store environemtn variable, such as your Hedera Account ID, your Private Key, and Topic IDs used by the app.

### Set Hedera Credentials

> .env
>
> ACCOUNT_ID= (set account id)
>
> PRIVATE_KEY= (set private key)
>
> VEHICLE_EVENT_TOPIC_ID= (set later)
>
> VERIFIED_SERVICER_TOPIC_ID= (set later)

This project writes messages to a Hedera pub/sub topic, so you will need to create a topic by executing the following command from the project root directory.

> go run hederaTopicCreation.go

This will create a Hedera pub/sub topic and will return the Topic Ids for the application.
Edit the .env again and set the TOPIC_IDs

> .env
>
> ACCOUNT_ID=
>
> PRIVATE_KEY=
>
> VEHICLE_EVENT_TOPIC_ID= (set topic id)
>
> VERIFIED_SERVICER_TOPIC_ID= (set topic id))

Finally, execute the project.

> go run server.go

This will start a local webserver that serves the REST API used to create and read Hedera Topic messages.
THe default URL will be http://localhost:8082/vehicleEvents

Update server.go file to change the port, if desired.

## End Points

GET /vehicleEvents - return all messages for the vehicle event topic

GET /vehicleEvents/[ :vin ] - return messages for the vehicle event topic filtered by VIN

POST /vehicleEvents/ - save vehicle event to topic

Expected JSON request format for POST
>
>      {
>
>        "vin": "GA94234351",
>  
>        "workdescription": "Oil Change & Tune Up",
>
>        "servicer": "Smith Auto Repair",
>
>        "technician": "Joe Smith",
>
>        "selectedfile": "receipt.jpg"
>
>      }

GET /verifiedServicer - return all messages for the verified servicer topic

POST /verifiedServicer/ - save vehicle event to topic

Expected JSON request format for POST
>
> {
>
>     "id": "1",
>
>     "name": "Jiffy Lube #241",
>
>     "streetaddress": "11 Oak Mill Rd.",
>
>     "city": "Mephis",
>
>     "postalcode": "54322",
>
>     "country": "USA",
>
>     "services": [
>
>         ["Oil Change", 20],
>         ["Brakes",120],
>         ["Trans Fluid Drain",75]
>
>     ],
>
>     "technicians": [
>
>         ["Alton Green", 1],
>         ["Raj Patel", 2],
>         ["Mary Cook", 3]
>
>     ]
>
> }

The expected "selectedfile" is an uploaded image for the receipt or work summary.  Eventually, this project will be expanded with functionality to upload this file to a distrubuted storage layer, like [IPFS](https://ipfs.io/). 
  
## The UI  
The code for the UI used to interact with this REST API is in the [node-demo-ledger-ui repository](https://github.com/droatl2000/node-demo-ledger-ui)
