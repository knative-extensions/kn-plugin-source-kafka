#!/usr/bin/env bash

# Copyright 2020 The Knative Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

source $(dirname $0)/common.sh


# Add local dir to have access to built kn
export PATH=$PATH:${REPO_ROOT_DIR}

# Will create and delete this namespace (used for all tests, modify if you want a different one used)
export KN_E2E_NAMESPACE=kne2etests

# Strimzi installation config template used for starting up Kafka clusters.
readonly STRIMZI_INSTALLATION_CONFIG_TEMPLATE="${REPO_ROOT_DIR}/test/config/100-strimzi-cluster-operator.yaml"
# Strimzi installation config.
readonly STRIMZI_INSTALLATION_CONFIG="$(mktemp)"
# Kafka cluster CR config file.
readonly KAFKA_INSTALLATION_CONFIG="${REPO_ROOT_DIR}/test/config/100-kafka-ephemeral.yaml"
readonly KAFKA_TOPIC_INSTALLATION_CONFIG="${REPO_ROOT_DIR}/test/config/100-kafka-topic.yaml"
# Kafka cluster URL for our installation
readonly KAFKA_CLUSTER_URL="my-cluster-kafka-bootstrap.kafka:9092"
# Kafka channel CRD config template directory.
readonly KAFKA_CRD_CONFIG_TEMPLATE_DIR="kafka/channel/config"
# Kafka channel CRD config template file. It needs to be modified to be the real config file.
readonly KAFKA_CRD_CONFIG_TEMPLATE="400-kafka-config.yaml"
# Real Kafka channel CRD config , generated from the template directory and modified template file.
readonly KAFKA_CRD_CONFIG_DIR="$(mktemp -d)"
# Resolve the latest release eventing-kafka-broker resources
# Nightly is used for 'main' branch, for any 'release-*' branch the corresponding eventing-kafka released version
readonly KAFKA_SOURCE_CONTROLLER_YAML="$(get_latest_knative_yaml_source "eventing-kafka-broker" "eventing-kafka-controller")"
readonly KAFKA_SOURCE_YAML="$(get_latest_knative_yaml_source "eventing-kafka-broker" "eventing-kafka-source")"


run() {
  # Create cluster
  initialize $@ --skip-istio-addon

  # Kafka setup
  eval plugin_test_setup || fail_test

  # Integration tests
  eval integration_test || fail_test

  success
}

integration_test() {
  header "Running kn-plugin-source-kafka e2e tests for Knative Serving $KNATIVE_SERVING_VERSION and Eventing $KNATIVE_EVENTING_VERSION"

  go_test_e2e -v -timeout=90m ./test/e2e || return 1
}

# Fire up
run $@
