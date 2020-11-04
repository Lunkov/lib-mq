package mq

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCheckMQ(t *testing.T) {
  var ns NatsInfo

  ok_conn := ns.NatsInit("nats://localhost:1001")
  
  assert.Equal(t, false, ok_conn)

  ok_send := ns.NatsSendMsg("device.new.metric11", "test message")
  assert.Equal(t, false, ok_conn)
  assert.Equal(t, false, ns.NatsConnected())
  
  
  ok_conn = ns.NatsInit("nats://localhost:4222")
  assert.Equal(t, true, ok_conn)

  ok_conn = ns.NatsInit("nats://localhost:4222")
  assert.Equal(t, true, ok_conn)

  assert.Equal(t, true, ns.NatsConnected())

  ok_send = ns.NatsSendMsg("device.new.metric11", "test message")
  assert.Equal(t, true, ok_send)
  
  ns.NatsClose()

  ok_send = ns.NatsSendMsg("device.new.metric11", "test message")
  assert.Equal(t, false, ok_send)

}
