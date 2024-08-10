package infrastructure

import (
	"time"
	"fmt"

	"github.com/google/uuid"
)

func GenerateID() string{
	id := uuid.New().String()
	timeStamp := time.Now().UnixNano()

	return fmt.Sprintf("%s-%d", id, timeStamp)
}