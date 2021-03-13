package helpers

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
	"strings"
)

type sysmsg struct {
	body        string
	withoutFunc bool
	fnName      string
}

func NewSysMsg(str string, params ...interface{}) sysmsg {
	m := sysmsg{
		body: fmt.Sprintf(str, params...),
	}
	m.fnName = m.getCallerFunc()
	return m
}

func (m sysmsg) getCallerFunc() string {
	pc, _, _, _ := runtime.Caller(2)
	funcObj := runtime.FuncForPC(pc)

	parts := strings.Split(funcObj.Name(), "/")
	fnName := parts[len(parts)-1]

	return fnName
}

func (m sysmsg) WithoutFunc() sysmsg {
	m.withoutFunc = true

	return m
}

func (m sysmsg) String() string {
	if m.withoutFunc {
		return m.body
	}

	return fmt.Sprintf("%s, %s", m.fnName, m.body)
}

func (m sysmsg) S() string {
	return m.String()
}

func (m sysmsg) WrapErr(err error) error {
	return errors.Wrap(err, m.String())
}

func (m sysmsg) ToErr() error {
	return errors.New(m.String())
}

func PrettyType(t string) string {
	newT := strings.Replace(t, "*", "", -1)

	return newT
}
