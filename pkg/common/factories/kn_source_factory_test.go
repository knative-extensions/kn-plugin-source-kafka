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
	"testing"

	"gotest.tools/v3/assert"
	"k8s.io/client-go/rest"
	"knative.dev/client/pkg/commands/flags"
	"knative.dev/kn-plugin-source-kafka/pkg/common/types"
	"knative.dev/kn-plugin-source-kafka/pkg/common/types/typesfakes"
)

func TestNewDefaultKnSourceFactory(t *testing.T) {
	knSourceFactory := newDefaultKnSourceFactory()

	assert.Assert(t, knSourceFactory != nil)
}

func TestCreateKnSourceParams(t *testing.T) {
	knSourceFactory := newDefaultKnSourceFactory()

	knSourceParams := knSourceFactory.CreateKnSourceParams()
	assert.Assert(t, knSourceParams != nil)
}

func TestCreateKnSourceClient(t *testing.T) {
	knSourceFactory := newDefaultKnSourceFactory()
	client := knSourceFactory.CreateKnSourceClient(&rest.Config{}, "fake-namespace")

	assert.Assert(t, client != nil)
	assert.Equal(t, client.Namespace(), "fake-namespace")
}

// Private

func newDefaultKnSourceFactory() types.KnSourceFactory {
	return &DefautKnSourceFactory{
		knSourceParams:     &types.KnSourceParams{SinkFlag: flags.SinkFlags{}},
		knSourceClientFunc: fakeKnSourceClientFunc,
	}
}

func fakeKnSourceClientFunc(knSourceParams *types.KnSourceParams, restConfig *rest.Config, namespace string) types.KnSourceClient {
	fakeKnSourceClient := &typesfakes.FakeKnSourceClient{}
	fakeKnSourceClient.KnSourceParamsReturns(knSourceParams)
	fakeKnSourceClient.NamespaceReturns(namespace)
	fakeKnSourceClient.RestConfigReturns(restConfig)
	return fakeKnSourceClient
}
