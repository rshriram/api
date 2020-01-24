// Code generated by protoc-gen-go. DO NOT EDIT.
// source: operator/v1alpha1/component.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	v2beta1 "k8s.io/api/autoscaling/v2beta1"
	v1 "k8s.io/api/core/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// IstioComponentSpec defines the desired installed state of Istio components.
type IstioComponentSetSpec struct {
	Base                 *BaseComponentSpec `protobuf:"bytes,29,opt,name=base,proto3" json:"base,omitempty"`
	Pilot                *ComponentSpec     `protobuf:"bytes,30,opt,name=pilot,proto3" json:"pilot,omitempty"`
	Proxy                *ComponentSpec     `protobuf:"bytes,31,opt,name=proxy,proto3" json:"proxy,omitempty"`
	SidecarInjector      *ComponentSpec     `protobuf:"bytes,32,opt,name=sidecar_injector,json=sidecarInjector,proto3" json:"sidecar_injector,omitempty"`
	Policy               *ComponentSpec     `protobuf:"bytes,33,opt,name=policy,proto3" json:"policy,omitempty"`
	Telemetry            *ComponentSpec     `protobuf:"bytes,34,opt,name=telemetry,proto3" json:"telemetry,omitempty"`
	Citadel              *ComponentSpec     `protobuf:"bytes,35,opt,name=citadel,proto3" json:"citadel,omitempty"`
	NodeAgent            *ComponentSpec     `protobuf:"bytes,36,opt,name=node_agent,json=nodeAgent,proto3" json:"node_agent,omitempty"`
	Galley               *ComponentSpec     `protobuf:"bytes,37,opt,name=galley,proto3" json:"galley,omitempty"`
	Cni                  *ComponentSpec     `protobuf:"bytes,38,opt,name=cni,proto3" json:"cni,omitempty"`
	IngressGateways      []*GatewaySpec     `protobuf:"bytes,40,rep,name=ingress_gateways,json=ingressGateways,proto3" json:"ingress_gateways,omitempty"`
	EgressGateways       []*GatewaySpec     `protobuf:"bytes,41,rep,name=egress_gateways,json=egressGateways,proto3" json:"egress_gateways,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IstioComponentSetSpec) Reset()         { *m = IstioComponentSetSpec{} }
func (m *IstioComponentSetSpec) String() string { return proto.CompactTextString(m) }
func (*IstioComponentSetSpec) ProtoMessage()    {}
func (*IstioComponentSetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{0}
}

func (m *IstioComponentSetSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IstioComponentSetSpec.Unmarshal(m, b)
}
func (m *IstioComponentSetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IstioComponentSetSpec.Marshal(b, m, deterministic)
}
func (m *IstioComponentSetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioComponentSetSpec.Merge(m, src)
}
func (m *IstioComponentSetSpec) XXX_Size() int {
	return xxx_messageInfo_IstioComponentSetSpec.Size(m)
}
func (m *IstioComponentSetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioComponentSetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IstioComponentSetSpec proto.InternalMessageInfo

func (m *IstioComponentSetSpec) GetBase() *BaseComponentSpec {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *IstioComponentSetSpec) GetPilot() *ComponentSpec {
	if m != nil {
		return m.Pilot
	}
	return nil
}

func (m *IstioComponentSetSpec) GetProxy() *ComponentSpec {
	if m != nil {
		return m.Proxy
	}
	return nil
}

func (m *IstioComponentSetSpec) GetSidecarInjector() *ComponentSpec {
	if m != nil {
		return m.SidecarInjector
	}
	return nil
}

func (m *IstioComponentSetSpec) GetPolicy() *ComponentSpec {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *IstioComponentSetSpec) GetTelemetry() *ComponentSpec {
	if m != nil {
		return m.Telemetry
	}
	return nil
}

func (m *IstioComponentSetSpec) GetCitadel() *ComponentSpec {
	if m != nil {
		return m.Citadel
	}
	return nil
}

func (m *IstioComponentSetSpec) GetNodeAgent() *ComponentSpec {
	if m != nil {
		return m.NodeAgent
	}
	return nil
}

func (m *IstioComponentSetSpec) GetGalley() *ComponentSpec {
	if m != nil {
		return m.Galley
	}
	return nil
}

func (m *IstioComponentSetSpec) GetCni() *ComponentSpec {
	if m != nil {
		return m.Cni
	}
	return nil
}

func (m *IstioComponentSetSpec) GetIngressGateways() []*GatewaySpec {
	if m != nil {
		return m.IngressGateways
	}
	return nil
}

func (m *IstioComponentSetSpec) GetEgressGateways() []*GatewaySpec {
	if m != nil {
		return m.EgressGateways
	}
	return nil
}

// Configuration for base component.
type BaseComponentSpec struct {
	// Selects whether this component is installed.
	Enabled              *BoolValueForPB `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BaseComponentSpec) Reset()         { *m = BaseComponentSpec{} }
