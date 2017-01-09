package main

//type function struct {
//	// join is a channel for containers wishing to join the function
//	join chan *container
//}
//
//// newFunction creates a new function that is ready to go.
//func newFunction() *function {
//	return &function{
//		join:       make(chan *container),
//		leave:      make(chan *container),
//		forward:    make(chan []byte),
//		containers: make(make[*container]bool),
//		tracer:     trace.Off(),
//	}
//}
//
//func (f *function) run() {
//	for {
//		select {
//		case container := <-f.join:
//			// joining
//			f.containers[container] = true
//			f.tracer.Trace("New container joined")
//		case container := <-f.leave:
//			// leaving
//			delete(f.containers, container)
//			close(container.send)
//			f.tracer.Trace("Container left")
//		case msg := <-f.forward:
//			f.tracer.Trace("Message received: ", string(msg))
//			// forward message to all containers
//			for container := range f.containers {
//				container.send <- msg
//				f.tracer.Trace(" -- sent to client")
//			}
//		}
//	}
//}
