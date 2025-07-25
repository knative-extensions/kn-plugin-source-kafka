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

package client

import (
	"k8s.io/client-go/rest"

	"knative.dev/eventing-kafka-broker/control-plane/pkg/client/clientset/versioned/typed/sources/v1beta1/fake"
	sourcetypes "knative.dev/kn-plugin-source-kafka/pkg/common/types"
	"knative.dev/kn-plugin-source-kafka/pkg/common/types/typesfakes"
	"knative.dev/kn-plugin-source-kafka/pkg/types"
)

// NewFakeKafkaSourceClient is to create a fake KafkaSourceClient to test
func NewFakeKafkaSourceClient(fakeClientTest *fake.FakeSourcesV1beta1, ns string) types.KafkaSourceClient {
	kafkaParams := NewFakeKafkaSourceParams()
	knFakeSourceClient := &typesfakes.FakeKnSourceClient{}
	knFakeSourceClient.KnSourceParamsReturns(kafkaParams.KnSourceParams)
	knFakeSourceClient.NamespaceReturns(ns)
	knFakeSourceClient.RestConfigReturns(&rest.Config{})

	return &kafkaSourceClient{
		namespace:         ns,
		kafkaSourceParams: kafkaParams,
		client:            fakeClientTest,
		knSourceClient:    knFakeSourceClient,
	}
}

// NewFakeKafkaSourceParams is to create a fake KafkaSourceParams to test
func NewFakeKafkaSourceParams() *types.KafkaSourceParams {
	return &types.KafkaSourceParams{
		KnSourceParams: &sourcetypes.KnSourceParams{},
	}
}
