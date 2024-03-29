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

	"knative.dev/kn-plugin-source-kafka/pkg/types"

	"gotest.tools/v3/assert"
)

func TestNewKafkaSourceFlagsFactory(t *testing.T) {
	flagsFactory := createKafkaSourceFlagsFactory()
	assert.Assert(t, flagsFactory != nil)
}

func TestFlagsFactory_KafkaSourceFactory(t *testing.T) {
	flagsFactory := createKafkaSourceFlagsFactory()
	assert.Equal(t, flagsFactory.KafkaSourceFactory(), flagsFactory.KafkaSourceFactory())
}

func TestCreateFlags(t *testing.T) {
	flagsFactory := createKafkaSourceFlagsFactory()
	flags := flagsFactory.CreateFlags()
	assert.Assert(t, flags != nil)

	assert.Assert(t, flags.Lookup("servers") != nil)
	assert.Assert(t, flags.Lookup("consumergroup") != nil)
	assert.Assert(t, flags.Lookup("topics") != nil)
	assert.Assert(t, flags.Lookup("ce-override") != nil)
	assert.Assert(t, flags.Lookup("label") != nil)
	assert.Assert(t, flags.ShorthandLookup("l") != nil)
	assert.Assert(t, flags.Lookup("annotation") != nil)
	assert.Assert(t, flags.ShorthandLookup("a") != nil)
}

func TestDeleteFlags(t *testing.T) {
	flagsFactory := createKafkaSourceFlagsFactory()
	flags := flagsFactory.DeleteFlags()
	assert.Assert(t, flags != nil)
}

func TestDescribeFlags(t *testing.T) {
	flagsFactory := createKafkaSourceFlagsFactory()
	flags := flagsFactory.DescribeFlags()
	assert.Assert(t, flags != nil)
}

func TestUpdateFlags(t *testing.T) {
	flagsFactory := createKafkaSourceFlagsFactory()
	flags := flagsFactory.UpdateFlags()
	assert.Assert(t, flags != nil)
}
func TestListFlags(t *testing.T) {
	flagsFactory := createKafkaSourceFlagsFactory()
	flags := flagsFactory.ListFlags()
	assert.Assert(t, flags != nil)
}

// Private

func createKafkaSourceFlagsFactory() types.KafkaSourceFlagsFactory {
	factory := NewKafkaSourceFactory()
	return NewKafkaSourceFlagsFactory(factory)
}
