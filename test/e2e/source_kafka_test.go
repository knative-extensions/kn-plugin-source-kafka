// Copyright 2020 The Knative Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build e2e && !serving
// +build e2e,!serving

package e2e

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	testcommon "github.com/maximilien/kn-source-pkg/test/e2e"
	"gotest.tools/v3/assert"
	"knative.dev/client-pkg/pkg/util"
	"knative.dev/client-pkg/pkg/util/lib/test"
)

const (
	kafkaBootstrapUrl     = "my-cluster-kafka-bootstrap.kafka.svc:9092"
	kafkaClusterName      = "my-cluster"
	kafkaClusterNamespace = "kafka"
	kafkaTopic            = "test-topic"
	ceo                   = "type=foo"
)

type e2eTest struct {
	it *testcommon.E2ETest
}

func newE2ETest() *e2eTest {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil
	}

	it, err := testcommon.NewE2ETest("kn-source-kafka", filepath.Join(currentDir, "../.."), false)
	if err != nil {
		return nil
	}

	e2eTest := &e2eTest{
		it: it,
	}
	return e2eTest
}

func TestSourceKafka(t *testing.T) {
	t.Parallel()

	e2eTest := newE2ETest()
	assert.Assert(t, e2eTest != nil)
	defer func() {
		assert.NilError(t, e2eTest.it.KnTest().Teardown())
	}()

	r := test.NewKnRunResultCollector(t, e2eTest.it.KnTest())
	defer r.DumpIfFailed()

	err := e2eTest.it.KnPlugin().Install()
	assert.NilError(t, err)
	defer func() {
		err = e2eTest.it.KnPlugin().Uninstall()
		assert.NilError(t, err)
	}()

	serviceCreate(r, "sinksvc")

	for name, tc := range map[string]struct {
		name        string
		labels      map[string]string
		annotations map[string]string
	}{
		"test-source-kafka": {
			name: "mykafka1",
		},
		"test-source-kafka-labels": {
			name: "mykafka-labels",
			labels: map[string]string{
				"app":  "e2e",
				"role": "labels",
			},
		},
		"test-source-kafka-annotations": {
			name: "mykafka-annotations",
			annotations: map[string]string{
				"app":  "e2e",
				"role": "annotations",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Logf("test kn-plugin-source-kafka create %s", tc.name)
			e2eTest.knSourceKafkaCreate(t, r, tc.name, "sinksvc", tc.labels, tc.annotations)

			t.Logf("test kn-plugin-source-kafka describe %s", tc.name)
			e2eTest.knSourceKafkaDescribe(t, r, tc.name, "sinksvc", "cloudevent")

			t.Log("test kn-plugin-source-kafka list")
			e2eTest.knSourceKafkaList(t, r, tc.name)

			t.Logf("test kn-plugin-source-kafka delete %s", tc.name)
			e2eTest.knSourceKafkaDelete(t, r, tc.name)
		})
	}
}

// Private

func (et *e2eTest) knSourceKafkaCreate(t *testing.T, r *test.KnRunResultCollector, sourceName, sinkName string, labels map[string]string, annotations map[string]string) {
	flags := []string{"create", sourceName,
		"--servers", kafkaBootstrapUrl,
		"--topics", kafkaTopic,
		"--consumergroup", "test-consumer-group",
		"--sink", sinkName,
		"--ce-override", ceo,
	}
	if labels != nil {
		for k, v := range labels {
			flags = append(flags, "--label", fmt.Sprintf("%s=%s", k, v))
		}
	}
	if annotations != nil {
		for k, v := range annotations {
			flags = append(flags, "--annotation", fmt.Sprintf("%s=%s", k, v))
		}
	}
	out := et.it.KnPlugin().Run(flags...)
	r.AssertNoError(out)
	assert.Check(t, util.ContainsAllIgnoreCase(out.Stdout, "create", sourceName))
}

func (et *e2eTest) knSourceKafkaDelete(t *testing.T, r *test.KnRunResultCollector, sourceName string) {
	out := et.it.KnPlugin().Run("delete", sourceName)
	r.AssertNoError(out)
	assert.Check(t, util.ContainsAllIgnoreCase(out.Stdout, "delete", sourceName))
}

func (et *e2eTest) knSourceKafkaDescribe(t *testing.T, r *test.KnRunResultCollector, sourceName, sinkName, cloudEvent string) {
	out := et.it.KnPlugin().Run("describe", sourceName)
	r.AssertNoError(out)
	assert.Check(t, util.ContainsAllIgnoreCase(out.Stdout, sourceName, sinkName, cloudEvent))
}

func serviceCreate(r *test.KnRunResultCollector, serviceName string) {
	out := r.KnTest().Kn().Run("service", "create", serviceName, "--image", "gcr.io/knative-samples/helloworld-go")
	r.AssertNoError(out)
	assert.Check(r.T(), util.ContainsAllIgnoreCase(out.Stdout, "service", serviceName, "creating", "namespace", r.KnTest().Kn().Namespace(), "ready"))
}

func (et *e2eTest) knSourceKafkaList(t *testing.T, r *test.KnRunResultCollector, sourceName string) {
	out := et.it.KnPlugin().Run("list")
	r.AssertNoError(out)
	assert.Check(t, util.ContainsAll(out.Stdout, "NAME", "AGE", "SINK", "BOOTSTRAPSERVERS", sourceName))
}
