// Copyright © 2019 The Knative Authors
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

package client

import (
	"testing"

	client_testing "k8s.io/client-go/testing"

	"gotest.tools/v3/assert"
	"k8s.io/apimachinery/pkg/runtime"
	v1alpha1 "knative.dev/eventing-kafka/pkg/apis/sources/v1alpha1"
	"knative.dev/eventing-kafka/pkg/client/clientset/versioned/typed/sources/v1alpha1/fake"
	"knative.dev/kn-plugin-source-kafka/pkg/types"
)

var testNamespace = "fake-namespace"

func setup() (fake.FakeSourcesV1alpha1, types.KafkaSourceClient) {
	fakeClient := fake.FakeSourcesV1alpha1{Fake: &client_testing.Fake{}}
	knSourceClient := NewFakeKafkaSourceClient(&fakeClient, testNamespace)
	return fakeClient, knSourceClient
}
func TestKafkaSourceClient(t *testing.T) {
	_, knSourceClient := setup()
	assert.Assert(t, knSourceClient != nil)
}

func TestClient_KnSourceParams(t *testing.T) {
	_, knSourceClient := setup()
	fakeKafkaParams := knSourceClient.KafkaSourceParams()
	assert.Equal(t, knSourceClient.KnSourceParams(), fakeKafkaParams.KnSourceParams)
}

func TestNamespace(t *testing.T) {
	_, knSourceClient := setup()
	assert.Equal(t, knSourceClient.Namespace(), testNamespace)
}
func TestCreateKafka(t *testing.T) {
	_, cli := setup()
	objNew := newKafkaSource("samplekafka")
	err := cli.CreateKafkaSource(objNew)
	assert.NilError(t, err)
}

func TestDeleteKafka(t *testing.T) {
	_, cli := setup()
	objNew := newKafkaSource("samplekafka")
	err := cli.CreateKafkaSource(objNew)
	assert.NilError(t, err)
	err = cli.DeleteKafkaSource("samplekafka")
	assert.NilError(t, err)
}

func TestCreateKafkaMultipleTopicsServers(t *testing.T) {
	_, cli := setup()
	objNew := NewKafkaSourceBuilder("samplekafka").
		BootstrapServers([]string{"test.server.org", "foo.server.org"}).
		Topics([]string{"foo", "bar"}).
		ConsumerGroup("mygroup").
		Build()
	err := cli.CreateKafkaSource(objNew)
	assert.NilError(t, err)
}

func TestGetKafkaSources(t *testing.T) {
	fakeClient, cli := setup()
	fakeClient.AddReactor("list", "kafkasources",
		func(action client_testing.Action) (handled bool, ret runtime.Object, err error) {
			kafkaSrc1 := newKafkaSource("foo")
			kafkaSrc2 := newKafkaSource("bar")
			return true, &v1alpha1.KafkaSourceList{Items: []v1alpha1.KafkaSource{*kafkaSrc1, *kafkaSrc2}}, err
		})
	sources, err := cli.ListKafkaSources()
	assert.NilError(t, err)
	assert.Assert(t, len(sources.Items) == 2)
}

func newKafkaSource(name string) *v1alpha1.KafkaSource {
	return NewKafkaSourceBuilder(name).
		BootstrapServers([]string{"test.server.org"}).
		Topics([]string{"topic"}).
		ConsumerGroup("mygroup").
		CloudEventOverrides(map[string]string{"type": "foo"}, []string{}).
		Build()
}
