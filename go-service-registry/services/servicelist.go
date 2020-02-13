package services

import (
	"encoding/json"
	"net/http"
)

type ServiceList []*Service

func (sl ServiceList) ToJSON() ([]byte, error) {
	if sl == nil {
		return []byte{}, nil
	}

	return json.Marshal(sl)
}

func (sl ServiceList) FromJSON(data []byte) error {
	return json.Unmarshal(data, &sl)
}

func (sl ServiceList) WriteHTTP(w http.ResponseWriter) error {
	data, err := sl.ToJSON()

	if err != nil {
		return err
	}

	w.Write(data)
	return nil
}
