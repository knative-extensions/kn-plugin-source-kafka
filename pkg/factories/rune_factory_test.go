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

	"gotest.tools/v3/assert"
	"k8s.io/client-go/rest"
	"knative.dev/client/pkg/printers"
	"knative.dev/client/pkg/util"
	v1alpha1 "knative.dev/eventing-kafka/pkg/apis/sources/v1alpha1"
	"knative.dev/kn-plugin-source-kafka/pkg/client"
	v1 "knative.dev/pkg/apis/duck/v1"
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
	row, err := printKafkaSource(obj, printers.PrintOptions{})
	assert.NilError(t, err)
	assert.Assert(t, len(row) == 1)
	assert.Check(t, util.ContainsAll(fmt.Sprint(row[0].Cells), "foo", "ksvc:mysvc", "test.server.org"))
}

func TestPrintKafkaSourceList(t *testing.T) {
	kafkaSource1 := newKafkaSource("foo")
	kafkaSource2 := newKafkaSource("bar")
	obj := &v1alpha1.KafkaSourceList{Items: []v1alpha1.KafkaSource{*kafkaSource1, *kafkaSource2}}
	row, err := printKafkaSourceList(obj, printers.PrintOptions{})
	assert.NilError(t, err)
	assert.Assert(t, len(row) == 2)
	assert.Check(t, util.ContainsAll(fmt.Sprint(row[0].Cells), "bar", "ksvc:mysvc", "test.server.org"))
	assert.Check(t, util.ContainsAll(fmt.Sprint(row[1].Cells), "foo", "ksvc:mysvc", "test.server.org"))
}

func TestTrunc(t *testing.T) {
	str := "my-cluster-kafka-bootstrap.kafka.svc:9092,my-cluster1-kafka-bootstrap.kafka.svc:9092"
	truncStr := trunc(str)
	assert.Assert(t, len(truncStr) == 50)
	assert.Check(t, util.ContainsAll(truncStr, "my-cluster-kafka-bootstrap.kafka.svc:9092,my-c ..."))
	str = "mykafkasrc"
	truncStr = trunc(str)
	assert.Check(t, util.ContainsAll(truncStr, str))
}

func newKafkaSource(name string) *v1alpha1.KafkaSource {
	return client.NewKafkaSourceBuilder(name).
		BootstrapServers([]string{"test.server.org"}).
		Topics([]string{"topic"}).
		ConsumerGroup("mygroup").
		Sink(&v1.Destination{Ref: &v1.KReference{Name: "mysvc", Kind: "Service"}}).
		Build()
}
