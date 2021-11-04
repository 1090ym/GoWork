package flow

func (dis *Distributor) NewFlow() *Flow {
	flow := &Flow{}
	return flow
}

// 运行step下的task
func (fl *Flow) RunStep(inputDataSet InputDataSet, nextStep int) {
	if nextStep < len(fl.Steps) {
		return
	}
	// 把数据分片到各个task的channel中
	fl.Steps[nextStep].InputDataToStep(inputDataSet)
	step := fl.Steps[nextStep]
	for i := 0; i < len(step.Tasks); i++ {
		step.RunTask(step.Tasks[i])
	}
}
