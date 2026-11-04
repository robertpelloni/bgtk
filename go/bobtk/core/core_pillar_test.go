package core

import (
	"sync/atomic"
	"testing"
)

func TestCorePillarInitialization(t *testing.T) {
	pillar := NewCorePillar()
	pillar.Initialize()
	if !pillar.initialized {
		t.Errorf("expected CorePillar to be initialized")
	}
}

func TestQuantumSchedulerTaskExecution(t *testing.T) {
	scheduler := NewQuantumScheduler(4)
	defer scheduler.Close()

	var counter int64
	for i := 0; i < 20; i++ {
		scheduler.Dispatch(func() {
			atomic.AddInt64(&counter, 1)
		})
	}

	scheduler.Wait()

	if atomic.LoadInt64(&counter) != 20 {
		t.Fatalf("expected 20 tasks executed, got %d", counter)
	}

	if scheduler.ExecutedCount() != 20 {
		t.Fatalf("expected executed count 20, got %d", scheduler.ExecutedCount())
	}
}

func TestReflectContextAndDynamicValue(t *testing.T) {
	rc := NewReflectContext()

	var callbackValue string
	rc.OnChange("app.theme", func(v Value) {
		callbackValue = v.String()
	})

	rc.SetProperty("app.theme", "Dark-Modern")

	val, found := rc.GetProperty("app.theme")
	if !found {
		t.Fatalf("expected property 'app.theme' to be found")
	}

	if val.String() != "Dark-Modern" {
		t.Errorf("expected 'Dark-Modern', got %s", val.String())
	}

	if callbackValue != "Dark-Modern" {
		t.Errorf("expected listener callback 'Dark-Modern', got %s", callbackValue)
	}

	rc.SetProperty("window.width", 1920)
	valInt, _ := rc.GetProperty("window.width")
	if valInt.Int() != 1920 {
		t.Errorf("expected integer 1920, got %d", valInt.Int())
	}

	if rc.PropertyCount() != 2 {
		t.Errorf("expected 2 properties, got %d", rc.PropertyCount())
	}
}

func TestEntityManager(t *testing.T) {
	em := NewEntityManager()

	ent := em.CreateEntity("MainWindow")
	if ent == nil {
		t.Fatalf("failed to create entity")
	}

	if em.EntityCount() != 1 {
		t.Fatalf("expected 1 entity, got %d", em.EntityCount())
	}

	type WindowBounds struct {
		X, Y, W, H int
	}

	err := em.AttachComponent(ent.ID, "bounds", WindowBounds{X: 100, Y: 100, W: 800, H: 600})
	if err != nil {
		t.Fatalf("failed to attach component: %v", err)
	}

	comp, ok := em.GetComponent(ent.ID, "bounds")
	if !ok {
		t.Fatalf("expected component 'bounds' to be found")
	}

	bounds, ok := comp.(WindowBounds)
	if !ok || bounds.W != 800 {
		t.Errorf("invalid component data: %+v", comp)
	}
}
