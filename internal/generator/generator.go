package generator

import (
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"runtime"
)

type Generator struct {
}

func (g *Generator) GenerateID() string {
	id, err := uuid.NewV7()
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.Info(logrus.WithFields(logrus.Fields{
			"file": file,
			"line": line - 2,
		}))
	}
	return id.String()
}