func (m *BaseComponentSpec) String() string { return proto.CompactTextString(m) }
func (*BaseComponentSpec) ProtoMessage()    {}
func (*BaseComponentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{1}
}

func (m *BaseComponentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseComponentSpec.Unmarshal(m, b)
}
func (m *BaseComponentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseComponentSpec.Marshal(b, m, deterministic)
}
func (m *BaseComponentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseComponentSpec.Merge(m, src)
}
func (m *BaseComponentSpec) XXX_Size() int {
	return xxx_messageInfo_BaseComponentSpec.Size(m)
}
func (m *BaseComponentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseComponentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_BaseComponentSpec proto.InternalMessageInfo


// Configuration for internal components.
type ComponentSpec struct {
	// Selects whether this component is installed.
	Enabled *BoolValueForPB `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Namespace for the component.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Hub for the component (overrides top level hub setting).
	Hub string `protobuf:"bytes,10,opt,name=hub,proto3" json:"hub,omitempty"`
	// Tag for the component (overrides top level tag setting).
	Tag string `protobuf:"bytes,11,opt,name=tag,proto3" json:"tag,omitempty"`
	// Arbitrary install time configuration for the component.
	Spec interface{} `protobuf:"bytes,30,opt,name=spec,proto3" json:"spec,omitempty"`
	// Kubernetes resource spec.
	K8S                  *KubernetesResourcesSpec `protobuf:"bytes,50,opt,name=k8s,proto3" json:"k8s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ComponentSpec) Reset()         { *m = ComponentSpec{} }
func (m *ComponentSpec) String() string { return proto.CompactTextString(m) }
func (*ComponentSpec) ProtoMessage()    {}
func (*ComponentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{2}
}

func (m *ComponentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentSpec.Unmarshal(m, b)
}
func (m *ComponentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentSpec.Marshal(b, m, deterministic)
}
func (m *ComponentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentSpec.Merge(m, src)
}
func (m *ComponentSpec) XXX_Size() int {
	return xxx_messageInfo_ComponentSpec.Size(m)
}
func (m *ComponentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentSpec proto.InternalMessageInfo


func (m *ComponentSpec) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ComponentSpec) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

func (m *ComponentSpec) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}


func (m *ComponentSpec) GetK8S() *KubernetesResourcesSpec {
	if m != nil {
		return m.K8S
	}
	return nil
}

// Configuration for external components.
type ExternalComponentSpec struct {
	// Selects whether this component is installed.
	Enabled *BoolValueForPB `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Namespace for the component.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Arbitrary install time configuration for the component.
	Spec interface{} `protobuf:"bytes,10,opt,name=spec,proto3" json:"spec,omitempty"`
	// Chart path for addon components.
	ChartPath string `protobuf:"bytes,30,opt,name=chart_path,json=chartPath,proto3" json:"chart_path,omitempty"`
	// Optional schema to validate spec against.
	Schema *any.Any `protobuf:"bytes,35,opt,name=schema,proto3" json:"schema,omitempty"`
	// Kubernetes resource spec.
	K8S                  *KubernetesResourcesSpec `protobuf:"bytes,50,opt,name=k8s,proto3" json:"k8s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ExternalComponentSpec) Reset()         { *m = ExternalComponentSpec{} }
func (m *ExternalComponentSpec) String() string { return proto.CompactTextString(m) }
func (*ExternalComponentSpec) ProtoMessage()    {}
func (*ExternalComponentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{3}
}

func (m *ExternalComponentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalComponentSpec.Unmarshal(m, b)
}
func (m *ExternalComponentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalComponentSpec.Marshal(b, m, deterministic)
}
func (m *ExternalComponentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalComponentSpec.Merge(m, src)
}
func (m *ExternalComponentSpec) XXX_Size() int {
	return xxx_messageInfo_ExternalComponentSpec.Size(m)
}
func (m *ExternalComponentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalComponentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalComponentSpec proto.InternalMessageInfo


func (m *ExternalComponentSpec) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}


func (m *ExternalComponentSpec) GetChartPath() string {
	if m != nil {
		return m.ChartPath
	}
	return ""
}

func (m *ExternalComponentSpec) GetSchema() *any.Any {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *ExternalComponentSpec) GetK8S() *KubernetesResourcesSpec {
	if m != nil {
		return m.K8S
	}
	return nil
}

// Configuration for gateways.
type GatewaySpec struct {
	// Selects whether this gateway is installed.
	Enabled *BoolValueForPB `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Namespace for the gateway.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name for the gateway.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Labels for the gateway.
	Label map[string]string `protobuf:"bytes,4,rep,name=label,proto3" json:"label,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Hub for the component (overrides top level hub setting).
	Hub string `protobuf:"bytes,10,opt,name=hub,proto3" json:"hub,omitempty"`
	// Tag for the component (overrides top level tag setting).
	Tag string `protobuf:"bytes,11,opt,name=tag,proto3" json:"tag,omitempty"`
	// Kubernetes resource spec.
	K8S                  *KubernetesResourcesSpec `protobuf:"bytes,50,opt,name=k8s,proto3" json:"k8s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GatewaySpec) Reset()         { *m = GatewaySpec{} }
func (m *GatewaySpec) String() string { return proto.CompactTextString(m) }
func (*GatewaySpec) ProtoMessage()    {}
func (*GatewaySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{4}
}

func (m *GatewaySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewaySpec.Unmarshal(m, b)
}
func (m *GatewaySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewaySpec.Marshal(b, m, deterministic)
}
func (m *GatewaySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewaySpec.Merge(m, src)
}
func (m *GatewaySpec) XXX_Size() int {
	return xxx_messageInfo_GatewaySpec.Size(m)
}
func (m *GatewaySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewaySpec.DiscardUnknown(m)
}

var xxx_messageInfo_GatewaySpec proto.InternalMessageInfo


func (m *GatewaySpec) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GatewaySpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GatewaySpec) GetLabel() map[string]string {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *GatewaySpec) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

func (m *GatewaySpec) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *GatewaySpec) GetK8S() *KubernetesResourcesSpec {
	if m != nil {
		return m.K8S
	}
	return nil
}

// KubernetesResourcesConfig is a common set of k8s resource configs for components.
type KubernetesResourcesSpec struct {
	// k8s affinity.
	// https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
	Affinity *Affinity `protobuf:"bytes,1,opt,name=affinity,proto3" json:"affinity,omitempty"`
	// Deployment environment variables.
	// https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/
	Env []*EnvVar `protobuf:"bytes,2,rep,name=env,proto3" json:"env,omitempty"`
	// k8s HorizontalPodAutoscaler settings.
	// https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/
	HpaSpec *v2beta1.HorizontalPodAutoscalerSpec `protobuf:"bytes,3,opt,name=hpa_spec,json=hpaSpec,proto3" json:"hpa_spec,omitempty"`
	// k8s imagePullPolicy.
	// https://kubernetes.io/docs/concepts/containers/images/
	ImagePullPolicy string `protobuf:"bytes,4,opt,name=image_pull_policy,json=imagePullPolicy,proto3" json:"image_pull_policy,omitempty"`
	// k8s nodeSelector.
	// https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector
	NodeSelector map[string]string `protobuf:"bytes,5,rep,name=node_selector,json=nodeSelector,proto3" json:"node_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// k8s PodDisruptionBudget settings.
	// https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#how-disruption-budgets-work
	PodDisruptionBudget *PodDisruptionBudgetSpec `protobuf:"bytes,6,opt,name=pod_disruption_budget,json=podDisruptionBudget,proto3" json:"pod_disruption_budget,omitempty"`
	// k8s pod annotations.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
	PodAnnotations map[string]string `protobuf:"bytes,7,rep,name=pod_annotations,json=podAnnotations,proto3" json:"pod_annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// k8s priority_class_name. Default for all resources unless overridden.
	// https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
	PriorityClassName string `protobuf:"bytes,8,opt,name=priority_class_name,json=priorityClassName,proto3" json:"priority_class_name,omitempty"`
	// k8s readinessProbe settings.
	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/
	// k8s.io.api.core.v1.Probe readiness_probe = 9;
	ReadinessProbe *ReadinessProbe `protobuf:"bytes,9,opt,name=readiness_probe,json=readinessProbe,proto3" json:"readiness_probe,omitempty"`
	// k8s Deployment replicas setting.
	// https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
	ReplicaCount uint32 `protobuf:"varint,10,opt,name=replica_count,json=replicaCount,proto3" json:"replica_count,omitempty"`
	// k8s resources settings.
	// https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#resource-requests-and-limits-of-pod-and-container
	Resources *Resources `protobuf:"bytes,11,opt,name=resources,proto3" json:"resources,omitempty"`
	// k8s Service settings.
	// https://kubernetes.io/docs/concepts/services-networking/service/
	Service *ServiceSpec `protobuf:"bytes,12,opt,name=service,proto3" json:"service,omitempty"`
	// k8s deployment strategy.
	// https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
	Strategy *DeploymentStrategy `protobuf:"bytes,13,opt,name=strategy,proto3" json:"strategy,omitempty"`
	// k8s toleration
	// https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
	Tolerations []*v1.Toleration `protobuf:"bytes,14,rep,name=tolerations,proto3" json:"tolerations,omitempty"`
	// k8s service annotations.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
	ServiceAnnotations map[string]string `protobuf:"bytes,15,rep,name=service_annotations,json=serviceAnnotations,proto3" json:"service_annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Overlays for k8s resources in rendered manifests.
	Overlays             []*K8SObjectOverlay `protobuf:"bytes,100,rep,name=overlays,proto3" json:"overlays,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *KubernetesResourcesSpec) Reset()         { *m = KubernetesResourcesSpec{} }
func (m *KubernetesResourcesSpec) String() string { return proto.CompactTextString(m) }
func (*KubernetesResourcesSpec) ProtoMessage()    {}
func (*KubernetesResourcesSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{5}
}

func (m *KubernetesResourcesSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesResourcesSpec.Unmarshal(m, b)
}
func (m *KubernetesResourcesSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesResourcesSpec.Marshal(b, m, deterministic)
}
func (m *KubernetesResourcesSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesResourcesSpec.Merge(m, src)
}
func (m *KubernetesResourcesSpec) XXX_Size() int {
	return xxx_messageInfo_KubernetesResourcesSpec.Size(m)
}
func (m *KubernetesResourcesSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesResourcesSpec.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesResourcesSpec proto.InternalMessageInfo

func (m *KubernetesResourcesSpec) GetAffinity() *Affinity {
	if m != nil {
		return m.Affinity
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetEnv() []*EnvVar {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetHpaSpec() *v2beta1.HorizontalPodAutoscalerSpec {
	if m != nil {
		return m.HpaSpec
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetImagePullPolicy() string {
	if m != nil {
		return m.ImagePullPolicy
	}
	return ""
}

func (m *KubernetesResourcesSpec) GetNodeSelector() map[string]string {
	if m != nil {
		return m.NodeSelector
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetPodDisruptionBudget() *PodDisruptionBudgetSpec {
	if m != nil {
		return m.PodDisruptionBudget
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetPodAnnotations() map[string]string {
	if m != nil {
		return m.PodAnnotations
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetPriorityClassName() string {
	if m != nil {
		return m.PriorityClassName
	}
	return ""
}

func (m *KubernetesResourcesSpec) GetReadinessProbe() *ReadinessProbe {
	if m != nil {
		return m.ReadinessProbe
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetReplicaCount() uint32 {
	if m != nil {
		return m.ReplicaCount
	}
	return 0
}

func (m *KubernetesResourcesSpec) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetService() *ServiceSpec {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetStrategy() *DeploymentStrategy {
	if m != nil {
		return m.Strategy
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetTolerations() []*v1.Toleration {
	if m != nil {
		return m.Tolerations
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetServiceAnnotations() map[string]string {
	if m != nil {
		return m.ServiceAnnotations
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetOverlays() []*K8SObjectOverlay {
	if m != nil {
		return m.Overlays
	}
	return nil
}

// Patch for an existing k8s resource.
type K8SObjectOverlay struct {
	// Resource API version.
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// Resource kind.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Name of resource.
	// Namespace is always the component namespace.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// List of patches to apply to resource.
	Patches              []*K8SObjectOverlay_PathValue `protobuf:"bytes,4,rep,name=patches,proto3" json:"patches,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *K8SObjectOverlay) Reset()         { *m = K8SObjectOverlay{} }
func (m *K8SObjectOverlay) String() string { return proto.CompactTextString(m) }
func (*K8SObjectOverlay) ProtoMessage()    {}
func (*K8SObjectOverlay) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{6}
}

func (m *K8SObjectOverlay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K8SObjectOverlay.Unmarshal(m, b)
}
func (m *K8SObjectOverlay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K8SObjectOverlay.Marshal(b, m, deterministic)
}
func (m *K8SObjectOverlay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K8SObjectOverlay.Merge(m, src)
}
func (m *K8SObjectOverlay) XXX_Size() int {
	return xxx_messageInfo_K8SObjectOverlay.Size(m)
}
func (m *K8SObjectOverlay) XXX_DiscardUnknown() {
	xxx_messageInfo_K8SObjectOverlay.DiscardUnknown(m)
}

var xxx_messageInfo_K8SObjectOverlay proto.InternalMessageInfo

func (m *K8SObjectOverlay) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *K8SObjectOverlay) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *K8SObjectOverlay) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *K8SObjectOverlay) GetPatches() []*K8SObjectOverlay_PathValue {
	if m != nil {
		return m.Patches
	}
	return nil
}

