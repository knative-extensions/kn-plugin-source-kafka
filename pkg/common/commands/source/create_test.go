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

package source

import (
	"testing"

	"gotest.tools/v3/assert"

	"knative.dev/kn-plugin-source-kafka/pkg/common/types"
)

func TestNewCreateCommand(t *testing.T) {
	createCmd := NewCreateCommand(&types.KnSourceParams{})
	assert.Assert(t, createCmd != nil)
	assert.Equal(t, createCmd.Use, "create NAME [flags]")
	assert.Equal(t, createCmd.Short, "create {{.Name}} source")
	assert.Equal(t, createCmd.Example, "{{.CreateExample}}")
}
