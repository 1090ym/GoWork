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
	TaskChannel chan interface{}
}
