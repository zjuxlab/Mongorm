package Mongorm

import (
	"context"
	"reflect"

	"github.com/sirupsen/logrus"
)

var ctx context.Context = context.TODO()

func logError[T any](method string, err error) {
	if err == nil {
		return
	}
	var v T
	logrus.Errorf("[Mongorm] Error While %s[%s] in DB: \n%s", method, reflect.TypeOf(v).Name(), err.Error())
}
