// Copyright 2016 Anjieych. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//Usage:
//
//	//Send
//	if err := utils.NewActiveMQ("localhost:61613").Send("/queue/test-1", "test from 1"); err != nil {
//	fmt.Println("AMQ ERROR:", err)
//	
//  //this func will handle the messges get from activeMQ server.
//	handler := func(err error, msg string) { fmt.Println("AMQ MSG:", err, msg) }
//	if err := utils.NewActiveMQ("localhost:61613").Subscribe("/queue/test-1", handler); err != nil {
//	  fmt.Println("AMQ ERROR:", err)
//	}
//
package activeMQ

import (
	stomp "github.com/go-stomp/stomp"
)

type ActiveMQ struct {
	Addr string
}

//New activeMQ with addr[eg:localhost:61613] as host address.
func NewActiveMQ(addr string) *ActiveMQ {
	if addr == "" {
		addr = "localhost:61613"
	}
	return &ActiveMQ{addr}
}

// Used for health check
func (this *ActiveMQ) Check() error {
	conn, err := this.Connect()
	if err == nil {
		defer conn.Disconnect()
		return nil
	} else {
		return err
	}
}

// Connect to activeMQ
func (this *ActiveMQ) Connect() (*stomp.Conn, error) {
	return stomp.Dial("tcp", this.Addr)
}

// Send msg to destination
func (this *ActiveMQ) Send(destination string, msg string) error {
	conn, err := this.Connect()
	if err != nil {
		return err
	}
	defer conn.Disconnect()
	return conn.Send(
		destination,  // destination
		"text/plain", // content-type
		[]byte(msg))  // body
}

// Subscribe Message from destination
// func handler handle msg reveived from destination
func (this *ActiveMQ) Subscribe(destination string, handler func(err error, msg string)) error {

	conn, err := this.Connect()
	if err != nil {
		return err
	}
	sub, err := conn.Subscribe(destination, stomp.AckAuto)
	if err != nil {
		return err
	}
	defer conn.Disconnect()
	defer sub.Unsubscribe()
	for {
		m := <-sub.C
		handler(m.Err, string(m.Body))
	}
	return err
}
