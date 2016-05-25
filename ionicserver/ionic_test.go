package ionic

import (
	"os/exec"
	"testing"
)

// Test that running gopherjs test <file> works as we expect
func TestTest(t *testing.T) {
	// Stub out a simple unit test and check its return val
	exec.Command("")
}

// Test that running gopherjs run <file> works as we expect
func TestRun(t *testing.T) {
	// Stub out a simple program and check its return val
	exec.Command("")
}

// Test that we can bootstrap a new Ionic app using their CLI
func TestIonicBoot(t *testing.T) {
	// Create a new Ionic package
	exec.Command("")
	// Delete the boilerplate JS
	exec.Command("")
	// Write new GopherJS file
	exec.Command("")
}

// Test that the GopherJS compiled JS works in Ionic
func TestIonicRun(t *testing.T) {
	// Run the compiled Ionic program
	exec.Command("")
}
