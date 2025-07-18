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
)

// KnSourceFactory

func TestKnSourceFactory_KnSourceParams(t *testing.T) {
	knSourceFactory := NewDefaultKnSourceFactory()

	knSourceParams := knSourceFactory.KnSourceParams()
	assert.Assert(t, knSourceParams != nil)

	knSourceParams = knSourceFactory.CreateKnSourceParams()
	assert.Equal(t, knSourceFactory.KnSourceParams(), knSourceParams)
}

// CommandFactory

func TestKnSourceFactory(t *testing.T) {
	knSourceFactory := NewDefaultKnSourceFactory()
	commandFactory := NewDefaultCommandFactory(knSourceFactory)

	assert.Equal(t, commandFactory.KnSourceFactory(), knSourceFactory)
}

// FlagsFactory

func TestFlagsFactory_KnSourceFactory(t *testing.T) {
	knSourceFactory := NewDefaultKnSourceFactory()
	flagsFactory := NewDefaultFlagsFactory(knSourceFactory)

	assert.Equal(t, flagsFactory.KnSourceFactory(), knSourceFactory)
}

// RunEFactory
func TestKnSourceClient(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	knSourceClient := runEFactory.KnSourceClient(&rest.Config{}, "fake_namespace")
	assert.Assert(t, knSourceClient != nil)
}

func TestRunEFactory_KnSourceParams(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory.KnSourceFactory().KnSourceParams() != nil)
}

func TestKnSourceClientFactory(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory.KnSourceFactory() != nil)
}
