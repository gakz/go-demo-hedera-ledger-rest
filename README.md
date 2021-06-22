# go-demo-ledger-rest

This sample app is a REST API that reades and writes messages the Hedera distrubuted ledger.  It is written in Go and utilizes the Hedera Go SDK.

This sample app assumes you have already installed the GO distribution.  If not, you can find instructions [here](https://golang.org/doc/install)

Adiitionally, you will need a Hedera Portal profile. To create your Hedera Portal profile register [here] (https://portal.hedera.com/register).  nce registered, you'll need to note your Account ID and your Private Key.  These credential will be used by the the app to access any Hedera network services uned in the demo.

Before starting the project, update the .env file with your Hedera Account ID and your Private Key.

This project writes messages to a Hedera pub/sub topic, so you will need to create a topic by executing the following command from the project root directory.

> go run hederaTopicCreation.go

This will create a Hedera pub/sub topic and will return the Topic Id.
Edit the .env again and set the TOPIC_ID

Finally, execute the project.

> go run server.go

This will start a local webserver that serves the REST API used to create and read Hedera Topic messages.
THe default URL will be http://localhost:8082

Update server.go file to change the port, if desired.

THe code for the UI used to interact with this REST API is in this [repository] (https://github.com/droatl2000/node-demo-ledger-ui)
