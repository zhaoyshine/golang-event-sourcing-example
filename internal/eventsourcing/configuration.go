package eventsourcing

import "reflect"

type Configuration struct {
	viewRepositoryNama      string
	eventRepositoryNama     string
	commandHandlers         map[string][]func(Command)
	eventAppliedHandlers    map[string][]func(Event, Aggregate)
	eventCommittingHandlers map[string][]func(Event, Aggregate)
	eventCommittedHandlers  map[string][]func(Event, Aggregate)
}

func NewConfiguration(view, event string) *Configuration {
	return &Configuration{
		viewRepositoryNama:      view,
		eventRepositoryNama:     event,
		commandHandlers:         make(map[string][]func(Command)),
		eventAppliedHandlers:    make(map[string][]func(Event, Aggregate)),
		eventCommittingHandlers: make(map[string][]func(Event, Aggregate)),
		eventCommittedHandlers:  make(map[string][]func(Event, Aggregate)),
	}
}

func (c *Configuration) ViewRepositoryNama() string {
	return c.viewRepositoryNama
}

func (c *Configuration) EventRepositoryNama() string {
	return c.eventRepositoryNama
}

func (c *Configuration) AddCommand(commands []Command) {
	for _, command := range commands {
		name := reflect.TypeOf(command).Name()
		c.commandHandlers[name] = make([]func(Command), 0)
	}
}

func (c *Configuration) AddEvent(events []Event) {
	for _, event := range events {
		name := reflect.TypeOf(event).Name()
		c.eventAppliedHandlers[name] = make([]func(Event, Aggregate), 0)
		c.eventCommittingHandlers[name] = make([]func(Event, Aggregate), 0)
		c.eventCommittedHandlers[name] = make([]func(Event, Aggregate), 0)
	}
}

func (c *Configuration) OnCommand(cfs []map[Command]func(Command)) {
	for _, cf := range cfs {
		for command, f := range cf {
			name := reflect.TypeOf(command).Name()
			c.commandHandlers[name] = append(c.commandHandlers[name], f)
		}
	}
}

func (c *Configuration) OnEventApplied(cfs []map[Event]func(Event, Aggregate)) {
	for _, cf := range cfs {
		for command, f := range cf {
			name := reflect.TypeOf(command).Name()
			c.eventAppliedHandlers[name] = append(c.eventAppliedHandlers[name], f)
		}
	}
}

func (c *Configuration) OnEventCommitting(cfs []map[Event]func(Event, Aggregate)) {
	for _, cf := range cfs {
		for command, f := range cf {
			name := reflect.TypeOf(command).Name()
			c.eventCommittingHandlers[name] = append(c.eventCommittingHandlers[name], f)
		}
	}
}

func (c *Configuration) OnEventCommitted(cfs []map[Event]func(Event, Aggregate)) {
	for _, cf := range cfs {
		for command, f := range cf {
			name := reflect.TypeOf(command).Name()
			c.eventCommittedHandlers[name] = append(c.eventCommittedHandlers[name], f)
		}
	}
}
