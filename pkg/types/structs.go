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
	clientv1beta1 "knative.dev/eventing-kafka-broker/control-plane/pkg/client/clientset/versioned/typed/sources/v1beta1"

	sourcetypes "knative.dev/kn-plugin-source-kafka/pkg/common/types"
)

type KafkaSourceParams struct {
	KnSourceParams   *sourcetypes.KnSourceParams
	BootstrapServers []string
	Topics           []string
	ConsumerGroup    string
	CeOverrides      []string
	Labels           []string
	Annotations      []string
}

func (p *KafkaSourceParams) NewSourcesClient() (*clientv1beta1.SourcesV1beta1Client, error) {
	restConfig, err := p.KnSourceParams.RestConfig()
	if err != nil {
		return nil, err
	}

	c, _ := clientv1beta1.NewForConfig(restConfig)
	return c, nil
}
