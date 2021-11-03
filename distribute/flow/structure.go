package flow

type Flow struct {
	Steps []*Step
}

type Step struct {
	Id       int
	Tasks    []*Task
	NextStep int
	Function func([]interface{}) error
}

type Task struct {
	Id          int
	Step        *Step
	TaskChannel chan interface{}
}

type InputDataSet struct {
	shards []*InputDataShard
}

type OutputDataSet struct {
	shards []*OutputDataShard
}

type InputDataShard struct {
	data []interface{}
}

type OutputDataShard struct {
	data []interface{}
}
