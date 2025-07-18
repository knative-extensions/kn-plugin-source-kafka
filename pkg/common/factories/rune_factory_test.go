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

	"knative.dev/kn-plugin-source-kafka/pkg/common/types"
)

func TestNewDefaultRunEFactory(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory != nil)
}

func TestListRunE(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory.ListRunE() != nil)
}

func TestCreateRunE(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory.CreateRunE() != nil)
}

func TestDeleteRunE(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory.DeleteRunE() != nil)
}

func TestUpdateRunE(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory.UpdateRunE() != nil)
}

func TestDescribeRunE(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory.DescribeRunE() != nil)
}

// Private

func createDefaultRunEFactory() types.RunEFactory {
	knSourceFactory := newDefaultKnSourceFactory()
	return NewDefaultRunEFactory(knSourceFactory)
}
