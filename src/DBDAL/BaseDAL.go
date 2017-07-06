// DBDAL project DBDAL.go
package DBDAL

// BaseDAL
type DALInterface interface {
	Add(model interface{}) int
	Update(model interface{})
	GetById(id int) interface{}
	GetList(whereStr string) interface{}
	Delete(id int)
	DeleteModel(interface{})
}

type BaseDAL struct {
	DriverName       string
	ConnectionString string
}

func (dal *BaseDAL) checkError(err error) {
	if err != nil {
		panic(err)
	}

}
