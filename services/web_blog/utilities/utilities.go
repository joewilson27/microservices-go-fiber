package utilities

/*
Our base response object
*/
func GetBaseResponseObject() map[string]interface{} {
	response := make(map[string]interface{})
	response["status"] = "fail"
	response["message"] = "something went wrong"
	return response
}