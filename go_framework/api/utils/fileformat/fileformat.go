package fileformat

import (
	"path"
	"strings"

	"github.com/twinj/uuid"
)

func UniqueFormat(fn string) string {
	fileName := strings.TrimSuffix(fn, path.Ext(fn))
	extension := path.Ext(fn)
	u := uuid.NewV4()
	newFileName := fileName + "-" + u.String() + extension

	return newFileName
}
