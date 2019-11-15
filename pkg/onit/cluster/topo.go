// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import (
	"time"
)

const (
	topoType    = "topo"
	topoImage   = "onosproject/onos-topo:latest"
	topoService = "onos-topo"
	topoPort    = 5150
	topoAddress = "onos-topo:5150"
	topoTimeout = 30 * time.Second
)

var topoSecrets = map[string]string{
	"onf.cacrt":     caCert,
	"onos-topo.crt": topoCert,
	"onos-topo.key": topoKey,
}

var topoArgs = []string{
	"-caPath=/certs/onf.cacrt",
	"-keyPath=/certs/onos-topo.key",
	"-certPath=/certs/onos-topo.crt",
}

func newTopo(client *client) *Topo {
	return &Topo{
		Service: newService(topoService, topoPort, getLabels(topoType), topoImage, topoSecrets, topoArgs, client),
	}
}

// Topo provides methods for managing the onos-topo service
type Topo struct {
	*Service
}