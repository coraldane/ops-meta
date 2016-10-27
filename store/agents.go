package store

import (
	"sync"
	"time"

	"github.com/toolkits/container/set"

	"github.com/coraldane/ops-common/model"
)

type AgentsMap struct {
	sync.RWMutex
	M map[string]*model.DesiredAgent
}

func NewAgentsMap() *AgentsMap {
	return &AgentsMap{M: make(map[string]*model.DesiredAgent)}
}

func (this *AgentsMap) Get(agentName string) (*model.DesiredAgent, bool) {
	this.RLock()
	defer this.RUnlock()
	val, exists := this.M[agentName]
	return val, exists
}

func (this *AgentsMap) Keys() []string {
	this.RLock()
	defer this.RUnlock()
	r := []string{}
	for key := range this.M {
		r = append(r, key)
	}
	return r
}

func (this *AgentsMap) Len() int {
	this.RLock()
	defer this.RUnlock()
	return len(this.M)
}

func (this *AgentsMap) Clean() {
	this.RLock()
	defer this.RUnlock()
	this.M = make(map[string]*model.DesiredAgent)
}

func (this *AgentsMap) Put(agentName string, desiredAgent *model.DesiredAgent) {
	this.Lock()
	defer this.Unlock()
	this.M[agentName] = desiredAgent
}

func (this *AgentsMap) Delete(key string) {
	this.Lock()
	defer this.Unlock()
	delete(this.M, key)
}

var DesiredAgentsMap = NewAgentsMap()

type HostAgentSet struct {
	Timestamp    int64
	AgentNameSet *set.SafeSet
}
type HostAgentsMap struct {
	sync.RWMutex
	M map[string]*HostAgentSet
}

func NewHostAgentsMap() *HostAgentsMap {
	return &HostAgentsMap{M: make(map[string]*HostAgentSet)}
}

var HostAgents = NewHostAgentsMap()

func (this *HostAgentsMap) Get(hostname string) (*HostAgentSet, bool) {
	this.RLock()
	defer this.RUnlock()
	has, exists := this.M[hostname]
	return has, exists
}

func (this *HostAgentsMap) Put(hostname string, agentNames []string) {
	this.Lock()
	defer this.Unlock()
	has, exists := this.M[hostname]
	if !exists {
		has = &HostAgentSet{Timestamp: time.Now().Unix(), AgentNameSet: set.NewSafeSet()}
	}
	for _, name := range agentNames {
		has.AgentNameSet.Add(name)
	}

	this.M[hostname] = has
}

func (this *HostAgentsMap) Delete(hostname string) {
	this.Lock()
	defer this.Unlock()
	delete(this.M, hostname)
}

func (this *HostAgentsMap) Keys() []string {
	this.RLock()
	defer this.RUnlock()
	r := []string{}
	for key := range this.M {
		r = append(r, key)
	}
	return r
}

func (this *HostAgentsMap) Status(hostname string) []*model.DesiredAgent {
	ret := []*model.DesiredAgent{}
	has, exists := this.M[hostname]
	if exists {
		agentNames := has.AgentNameSet.ToSlice()
		for _, name := range agentNames {
			if da, ok := DesiredAgentsMap.Get(name); ok {
				ret = append(ret, da)
			}
		}
	}
	return ret
}
