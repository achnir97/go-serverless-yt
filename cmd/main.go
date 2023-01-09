package main 
import(
	"github.com/akhil/go-serverless-yt/pkg/handlers"
	"os"
	"github.com/aws/aws-lamda-go/events"
	"github.com/aws/aws-lamda-go/lamda"
	"github.com/aws-sdk-go/aws"
	"github.com/aws/aws-skd-go/session"
	"github.com/aws/aws-skd-go/service"
	"github.com/aws/aws-skd-go/dynamodb"
)
var (
	dyanClient dynamodbface.DynamoDBAPI
)
fun main(){
region:=os.Getenv("AWS_REGION")
awsSession, err:=session.NewSession(&aws.Config{
	Region:awas.String(region)},)
	if err!=nil {
		return 
	}
	dyanClient=dynnamodb.New(awsSession)
	lamda.Start(handler)
}

const tableName="LamdaInGoUser"
func handler(req events.APIGateProxyRequest)(*events.APIGateProxyRequest, error){
switch req.HTTPMethod
{case "GET":
	return handlers.GetUser(req, tableName, dyanClient)
case "POST":
	return handlers.CreateUser(req, tableName, dyanClient)
case "PUT":
	return handlers.UpdateUser(req, tableName, dyanClient)
case "DELETE":
	return handlers.DeleteUser(req, tableName, dyanClient)
}
default:
	return handlers.UnhandleMethod()
}