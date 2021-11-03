package flow

import "GoWork/distribute/conf"

func (fl *Flow) NewStep() *Step {
	step := &Step{
		Id: len(fl.Steps),
	}
	fl.Steps = append(fl.Steps, step)
	return step
}

func (step *Step) NewTask() *Task {
	task := &Task{Step: step, Id: len(step.Tasks), TaskChannel: make(chan interface{}, conf.TASK_CHANNEL_SIZE)}
	step.Tasks = append(step.Tasks, task)
	return task
}

func (step *Step) RunTask(task *Task) {
	taskChan := task.TaskChannel
	inputData := make([]interface{}, conf.TASK_CHANNEL_SIZE)
	for len(taskChan) != 0 {
		input := <-taskChan
		inputData = append(inputData, input)
	}
	go step.Function(inputData)
}

func (step *Step) InputDataToStep(inputDataSet InputDataSet) {
	for index, shard := range inputDataSet.shards {
		if index >= len(step.Tasks) {
			step.Tasks = append(step.Tasks, step.NewTask())
		}
		step.Tasks[index].InputDataToTask(*shard)
	}
}

func (task *Task) InputDataToTask(input InputDataShard) {
	for data := range input.data {
		task.TaskChannel <- data
	}
}

func (task *Task) OutputDataToTask(output OutputDataShard) {
	for data := range output.data {
		task.TaskChannel <- data
	}
}
