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
	"errors"
	"fmt"
	"sort"
	"strings"

	v1alpha1 "knative.dev/eventing-kafka/pkg/apis/sources/v1alpha1"
	"knative.dev/kn-plugin-source-kafka/pkg/client"
	"knative.dev/kn-plugin-source-kafka/pkg/types"

	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"knative.dev/client/pkg/kn/commands"
	"knative.dev/client/pkg/kn/commands/flags"
	"knative.dev/client/pkg/printers"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// Max column size
const listColumnMaxLength = 50

type kafkaSourceRunEFactory struct {
	kafkaSourceClient  types.KafkaSourceClient
	kafkaSourceFactory types.KafkaSourceFactory
}

func NewKafkaSourceRunEFactory(kafkaFactory types.KafkaSourceFactory) types.KafkaSourceRunEFactory {
	return &kafkaSourceRunEFactory{
		kafkaSourceFactory: kafkaFactory,
		kafkaSourceClient:  kafkaFactory.KafkaSourceClient(),
	}
}

func NewFakeKafkaSourceRunEFactory(ns string) types.KafkaSourceRunEFactory {
	kafkaFactory := NewFakeKafkaSourceFactory(ns)
	return &kafkaSourceRunEFactory{
		kafkaSourceFactory: kafkaFactory,
		kafkaSourceClient:  kafkaFactory.KafkaSourceClient(),
	}
}

func (f *kafkaSourceRunEFactory) KafkaSourceClient(restConfig *rest.Config, namespace string) (types.KafkaSourceClient, error) {
	var err error
	f.kafkaSourceClient, err = f.KafkaSourceFactory().CreateKafkaSourceClient(restConfig, namespace)
	return f.kafkaSourceClient, err
}

func (f *kafkaSourceRunEFactory) KafkaSourceFactory() types.KafkaSourceFactory {
	return f.kafkaSourceFactory
}

func (f *kafkaSourceRunEFactory) CreateRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		namespace, err := f.KnSourceParams().GetNamespace(cmd)
		if err != nil {
			return err
		}

		restConfig, err := f.KnSourceParams().KnParams.RestConfig()
		if err != nil {
			return err
		}

		f.kafkaSourceClient, err = f.KafkaSourceClient(restConfig, namespace)
		if err != nil {
			return err
		}

		if len(args) != 1 {
			return errors.New("requires the name of the source to create as single argument")
		}
		name := args[0]

		dynamicClient, err := f.KnSourceParams().KnParams.NewDynamicClient(f.kafkaSourceClient.Namespace())
		if err != nil {
			return err
		}
		objectRef, err := f.KnSourceParams().SinkFlag.ResolveSink(dynamicClient, f.kafkaSourceClient.Namespace())
		if err != nil {
			return fmt.Errorf(
				"cannot create kafka '%s' in namespace '%s' "+
					"because: %s", name, f.kafkaSourceClient.Namespace(), err)
		}

		b := client.NewKafkaSourceBuilder(name).
			BootstrapServers(f.kafkaSourceFactory.KafkaSourceParams().BootstrapServers).
			Topics(f.kafkaSourceFactory.KafkaSourceParams().Topics).
			ConsumerGroup(f.kafkaSourceFactory.KafkaSourceParams().ConsumerGroup).
			Sink(objectRef)

		err = f.kafkaSourceClient.CreateKafkaSource(b.Build())

		if err != nil {
			return fmt.Errorf(
				"cannot create KafkaSource '%s' in namespace '%s' "+
					"because: %s", name, f.kafkaSourceClient.Namespace(), err)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Kafka source '%s' created in namespace '%s'.\n", name, f.kafkaSourceClient.Namespace())
		return nil
	}
}

func (f *kafkaSourceRunEFactory) DeleteRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		namespace, err := f.KnSourceParams().GetNamespace(cmd)
		if err != nil {
			return err
		}

		restConfig, err := f.KnSourceParams().KnParams.RestConfig()
		if err != nil {
			return err
		}

		f.kafkaSourceClient, err = f.KafkaSourceClient(restConfig, namespace)
		if err != nil {
			return err
		}

		if len(args) != 1 {
			return errors.New("requires the name of the source to create as single argument")
		}
		name := args[0]

		err = f.kafkaSourceClient.DeleteKafkaSource(name)

		if err != nil {
			return fmt.Errorf(
				"cannot delete KafkaSource '%s' in namespace '%s' "+
					"because: %s", name, f.kafkaSourceClient.Namespace(), err)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Kafka source '%s' deleted in namespace '%s'.\n", name, f.kafkaSourceClient.Namespace())
		return nil
	}
}

func (f *kafkaSourceRunEFactory) UpdateRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Kafka source update is not supported because kafka source spec is immutable.\n")
		return nil
	}
}

