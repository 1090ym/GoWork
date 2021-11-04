package flow

import "strconv"

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
	flow.RunStep(inputDataSet, step)
}

// 为每个vid对应的flow添加一个step
func (dis *Distributor) AddStepToFlow(Function func(shard InputDataShard)) {
	for _, flow := range dis.Flows {
		step := flow.NewStep()
		step.Function = Function
		flow.Steps = append(flow.Steps, step)
	}
}

// 数据分区策略
func DataPartition(row []interface{}) InputDataSet {
	// Initial InputData
	inputDataShard := InputDataSet{
		shards: make([]*InputDataShard, TASK_SLOT_SIZE),
	}
	for i := 0; i < TASK_SLOT_SIZE; i++ {
		shard := &InputDataShard{
			data: make([]interface{}, 0),
		}
		inputDataShard.shards = append(inputDataShard.shards, shard)
	}

	// row partition by hash % TASK_SLOT_SIZE
	for r := range row {
		hash := CalcHash(strconv.Itoa(int(r)))
		index := hash % TASK_SLOT_SIZE
		inputDataShard.shards[index].data = append(inputDataShard.shards[index].data, r)
	}
	return inputDataShard
}
