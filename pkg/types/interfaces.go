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

package types

import (
	"context"

	v1beta1 "knative.dev/eventing-kafka-broker/control-plane/pkg/apis/sources/v1beta1"

	"k8s.io/client-go/rest"
	sourcetypes "knative.dev/client-pkg/pkg/kn-source-pkg/pkg/types"
)

type KafkaSourceClient interface {
	sourcetypes.KnSourceClient
	KafkaSourceParams() *KafkaSourceParams
	CreateKafkaSource(ctx context.Context, kafkaSource *v1beta1.KafkaSource) error
	DeleteKafkaSource(ctx context.Context, name string) error
	GetKafkaSource(ctx context.Context, name string) (*v1beta1.KafkaSource, error)
	ListKafkaSources(ctx context.Context) (*v1beta1.KafkaSourceList, error)
}

type KafkaSourceFactory interface {
	sourcetypes.KnSourceFactory

	KafkaSourceParams() *KafkaSourceParams
	KafkaSourceClient() KafkaSourceClient

	CreateKafkaSourceClient(restConfig *rest.Config, namespace string) (KafkaSourceClient, error)
	CreateKafkaSourceParams() *KafkaSourceParams
}

type KafkaSourceCommandFactory interface {
	sourcetypes.CommandFactory

	KafkaSourceFactory() KafkaSourceFactory
}

type KafkaSourceFlagsFactory interface {
	sourcetypes.FlagsFactory

	KafkaSourceFactory() KafkaSourceFactory
}

type KafkaSourceRunEFactory interface {
	sourcetypes.RunEFactory

	KafkaSourceFactory() KafkaSourceFactory
	KafkaSourceClient(restConfig *rest.Config, namespace string) (KafkaSourceClient, error)
}
