package flow

import "GoWork/distribute/conf"

type Distributor struct {
	Flows     []*Flow
	VidToFlow map[int]int
}

func (dis *Distributor) Init(vids []int) {
	dis.VidToFlow = make(map[int]int, len(vids))
	for index, vid := range vids {
		dis.VidToFlow[vid] = index
		dis.Flows = append(dis.Flows, dis.NewFlow())
	}
}

func ReceiveRpcMsg(line []byte, vid int, step int) {

}

func DataPartition(row []interface{}) InputDataSet {
	inputDataShard := InputDataSet{
		shards: make([]*InputDataShard, conf.TASK_SLOT_SIZE),
	}
	for i := 0; i < conf.TASK_SLOT_SIZE; i++ {
		shard := &InputDataShard{
			data: make([]interface{}, 0),
		}
		inputDataShard.shards = append(inputDataShard.shards, shard)
	}

	// row partition

	return inputDataShard
}
