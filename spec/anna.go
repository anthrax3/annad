package spec

// Anna represents the main object, the programm root. It contains all further
// objects and manages the whole control flow around Anna's implementations.
type Anna interface {
	// Boot initializes and starts Anna like booting a machine. The call to Boot
	// blocks until Anna is completely initialized, so you might want to call it
	// in a separate goroutine.
	Boot()

	Object

	// Shutdown ends all processes of Anna like shutting down a machine. The call
	// to Shutdown blocks until Anna is completely shut down, so you might want
	// to call it in a separate goroutine.
	Shutdown()
}
