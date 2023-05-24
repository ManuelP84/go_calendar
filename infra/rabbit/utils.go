package rabbit

import (
	"bytes"
	"encoding/json"

	"github.com/ManuelP84/calendar/domain/task/events"
)

func Serialize(event events.TaskEvent) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(event)
	return b.Bytes(), err
}
