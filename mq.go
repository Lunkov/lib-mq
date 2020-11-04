package mq

import (
  "github.com/golang/glog"
  "github.com/nats-io/nats.go"
)

type NatsInfo struct {
  url    string
  ncNats *nats.Conn
  ecNats *nats.EncodedConn
}


func (n *NatsInfo) NatsInit(NatsUrl string) bool {
  if glog.V(2) {
    glog.Infof("LOG: InitNats(%s)", NatsUrl)
  }
  var err error
  n.url = NatsUrl
  natInit := false
  if n.ncNats != nil && n.ncNats.Status() == nats.CONNECTED {
    natInit = true
  } else {
    n.ncNats, err = nats.Connect(NatsUrl, nats.MaxReconnects(-1))
    if err != nil {
      glog.Errorf("ERR: NATS Connect(%s): %v", NatsUrl, err)
    } else {
      n.ecNats, err = nats.NewEncodedConn(n.ncNats, nats.JSON_ENCODER)
      if err != nil {
        glog.Errorf("ERR: NATS NewEncodedConn: %v", err)
      } else {
        natInit = true
      }
    }
  }
  return natInit
}

func (n *NatsInfo) NatsConnected() bool {
  return n.ncNats != nil && n.ecNats != nil
}

func (n *NatsInfo) NatsClose() {
  if n.ncNats != nil {
    n.ncNats.Close()
  }
}

func (n *NatsInfo) NatsSendMsg(subject string, data interface{}) bool {
  if !n.NatsConnected() {
    n.NatsInit(n.url) 
  }
  if n.NatsConnected() {
    if glog.V(9) {
      glog.Infof("LOG: MQ: sendNatsMsg(%s) data='%v'", subject, data)
    }
    err := n.ecNats.Publish(subject, data)
    if err == nil {
      return true
    }
    glog.Errorf("ERR: MQ: sendNatsMsg(%s) err=%v", subject, err)
    n.NatsClose()
    n.NatsInit(n.url) 
  } else {
    glog.Errorf("ERR: NATS not connected: for sendNatsMsg(%s)", subject)
  }
  return false
}

