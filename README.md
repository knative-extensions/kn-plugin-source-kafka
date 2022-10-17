# kn-plugin-source-kafka

**[This component is GA](https://github.com/knative/community/tree/main/mechanics/MATURITY-LEVELS.md)**

`kn-plugin-source-kafka` is a plugin of Knative Client, for management of kafka event
source interactively from the command line.

## Description

`kn-plugin-source-kafka` is a plugin of Knative Client. You can create, describe and
delete kafka event sources. Go to
[Knative Eventing document](https://knative.dev/docs/eventing/samples/kafka/source/)
to understand more about kafka event sources.

## Build and Install

You must
[set up your development environment](https://github.com/knative/client/blob/master/docs/DEVELOPMENT.md#prerequisites)
before you build `kn-plugin-source-kafka`.

**Building:**

Once you've set up your development environment, let's build `kn-plugin-source-kafka`.
Run below command under the root directory of this repository.

```sh
$ hack/build.sh
```

**Installing:**

You will get an executable file `kn-plugin-source-kafka` under the directory of
`client-contrib/plugins/source-kafka` after you run the build command. Then
let's install it to become a Knative Client `kn` plugin.

Install the plugin by simply copying the executable file `kn-plugin-source-kafka` to
the folder of the `kn` plugins directory. You will be able to invoke it by
`kn source kafka`.

## Usage

### kafka

Knative eventing kafka source plugin

#### Synopsis

Manage Knative kafka eventing sources

#### Options

```
  -h, --help   help for kafka
```

#### SEE ALSO

* [kafka create](#kafka-create)	 - Create a kafka source
* [kafka delete](#kafka-delete)	 - Delete a kafka source
* [kafka describe](#kafka-describe)	 - Describe a kafka source
* [kafka list](#kafka-list)	 - List kafka sources

### kafka create

Create a kafka source

```
kafka create NAME --servers SERVERS --topics TOPICS --sink SINK [flags]
```

#### Examples

```
# Create a new kafka source 'mykafkasrc' which subscribes a kafka server 'my-cluster-kafka-bootstrap.kafka.svc:9092' at topic 'test-topic' and sends the events to service 'event-display'
kn source kafka create mykafkasrc --servers my-cluster-kafka-bootstrap.kafka.svc:9092 --topics test-topic --sink svc:event-display

# Create a new kafka source 'mykafkasrc' which subscribes a kafka server 'my-cluster-kafka-bootstrap.kafka.svc:9092' at topic 'test-topic' using the consumer group ID 'test-consumer-group' and sends the events to service 'event-display'
kn source kafka create mykafkasrc --servers my-cluster-kafka-bootstrap.kafka.svc:9092 --topics test-topic --consumergroup test-consumer-group --sink svc:event-display --ce-override "sink=bound"

```

#### Options

```
  -A, --all-namespaces            If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -a, --annotation stringArray    Metadata annotations to set on the resources. Example: '--annotation key=value' You may be provide this flag multiple times.
      --ce-override stringArray   Cloud Event overrides to apply before sending event to sink. Example: '--ce-override key=value' You may be provide this flag multiple times. To unset, append "-" to the key (e.g. --ce-override key-).
      --consumergroup string      the consumer group ID
  -h, --help                      help for create
  -l, --label stringArray         Metadata labels to set on the resources. Example: '--label key=value' You may be provide this flag multiple times.
  -n, --namespace string          Specify the namespace to operate in.
      --servers stringArray       Kafka bootstrap servers that the consumer will connect to, consist of a hostname plus a port pair, e.g. my-kafka-bootstrap.kafka:9092. Flag can be used multiple times.
  -s, --sink string               Addressable sink for events. You can specify a broker, channel, Knative service or URI. Examples: '--sink broker:nest' for a broker 'nest', '--sink channel:pipe' for a channel 'pipe', '--sink ksvc:mysvc:mynamespace' for a Knative service 'mysvc' in another namespace 'mynamespace', '--sink https://event.receiver.uri' for an URI with an 'http://' or 'https://' schema, '--sink ksvc:receiver' or simply '--sink receiver' for a Knative service 'receiver' in the current namespace. If a prefix is not provided, it is considered as a Knative service in the current namespace. If referring to a Knative service in another namespace, 'ksvc:name:namespace' combination must be provided explicitly.
      --topics stringArray        Topics to consume messages from. Flag can be used multiple times.
```

#### SEE ALSO

* [kafka](#kafka)	 - Knative eventing kafka source plugin

### kafka delete

Delete a kafka source

```
kafka delete NAME [flags]
```

#### Examples

```
# Delete a kafka source with name 'mykafkasrc'
kn source kafka delete mykafkasrc
```

#### Options

```
  -A, --all-namespaces     If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -h, --help               help for delete
  -n, --namespace string   Specify the namespace to operate in.
```

#### SEE ALSO

* [kafka](#kafka)	 - Knative eventing kafka source plugin

### kafka describe

Describe a kafka source

```
kafka describe NAME [flags]
```

#### Examples

```
# Describe a kafka source with NAME
kn source kafka describe kafka-name
```

#### Options

```
  -A, --all-namespaces     If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -h, --help               help for describe
  -n, --namespace string   Specify the namespace to operate in.
```

#### SEE ALSO

* [kafka](#kafka)	 - Knative eventing kafka source plugin

### kafka list

List kafka sources

```
kafka list [flags]
```

#### Examples

```
# List the available kafka sources
kn source kafka list
```

#### Options

```
  -A, --all-namespaces     If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -h, --help               help for list
  -n, --namespace string   Specify the namespace to operate in.
```

#### SEE ALSO

* [kafka](#kafka)	 - Knative eventing kafka source plugin

## More information
	
* [Knative Client](https://github.com/knative/client)
* [How to contribute a plugin](https://github.com/knative/client-contrib#how-to-contribute-a-plugin)

