import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func (conn *Conn) RegisterQueues(queues []string) {
	gv := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "redis_queue_count",
			Help: "Number of samples in some queue",
		},
		[]string{
			"queue",
		},
	)
	prometheus.Register(gv)

	go func() {
		ticker := time.NewTicker(time.Second * 15)
		for range ticker.C {
			gaugeQueueLevels(qv, queues)
		}
	}()
}

func (conn *Conn) gaugeQueueLevels(gv *prometheus.GaugeVec, queues []string) {
	rc := conn.pool.Get()
	defer rc.Close()

	for _, queue := range queues {
		count, err := redis.Int64(rc.Do(`llen`, queue))
		if err != nil {
			logger.LogError(err)
		}
		gv.WithLabelValues(queue, vm).Set(float64(count))
	}
}

