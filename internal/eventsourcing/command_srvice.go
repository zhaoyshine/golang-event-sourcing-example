package eventsourcing

import "reflect"

type CommandService struct {
	Configuration *Configuration
}

func (c *CommandService) Execute(command Command) {
	name := reflect.TypeOf(command).Name()
	fs, ok := c.Configuration.commandHandlers[name]
	if ok {
		for _, f := range fs {
			f(command)
		}
	}
}
