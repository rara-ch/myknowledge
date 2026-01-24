package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	tcs := []struct {
		name        string
		handler     handler
		expectError bool
	}{
		{
			name: "commandName",
			handler: func(s *state, cmd inputCommand) error {
				return nil
			},
			expectError: false,
		},
		{
			name: "errorCommandName",
			handler: func(s *state, cmd inputCommand) error {
				return errors.New("manufactured error in tests")
			},
			expectError: true,
		},
	}

	cmds := newCLICommands()
	s := &state{}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("test #%d", i), func(t *testing.T) {
			cmds.register(tc.name, tc.handler)
			err := cmds.run(s, inputCommand{name: tc.name})
			if tc.expectError {
				if err == nil {
					t.Errorf("did not get an expected error")
				}
			} else {
				if err != nil {
					t.Errorf("got an error when not expecting")
				}
			}
		})
	}
}
