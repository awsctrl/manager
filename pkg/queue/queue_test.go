package queue

import (
	"testing"

	"awsctrl.io/pkg/event"
)

type queueTest struct{}

func (q queueTest) Start(<-chan struct{}) error {
	return nil
}
func (q queueTest) Reconcile(evt *event.Event) error {
	return nil
}

func TestReconcile(t *testing.T) {
	evt := event.Event{}
	q := queueTest{}

	if err := q.Reconcile(&evt); err != nil {
		t.Errorf("reconcile failed with error %s", err)
	}
}
