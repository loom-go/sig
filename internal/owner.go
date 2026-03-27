package internal

import (
	"iter"
	"sync"
)

type Owner struct {
	mu  sync.RWMutex
	run sync.Mutex

	// cleanup functions to be called when the node is disposed
	cleanups []func()

	errorListeners   []func(any)
	disposeListeners []func()

	// the context values of this owner
	context map[uint64]any

	parent       *Owner
	prevSibling  *Owner
	nextSibling  *Owner
	childrenHead *Owner
}

func (r *Runtime) NewOwner() *Owner {
	o := &Owner{
		cleanups: make([]func(), 0),
		context:  make(map[uint64]any),
	}

	if parent := r.CurrentOwner(); parent != nil {
		parent.AddChild(o)
	}

	return o
}

func (o *Owner) Run(fn func() error) (err error) {
	o.run.Lock()
	defer o.run.Unlock()

	r := GetRuntime()
	r.tracker.RunWithOwner(o, func() { err = fn() })

	return err
}

func (parent *Owner) AddChild(child *Owner) {
	parent.mu.Lock()
	defer parent.mu.Unlock()

	child.parent = parent
	child.prevSibling = nil
	child.nextSibling = parent.childrenHead

	if parent.childrenHead != nil {
		parent.childrenHead.prevSibling = child
	}

	parent.childrenHead = child
}

func (n *Owner) Children() iter.Seq[*Owner] {
	return func(yield func(*Owner) bool) {
		// Snapshot the head under lock. The walk follows nextSibling
		// pointers which are set once per node (in AddChild) and never
		// modified, so traversal is safe without holding the lock.
		// New children prepended after the snapshot are simply missed.
		n.mu.RLock()
		child := n.childrenHead
		n.mu.RUnlock()

		for child != nil {
			next := child.nextSibling
			if !yield(child) {
				return
			}
			child = next
		}
	}
}

func (n *Owner) Cleanup() {
	n.run.Lock()
	defer n.run.Unlock()

	defer n.recover()

	n.cleanup()
}

func (n *Owner) cleanup() {
	n.DisposeChildren()

	n.mu.RLock()
	fns := n.cleanups
	n.cleanups = nil
	n.mu.RUnlock()

	for _, fn := range fns {
		fn()
	}
}

func (n *Owner) Dispose() {
	n.run.Lock()
	defer n.run.Unlock()

	defer n.recover()

	n.dispose()
}

func (n *Owner) dispose() {
	n.cleanup()

	n.mu.RLock()
	fns := n.disposeListeners
	n.disposeListeners = nil
	n.mu.RUnlock()

	for _, fn := range fns {
		fn()
	}
}

func (n *Owner) DisposeChildren() {
	// Detach the child list under lock. The detached list is safe to walk
	// via nextSibling without holding the lock — nextSibling is set once
	// per node in AddChild and never modified after.
	n.mu.RLock()
	head := n.childrenHead
	n.childrenHead = nil
	n.mu.RUnlock()

	for child := head; child != nil; {
		next := child.nextSibling
		child.Dispose()
		// Break reference chain to allow GC of disposed subtrees.
		child.nextSibling = nil
		child.prevSibling = nil
		child = next
	}
}

// OnCleanup registers a function to be called ONCE when this node recomputes (for Computed) AND when the owner is disposed
func (n *Owner) OnCleanup(fn func()) {
	n.mu.Lock()
	n.cleanups = append(n.cleanups, fn)
	n.mu.Unlock()
}

// OnDispose registers a function to be called ONCE, only when the owner is disposed (not when this Computed node is recomputed)
func (n *Owner) OnDispose(fn func()) {
	n.mu.Lock()
	n.disposeListeners = append(n.disposeListeners, fn)
	n.mu.Unlock()
}

func (n *Owner) OnError(fn func(any)) {
	n.mu.Lock()
	n.errorListeners = append(n.errorListeners, fn)
	n.mu.Unlock()
}

func (n *Owner) recover() {
	r := recover()
	if r == nil {
		return
	}

	for owner := n; owner != nil; owner = owner.parent {
		owner.mu.RLock()
		listeners := owner.errorListeners
		owner.mu.RUnlock()

		if len(listeners) > 0 {
			for _, fn := range listeners {
				fn(r)
			}
			return
		}
	}

	panic(r)
}

func (n *Owner) setContext(id uint64, value any) {
	n.mu.Lock()
	n.context[id] = value
	n.mu.Unlock()
}

func (n *Owner) getContext(id uint64) (any, bool) {
	n.mu.RLock()
	val, ok := n.context[id]
	n.mu.RUnlock()
	return val, ok
}
