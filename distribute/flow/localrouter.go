package flow

var LocalRt *LocalRouter

type LocalRouter struct {
	MessageChannel chan interface{}
}

type Message struct {
	InputDataShard InputDataShard
	Step           int
}

func (router *LocalRouter) Process() {
	for message := range router.MessageChannel {

		/*	}
			for message := <- router.MessageChannel {*/
		router.SendMsg(message.(Message).InputDataShard, message.(Message).Step)
	}
}

func (router *LocalRouter) SendMsg(inputDataShard InputDataShard, step int) {

}

func (router *LocalRouter) InputDataToChan(inputDataShard InputDataShard, step int) {
	router.MessageChannel <- Message{InputDataShard: inputDataShard, Step: step}
}
