package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type cliente struct {
	ID          string
	sendMessage chan EventMessage
}
type EventMessage struct {
	EventName string
	Data      any
}

func NewCliente(id string) *cliente {
	return &cliente{
		ID: id,
	}
}

type HandlerEvent struct {
	m        sync.Mutex
	clientes map[string]*cliente
}

func NewHandlerEvent() *HandlerEvent {
	return &HandlerEvent{
		clientes: make(map[string]*cliente),
	}
}

func (h *HandlerEvent) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	id := r.URL.Query().Get("id")

	flusher, ok := w.(http.Flusher)
	if !ok {
		//Error 400
		return
	}
	c := NewCliente(id)

	h.register(c)

	c.Online(r.Context(), w, flusher)

	h.remove(c)
}

func (h *HandlerEvent) register(cliente *cliente) {
	h.m.Lock()
	defer h.m.Unlock()
	h.clientes[cliente.ID] = cliente
}

func (h *HandlerEvent) remove(cliente *cliente) {
	h.m.Lock()
	delete(h.clientes, cliente.ID)
	defer h.m.Unlock()

}
func (h *HandlerEvent) Broadcast(m EventMessage) {
	h.m.Lock()
	defer h.m.Unlock()
	for _, cliente := range h.clientes {
		cliente.sendMessage <- m
	}
}

func (c *cliente) Online(ctx context.Context, w io.Writer, flusher http.Flusher) {
	for {
		select {
		case m := <-c.sendMessage:
			data, err := json.Marshal(m.Data)
			if err != nil {
				log.Println(err)
			}
			const format = "event:%s\ndata:%s\n\n"
			fmt.Fprintf(w, format, m.EventName, string(data))
			flusher.Flush()

		case <-ctx.Done():
			return
		}
	}
}
