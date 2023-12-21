package day20

type Module interface {
	ReceivePulse(bool) []Event
	AddDestModules([]Module)
}

type Broadcaster struct {
	name               string
	destinationModules []Module
}

func (b *Broadcaster) ReceivePulse(highPulse bool) []Event {
	var events = []Event{}
	for _, mod := range b.destinationModules {
		events = append(events, Event{pulse: highPulse, moduleTargeted: mod, moduleSource: b})
	}
	return events
}

func (b *Broadcaster) AddDestModules(modules []Module) {
	b.destinationModules = modules
}

type Conjunction struct {
	name               string
	inputModulesPulses map[Module]bool
	destinationModules []Module
}

func (c *Conjunction) ReceivePulse(highPulse bool) []Event {
	var events = []Event{}

	return events
}

func (c *Conjunction) ReceivePulseInput(highPulse bool, sourceModule Module) []Event {
	var events = []Event{}
	c.inputModulesPulses[sourceModule] = highPulse

	var pulseToSend bool
	for _, allHighPulse := range c.inputModulesPulses { // if all high
		if !allHighPulse {
			pulseToSend = true
		}
	}

	for _, mod := range c.destinationModules {
		events = append(events, Event{pulse: pulseToSend, moduleTargeted: mod, moduleSource: c})
	}

	return events
}  

func (c *Conjunction) AddDestModules(modules []Module) {
	c.destinationModules = modules
}

func (c *Conjunction) AddInputModule(module Module) {
	c.inputModulesPulses[module] = false
}

type Flipflop struct {
	name               string
	onState            bool
	destinationModules []Module
}

func (f *Flipflop) ReceivePulse(highPulse bool) []Event {
	var events = []Event{}

	if highPulse { // nothing happens
		return events
	}

	var pulseToSend bool
	if !f.onState {
		pulseToSend = true
	}

	for _, mod := range f.destinationModules {
		events = append(events, Event{pulse: pulseToSend, moduleTargeted: mod, moduleSource: f})
	}

	f.onState = !f.onState

	return events
}

func (f *Flipflop) AddDestModules(modules []Module) {
	f.destinationModules = modules
}
