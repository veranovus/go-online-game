package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
)

type MessageEvent struct {
	msg string
}

type UserJoinedEvent struct {
}

type ClientInput struct {
	user  *User
	event interface{}
}

type User struct {
	name    string
	session *Session
}

type Session struct {
	conn net.Conn
}

func (s *Session) WriteLine(str string) error {
	_, err := s.conn.Write([]byte(str + "\r\n"))
	return err
}

type World struct {
	users []*User
}

func generateName(world *World) string {

	var name string
	validName := false

	for !validName {

		name = fmt.Sprintf("USER-%06d", rand.Intn(100000)+1)
		validName = true

		for _, user := range world.users {
			if user.name == name {
				validName = false
			}
		}
	}

	return name
}

func handleConnection(w *World, conn net.Conn, inputChannel chan ClientInput) error {
	log.Println("Successfully connected to server.")

	buff := make([]byte, 4096)

	session := &Session{conn}
	user := &User{generateName(w), session}

	inputChannel <- ClientInput{
		user,
		&UserJoinedEvent{},
	}

	for {
		n, err := conn.Read(buff)
		if err != nil {
			return err
		}

		if n < 2 { // Old code was `n == 0`
			log.Println("Zero bytes. Closing the connection.")
			return nil
		}

		msg := string(buff)
		msg = msg[:n-2]

		e := ClientInput{user, &MessageEvent{msg}}
		inputChannel <- e
	}
}

func startServer(w *World, eventChannel chan ClientInput) error {
	log.Println("Starting the server.")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Failed to accept the connection", err)
			continue
		}

		go func() {
			if err := handleConnection(w, conn, eventChannel); err != nil {
				log.Println("Error while handling the connection", err)
				return
			}
		}()
	}
}

func startGameLoop(w *World, clientInputChannel <-chan ClientInput) {
	for input := range clientInputChannel {
		switch event := input.event.(type) {
		case *MessageEvent:
			log.Println("Received message event:", event.msg)

			resp := fmt.Sprintf("[You] %s", event.msg)
			_ = input.user.session.WriteLine(resp)

			for _, user := range w.users {
				if user == input.user {
					continue
				}
				resp = fmt.Sprintf("[%s] %s", input.user.name, event.msg)
				_ = user.session.WriteLine(resp)
			}
		case *UserJoinedEvent:
			log.Println("User joined:", input.user.name)
			w.users = append(w.users, input.user)

			resp := fmt.Sprintf("Welcome '%s'", input.user.name)
			_ = input.user.session.WriteLine(resp)
		default:
			log.Println("Unknown event:", event)
		}
	}
}

func main() {

	ch := make(chan ClientInput)
	w := &World{}

	go startGameLoop(w, ch)

	err := startServer(w, ch)
	if err != nil {
		log.Fatal(err)
	}
}
