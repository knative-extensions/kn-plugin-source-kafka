/*
 * Copyright 2021 The Knative Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	sourcesv1beta1 "knative.dev/eventing-kafka-broker/control-plane/pkg/apis/sources/v1beta1"
	scheme "knative.dev/eventing-kafka-broker/control-plane/pkg/client/clientset/versioned/scheme"
)

// KafkaSourcesGetter has a method to return a KafkaSourceInterface.
// A group's client should implement this interface.
type KafkaSourcesGetter interface {
	KafkaSources(namespace string) KafkaSourceInterface
}

// KafkaSourceInterface has methods to work with KafkaSource resources.
type KafkaSourceInterface interface {
	Create(ctx context.Context, kafkaSource *sourcesv1beta1.KafkaSource, opts v1.CreateOptions) (*sourcesv1beta1.KafkaSource, error)
	Update(ctx context.Context, kafkaSource *sourcesv1beta1.KafkaSource, opts v1.UpdateOptions) (*sourcesv1beta1.KafkaSource, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, kafkaSource *sourcesv1beta1.KafkaSource, opts v1.UpdateOptions) (*sourcesv1beta1.KafkaSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*sourcesv1beta1.KafkaSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*sourcesv1beta1.KafkaSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *sourcesv1beta1.KafkaSource, err error)
	GetScale(ctx context.Context, kafkaSourceName string, options v1.GetOptions) (*autoscalingv1.Scale, error)
	UpdateScale(ctx context.Context, kafkaSourceName string, scale *autoscalingv1.Scale, opts v1.UpdateOptions) (*autoscalingv1.Scale, error)

	KafkaSourceExpansion
}

// kafkaSources implements KafkaSourceInterface
type kafkaSources struct {
	*gentype.ClientWithList[*sourcesv1beta1.KafkaSource, *sourcesv1beta1.KafkaSourceList]
}

// newKafkaSources returns a KafkaSources
func newKafkaSources(c *SourcesV1beta1Client, namespace string) *kafkaSources {
	return &kafkaSources{
		gentype.NewClientWithList[*sourcesv1beta1.KafkaSource, *sourcesv1beta1.KafkaSourceList](
			"kafkasources",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *sourcesv1beta1.KafkaSource { return &sourcesv1beta1.KafkaSource{} },
			func() *sourcesv1beta1.KafkaSourceList { return &sourcesv1beta1.KafkaSourceList{} },
		),
	}
}

// GetScale takes name of the kafkaSource, and returns the corresponding autoscalingv1.Scale object, and an error if there is any.
func (c *kafkaSources) GetScale(ctx context.Context, kafkaSourceName string, options v1.GetOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.GetClient().Get().
		Namespace(c.GetNamespace()).
		Resource("kafkasources").
		Name(kafkaSourceName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *kafkaSources) UpdateScale(ctx context.Context, kafkaSourceName string, scale *autoscalingv1.Scale, opts v1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.GetClient().Put().
		Namespace(c.GetNamespace()).
		Resource("kafkasources").
		Name(kafkaSourceName).
		SubResource("scale").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scale).
		Do(ctx).
		Into(result)
	return
}
