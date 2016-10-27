package store

import (
	"math"
	"regexp"
	"time"

	"github.com/toolkits/container/set"

	cmodel "github.com/coraldane/ops-common/model"
	"github.com/coraldane/ops-meta/g"
	"github.com/coraldane/ops-meta/models"
)

func RefreshDesiredAgent() {
	d := time.Duration(g.Config().RefreshInterval) * time.Second
	for {
		refreshDesiredAgent()
		time.Sleep(d)
	}
}

func refreshDesiredAgent() {
	pageInfo := &models.PageInfo{PageIndex: 1, PageSize: math.MaxInt32}
	agentList, _ := models.QueryAgentList(models.QueryAgentDto{}, pageInfo)
	if nil == agentList || 0 == len(agentList) {
		return
	}

	refreshAgentsMap(agentList)

	hostGroupAgentMap := make(map[int64]*set.SafeSet)
	usedHostGroupIds := set.NewSafeInt64Set()

	//获取Agent关联到的HostGroupId
	for _, agent := range agentList {
		relAgentGroups, _ := models.QueryRelAgentGroupList(agent.Id)
		if nil == relAgentGroups || 0 == len(relAgentGroups) {
			continue
		}

		for _, rag := range relAgentGroups {
			usedHostGroupIds.Add(rag.HostGroupId)
			nameSet, exists := hostGroupAgentMap[rag.HostGroupId]
			if !exists {
				nameSet = set.NewSafeSet()
			}
			nameSet.Add(agent.Name)
			hostGroupAgentMap[rag.HostGroupId] = nameSet
		}
	}

	//获取每个HostGroup下属Endpoint
	endpointList, _ := models.QueryEndpointList(models.QueryEndpointDto{}, pageInfo)
	if nil == endpointList || 0 == len(endpointList) {
		return
	}
	for _, hostGroupId := range usedHostGroupIds.Slice() {
		relEndpointGroupList, _ := models.QueryRelEndpointGroupList(models.QueryRelEndpointGroupDto{HostGroupId: hostGroupId}, pageInfo)
		if nil == relEndpointGroupList || 0 == len(relEndpointGroupList) {
			continue
		}
		for _, reg := range relEndpointGroupList {
			if "fixed" == reg.RelType && "hostname" == reg.PropName {
				HostAgents.Put(reg.PropValue, hostGroupAgentMap[hostGroupId].ToSlice())
			} else {
				pattern := regexp.MustCompile(reg.PropValue)
				for _, endpoint := range endpointList {
					if "hostname" == reg.PropName {
						if pattern.MatchString(endpoint.Hostname) {
							HostAgents.Put(endpoint.Hostname, hostGroupAgentMap[hostGroupId].ToSlice())
						}
					} else if "ip" == reg.PropValue {
						if pattern.MatchString(endpoint.Ip) {
							HostAgents.Put(endpoint.Hostname, hostGroupAgentMap[hostGroupId].ToSlice())
						}
					}
				}
			}
		}
	}
}

func refreshAgentsMap(agentList []models.Agent) {
	keys := DesiredAgentsMap.Keys()
	for _, key := range keys {
		exists := false
		for _, agent := range agentList {
			if agent.Name == key {
				exists = true
				break
			}
		}

		if !exists {
			DesiredAgentsMap.Delete(key)
		}
	}

	for _, agent := range agentList {
		da := cmodel.DesiredAgent{}
		da.Name = agent.Name
		da.Version = agent.Version
		da.Tarball = agent.Tarball
		da.Md5 = agent.Md5
		da.Cmd = agent.Cmd
		da.RunUser = agent.RunUser
		da.WorkDir = agent.WorkDir
		da.ConfigFileName = agent.ConfigFileName
		da.ConfigRemoteUrl = agent.ConfigRemoteUrl
		DesiredAgentsMap.Put(agent.Name, &da)
	}
}
