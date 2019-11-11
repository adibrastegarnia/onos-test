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

package topo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/onosproject/onos-test/pkg/new/onit"
	"github.com/onosproject/onos-topo/pkg/northbound/device"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TopoBenchmarkSuite is a benchmark suite for the topo service
type TopoBenchmarkSuite struct {
	onit.BenchmarkSuite
}

func (s *TopoBenchmarkSuite) SetupBenchmarkSuite() {
	setup := s.Setup()
	setup.Database().
		Partitions(3).
		Nodes(3)
	setup.Topo().Nodes(2)
	setup.SetupOrDie()
}

func (s *TopoBenchmarkSuite) BenchmarkDeviceService(b *testing.B) {
	env := s.Env()
	conn, err := env.Topo().Connect()
	assert.NoError(b, err)
	defer conn.Close()
	client := device.NewDeviceServiceClient(conn)
	for i := 0; i < b.N; i++ {
		id := uuid.New().String()
		_, _ = client.Add(context.Background(), &device.AddRequest{
			Device: &device.Device{
				ID:      device.ID(id),
				Address: fmt.Sprintf("%s:5150", id),
				Version: "1.0.0",
				Type:    "stratum",
			},
		})
	}
}
