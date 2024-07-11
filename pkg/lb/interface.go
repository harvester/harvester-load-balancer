package lb

import (
	"errors"

	"k8s.io/apimachinery/pkg/types"

	lbv1 "github.com/harvester/harvester-load-balancer/pkg/apis/loadbalancer.harvesterhci.io/v1beta1"
)

type HealthCheckHandler func(namespace, name string) error

type Manager interface {
	// Step 1. Ensure loadbalancer
	EnsureLoadBalancer(lb *lbv1.LoadBalancer) error
	DeleteLoadBalancer(lb *lbv1.LoadBalancer) error

	// Step 2. Ensure loadbalancer external IP
	EnsureLoadBalancerServiceIP(lb *lbv1.LoadBalancer) (string, error)

	// Step 3. Ensure service backend servers
	EnsureBackendServers(lb *lbv1.LoadBalancer) ([]BackendServer, error)

	ListBackendServers(lb *lbv1.LoadBalancer) ([]BackendServer, error)

	// return the count of endpoints which are probed as Ready
	// if probe is disabled, then return the count of all endpoints
	GetProbeReadyBackendServerCount(lb *lbv1.LoadBalancer) (int, error)

	// register a handler to get which lb is happending changes per health check
	RegisterHealthCheckHandler(handler HealthCheckHandler) error
}

type BackendServer interface {
	GetUID() types.UID
	GetNamespace() string
	GetName() string
	GetAddress() (string, bool)
}

var (
	ErrWaitExternalIP = errors.New("service is waiting for external IP")
)
