package flow

type Flow struct {
	Steps []*Step
}

type Step struct {
	Id       int
	Tasks    []*Task
	NextStep int
	Function func(shard InputDataShard)
}

type Task struct {
	Id          int
	Step        *Step
	TaskChannel chan interface{}
	State       State
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
