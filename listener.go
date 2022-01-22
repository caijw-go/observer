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
            eventListenerMap[eventName] = append(eventListenerMap[eventName], listener)
        }
    }
}
