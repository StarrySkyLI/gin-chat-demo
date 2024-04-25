package script

import (
	"chat/mq/task"
	"context"
)

func TaskCreateSync(ctx context.Context) {
	tSync := new(task.SyncTask)
	_ = tSync.RunTaskCreate(ctx)

}
