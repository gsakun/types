package v1alpha2

import (
	client "istio.io/api/mixer/v1/config/client"
	v1beta1 "istio.io/api/policy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AttributeManifest struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1beta1.AttributeManifest `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AttributeManifestList is a collection of AttributeManifests.
type AttributeManifestList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []AttributeManifest `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HTTPAPISpec struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec client.HTTPAPISpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPAPISpecList is a collection of HTTPAPISpecs.
type HTTPAPISpecList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []HTTPAPISpec `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPAPISpecBinding defines the binding between HTTPAPISpecs and one or more
// IstioService. For example, the following establishes a binding
// between the HTTPAPISpec `petstore` and service `foo` in namespace `bar`.
type HTTPAPISpecBinding struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec client.HTTPAPISpecBinding `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPAPISpecBindingList is a collection of HTTPAPISpecBindings.
type HTTPAPISpecBindingList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []HTTPAPISpecBinding `json:"items" protobuf:"bytes,2,rep,name=items"`
}


// Each adapter implementation defines its own `params` proto.
//
// In the following example we define a `metrics` handler for the `prometheus` adapter.
// The example is in the form of a Kubernetes resource:
// * The `metadata.name` is the name of the handler
// * The `kind` refers to the adapter name
// * The `spec` block represents adapter-specific configuration as well as the connection information
//
// ```yaml
// # Sample-1: No connection specified (for compiled in adapters)
// # Note: if connection information is not specified, the adapter configuration is directly inside
// # `spec` block. This is going to be DEPRECATED in favor of Sample-2
// apiVersion: "config.istio.io/v1alpha2"
// kind: handler
// metadata:
//   name: requestcount
//   namespace: istio-system
// spec:
//   compiledAdapter: prometheus
//   params:
//     metrics:
//     - name: request_count
//       instance_name: requestcount.metric.istio-system
//       kind: COUNTER
//       label_names:
//       - source_service
//       - source_version
//       - destination_service
//       - destination_version
// ---
// # Sample-2: With connection information (for out-of-process adapters)
// # Note: Unlike sample-1, the adapter configuration is parallel to `connection` and is nested inside `param` block.
// apiVersion: "config.istio.io/v1alpha2"
// kind: handler
// metadata:
//   name: requestcount
//   namespace: istio-system
// spec:
//   compiledAdapter: prometheus
//   params:
//     param:
//       metrics:
//       - name: request_count
//         instance_name: requestcount.metric.istio-system
//         kind: COUNTER
//         label_names:
//         - source_service
//         - source_version
//         - destination_service
//         - destination_version
//     connection:
//       address: localhost:8090
// ---
// ```
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Handler allows the operator to configure a specific adapter implementation.
type Handler struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	Spec HandlerSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type HandlerSpec struct {
	// Must be unique in the entire Mixer configuration. Used by [Actions][istio.policy.v1beta1.Action.handler]
	// to refer to this handler.
	Name string `json:"name,omitempty"`
	// The name of the compiled in adapter this handler instantiates. For referencing non compiled-in
	// adapters, use the `adapter` field instead.
	//
	// The value must match the name of the available adapter Mixer is built with. An adapter's name is typically a
	// constant in its code.
	CompiledAdapter string `json:"compiledAdapter,omitempty"`
	// The name of a specific adapter implementation. For referencing compiled-in
	// adapters, use the `compiled_adapter` field instead.
	//
	// An adapter's implementation name is typically a constant in its code.
	Adapter string `json:"adapter,omitempty"`
	// Depends on adapter implementation. Struct representation of a
	// proto defined by the adapter implementation; this varies depending on the value of field `adapter`.
	Params HandlerParams `json:"params,omitempty"`
	// Information on how to connect to the out-of-process adapter.
	// This is used if the adapter is not compiled into Mixer binary and is running as a separate process.
	//Connection *Connection `protobuf:"bytes,4,opt,name=connection,proto3" json:"connection,omitempty"`
}

type HandlerParams struct {
	RedisServerUrl string `json:"redisServerUrl,omitempty"`
	ConnectionPoolSize int32 `json:"redisServerUrl,omitempty"`
	Quotas  []Quota  `json:"quotas,omitempty"`
}

type Quota struct {
	Name   string  `json:"name"`
	MaxAmount  int32  `json:"maxAmount"`
	ValidDuration string `json:"validDuration"`
	BucketDuration  string  `json:"bucketDuration"`
	RateLimitAlgorithm RateLimitAlgorithm `json:"rateLimitAlgorithm"`
}

type RateLimitAlgorithm string

const (
	ROLLING RateLimitAlgorithm = ROLLING_WINDOW
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HandlerList is a collection of Handlers.
type HandlerList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Handler `json:"items" protobuf:"bytes,2,rep,name=items"`
}


// Instance is defined by the operator. Instance is defined relative to a known
// template. Their purpose is to tell Mixer how to use attributes or literals to produce
// instances of the specified template at runtime.
//
// The following example instructs Mixer to construct an instance associated with template
// 'istio.mixer.adapter.metric.Metric'. It provides a mapping from the template's fields to expressions.
// Instances produced with this instance can be referenced by [Actions][istio.policy.v1beta1.Action] using name
// 'RequestCountByService'
//
// ```yaml
// - name: RequestCountByService
//   template: istio.mixer.adapter.metric.Metric
//   params:
//     value: 1
//     dimensions:
//       source: source.name
//       destination_ip: destination.ip
// ```

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// An Instance tells Mixer how to create instances for particular template.
type Instance struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1beta1.Instance `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InstanceList is a collection of Instances.
type InstanceList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Instance `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Determines the quotas used for individual requests.
type QuotaSpec struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec client.QuotaSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// QuotaSpecList is a collection of QuotaSpecs.
type QuotaSpecList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []QuotaSpec `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type QuotaSpecBinding struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec client.QuotaSpecBinding `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// QuotaSpecBindingList is a collection of QuotaSpecBindings.
type QuotaSpecBindingList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []QuotaSpecBinding `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// The following example instructs Mixer to invoke `prometheus-handler` handler for all services and pass it the
// instance constructed using the 'RequestCountByService' instance.
//
// ```yaml
// - match: match(destination.service.host, "*")
//   actions:
//   - handler: prometheus-handler
//     instances:
//     - RequestCountByService
// ```

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// A Rule is a selector and a set of intentions to be executed when the selector is `true`
type Rule struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1beta1.Rule `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RuleList is a collection of Rules.
type RuleList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Rule `json:"items" protobuf:"bytes,2,rep,name=items"`
}
