package actions

import (
	"encoding/json"
	"testing"

	keptnevents "github.com/keptn/go-utils/pkg/events"
	"github.com/stretchr/testify/assert"
)

const virtualService = `apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  creationTimestamp: null
  name: carts
spec:
  gateways:
  - sockshop-production-gateway.sockshop-production
  - mesh
  hosts:
  - carts.sockshop-production.1.1.1.1.xip.io
  - carts
  http:
  - route:
    - destination:
        host: carts-canary.sockshop-production.svc.cluster.local
    - destination:
        host: carts-primary.sockshop-production.svc.cluster.local
      weight: 100
`

func TestGetIP(t *testing.T) {

	const problemData = `{
        "Tags":"environment:testenv",
        "ProblemID":"1337",
        "PID":"18416",
        "ProblemTitle":"CMD Injection",
        "ProblemDetails":{ "ClientIP":"127.0.0.1" }
    }`

	problem := keptnevents.ProblemEventData{}
	err := json.Unmarshal([]byte(problemData), &problem)
	assert.Nil(t, err)

	assert.Equal(t, "environment:testenv", problem.Tags)
	assert.Equal(t, "1337", problem.ProblemID)
	assert.Equal(t, "18416", problem.PID)
	assert.Equal(t, "CMD Injection", problem.ProblemTitle)

	ip, err := getIP(&problem)
	assert.Nil(t, err)
	assert.Equal(t, "127.0.0.1", ip)
}
