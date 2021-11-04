package flow

import "time"

const ETCD_RESPONSE_TIMEOUT = time.Second * 5

const TASK_CHANNEL_SIZE = 1024

const TASK_SLOT_SIZE = 100

type State int

const (
	PENDING State = iota
	READY
	RUNNING
	SUCCESS
	FAIL
)
