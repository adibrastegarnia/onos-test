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

package testapi

import (
	"os"
	"testing"

	"github.com/onosproject/onos-test/pkg/onit/setup"
	"gotest.tools/assert"

	"github.com/onosproject/onos-test/pkg/runner"
	"github.com/onosproject/onos-test/test"
)

func init() {
	test.Registry.RegisterTest("add-simulator", addSimulator, []*runner.TestSuite{})
}

// TestSetup tests k8s setup interface
func addSimulator(t *testing.T) {
	testSetupBuilder := setup.New()
	clusterID := os.Getenv("TEST_NAMESPACE")
	testSetupBuilder.SetClusterID(clusterID)
	testSetupBuilder.SetSimulatorName("simulator-1")
	testSetup := testSetupBuilder.Build()
	testSetup.AddSimulator()
	simulators, _ := testSetup.GetSimulators()
	assert.Equal(t, len(simulators), 1)
	testSetup.RemoveSimulator()
}