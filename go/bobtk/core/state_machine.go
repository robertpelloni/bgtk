package core

import (
	"fmt"
	"sync"
)

// StateNode represents an individual state in the state hierarchy (bobguistate.h parity).
type StateNode struct {
	Name        string
	Transitions map[string]string // Event -> TargetState
	OnEnter     func()
	OnExit      func()
}

// NewStateNode creates a new state node.
func NewStateNode(name string) *StateNode {
	return &StateNode{
		Name:        name,
		Transitions: make(map[string]string),
	}
}

func (s *StateNode) AddTransition(event, targetState string) {
	s.Transitions[event] = targetState
}

// StateMachine manages states, active state, and event-driven transitions (bobguistate.h parity).
type StateMachine struct {
	mu           sync.RWMutex
	states       map[string]*StateNode
	currentState *StateNode
	initialState *StateNode
}

// NewStateMachine instantiates a state machine with an initial state.
func NewStateMachine(initial *StateNode) *StateMachine {
	sm := &StateMachine{
		states:       make(map[string]*StateNode),
		currentState: initial,
		initialState: initial,
	}
	if initial != nil {
		sm.states[initial.Name] = initial
		if initial.OnEnter != nil {
			initial.OnEnter()
		}
	}
	return sm
}

// AddState registers an additional state node.
func (sm *StateMachine) AddState(state *StateNode) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.states[state.Name] = state
}

// CurrentState returns the active state node name.
func (sm *StateMachine) CurrentState() string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	if sm.currentState != nil {
		return sm.currentState.Name
	}
	return ""
}

// PostEvent processes an event and transitions if a matching transition exists.
func (sm *StateMachine) PostEvent(event string) (string, error) {
	sm.mu.Lock()
	if sm.currentState == nil {
		sm.mu.Unlock()
		return "", fmt.Errorf("no active state configured")
	}

	targetName, ok := sm.currentState.Transitions[event]
	if !ok {
		sm.mu.Unlock()
		return sm.currentState.Name, fmt.Errorf("no transition for event %q in state %q", event, sm.currentState.Name)
	}

	targetState, exists := sm.states[targetName]
	if !exists {
		sm.mu.Unlock()
		return sm.currentState.Name, fmt.Errorf("target state %q not found", targetName)
	}

	exitCb := sm.currentState.OnExit
	enterCb := targetState.OnEnter
	sm.currentState = targetState
	newStateName := targetState.Name
	sm.mu.Unlock()

	if exitCb != nil {
		exitCb()
	}
	if enterCb != nil {
		enterCb()
	}

	return newStateName, nil
}

// StateCount returns registered state count.
func (sm *StateMachine) StateCount() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.states)
}
