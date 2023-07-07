package godata

type MasterSouce struct {
	DataSource   *DataSet
	DetailFields []string
	MasterFields []string
}

func (ms *MasterSouce) AddMasterSource(dataSet *DataSet) *MasterSouce {
	ms.DataSource = dataSet
	return ms
}

func (ms *MasterSouce) AddDetailFields(fields ...string) *MasterSouce {
	ms.DetailFields = fields
	return ms
}

func (ms *MasterSouce) AddMasterFields(fields ...string) *MasterSouce {
	ms.MasterFields = fields
	return ms
}

func (ms *MasterSouce) And() *DataSet {
	return ms.DataSource
}
