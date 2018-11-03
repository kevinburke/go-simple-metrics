package metrics

import (
	"testing"
	"time"

	"github.com/rcrowley/go-metrics"
)

func TestNamespace(t *testing.T) {
	originalNamespace := Namespace
	defer func() {
		Namespace = originalNamespace
	}()
	Namespace = "foo"
	if getWithNamespace("bar") != "foo.bar" {
		t.Errorf("expected getWithNamespace(bar) to be foo.bar, was %s", getWithNamespace("bar"))
	}

	Namespace = ""
	if got := getWithNamespace("bar"); got != "bar" {
		t.Errorf("getWithNamespace(bar) should be 'bar', got %q", got)
	}
}

func TestIncrementIncrements(t *testing.T) {
	Increment("bar")
	Increment("bar")
	Increment("bar")
	mn := metrics.GetOrRegisterCounter("bar", nil)
	if mn.Count() != 3 {
		t.Errorf("expected Count() to be 3, was %d", mn.Count())
	}
}

func ExampleIncrement() {
	Start("web", "test@example.com")
	Increment("dequeue.success")
}

func ExampleMeasure() {
	Start("web", "test@example.com")
	Measure("workers.active", 6)
}

func ExampleTime() {
	Start("web", "test@example.com")
	start := time.Now()
	time.Sleep(3)
	Time("auth.latency", time.Since(start))
}
