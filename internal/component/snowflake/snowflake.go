package snowflake

import (
	"context"
	"sdxx/server/internal/component/redis"
	"sync"

	"github.com/bwmarrin/snowflake"
	"github.com/dobyte/due/v2/log"
)

var (
	once     sync.Once
	instance *Node
)

type Node struct {
	node *snowflake.Node
}

func (n *Node) Generate() int64 {
	return n.node.Generate().Int64()
}

func (n *Node) BatchGenerate(count int) []int64 {
	var result []int64
	for i := 0; i < count; i++ {
		result = append(result, n.Generate())
	}
	return result
}

func Instance() *Node {
	once.Do(func() {
		rdb := redis.Instance()
		id, err := rdb.Incr(context.Background(), "snowflake:worker_id").Result()
		if err != nil {
			log.Fatalf("snowflake failed to allocate id: %v", err)
		}
		node, err := snowflake.NewNode(id % 1024)
		if err != nil {
			log.Fatalf("snowflake new node err: %v", err)
		}
		instance = &Node{node: node}
	})
	return instance
}

func Generate() uint64 {
	return uint64(Instance().Generate())
}

func BatchGenerate(count int) []uint64 {
	values := Instance().BatchGenerate(count)
	var results []uint64
	for _, value := range values {
		results = append(results, uint64(value))
	}
	return results
}
