package flow

import (
	"fmt"
	"strconv"
)

var DisManager Distributor

type Distributor struct {
	Flows     []*Flow
	VidToFlow map[int]int
}

// 为每个vid初始化flow
func (dis *Distributor) InitDistributor(vids []int) {
	dis.VidToFlow = make(map[int]int, len(vids))
	for index, vid := range vids {
		dis.VidToFlow[vid] = index
		dis.Flows = append(dis.Flows, dis.NewFlow())
	}
}

// 接收rpc消息
func (dis *Distributor) ReceiveRpcMsg(row []interface{}, vid int, step int) {
	index := dis.VidToFlow[vid]
	flow := dis.Flows[index]
	inputDataSet := DataPartition(row)
	PrintInputDataSet(inputDataSet)
	flow.RunStep(inputDataSet, step)
}

// 为每个vid对应的flow添加一个step
func (dis *Distributor) AddStepToFlow(Function func(input InputDataShard)) {
	for _, flow := range dis.Flows {
		step := flow.NewStep()
		step.Function = Function
		flow.Steps = append(flow.Steps, step)
	}
}

// 对每个step接收到的数据分片到每个task
func DataPartition(row []interface{}) InputDataSet {
	// Initial InputData
	inputData := InputDataSet{
		shards: make([]*InputDataShard, 0),
	}
	for i := 0; i < TASK_SLOT_SIZE; i++ {
		shard := &InputDataShard{
			data: make([]interface{}, 0),
		}
		inputData.shards = append(inputData.shards, shard)
	}
	// row partition by hash % TASK_SLOT_SIZE
	for i, r := range row {
		hash := CalcHash(strconv.Itoa(int(i)))
		index := hash % TASK_SLOT_SIZE
		inputData.shards[index].data = append(inputData.shards[index].data, r)
	}
	return inputData
}

func PrintInputDataSet(inputDataSet InputDataSet) {
	fmt.Print("inputDataSet: ")
	for _, shard := range inputDataSet.shards {
		fmt.Print(*shard, " ")
	}
	fmt.Println()
}
