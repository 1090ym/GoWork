package flow

// 创建一个step
func (fl *Flow) NewStep() *Step {
	step := &Step{
		Id: len(fl.Steps),
	}
	fl.Steps = append(fl.Steps, step)
	return step
}

// 创建一个task
func (step *Step) NewTask() *Task {
	task := &Task{Step: step, Id: len(step.Tasks), TaskChannel: make(chan interface{}, TASK_CHANNEL_SIZE)}
	step.Tasks = append(step.Tasks, task)
	return task
}

// 运行该task
func (step *Step) RunTask(task *Task) {
	taskChan := task.TaskChannel
	inputData := make([]interface{}, TASK_CHANNEL_SIZE)
	for len(taskChan) != 0 {
		input := <-taskChan
		inputData = append(inputData, input)
	}
	go step.Function(inputData)
}

// 将数据分发到step的每个task的channel中
func (step *Step) InputDataToStep(inputDataSet InputDataSet) {
	for index, shard := range inputDataSet.shards {
		if index >= len(step.Tasks) {
			step.Tasks = append(step.Tasks, step.NewTask())
		}
		step.Tasks[index].InputDataToTask(*shard)
	}
}

// 将输入数据发送到task的channel中
func (task *Task) InputDataToTask(input InputDataShard) {
	for data := range input.data {
		task.TaskChannel <- data
	}
}

// 将输出数据发送到task的channel中
func (task *Task) OutputDataToTask(output OutputDataShard) {
	for data := range output.data {
		task.TaskChannel <- data
	}
}
