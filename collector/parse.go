package collector

import (
	"log"
	"strconv"

	"github.com/docker/docker/api/types/swarm"

	"github.com/tpbowden/swarm-ingress-router/types"
)

func parseServices(swarmServices []swarm.Service) []types.Service {
	result := []types.Service{}
	for _, s := range swarmServices {
		name := s.Spec.Annotations.Name
		port, err := strconv.Atoi(s.Spec.Annotations.Labels["ingress.targetport"])
		if err != nil {
			log.Printf("Failed to parse port for service %s: %v", name, err)
			continue
		}

		secure := s.Spec.Annotations.Labels["ingress.tls"] == "true"
		forceTLS := s.Spec.Annotations.Labels["ingress.forcetls"] == "true"
		cert := s.Spec.Annotations.Labels["ingress.cert"]
		key := s.Spec.Annotations.Labels["ingress.key"]

		result = append(result, types.Service{
			Name:        name,
			Port:        port,
			Secure:      secure,
			ForceTLS:    forceTLS,
			Certificate: cert,
			Key:         key,
		})
	}
	return result
}