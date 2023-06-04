package utils

import "github.com/google/uuid"

type IDGenerator struct{}

func (g *IDGenerator) GenerateID() string {
	return uuid.New().String()
}
