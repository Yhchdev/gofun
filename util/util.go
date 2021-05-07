package util

import (
	"github.com/gofrs/uuid"
	"strings"
)

// Uuid 生成uuid
func Uuid() string {
	u := uuid.Must(uuid.NewV4()).String()
	return strings.ReplaceAll(u, "-", "")
}
