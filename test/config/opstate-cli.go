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

package config

import (
	"github.com/onosproject/onos-test/pkg/runner"
	"github.com/onosproject/onos-test/test"
	"github.com/onosproject/onos-test/test/env"
	"github.com/onosproject/onos-test/test/integration"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type testCase struct {
	path          string
	expectedValue string
}

func init() {
	test.Registry.RegisterTest("config-opstate-cli", TestConfigGetCLI, []*runner.TestSuite{ConfigTests, integration.IntegrationTests})
}

func makeDescription(path string) string {
	if len(path) <= 25 {
		return path
	}
	return "..." + path[len(path)-22:]
}

func parseOpstateCommandOutput(t *testing.T, output []string) map[string]string {
	t.Helper()
	combinedLines := make([]string, 0)
	thisLine := ""
	for _, line := range output {
		if strings.HasPrefix(line, "OPSTATE CACHE") ||
			strings.HasPrefix(line, "PATH") {
			//  Skip column headers
			continue
		}
		line = strings.TrimPrefix(line, "  ")
		if strings.Contains(line, "|") {
			combinedLines = append(combinedLines, thisLine+line)
			thisLine = ""
		} else {
			thisLine = thisLine + line
		}
	}

	opState := make(map[string]string)
	for _, combinedLine := range combinedLines {
		tokens := strings.Split(strings.ReplaceAll(combinedLine, " ", ""), "|")
		opState[tokens[0]] = tokens[1]
	}

	return opState
}

// TestConfigGetCLI tests the topo service's device CLI commands
func TestConfigGetCLI(t *testing.T) {
	device1 := env.GetDevices()[0]

	output, code := env.ExecuteCLI("onos config get opstate " + device1)
	assert.Equal(t, 0, code)

	opState := parseOpstateCommandOutput(t, output)

	testCases := []testCase{
		{
			path:          "/system/openflow/controllers/controller[name=second]/connections/connection[aux-id=1]/state/address",
			expectedValue: "(STRING)192.0.3.11",
		},
		{
			path:          "/system/openflow/controllers/controller[name=main]/connections/connection[aux-id=1]/state/transport",
			expectedValue: "(STRING)TLS",
		},
		{
			path:          "/system/openflow/controllers/controller[name=main]/connections/connection[aux-id=1]/state/priority",
			expectedValue: "(UINT)2",
		},
		{
			path:          "/system/openflow/controllers/controller[name=main]/connections/connection[aux-id=0]/state/source-interface",
			expectedValue: "(STRING)admin",
		},
		{
			path:          "/system/openflow/controllers/controller[name=main]/connections/connection[aux-id=0]/state/port",
			expectedValue: "(STRING)6633",
		},
		{
			path:          "/system/openflow/controllers/controller[name=second]/connections/connection[aux-id=0]/state/aux-id",
			expectedValue: "(STRING)0",
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		description := makeDescription(testCase.path)
		t.Run(description,
			func(t *testing.T) {
				path := testCase.path
				expectedValue := testCase.expectedValue
				t.Parallel()
				assert.Equal(t, expectedValue, opState[path])
			})
	}
}
