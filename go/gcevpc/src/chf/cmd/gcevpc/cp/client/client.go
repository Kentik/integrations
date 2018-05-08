package client

import (
	"fmt"

	"github.com/kentik/libkflow"
	"github.com/kentik/libkflow/api"
)

type FlowClient struct {
	Sender          *libkflow.Sender
	SetSrcHostTags  map[string]bool
	SetDestHostTags map[string]bool
	interfaces      map[string]api.InterfaceUpdate
	idsByAlias      map[string]uint32
	doneInit        bool
	nextInterface   uint32
}

func NewFlowClient(client *libkflow.Sender) *FlowClient {
	return &FlowClient{
		Sender:          client,
		SetSrcHostTags:  map[string]bool{},
		SetDestHostTags: map[string]bool{},
		idsByAlias: map[string]uint32{
			"": 1,
		},
		interfaces: map[string]api.InterfaceUpdate{
			"eth0": api.InterfaceUpdate{ // Pre-populate this with eth0 for now.
				Index:   1,
				Desc:    "eth0",
				Alias:   "",
				Address: "127.0.0.1",
			},
		},
		nextInterface: 2,
	}
}

func (c *FlowClient) ResetTags() {
	c.SetSrcHostTags = map[string]bool{}
	c.SetDestHostTags = map[string]bool{}
}

func (c *FlowClient) GetInterfaceID(host string) uint32 {
	return c.idsByAlias[host]
}

func (c *FlowClient) UpdateInterfaces(isFromInterfaceUpdate bool) error {

	// Only run from not interfaces once
	if c.doneInit && !isFromInterfaceUpdate {
		return nil
	}
	c.doneInit = true

	client := c.Sender.GetClient()
	if client != nil {
		err := client.UpdateInterfacesDirectly(c.Sender.Device, c.interfaces)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *FlowClient) AddInterface(intf *api.InterfaceUpdate) {
	intf.Index = uint64(c.nextInterface)
	c.idsByAlias[intf.Alias] = c.nextInterface
	intf.Desc = fmt.Sprintf("kentik.%d", c.nextInterface)
	c.nextInterface++

	c.interfaces[intf.Desc] = *intf
}