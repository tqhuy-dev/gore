package utilities

import (
	"errors"
	"fmt"
)

var ErrDBCrack = errors.New("errDBCrack")

type ErrorAppTier string

const (
	DefaultTier ErrorAppTier = ""
	StoreTier   ErrorAppTier = "Store"
	ServiceTier ErrorAppTier = "Service"
)

func NewErrorApp(errorTier ErrorAppTier, funcName string, err error) error {
	return fmt.Errorf("[%s] %s failed \n -> %w", errorTier, funcName, err)
}

func RunErrorDB() error {
	return NewErrorApp(StoreTier, "RunErrorDB", ErrDBCrack)
}

func RunErrorService() error {
	err := RunErrorDB()
	return NewErrorApp(ServiceTier, "RunErrorDB", err)
}
