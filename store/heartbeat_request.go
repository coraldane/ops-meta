package store

import (
	"time"

	cmodel "github.com/coraldane/ops-common/model"
	"github.com/coraldane/ops-meta/models"
)

func ParseHeartbeatRequest(req *cmodel.HeartbeatRequest, strIp string) {
	//保存Endpoint信息
	convertHeartbeatReq2Endpoint(req, strIp).Insert()

	//保存Endpoint对应的Agents
	saveEndpointAgents(req)
}

func convertHeartbeatReq2Endpoint(req *cmodel.HeartbeatRequest, strIp string) *models.Endpoint {
	endpoint := models.Endpoint{}
	endpoint.Hostname = req.Hostname
	if "" == req.Ip {
		endpoint.Ip = strIp
	} else {
		endpoint.Ip = req.Ip
	}
	endpoint.UpdaterVersion = req.UpdaterVersion
	endpoint.RunUser = req.RunUser
	return &endpoint
}

func saveEndpointAgents(req *cmodel.HeartbeatRequest) {
	ea := models.EndpointAgent{Hostname: req.Hostname}
	var excludeAgents []string

	if req.RealAgents == nil || len(req.RealAgents) == 0 {
		ea.Delete(excludeAgents)
	} else {
		for _, agent := range req.RealAgents {
			excludeAgents = append(excludeAgents, agent.Name)

			endpointAgent := models.EndpointAgent{Hostname: req.Hostname}
			endpointAgent.GmtModified = time.Now()
			endpointAgent.AgentName = agent.Name
			endpointAgent.AgentVersion = agent.Version
			endpointAgent.RunUser = agent.RunUser
			endpointAgent.WorkDir = agent.WorkDir
			endpointAgent.Status = agent.Status
			endpointAgent.GmtReport = time.Unix(agent.Timestamp, 0)
			endpointAgent.Insert()
		}

		ea.Delete(excludeAgents)
	}
}
