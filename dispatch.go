package observer

func Dispatch(event interface{}) {
    for _, listener := range eventListenerMap[getEventName(event)] {
        listener.Process(event)
    }
}
