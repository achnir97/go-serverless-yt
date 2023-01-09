package handlers 
import ( 
"net/http"
"github.com/achnir97/go-serverless-yt/pkg/users"
"github.com/aws/aws-lamda-go/events"
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/dyanomodb/dynamodbiface"
)


var ErrorMethodNotAllowed="method not allowed"
type ErrorBody struct {
	Error *string `json:"error, omitempty"`
}
func GetUser(req events.APIGateProxyRequest, tableName string, dyanClient dynamodbiface.DynamoDBAPI)(
	*events.APIGateProxyResponce, error 
){ 
	email:=req.QueryString["email"]
	if len(email)>0{
	result, err:= user.fetchUser(email, tableName, dyanClient)
	if err!=nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatuysOK, result)
	}
	user.fetchUsers(tableName, dyanClient)
	if err!=nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
return apiResponse(http.StatuysOK, result)
}


func CreateUser(req events.APIGateProxyRequest, tableName string, dyanClient dynamodbiface.DynamoDBAPI)(
	*events.APIGateProxyResponce, error 
){
result, err:=user.CreateUser(req, tableName, dyanClient) 
if err!=nil {
return apiResponse(http.StatusBadRequest,ErrorBody{
	aws.String(err.Error()),
})
}
return apiResponse(http.StatusCreated, result)
}


func Updateser(req events.APIGateProxyRequest, tableName string, dyanClient dynamodbiface.DynamoDBAPI)(
	*events.APIGateProxyResponce, error 
){
	result, err:=user.UpdateUser(req, tableName, dyanClient)
	if err!=nil {
		return apiResponse(http.StatusBadRequest,ErrorBody{
			aws.String(err.Error()),
		})
		}
		return apiResponse(http.StatusCreated, result)
		}
		

func DeleteUser(req events.APIGateProxyRequest, tableName string, dyanClient dynamodbiface.DynamoDBAPI)(
	*events.APIGateProxyResponce, error 
){
	err:=user.DeleteUser(req, tableName, dyanClient)
	if err!=nil {
		return apiResponse(http.StatusBadRequest,ErrorBody{
			aws.String(err.Error()),
		})
		}
		return apiResponse(http.StatusCreated, mil)
		}

func UnhandleMethod()(*even.APIGateProxyResponce, error) {
return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}