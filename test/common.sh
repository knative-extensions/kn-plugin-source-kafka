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

source $(dirname $0)/../vendor/knative.dev/hack/e2e-tests.sh

function cluster_setup() {
  header "Installing client"
  local kn_build=$(mktemp -d)
  local failed=""
  pushd "$kn_build"
  git clone https://github.com/knative/client . || failed="Cannot clone kn githup repo"
  hack/build.sh -f || failed="error while builing kn"
  cp kn /usr/local/bin/kn || failed="can't copy kn to /usr/local/bin"
  chmod a+x /usr/local/bin/kn || failed="can't chmod kn"
  if [ -n "$failed" ]; then
     echo "ERROR: $failed"
     exit 1
  fi
  popd
  rm -rf "$kn_build"

  header "Building plugin"
  ${REPO_ROOT_DIR}/hack/build.sh -f || return 1
}

function knative_setup() {
  local serving_version=${KNATIVE_SERVING_VERSION:-latest}
  header "Installing Knative Serving (${serving_version})"

  if [ "${serving_version}" = "latest" ]; then
    start_latest_knative_serving
  else
    start_release_knative_serving "${serving_version}"
  fi

  local eventing_version=${KNATIVE_EVENTING_VERSION:-latest}
  header "Installing Knative Eventing (${eventing_version})"

  if [ "${eventing_version}" = "latest" ]; then
    start_latest_knative_eventing

    subheader "Installing eventing extension: sugar-controller (${eventing_version})"
    # install the sugar controller
    kubectl apply --filename https://storage.googleapis.com/knative-nightly/eventing/latest/eventing-sugar-controller.yaml
    wait_until_pods_running knative-eventing || return 1

  else
    start_release_knative_eventing "${eventing_version}"

    subheader "Installing eventing extension: sugar-controller (${eventing_version})"
    # install the sugar controller
    kubectl apply --filename https://storage.googleapis.com/knative-releases/eventing/previous/v${eventing_version}/eventing-sugar-controller.yaml
    wait_until_pods_running knative-eventing || return 1
  fi
}

function kafka_setup() {
  subheader "Installing Kafka 2.6.0 using strimzi 0.20.0"
  kubectl create namespace kafka || return 1
  sed 's/namespace: .*/namespace: kafka/' ${STRIMZI_INSTALLATION_CONFIG_TEMPLATE} > ${STRIMZI_INSTALLATION_CONFIG}
  kubectl apply -f ${STRIMZI_INSTALLATION_CONFIG} -n kafka
  kubectl apply -f ${KAFKA_INSTALLATION_CONFIG} -n kafka
  kubectl apply -f ${KAFKA_TOPIC_INSTALLATION_CONFIG} -n kafka
  wait_until_pods_running kafka || fail_test "Failed to start up a Kafka cluster"
}

function install_sources_crds() {
  subheader "Installing Kafka Source CRD"
  kubectl apply -f ${KAFKA_SOURCE_CRD_YAML}
  wait_until_pods_running knative-eventing || fail_test "Failed to install the Kafka Source CRD"
}

function plugin_test_setup() {
  header "🧪  Setup Kafka and source CRDs"
  kafka_setup || return 1
  install_sources_crds || return 1
}

function uninstall_sources_crds() {
  echo "Uninstalling Kafka Source CRD"
  kubectl delete -f ${KAFKA_SOURCE_CRD_YAML}
}

function kafka_teardown() {
  echo "Uninstalling Kafka cluster"
  kubectl delete -f ${KAFKA_TOPIC_INSTALLATION_CONFIG} -n kafka
  kubectl delete -f ${KAFKA_INSTALLATION_CONFIG} -n kafka
  kubectl delete -f "${STRIMZI_INSTALLATION_CONFIG}" -n kafka
  kubectl delete namespace kafka
}

function plugin_test_teardown() {
  kafka_teardown
  uninstall_sources_crds
}
