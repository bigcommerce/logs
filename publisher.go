package logs

type Publisher []Subscriber

func (p Publisher) Log(event *Event) {
	for _, subscriber := range p {
		subscriber.Log(event)
	}
}

func NewPublisher(subscribers ...Subscriber) Publisher {
	return subscribers
}

func (p *Publisher) Add(subscriber Subscriber) {
	*p = append(*p, subscriber)
}