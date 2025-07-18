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

package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"knative.dev/kn-plugin-source-kafka/pkg/common/factories"
	"knative.dev/kn-plugin-source-kafka/pkg/common/types"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func TestNewKnSourceCommand(t *testing.T) {
	knSourceFactory := factories.NewDefaultKnSourceFactory()
	commandFactory := factories.NewDefaultCommandFactory(knSourceFactory)
	flagsFactory := factories.NewDefaultFlagsFactory(knSourceFactory)
	runEFactory := factories.NewDefaultRunEFactory(knSourceFactory)
	knSourceCmd := NewKnSourceCommand(knSourceFactory, commandFactory, flagsFactory, runEFactory)

	assert.Assert(t, knSourceCmd != nil)
	assert.Equal(t, knSourceCmd.Use, "source")
	assert.Equal(t, knSourceCmd.Short, "Knative eventing {{.Name}} source plugin")
	assert.Equal(t, knSourceCmd.Long, "Manage your Knative {{.Name}} eventing sources")
	assert.Equal(t, knSourceCmd.Example, "")
	assert.Equal(t, knSourceCmd.DisableAutoGenTag, true)
	assert.Equal(t, knSourceCmd.SilenceUsage, true)
	assert.Equal(t, knSourceCmd.SilenceErrors, true)

	testSubCommands(t, knSourceCmd, []string{"list", "create", "update", "describe", "delete", "help"})
	testSubCommandsRunE(t, knSourceCmd, []string{"list", "create", "update", "describe", "delete", "help"})
	testSubCommandsFlags(t, knSourceCmd, []string{"list", "create", "update", "describe", "delete", "help"})
}

// Private test methods

func testSubCommands(t *testing.T, cmd *cobra.Command, subCommandNames []string) {
	assert.Equal(t, len(cmd.Commands()), len(subCommandNames))
	childrenSubCommandsMap := map[string]bool{}
	for _, childCmd := range cmd.Commands() {
		childrenSubCommandsMap[childCmd.Name()] = false
		for _, name := range subCommandNames {
			if childCmd.Name() == name {
				childrenSubCommandsMap[childCmd.Name()] = true
			}
		}
	}
	assert.Equal(t, len(childrenSubCommandsMap), len(subCommandNames))
	for _, childCmd := range cmd.Commands() {
		if childrenSubCommandsMap[childCmd.Name()] == false {
			t.Errorf("did not find command %s as child of root command", childCmd.Name())
		}
	}
}

func testSubCommandsRunE(t *testing.T, cmd *cobra.Command, subCommandNames []string) {
	assert.Equal(t, len(cmd.Commands()), len(subCommandNames))
	childrenSubCommandsMap := map[string]types.RunE{}
	for _, childCmd := range cmd.Commands() {
		childrenSubCommandsMap[childCmd.Name()] = nil
		for _, name := range subCommandNames {
			if childCmd.Name() == name {
				childrenSubCommandsMap[childCmd.Name()] = childCmd.RunE
			}
		}
	}
	assert.Equal(t, len(childrenSubCommandsMap), len(subCommandNames))
	for _, childCmd := range cmd.Commands() {
		if _, ok := childrenSubCommandsMap[childCmd.Name()]; childCmd.Name() != "help" && !ok {
			t.Errorf("did not find RunE func for command %s", childCmd.Name())
		}
	}
}

func testSubCommandsFlags(t *testing.T, cmd *cobra.Command, subCommandNames []string) {
	assert.Equal(t, len(cmd.Commands()), len(subCommandNames))
	childrenSubCommandsMap := map[string]*pflag.FlagSet{}
	for _, childCmd := range cmd.Commands() {
		childrenSubCommandsMap[childCmd.Name()] = nil
		for _, name := range subCommandNames {
			if childCmd.Name() == name {
				childrenSubCommandsMap[childCmd.Name()] = childCmd.Flags()
			}
		}
	}
	assert.Equal(t, len(childrenSubCommandsMap), len(subCommandNames))
	for _, childCmd := range cmd.Commands() {
		if childrenSubCommandsMap[childCmd.Name()] == nil {
			t.Errorf("did not find FlagSet for command %s", childCmd.Name())
		}
	}
}
