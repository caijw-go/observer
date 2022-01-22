package observer

import "fmt"

type Listener interface {
    Listen() []interface{}
    Process(event interface{})
}

var eventListenerMap = make(map[string][]Listener)

func getEventName(event interface{}) string {
    return fmt.Sprintf("%T", event)
}

func RegisterListener(listeners ...Listener) {
    for _, listener := range listeners {
        for _, event := range listener.Listen() {
            eventName := getEventName(event)
            if _, ok := eventListenerMap[eventName]; !ok {
                eventListenerMap[eventName] = []Listener{listener}
            } else {
                eventListenerMap[eventName] = append(eventListenerMap[eventName], listener)
            }
        }
    }
}
