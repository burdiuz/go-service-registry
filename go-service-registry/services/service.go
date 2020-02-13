package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Service struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	RemoteAddr string `json:"addr"`
	RemotePort string `json:"port"`
	lastSeenAt time.Time
}

func NewService(name, version, remoteAddr, remotePort string) *Service {
	return &Service{name, version, remoteAddr, remotePort, time.Now()}
}

func (s *Service) Touch() {
	s.lastSeenAt = time.Now()
}

func (s *Service) Remote() string {
	return fmt.Sprintf("%v:%v", s.RemoteAddr, s.RemotePort)
}

func (s *Service) IsExpired(dur time.Duration) bool {
	return time.Now().Before(s.lastSeenAt.Add(dur))
}

func (s *Service) ToJSON() ([]byte, error) {
	if s == nil {
		return []byte{}, nil
	}

	return json.Marshal(s)
}

func (s *Service) FromJSON(data []byte) error {
	return json.Unmarshal(data, s)
}

func (s *Service) WriteHTTP(w http.ResponseWriter) error {
	data, err := s.ToJSON()

	if err != nil {
		return err
	}

	w.Write(data)
	return nil
}
