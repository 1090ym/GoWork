package flow

func (dis *Distributor) NewFlow() *Flow {
	flow := &Flow{}
	return flow
}

func (fl *Flow) RunStep(nextStep int, inputDataSet InputDataSet) {
	if nextStep < len(fl.Steps) {
		return
	}
	fl.Steps[nextStep].InputDataToStep(inputDataSet)
	step := fl.Steps[nextStep]
	for i := 0; i < len(step.Tasks); i++ {
		step.RunTask(step.Tasks[i])
	}
}
