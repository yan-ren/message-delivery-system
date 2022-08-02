package main

import (
	"io"
	"log"
	"message-delivery-system/proto/pb"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	port := ":8080"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to Listen: %v", err)
	}

	server := NewServer()

	s := grpc.NewServer()

	pb.RegisterDeliveryServiceServer(s, server)

	log.Println("Starting Server On Port:" + port)

	e := s.Serve(lis)
	if e != nil {
		log.Fatalf("failed to Serve: %v", e)
	}
}

type Server struct {
	mutex   *sync.Mutex
	clients map[pb.DeliveryService_DeliveryServer]uint64
}

func (s *Server) assignUserId() uint64 {
	return uint64(time.Now().Unix())
}

func NewServer() *Server {
	return &Server{mutex: &sync.Mutex{}, clients: make(map[pb.DeliveryService_DeliveryServer]uint64)}
}

func (s *Server) HandleList(m pb.DeliveryService_DeliveryServer, req *pb.Request) {
	// if the stream is new, store the stream
	s.mutex.Lock()
	if userId, contained := s.clients[m]; !contained {
		userId = s.assignUserId()
		s.clients[m] = userId
	}
	s.mutex.Unlock()

	users := []uint64{}
	s.mutex.Lock()
	for k, v := range s.clients {
		// excluding the requesting client
		if k != m {
			users = append(users, v)
		}
	}
	s.mutex.Unlock()

	// send it to stream
	resp := pb.Response{Users: users}
	if err := m.Send(&resp); err != nil {
		log.Println(err.Error())
	}
}

func (s *Server) HandleIdentity(m pb.DeliveryService_DeliveryServer, req *pb.Request) {
	// if the stream is new, store the stream
	s.mutex.Lock()
	var userId uint64
	var contained bool
	userId, contained = s.clients[m]
	if !contained {
		userId = s.assignUserId()
		s.clients[m] = userId
	}
	s.mutex.Unlock()

	resp := pb.Response{Users: []uint64{userId}, Message: "register!"}
	if err := m.Send(&resp); err != nil {
		log.Println(err.Error())
	}
}

func (s *Server) HandleRelay(m pb.DeliveryService_DeliveryServer, req *pb.Request) {
	s.mutex.Lock()
	if userId, contained := s.clients[m]; !contained {
		userId = s.assignUserId()
		s.clients[m] = userId
	}
	s.mutex.Unlock()

	// no user provided in relay message
	if len(req.Users) == 0 {
		log.Printf("no user exists")
		return
	}
	for k, v := range s.clients {
		for _, UserIDs := range req.Users {
			if v == UserIDs {
				// send it to stream
				resp := pb.Response{Message: req.Message}
				if err := k.Send(&resp); err != nil {
					log.Println(err.Error())
				}
			}
		}
	}
}

// Delivery Listens for stream from the server
func (s *Server) Delivery(m pb.DeliveryService_DeliveryServer) error {
	ctx := m.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := m.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			delete(s.clients, m)
			continue
		}

		switch req.Type {
		case "relay":
			s.HandleRelay(m, req)
		case "identity":
			s.HandleIdentity(m, req)
		case "list":
			s.HandleList(m, req)
		}
	}
}
