package main

import "time"

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) put(keyValuePair KeyValuePair) error {
	if keyValuePair.Ttl > 0 {
		keyValuePair.Ttl = int(time.Now().Unix()) + keyValuePair.Ttl
	}
	return s.repository.put(keyValuePair)
}
func (s *Service) get(key string) (string, error) {
	KeyValuePairModel, err := s.repository.get(key)
	if err != nil {
		return "", err
	} else {
		return KeyValuePairModel.Value, nil
	}
}
func (s *Service) delete(key string) error {
	err := s.repository.delete(key)
	return err
}