type K8SObjectOverlay_PathValue struct {
	// Path of the form a.[key1:value1].b.[:value2]
	// Where [key1:value1] is a selector for a key-value pair to identify a list element and [:value] is a value
	// selector to identify a list element in a leaf list.
	// All path intermediate nodes must exist.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Value to add, delete or replace.
	// For add, the path should be a new leaf.
	// For delete, value should be unset.
	// For replace, path should reference an existing node.
	// All values are strings but are converted into appropriate type based on schema.
	Value                interface{} `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *K8SObjectOverlay_PathValue) Reset()         { *m = K8SObjectOverlay_PathValue{} }
func (m *K8SObjectOverlay_PathValue) String() string { return proto.CompactTextString(m) }
func (*K8SObjectOverlay_PathValue) ProtoMessage()    {}
func (*K8SObjectOverlay_PathValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{6, 0}
}

func (m *K8SObjectOverlay_PathValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K8SObjectOverlay_PathValue.Unmarshal(m, b)
}
func (m *K8SObjectOverlay_PathValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K8SObjectOverlay_PathValue.Marshal(b, m, deterministic)
}
func (m *K8SObjectOverlay_PathValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K8SObjectOverlay_PathValue.Merge(m, src)
}
func (m *K8SObjectOverlay_PathValue) XXX_Size() int {
	return xxx_messageInfo_K8SObjectOverlay_PathValue.Size(m)
}
func (m *K8SObjectOverlay_PathValue) XXX_DiscardUnknown() {
	xxx_messageInfo_K8SObjectOverlay_PathValue.DiscardUnknown(m)
}

var xxx_messageInfo_K8SObjectOverlay_PathValue proto.InternalMessageInfo

func (m *K8SObjectOverlay_PathValue) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}





func init() {
	proto.RegisterType((*IstioComponentSetSpec)(nil), "istio.operator.v1alpha1.IstioComponentSetSpec")
	proto.RegisterType((*BaseComponentSpec)(nil), "istio.operator.v1alpha1.BaseComponentSpec")
	proto.RegisterType((*ComponentSpec)(nil), "istio.operator.v1alpha1.ComponentSpec")
	proto.RegisterType((*ExternalComponentSpec)(nil), "istio.operator.v1alpha1.ExternalComponentSpec")
	proto.RegisterType((*GatewaySpec)(nil), "istio.operator.v1alpha1.GatewaySpec")
	proto.RegisterMapType((map[string]string)(nil), "istio.operator.v1alpha1.GatewaySpec.LabelEntry")
	proto.RegisterType((*KubernetesResourcesSpec)(nil), "istio.operator.v1alpha1.KubernetesResourcesSpec")
	proto.RegisterMapType((map[string]string)(nil), "istio.operator.v1alpha1.KubernetesResourcesSpec.NodeSelectorEntry")
	proto.RegisterMapType((map[string]string)(nil), "istio.operator.v1alpha1.KubernetesResourcesSpec.PodAnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "istio.operator.v1alpha1.KubernetesResourcesSpec.ServiceAnnotationsEntry")
	proto.RegisterType((*K8SObjectOverlay)(nil), "istio.operator.v1alpha1.K8sObjectOverlay")
	proto.RegisterType((*K8SObjectOverlay_PathValue)(nil), "istio.operator.v1alpha1.K8sObjectOverlay.PathValue")
}

func init() { proto.RegisterFile("operator/v1alpha1/component.proto", fileDescriptor_6ed34a579e9b43a2) }

var fileDescriptor_6ed34a579e9b43a2 = []byte{
	// 1261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x97, 0xcf, 0x72, 0x13, 0x47,
	0x13, 0xc0, 0xcb, 0x92, 0xff, 0xa9, 0x85, 0x2d, 0x7b, 0x0c, 0x1f, 0xfb, 0xb9, 0x3e, 0xc0, 0x08,
	0x3e, 0x62, 0x48, 0x6a, 0x15, 0x9b, 0x8b, 0x8a, 0x24, 0x04, 0x0b, 0x2b, 0x40, 0x25, 0x80, 0xb2,
	0xa6, 0x7c, 0xa0, 0x2a, 0xa5, 0x1a, 0xed, 0xb6, 0x57, 0x13, 0xaf, 0x66, 0xa6, 0x66, 0x46, 0x0a,
	0x9b, 0x37, 0xc8, 0x25, 0xe7, 0xbc, 0x52, 0x4e, 0x79, 0x94, 0xbc, 0x42, 0x6a, 0x66, 0x57, 0x92,
	0x8d, 0x50, 0xf0, 0xa6, 0x48, 0x72, 0x9b, 0xed, 0xe9, 0x5f, 0xcf, 0x74, 0x6f, 0xf7, 0x76, 0x2f,
	0xdc, 0x14, 0x12, 0x15, 0x35, 0x42, 0x35, 0x46, 0x7b, 0x34, 0x91, 0x7d, 0xba, 0xd7, 0x08, 0xc5,
	0x40, 0x0a, 0x8e, 0xdc, 0xf8, 0x52, 0x09, 0x23, 0xc8, 0x55, 0xa6, 0x0d, 0x13, 0xfe, 0x58, 0xd1,
	0x1f, 0x2b, 0x6e, 0xff, 0x37, 0x16, 0x22, 0x4e, 0xb0, 0xe1, 0xd4, 0x7a, 0xc3, 0x93, 0x06, 0xe5,
	0x69, 0xc6, 0x6c, 0xfb, 0xa7, 0x4d, 0xed, 0x33, 0xd1, 0xa0, 0x92, 0x35, 0xe8, 0xd0, 0x08, 0x1d,
	0xd2, 0x84, 0xf1, 0xb8, 0x31, 0xda, 0xef, 0xa1, 0xa1, 0x7b, 0x8d, 0x18, 0xb9, 0xb5, 0x86, 0x51,
	0xae, 0x5f, 0x3f, 0xa3, 0x1f, 0x0a, 0x85, 0x8d, 0xd1, 0x3b, 0x74, 0x66, 0xaf, 0x7a, 0x3a, 0xec,
	0xa1, 0xe2, 0x68, 0x50, 0x67, 0x3a, 0xf5, 0xdf, 0x96, 0xe1, 0xca, 0x33, 0x7b, 0xdd, 0xc7, 0x63,
	0x27, 0x8e, 0xd0, 0x1c, 0x49, 0x0c, 0xc9, 0x43, 0x58, 0xec, 0x51, 0x8d, 0xde, 0xb5, 0x9d, 0x85,
	0xdd, 0xea, 0xfe, 0x3d, 0x7f, 0x8e, 0x53, 0x7e, 0x8b, 0x6a, 0x9c, 0xc2, 0x12, 0xc3, 0xc0, 0x71,
	0xe4, 0x73, 0x58, 0x92, 0x2c, 0x11, 0xc6, 0xbb, 0xee, 0x0c, 0xdc, 0x99, 0x6b, 0xe0, 0x3c, 0x9c,
	0x41, 0x8e, 0x56, 0xe2, 0x4d, 0xea, 0xdd, 0x28, 0x48, 0x5b, 0x88, 0x7c, 0x0b, 0x1b, 0x9a, 0x45,
	0x18, 0x52, 0xd5, 0x65, 0xfc, 0x7b, 0x0c, 0x8d, 0x50, 0xde, 0x4e, 0x21, 0x43, 0xb5, 0x9c, 0x7f,
	0x96, 0xe3, 0xe4, 0x21, 0x2c, 0x4b, 0x91, 0xb0, 0x30, 0xf5, 0x6e, 0x16, 0x32, 0x94, 0x53, 0xe4,
	0x10, 0x2a, 0x06, 0x13, 0x1c, 0xa0, 0x51, 0xa9, 0x57, 0x2f, 0x64, 0x62, 0x0a, 0x92, 0x47, 0xb0,
	0x12, 0x32, 0x43, 0x23, 0x4c, 0xbc, 0x5b, 0x85, 0x6c, 0x8c, 0x31, 0xd2, 0x06, 0xe0, 0x22, 0xc2,
	0x2e, 0x8d, 0x91, 0x1b, 0xef, 0x76, 0xb1, 0x8b, 0x58, 0xf2, 0xc0, 0x82, 0x36, 0x1c, 0x31, 0x4d,
	0x12, 0x4c, 0xbd, 0xff, 0x17, 0x0b, 0x47, 0x46, 0x91, 0x26, 0x94, 0x43, 0xce, 0xbc, 0x3b, 0x85,
	0x60, 0x8b, 0x90, 0x97, 0xb0, 0xc1, 0x78, 0xac, 0x50, 0xeb, 0x6e, 0x4c, 0x0d, 0xfe, 0x40, 0x53,
	0xed, 0xed, 0xee, 0x94, 0x77, 0xab, 0xfb, 0xb7, 0xe7, 0x9a, 0x79, 0x92, 0x29, 0x66, 0x6f, 0x36,
	0xa7, 0x73, 0x99, 0x26, 0xcf, 0xa1, 0x86, 0x6f, 0xd9, 0xbb, 0x5b, 0xc0, 0xde, 0x3a, 0x9e, 0x33,
	0x57, 0x7f, 0x0d, 0x9b, 0x33, 0x25, 0x41, 0xda, 0xb0, 0x82, 0x9c, 0xf6, 0x12, 0x8c, 0xbc, 0x05,
	0xe7, 0xf2, 0xc7, 0x73, 0x6d, 0xbf, 0x4a, 0x25, 0xb6, 0x84, 0x48, 0x8e, 0x69, 0x32, 0xc4, 0xaf,
	0x84, 0xea, 0xb4, 0x82, 0x31, 0x5b, 0xff, 0xa5, 0x04, 0x6b, 0x7f, 0x87, 0x61, 0xf2, 0x3f, 0xa8,
	0x70, 0x3a, 0x40, 0x2d, 0x69, 0x88, 0x5e, 0x69, 0x67, 0x61, 0xb7, 0x12, 0x4c, 0x05, 0x64, 0x03,
	0xca, 0xfd, 0x61, 0xcf, 0x03, 0x27, 0xb7, 0x4b, 0x2b, 0x31, 0x34, 0xf6, 0xaa, 0x99, 0xc4, 0xd0,
	0x98, 0x3c, 0x80, 0x45, 0x2d, 0x31, 0x7c, 0x6f, 0xb5, 0xdb, 0x5b, 0x3c, 0xe3, 0x06, 0xd5, 0x09,
	0x0d, 0x31, 0x70, 0x0c, 0x69, 0x41, 0xf9, 0xb4, 0xa9, 0xbd, 0x7d, 0x87, 0x7e, 0x3a, 0x17, 0xfd,
	0x7a, 0xf2, 0xf1, 0x0a, 0x50, 0x8b, 0xa1, 0x0a, 0x51, 0x67, 0x69, 0x71, 0xda, 0xd4, 0xf5, 0x5f,
	0x4b, 0x70, 0xa5, 0xfd, 0xc6, 0xa0, 0xe2, 0x34, 0xf9, 0x17, 0x42, 0x34, 0x76, 0x1f, 0xfe, 0x82,
	0xfb, 0xd7, 0x00, 0xc2, 0x3e, 0x55, 0xa6, 0x2b, 0xa9, 0xe9, 0xbb, 0x00, 0x56, 0x82, 0x8a, 0x93,
	0x74, 0xa8, 0xe9, 0x93, 0x4f, 0x60, 0x59, 0x87, 0x7d, 0x1c, 0xd0, 0xbc, 0xe4, 0x2f, 0xfb, 0x59,
	0x1b, 0xf1, 0xc7, 0x6d, 0xc4, 0x3f, 0xe0, 0x69, 0x90, 0xeb, 0x7c, 0x90, 0x58, 0xfe, 0x5e, 0x82,
	0xea, 0x99, 0x14, 0xff, 0x67, 0x22, 0x48, 0x60, 0xd1, 0x3e, 0x78, 0x65, 0xb7, 0xe1, 0xd6, 0xa4,
	0x0d, 0x4b, 0x09, 0xed, 0x61, 0xe2, 0x2d, 0xba, 0x82, 0x6c, 0x5c, 0xa4, 0x20, 0xfd, 0x6f, 0x2c,
	0xd1, 0xe6, 0x46, 0xa5, 0x41, 0x46, 0x5f, 0x28, 0x7f, 0x3f, 0x40, 0xdc, 0xb6, 0x9b, 0x00, 0xd3,
	0xc3, 0xed, 0x19, 0xa7, 0x98, 0xba, 0x88, 0x55, 0x02, 0xbb, 0x24, 0x97, 0x61, 0x69, 0x64, 0xe3,
	0x92, 0x3b, 0x9f, 0x3d, 0x3c, 0x28, 0x35, 0x17, 0xea, 0x3f, 0x57, 0xe1, 0xea, 0x1c, 0xd3, 0xe4,
	0x0b, 0x58, 0xa5, 0x27, 0x27, 0x8c, 0x33, 0x93, 0xe6, 0xe1, 0xbf, 0x39, 0xf7, 0x7a, 0x07, 0xb9,
	0x62, 0x30, 0x41, 0xc8, 0x1e, 0x94, 0x91, 0x8f, 0xbc, 0x92, 0x8b, 0xe0, 0x8d, 0xb9, 0x64, 0x9b,
	0x8f, 0x8e, 0xa9, 0x0a, 0xac, 0x2e, 0x39, 0x86, 0xd5, 0xbe, 0xa4, 0x5d, 0x97, 0xd0, 0x65, 0x77,
	0xe2, 0x67, 0xf9, 0x7c, 0xe2, 0x53, 0xc9, 0xfc, 0x33, 0xf3, 0x89, 0x9f, 0xcf, 0x27, 0xfe, 0x53,
	0xa1, 0xd8, 0x8f, 0x82, 0x1b, 0x9a, 0x74, 0x44, 0x74, 0x90, 0x2b, 0xa0, 0xca, 0x7a, 0x4f, 0x5f,
	0x52, 0xe7, 0xc9, 0x3d, 0xd8, 0x64, 0x03, 0x1a, 0x63, 0x57, 0x0e, 0x93, 0xa4, 0x9b, 0xb7, 0xd3,
	0x45, 0x17, 0x8b, 0x9a, 0xdb, 0xe8, 0x0c, 0x93, 0xa4, 0x93, 0xf5, 0xcb, 0x18, 0xd6, 0x5c, 0x9f,
	0xd2, 0x98, 0x64, 0xfd, 0x7b, 0xc9, 0x39, 0xd0, 0x2a, 0xfa, 0x66, 0xfc, 0x17, 0x22, 0xc2, 0xa3,
	0xdc, 0x48, 0x96, 0x15, 0x97, 0xf8, 0x19, 0x11, 0x89, 0xe0, 0x8a, 0x14, 0x51, 0x37, 0x62, 0x5a,
	0x0d, 0xa5, 0x61, 0x82, 0x77, 0x7b, 0xc3, 0x28, 0x46, 0xe3, 0x2d, 0xbf, 0x27, 0x15, 0x3a, 0x22,
	0x3a, 0x9c, 0x40, 0x2d, 0xc7, 0x38, 0x77, 0xb7, 0xe4, 0xec, 0x06, 0x19, 0x40, 0xcd, 0x9e, 0x42,
	0x39, 0x17, 0x86, 0x5a, 0xb9, 0xf6, 0x56, 0x9c, 0x43, 0x87, 0x85, 0x1d, 0xb2, 0x01, 0x9e, 0x9a,
	0xc9, 0x5c, 0x5a, 0x97, 0xe7, 0x84, 0xc4, 0x87, 0x2d, 0xa9, 0x98, 0x50, 0xcc, 0xa4, 0xdd, 0x30,
	0xa1, 0x5a, 0x77, 0x5d, 0x6d, 0xad, 0xba, 0x58, 0x6f, 0x8e, 0xb7, 0x1e, 0xdb, 0x9d, 0x17, 0xb6,
	0xd0, 0x3a, 0x50, 0x53, 0x48, 0x23, 0xc6, 0x6d, 0x1b, 0x94, 0x4a, 0xf4, 0xd0, 0xab, 0x38, 0xf7,
	0x3f, 0x9a, 0x7b, 0xbd, 0x60, 0xac, 0xdf, 0xb1, 0xea, 0xc1, 0xba, 0x3a, 0xf7, 0x4c, 0x6e, 0xc1,
	0x9a, 0x42, 0x99, 0xb0, 0x90, 0x76, 0x43, 0x31, 0xe4, 0xc6, 0x55, 0xdf, 0x5a, 0x70, 0x29, 0x17,
	0x3e, 0xb6, 0x32, 0xf2, 0x08, 0x2a, 0x6a, 0xec, 0x9b, 0x2b, 0xc6, 0xea, 0x7e, 0xfd, 0x4f, 0x0e,
	0xcc, 0x35, 0x83, 0x29, 0x44, 0x1e, 0xc2, 0x8a, 0x46, 0x35, 0x62, 0x21, 0x7a, 0x97, 0x1c, 0x3f,
	0xbf, 0x69, 0x1f, 0x65, 0x7a, 0x59, 0x4a, 0xe6, 0x10, 0x79, 0x02, 0xab, 0xda, 0xd8, 0xa1, 0x39,
	0x4e, 0xbd, 0xb5, 0xf7, 0x7c, 0xdb, 0x0e, 0x51, 0x26, 0x22, 0x1d, 0xd8, 0xbe, 0x92, 0x23, 0xc1,
	0x04, 0x26, 0x8f, 0xa0, 0x6a, 0x44, 0x62, 0x11, 0xf7, 0x72, 0xd7, 0xdd, 0xcb, 0xbd, 0x7e, 0xb6,
	0x6c, 0xec, 0x98, 0xee, 0x8f, 0xf6, 0xfc, 0x57, 0x13, 0xb5, 0xe0, 0x2c, 0x42, 0x52, 0xd8, 0xca,
	0x6f, 0x75, 0x2e, 0x4d, 0x6a, 0xce, 0xd2, 0xd3, 0xc2, 0x69, 0x92, 0xbb, 0x3b, 0x93, 0x2a, 0x44,
	0xcf, 0x6c, 0x90, 0x36, 0xac, 0x8a, 0x11, 0xaa, 0xc4, 0xce, 0x3e, 0x91, 0x3b, 0xef, 0xee, 0xfc,
	0xf3, 0x9a, 0xfa, 0x65, 0xcf, 0x0e, 0xc5, 0x2f, 0x33, 0x22, 0x98, 0xa0, 0xdb, 0x5f, 0xc2, 0xe6,
	0x4c, 0xb5, 0x15, 0xf9, 0x0c, 0x6e, 0x1f, 0xc0, 0xd6, 0x3b, 0xb2, 0xbb, 0x90, 0x89, 0x36, 0x5c,
	0x9d, 0xe3, 0x79, 0xa1, 0x0f, 0xf2, 0x4f, 0x25, 0xd8, 0x78, 0xdb, 0x53, 0x72, 0x03, 0xaa, 0x54,
	0xb2, 0xee, 0x08, 0x95, 0x66, 0x82, 0xe7, 0x86, 0x80, 0x4a, 0x76, 0x9c, 0x49, 0x6c, 0x0f, 0x3b,
	0x65, 0x3c, 0xca, 0xcd, 0xb9, 0xf5, 0x3b, 0xfb, 0xda, 0x73, 0x58, 0x91, 0xd4, 0x84, 0x7d, 0xd4,
	0x79, 0x67, 0xbb, 0x7f, 0xe1, 0x70, 0xfb, 0x76, 0x26, 0x70, 0xdd, 0x35, 0x18, 0xdb, 0xd8, 0xfe,
	0x0e, 0x2a, 0x13, 0xa9, 0x3d, 0xcf, 0xcd, 0x11, 0xd9, 0xed, 0xdc, 0xda, 0xfe, 0x4d, 0x4d, 0xfd,
	0xbc, 0xf8, 0x78, 0x92, 0x41, 0x75, 0x0f, 0xfe, 0x63, 0xe5, 0xcf, 0xa9, 0x3c, 0x32, 0x8a, 0xf1,
	0x78, 0xa2, 0x50, 0xaf, 0xc1, 0xda, 0x39, 0xa2, 0x7e, 0x19, 0xc8, 0xec, 0x04, 0xd0, 0xda, 0x79,
	0x7d, 0x3d, 0x3b, 0x30, 0xff, 0x61, 0x9d, 0xf9, 0x2f, 0xed, 0x2d, 0xbb, 0x59, 0xe6, 0xfe, 0x1f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xbe, 0xf8, 0x50, 0x5e, 0x0f, 0x00, 0x00,
}
