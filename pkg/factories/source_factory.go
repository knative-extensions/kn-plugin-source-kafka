// Copyright © 2018 The Knative Authors
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
	"k8s.io/client-go/rest"
	clienttesting "k8s.io/client-go/testing"
	"knative.dev/eventing-kafka-broker/control-plane/pkg/client/clientset/versioned/typed/sources/v1beta1/fake"
	"knative.dev/kn-plugin-source-kafka/pkg/client"
	sourcefactories "knative.dev/kn-plugin-source-kafka/pkg/common/factories"
	sourcetypes "knative.dev/kn-plugin-source-kafka/pkg/common/types"
	"knative.dev/kn-plugin-source-kafka/pkg/types"
)

type kafkaClientFactory struct {
	kafkaSourceParams *types.KafkaSourceParams
	kafkaSourceClient types.KafkaSourceClient
	knSourceFactory   sourcetypes.KnSourceFactory
}

func NewKafkaSourceFactory() types.KafkaSourceFactory {
	return &kafkaClientFactory{
		kafkaSourceParams: nil,
		kafkaSourceClient: nil,
		knSourceFactory:   sourcefactories.NewDefaultKnSourceFactory(),
	}
}

func NewFakeKafkaSourceFactory(ns string) types.KafkaSourceFactory {
	fakeClient := fake.FakeSourcesV1beta1{Fake: &clienttesting.Fake{}}
	fakeSourceClient := client.NewFakeKafkaSourceClient(&fakeClient, ns)
	fakeParams := fakeSourceClient.KafkaSourceParams()
	return &kafkaClientFactory{
		kafkaSourceParams: fakeParams,
		kafkaSourceClient: fakeSourceClient,
		knSourceFactory:   sourcefactories.NewDefaultKnSourceFactory(),
	}
}

func (f *kafkaClientFactory) CreateKafkaSourceClient(restConfig *rest.Config, namespace string) (types.KafkaSourceClient, error) {
	var err error
	if f.kafkaSourceClient == nil {
		f.kafkaSourceClient, err = client.NewKafkaSourceClient(f.KafkaSourceParams(), restConfig, namespace)
		if err != nil {
			return nil, err
		}
	}
	return f.kafkaSourceClient, nil
}

func (f *kafkaClientFactory) KafkaSourceParams() *types.KafkaSourceParams {
	if f.kafkaSourceParams == nil {
		f.initKafkaSourceParams()
	}
	return f.kafkaSourceParams
}

func (f *kafkaClientFactory) CreateKafkaSourceParams() *types.KafkaSourceParams {
	if f.kafkaSourceParams == nil {
		f.initKafkaSourceParams()
	}
	return f.kafkaSourceParams
}

// Private

func (f *kafkaClientFactory) initKafkaSourceParams() {
	f.kafkaSourceParams = &types.KafkaSourceParams{
		KnSourceParams: f.knSourceFactory.CreateKnSourceParams(),
	}
	f.kafkaSourceParams.KnSourceParams.Initialize()
}
