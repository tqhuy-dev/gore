package utilities

import "encoding/json"

func JSONStringToStruct(jsonString string, structData interface{}) error {

	jsonBytes, err := json.Marshal(jsonString)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonBytes, structData)
	if err != nil {
		return err
	}

	return nil
}

func StructToJSONString(structData interface{}) (jsonString string, err error) {

	jsonBytes, err := json.Marshal(structData)
	if err != nil {
		return
	}
	jsonString = string(jsonBytes)
	return
}

func InterfaceToStruct(source interface{}, des interface{}) error {

	jsonBytes, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonBytes, des)
	return err
}
