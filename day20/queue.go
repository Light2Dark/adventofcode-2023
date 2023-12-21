package day20

type Queue struct {
	events []Event
}

type Event struct {
	pulse bool // low is false, high is true
	moduleTargeted Module
	moduleSource Module
}


func (q *Queue) enqueue(events ...Event) {
	q.events = append(q.events, events...)
}

func (q *Queue) dequeue() (Event, bool) {
	if q.length() == 0 {
		return Event{}, false
	}
	elem := q.events[0]
	if q.length() == 1 {
		q.events = []Event{}
	} else {
		q.events = q.events[1:]
	}
	return elem, true
}

func (q *Queue) length() int {
	return len(q.events)
}


