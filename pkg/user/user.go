package users 

import (
	
	"encoding/json"
	"errors"
	"github.com/aws/aws-lamda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/dynomodb"
	"github.com/aws/aws-sdk-go/dynomodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
    ErrorFailedToUnmarshalRecord="failed to unmarshal record"
	ErrorFailedToFetchRecord="failed to fetch record"
	ErrorInvalidUserData="invalid user data "
	ErrorInvalidEmail="Invalid email"
	ErrorCouldNotMarhsalItem="could not Marhshal Item"
	ErrorCouldDeleteItem="could bit dynamo put Item"
	ErrorUserAlreadyExists="user.User already Exists"
	ErrorUserDoesNotExists="user.User doesnot Exist"
)
type User struct {
	Email string `json:"email"`
	firstName string  `json:"firstName"`
	LastName string    `json:"lastName"`
}
func fetchUser(email, tableName string , dyanClient, dynamodbiface.DynamoDBAPI)(*User, error){
  input := &dynamodb.GetItemInput {
	Key:map[string]*dynamodb.AttributeValue{
		"email":{
			S:aws.String(email),
		},
	},
	TableName:aws.String(tableName),
  }
  resutl, er:=dyanClient.GetItem(input)
  if err!=nil {
	return nil, errors.New(ErrorFailedToFetchRecord)
  }
 item:new(User)
 err=dynamodbattribute.UnmarshalMap(result.Item, item)
 if err!=nil {
	return nil, errors.New(ErrorFailedToUnmarshalRecord)
 }
 return item, nil 
}

func fetchUsers(tableName string, dyanClient dyanClientface.DynamoDBAPI)(*[]User, error ){
 input:= &dynamodb.ScanInput {
	TableName:aws.String(tableName)
 }
 result,err:=dyanClient.Scan(input) {
	if err!=nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item:=new([]User)
	err = dynamodbattribute.Unmarshal(result.Items, item)
	return item, nil 
 }
}

func CreateUser(req events.APIGateProxyRequest, tableName string , dyanClient dyanClientface.DynamoDBAPI)(*User, error){
var u User 
if err:= json.Unmarshal([]byte(req.body), &u); err!=nil {
}
if !validators.IsEmailValid(u.Email){
	return nil, errors.New(ErrorInvalidEmail)
}
currentUser, _= fetchUser(u.Email, tableName,dynaClient) 
if currentUser!=nil && len(currentUser.Email) != 0{
	return nil, errors.new(ErrorUserAlreadyExists)
}  
av, err:dynamodbattribute.marshalMap(u)
if err!=nil {
	return nil, errors.New(ErrorCouldNotMarhsalItem)
}
input:&dynampdb.PutItemInput{
	Item:av, 
	TableName:aws.String(tableName)
}
_, err=dyanClientPutItem(input)
irr err!=nil {
	return nil, errors.New(ErrorCouldNotDynamoPutItem)
}
return &u, nil 
}  

func UpdateUser(req events.APIGateProxyRequest, tableName string, dyanClient dynamodbiface.DynamoDBAPI)(*User, error){
     var u User 
	 if err:=json.Unmarshal([]byte(req.Body), &u); err:nil {
		return nil, error.New(ErrorInvalidEmail)
	 }
	 currentUser, _:=fetchUser(u.Email, tableName, dynaClient)
	 if currentUser!=nil && len(currentUser.email)==0 {
		return nil , error.New(ErrorUserDoesNotExists)
	 }
	av, err:= dynomodbattribute.marshalMap(u)
	if err!=nil {
		return nil, errors.New(ErrorCouldNotMarhsalItem)
	}
	input:=&dynamodb.PutItemInput{
		Item:av, 
		TableName:aws.String(tableName),
	}
	_, err:dynaClient.PutItem(input) 
	if err!=nil{
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
}

func DeleteUser(req events.APIGateProxyRequest, tableName string, dyanClient dynamodbiface.DynamoDBAPI) error {
 email:=req.QueryStringParameters["email"]
 input:=&dynamodb.DeleteItemInput{
	Key:map[string]*dynnamodb.AttributeValue{
		"email":{
			S:aws.String(email),
		},
	}
	TableName:aws.String(tableName),
 }
 _, err:=dynaClient.DeleteItem(input)
 if err!=nil {
	return errors.New(ErrorCouldnotDeleteItem)
 }
 return nil 
}