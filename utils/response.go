package utils

//GenericResponse response that is sent to rest call at all the times
type GenericResponse struct {
	Code        int64          `json:"CODE"`
	Description string         `json:"DESCRIPTION"`
	Datalist    *[]interface{} `json:"DATALIST"`
}

//Method to Build Generic Response
func BldGnrRsp(cusCode int64, cusMsg string, cusBody *[]interface{}) (res *GenericResponse) {
	res = new(GenericResponse)
	res.Code = cusCode
	res.Description = cusMsg
	res.Datalist = cusBody
	return res
}

func ToInterfaceArrayFromList(inputList *[]interface{}) *([]interface{}) {
	var retArray ([]interface{}) = nil
	if *inputList != nil && len(*inputList) > 0 {
		retArray = make([]interface{}, len(*inputList))
		for i, d := range *inputList {
			retArray[i] = d
		}
	}
	return &retArray
}
func ToInterfaceArray(input *interface{}) *([]interface{}) {
	var retArray ([]interface{}) = nil
	retArray = make([]interface{}, 1)
	retArray[0] = input
	return &retArray
}

func ToInterfaceArrayFromString(input string) *([]interface{}) {
	var retArray ([]interface{}) = nil
	retArray = make([]interface{}, 1)
	retArray[0] = input
	return &retArray
}
