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

package types

import (
	"testing"

	"github.com/spf13/cobra"
	"gotest.tools/v3/assert"
	"knative.dev/client/pkg/commands/flags"
)

func TestAddCommonFlags(t *testing.T) {
	fakeCmd := &cobra.Command{}
	fakeKnSourceParams := newFakeKnSourceParams()

	fakeKnSourceParams.AddCommonFlags(fakeCmd)

	flag := fakeCmd.Flags().Lookup("namespace")
	assert.Assert(t, flag != nil)
	assert.Assert(t, flag.Name == "namespace")

	flag = fakeCmd.Flags().Lookup("all-namespaces")
	assert.Assert(t, flag != nil)
	assert.Assert(t, flag.Name == "all-namespaces")
}

func TestAddCreateUpdateFlags(t *testing.T) {
	fakeCmd := &cobra.Command{}
	fakeKnSourceParams := newFakeKnSourceParams()

	fakeKnSourceParams.AddCreateUpdateFlags(fakeCmd)

	flag := fakeCmd.Flags().Lookup("sink")
	assert.Assert(t, flag != nil)
	assert.Assert(t, flag.Name == "sink")
}

// Private

func newFakeKnSourceParams() *KnSourceParams {
	return &KnSourceParams{
		SinkFlag: flags.SinkFlags{},
	}
}