func (f *kafkaSourceRunEFactory) DescribeRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		namespace, err := f.KnSourceParams().GetNamespace(cmd)
		if err != nil {
			return err
		}

		restConfig, err := f.KnSourceParams().KnParams.RestConfig()
		if err != nil {
			return err
		}

		f.kafkaSourceClient, err = f.KafkaSourceClient(restConfig, namespace)
		if err != nil {
			return err
		}

		if len(args) != 1 {
			return errors.New("requires the name of the source to create as single argument")
		}
		name := args[0]

		kafkaSource, err := f.kafkaSourceClient.GetKafkaSource(name)

		if err != nil {
			return fmt.Errorf(
				"cannot describe KafkaSource '%s' in namespace '%s' "+
					"because: %s", name, f.kafkaSourceClient.Namespace(), err)
		}

		out := cmd.OutOrStdout()
		dw := printers.NewPrefixWriter(out)

		writeKafkaSource(dw, kafkaSource)
		dw.WriteLine()
		if err := dw.Flush(); err != nil {
			return err
		}

		if kafkaSource.Spec.Sink != nil {
			writeSink(dw, kafkaSource.Spec.Sink)
			dw.WriteLine()
			if err := dw.Flush(); err != nil {
				return err
			}
		}

		commands.WriteConditions(dw, kafkaSource.Status.Conditions, true)
		if err := dw.Flush(); err != nil {
			return err
		}
		return nil
	}
}

func (f *kafkaSourceRunEFactory) ListRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		namespace, err := f.KnSourceParams().GetNamespace(cmd)
		if err != nil {
			return err
		}
		restConfig, err := f.KnSourceParams().KnParams.RestConfig()
		if err != nil {
			return err
		}

		f.kafkaSourceClient, err = f.KafkaSourceClient(restConfig, namespace)
		if err != nil {
			return err
		}

		kafkaSourceList, err := f.kafkaSourceClient.ListKafkaSources()

		if err != nil {
			return fmt.Errorf(
				"cannot list KafkaSources in namespace '%s' "+
					"because: %s", f.kafkaSourceClient.Namespace(), err)
		}
		if len(kafkaSourceList.Items) == 0 {
			return fmt.Errorf("no kafka source found in namespace '%s'", f.kafkaSourceClient.Namespace())
		}

		printer := printers.NewTableGenerator()
		kafkaSourceListColumnDefinitions := []metav1.TableColumnDefinition{
			{Name: "Name", Type: "string", Description: "Name of the created kafka source", Priority: 1},
			{Name: "Age", Type: "string", Description: "Age of the kafka source", Priority: 1},
			{Name: "Sink", Type: "string", Description: "Sink of the kafka source", Priority: 1},
			{Name: "BootstrapServers", Type: "string", Description: "Kafka bootstrap server", Priority: 1},
		}
		err = printer.TableHandler(kafkaSourceListColumnDefinitions, printKafkaSource)
		if err != nil {
			return err
		}
		err = printer.TableHandler(kafkaSourceListColumnDefinitions, printKafkaSourceList)
		if err != nil {
			return err
		}
		err = printer.PrintObj(kafkaSourceList, cmd.OutOrStdout())
		if err != nil {
			return err
		}
		return nil
	}
}

func writeSink(dw printers.PrefixWriter, sink *duckv1.Destination) {
	subWriter := dw.WriteAttribute("Sink", "")
	ref := sink.Ref
	if ref != nil {
		subWriter.WriteAttribute("Name", sink.Ref.Name)
		if sink.Ref.Namespace != "" {
			subWriter.WriteAttribute("Namespace", sink.Ref.Namespace)
		}
		subWriter.WriteAttribute("Resource", fmt.Sprintf("%s (%s)", sink.Ref.Kind, sink.Ref.APIVersion))
	}
	uri := sink.URI
	if uri != nil {
		subWriter.WriteAttribute("URI", uri.String())
	}
}

func writeKafkaSource(dw printers.PrefixWriter, source *v1alpha1.KafkaSource) {
	commands.WriteMetadata(dw, &source.ObjectMeta, true)
	dw.WriteAttribute("BootstrapServers", strings.Join(source.Spec.BootstrapServers, ", "))
	dw.WriteAttribute("Topics", strings.Join(source.Spec.Topics, ","))
	dw.WriteAttribute("ConsumerGroup", source.Spec.ConsumerGroup)
}

// printKafkaSource populates a single row of kafka source list table
func printKafkaSource(kafkaSource *v1alpha1.KafkaSource, options printers.PrintOptions) ([]metav1.TableRow, error) {
	row := metav1.TableRow{
		Object: runtime.RawExtension{Object: kafkaSource},
	}

	if options.AllNamespaces {
		row.Cells = append(row.Cells, kafkaSource.ObjectMeta.Namespace)
	}

	row.Cells = append(row.Cells,
		trunc(kafkaSource.ObjectMeta.Name),
		commands.Age(kafkaSource.ObjectMeta.CreationTimestamp.Time),
		trunc(flags.SinkToString(*kafkaSource.Spec.Sink)),
		trunc(strings.Join(kafkaSource.Spec.BootstrapServers, ",")),
	)
	return []metav1.TableRow{row}, nil
}

// printKafkaSourceList populates the kafka source list table rows
func printKafkaSourceList(sourceList *v1alpha1.KafkaSourceList, options printers.PrintOptions) ([]metav1.TableRow, error) {
	rows := make([]metav1.TableRow, 0, len(sourceList.Items))

	sort.SliceStable(sourceList.Items, func(i, j int) bool {
		return sourceList.Items[i].ObjectMeta.Name < sourceList.Items[j].ObjectMeta.Name
	})
	for _, source := range sourceList.Items {
		row, err := printKafkaSource(&source, options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, row...)
	}
	return rows, nil
}

func trunc(txt string) string {
	if len(txt) <= listColumnMaxLength {
		return txt
	}
	return txt[:listColumnMaxLength-4] + " ..."
}
