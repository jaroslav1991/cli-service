//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=interfaces_mock.go

package service

import (
	"github.com/jaroslav1991/cli-service/internal/model"
)

type CLIService struct {
	repo     Repository
	httpAddr string
	authKey  string
}

func NewCLIService(repo Repository, httpAddr, authKey string) *CLIService {
	return &CLIService{repo: repo, httpAddr: httpAddr, authKey: authKey}
}

type Repository interface {
	Create(events model.Events) error
	Get(authKey []string) (model.EventsByAuthKey, error)
	GetAuthKeys() ([]string, error)
	Update() error
	Drop() error
}
