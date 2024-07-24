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

# Copied from knative/serving setup script:
# https://github.com/knative/serving/blob/main/test/e2e-networking-library.sh#L17
function install_istio() {
  if [[ -z "${ISTIO_VERSION:-}" ]]; then
    readonly ISTIO_VERSION="latest"
  fi
  header "Installing Istio ${ISTIO_VERSION}"
  local LATEST_NET_ISTIO_RELEASE_VERSION=$(curl -L --silent "https://api.github.com/repos/knative/net-istio/releases" | \
    jq -r '[.[].tag_name] | sort_by( sub("knative-";"") | sub("v";"") | split(".") | map(tonumber) ) | reverse[0]')
  # And checkout the setup script based on that release
  local NET_ISTIO_DIR=$(mktemp -d)
  (
    cd $NET_ISTIO_DIR \
      && git init \
      && git remote add origin https://github.com/knative-sandbox/net-istio.git \
      && git fetch --depth 1 origin $LATEST_NET_ISTIO_RELEASE_VERSION \
      && git checkout FETCH_HEAD
  )

  if [[ -z "${ISTIO_PROFILE:-}" ]]; then
    readonly ISTIO_PROFILE="istio-ci-no-mesh.yaml"
  fi

  if [[ -n "${CLUSTER_DOMAIN:-}" ]]; then
    sed -ie "s/cluster\.local/${CLUSTER_DOMAIN}/g" ${NET_ISTIO_DIR}/third_party/istio-${ISTIO_VERSION}/${ISTIO_PROFILE}
  fi

  echo ">> Installing Istio"
  echo "Istio version: ${ISTIO_VERSION}"
  echo "Istio profile: ${ISTIO_PROFILE}"
  kubectl apply -f ${NET_ISTIO_DIR}/third_party/istio-${ISTIO_VERSION}/${ISTIO_PROFILE%%.*}/istio.yaml

}

function knative_setup() {
  install_istio

  header "Installing Knative Serving"
  # Defined by knative/hack/library.sh
  kubectl apply --filename ${KNATIVE_SERVING_RELEASE_CRDS}
  kubectl apply --filename ${KNATIVE_SERVING_RELEASE_CORE}
  kubectl apply --filename ${KNATIVE_NET_ISTIO_RELEASE}
  wait_until_pods_running knative-serving || return 1

  local eventing_version=${KNATIVE_EVENTING_VERSION:-latest}
  header "Installing Knative Eventing"
  # Defined by knative/hack/library.sh
  kubectl apply --filename ${KNATIVE_EVENTING_RELEASE}
  kubectl apply --filename ${KNATIVE_EVENTING_SUGAR_CONTROLLER_RELEASE}
  wait_until_pods_running knative-eventing || return 1
}

function kafka_setup() {
  subheader "Installing Kafka using Strimzi config file"
  kubectl create namespace kafka || return 1
  sed 's/namespace: .*/namespace: kafka/' ${STRIMZI_INSTALLATION_CONFIG_TEMPLATE} > ${STRIMZI_INSTALLATION_CONFIG}
  kubectl apply -f ${STRIMZI_INSTALLATION_CONFIG} -n kafka
  kubectl apply -f ${KAFKA_INSTALLATION_CONFIG} -n kafka
  kubectl apply -f ${KAFKA_TOPIC_INSTALLATION_CONFIG} -n kafka
  wait_until_pods_running kafka || fail_test "Failed to start up a Kafka cluster"
}

function install_sources_crds() {
  subheader "Installing Kafka Source CRD"
  kubectl apply -f ${KAFKA_SOURCE_CONTROLLER_YAML}
  kubectl apply -f ${KAFKA_SOURCE_YAML}
  wait_until_pods_running knative-eventing || fail_test "Failed to install the Kafka Source CRD"
}

function plugin_test_setup() {
  header "🧪  Setup Kafka and source CRDs"
  kafka_setup || return 1
  install_sources_crds || return 1
}

function uninstall_sources_crds() {
  echo "Uninstalling Kafka Source CRD"
  kubectl delete -f ${KAFKA_SOURCE_YAML}
  kubectl delete -f ${KAFKA_SOURCE_CONTROLLER_YAML}
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
