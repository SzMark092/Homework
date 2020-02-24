package db

import (
	"encoding/json"
	"errors"
)

func makeJason(CodeOfType int, actHandler Handler) ([]byte, error) {

	var actTypeData []interface{} = nil
	var err []error = nil
	var result []byte

	switch CodeOfType {
	case 1:
		actTypeData, err[0] = actHandler.SelectAllData(&DataPointDescription{})
	case 2:
		actTypeData, err[0] = actHandler.SelectAllData(&DataPoint{})
	case 3:
		actTypeData, err[0] = actHandler.SelectAllData(&Module{})
	default:

		err[0] = errors.New("Wrong type code.")

	}

	result, err[1] = json.Marshal(actTypeData)

	if err[0] != nil {

		return result, err[0]
	}

	return result, err[1]

}
