
package workers

import (
    "context"
    "github.com/redis/go-redis/v9"
)

func StartWorker(rdb *redis.Client) {
    for {
        rdb.BLPop(context.Background(), 0, "jobs")
    }
}
