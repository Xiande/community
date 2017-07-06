package DBBLL

import (
	"DBDAL"
)

//DBBLL

type BaseBLL struct {
	dal DBDAL.DALInterface
}

type ParseInterface interface {
	Scan(dest ...interface{}) error
}
