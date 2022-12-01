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
	sourcefactories "knative.dev/client-pkg/pkg/kn-source-pkg/pkg/factories"
	sourcetypes "knative.dev/client-pkg/pkg/kn-source-pkg/pkg/types"

	"knative.dev/kn-plugin-source-kafka/pkg/types"

	"github.com/spf13/pflag"
)

type kafkaSourceFlagsFactory struct {
	defaultFlagsFactory sourcetypes.FlagsFactory
	kafkaSourceFactory  types.KafkaSourceFactory
}

func NewKafkaSourceFlagsFactory(kafkaFactory types.KafkaSourceFactory) types.KafkaSourceFlagsFactory {
	return &kafkaSourceFlagsFactory{
		defaultFlagsFactory: sourcefactories.NewDefaultFlagsFactory(kafkaFactory),
		kafkaSourceFactory:  kafkaFactory,
	}
}

func (f *kafkaSourceFlagsFactory) KafkaSourceParams() *types.KafkaSourceParams {
	return f.kafkaSourceFactory.KafkaSourceParams()
}

func (f *kafkaSourceFlagsFactory) KafkaSourceFactory() types.KafkaSourceFactory {
	return f.kafkaSourceFactory
}

func (f *kafkaSourceFlagsFactory) CreateFlags() *pflag.FlagSet {
	flagSet := pflag.NewFlagSet("create", pflag.ExitOnError)
	flagSet.StringArrayVar(&f.KafkaSourceParams().BootstrapServers, "servers", []string{}, "Kafka bootstrap servers that the consumer will connect to, consist of a hostname plus a port pair, e.g. my-kafka-bootstrap.kafka:9092. Flag can be used multiple times.")
	flagSet.StringArrayVar(&f.KafkaSourceParams().Topics, "topics", []string{}, "Topics to consume messages from. Flag can be used multiple times.")
	flagSet.StringVar(&f.KafkaSourceParams().ConsumerGroup, "consumergroup", "", "the consumer group ID")
	flagSet.StringArrayVar(&f.KafkaSourceParams().CeOverrides, "ce-override", []string{}, "Cloud Event overrides to apply before sending event to sink. "+
		"Example: '--ce-override key=value' "+
		"You may be provide this flag multiple times. "+
		"To unset, append \"-\" to the key (e.g. --ce-override key-).")
	flagSet.StringArrayVarP(&f.KafkaSourceParams().Labels, "label", "l", []string{}, "Metadata labels to set on the resources. "+
		"Example: '--label key=value' "+
		"You may be provide this flag multiple times.")
	flagSet.StringArrayVarP(&f.KafkaSourceParams().Annotations, "annotation", "a", []string{}, "Metadata annotations to set on the resources. "+
		"Example: '--annotation key=value' "+
		"You may be provide this flag multiple times.")
	return flagSet
}

func (f *kafkaSourceFlagsFactory) DeleteFlags() *pflag.FlagSet {
	return pflag.NewFlagSet("delete", pflag.ExitOnError)
}

func (f *kafkaSourceFlagsFactory) UpdateFlags() *pflag.FlagSet {
	flagSet := pflag.NewFlagSet("update", pflag.ExitOnError)
	return flagSet
}

func (f *kafkaSourceFlagsFactory) DescribeFlags() *pflag.FlagSet {
	return pflag.NewFlagSet("describe", pflag.ExitOnError)
}

func (f *kafkaSourceFlagsFactory) ListFlags() *pflag.FlagSet {
	return pflag.NewFlagSet("list", pflag.ExitOnError)
}
