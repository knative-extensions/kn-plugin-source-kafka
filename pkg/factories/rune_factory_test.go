// Copyright © 2020 The Knative Authors
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

package factories

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
	"k8s.io/client-go/rest"
	"knative.dev/client/pkg/printers"
	"knative.dev/client/pkg/util"
	v1alpha1 "knative.dev/eventing-kafka/pkg/apis/sources/v1alpha1"
	"knative.dev/kn-plugin-source-kafka/pkg/client"
)

func TestNewKafkaSourceRunEFactory(t *testing.T) {
	runEFactory := NewFakeKafkaSourceRunEFactory("fake_namespace")
	assert.Assert(t, runEFactory != nil)
}

func TestRunEFactory_KafkaSourceParams(t *testing.T) {
	runEFactory := NewFakeKafkaSourceRunEFactory("fake_namespace")
	assert.Assert(t, runEFactory.KafkaSourceFactory().KafkaSourceParams() != nil)
}

func TestRunEFactory_KafkaSourceFactory(t *testing.T) {
	runEFactory := NewFakeKafkaSourceRunEFactory("fake_namespace")
	assert.Assert(t, runEFactory.KafkaSourceFactory() != nil)
}

func TestRunEFactory_KafkaSourceClient(t *testing.T) {
	runEFactory := NewFakeKafkaSourceRunEFactory("fake_namespace")
	knSourceClient, err := runEFactory.KafkaSourceClient(&rest.Config{}, "fake_namespace")
	assert.Assert(t, knSourceClient != nil)
	assert.Assert(t, err == nil)
}

func TestCreateRunE(t *testing.T) {
	runEFactory := NewFakeKafkaSourceRunEFactory("fake_namespace")
	function := runEFactory.CreateRunE()
	assert.Assert(t, function != nil)
}

func TestDeleteRunE(t *testing.T) {
	runEFactory := NewFakeKafkaSourceRunEFactory("fake_namespace")
	function := runEFactory.DeleteRunE()
	assert.Assert(t, function != nil)
}

func TestUpdateRunE(t *testing.T) {
	runEFactory := NewFakeKafkaSourceRunEFactory("fake_namespace")
	function := runEFactory.UpdateRunE()
	assert.Assert(t, function != nil)
}

func TestDescribeRunE(t *testing.T) {
	runEFactory := NewFakeKafkaSourceRunEFactory("fake_namespace")
	function := runEFactory.DescribeRunE()
	assert.Assert(t, function != nil)
}

func TestListRunE(t *testing.T) {
	runEFactory := NewFakeKafkaSourceRunEFactory("fake_namespace")
	function := runEFactory.ListRunE()
	assert.Assert(t, function != nil)
}

func TestPrintKafkaSource(t *testing.T) {
	obj := newKafkaSource("foo")
	row := printKafkaSource(obj, printers.PrintOptions{})
	assert.Assert(t, len(row) == 1)
	assert.Check(t, util.ContainsAll(fmt.Sprint(row[0].Cells), "foo", "test.server.org", "topic", "mygroup"))
}

func TestPrintKafkaSourceList(t *testing.T) {
	kafkaSource1 := newKafkaSource("foo")
	kafkaSource2 := newKafkaSource("bar")
	obj := &v1alpha1.KafkaSourceList{Items: []v1alpha1.KafkaSource{*kafkaSource1, *kafkaSource2}}
	row := printKafkaSourceList(obj, printers.PrintOptions{})
	assert.Assert(t, len(row) == 2)
	assert.Check(t, util.ContainsAll(fmt.Sprint(row[0].Cells), "bar", "test.server.org", "topic", "mygroup"))
	assert.Check(t, util.ContainsAll(fmt.Sprint(row[1].Cells), "foo", "test.server.org", "topic", "mygroup"))
}

func newKafkaSource(name string) *v1alpha1.KafkaSource {
	return client.NewKafkaSourceBuilder(name).
		BootstrapServers([]string{"test.server.org"}).
		Topics([]string{"topic"}).
		ConsumerGroup("mygroup").
		Build()
}
