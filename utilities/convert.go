package utilities

import "encoding/json"

func BytesToStruct(jsonBytes []byte, structData interface{}) error {

	err := json.Unmarshal(jsonBytes, structData)
	if err != nil {
		return err
	}

	return nil
}
func MapToStruct(data map[string]interface{}, obj interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, &obj)
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
