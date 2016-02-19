// Code generated by protoc-gen-go.
// source: google.golang.org/cloud/bigtable/internal/cluster_service_proto/bigtable_cluster_service_messages.proto
// DO NOT EDIT!

/*
Package google_bigtable_admin_cluster_v1 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/cloud/bigtable/internal/cluster_service_proto/bigtable_cluster_service_messages.proto
	google.golang.org/cloud/bigtable/internal/cluster_service_proto/bigtable_cluster_service.proto

It has these top-level messages:
	ListZonesRequest
	ListZonesResponse
	GetClusterRequest
	ListClustersRequest
	ListClustersResponse
	CreateClusterRequest
	CreateClusterMetadata
	UpdateClusterMetadata
	DeleteClusterRequest
	UndeleteClusterRequest
	UndeleteClusterMetadata
*/
package google_bigtable_admin_cluster_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_bigtable_admin_cluster_v11 "google.golang.org/cloud/bigtable/internal/cluster_data_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Request message for BigtableClusterService.ListZones.
type ListZonesRequest struct {
	// The unique name of the project for which a list of supported zones is
	// requested.
	// Values are of the form projects/<project>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ListZonesRequest) Reset()         { *m = ListZonesRequest{} }
func (m *ListZonesRequest) String() string { return proto.CompactTextString(m) }
func (*ListZonesRequest) ProtoMessage()    {}

// Response message for BigtableClusterService.ListZones.
type ListZonesResponse struct {
	// The list of requested zones.
	Zones []*google_bigtable_admin_cluster_v11.Zone `protobuf:"bytes,1,rep,name=zones" json:"zones,omitempty"`
}

func (m *ListZonesResponse) Reset()         { *m = ListZonesResponse{} }
func (m *ListZonesResponse) String() string { return proto.CompactTextString(m) }
func (*ListZonesResponse) ProtoMessage()    {}

func (m *ListZonesResponse) GetZones() []*google_bigtable_admin_cluster_v11.Zone {
	if m != nil {
		return m.Zones
	}
	return nil
}

// Request message for BigtableClusterService.GetCluster.
type GetClusterRequest struct {
	// The unique name of the requested cluster.
	// Values are of the form projects/<project>/zones/<zone>/clusters/<cluster>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetClusterRequest) Reset()         { *m = GetClusterRequest{} }
func (m *GetClusterRequest) String() string { return proto.CompactTextString(m) }
func (*GetClusterRequest) ProtoMessage()    {}

// Request message for BigtableClusterService.ListClusters.
type ListClustersRequest struct {
	// The unique name of the project for which a list of clusters is requested.
	// Values are of the form projects/<project>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ListClustersRequest) Reset()         { *m = ListClustersRequest{} }
func (m *ListClustersRequest) String() string { return proto.CompactTextString(m) }
func (*ListClustersRequest) ProtoMessage()    {}

// Response message for BigtableClusterService.ListClusters.
type ListClustersResponse struct {
	// The list of requested Clusters.
	Clusters []*google_bigtable_admin_cluster_v11.Cluster `protobuf:"bytes,1,rep,name=clusters" json:"clusters,omitempty"`
	// The zones for which clusters could not be retrieved.
	FailedZones []*google_bigtable_admin_cluster_v11.Zone `protobuf:"bytes,2,rep,name=failed_zones" json:"failed_zones,omitempty"`
}

func (m *ListClustersResponse) Reset()         { *m = ListClustersResponse{} }
func (m *ListClustersResponse) String() string { return proto.CompactTextString(m) }
func (*ListClustersResponse) ProtoMessage()    {}

func (m *ListClustersResponse) GetClusters() []*google_bigtable_admin_cluster_v11.Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *ListClustersResponse) GetFailedZones() []*google_bigtable_admin_cluster_v11.Zone {
	if m != nil {
		return m.FailedZones
	}
	return nil
}

// Request message for BigtableClusterService.CreateCluster.
type CreateClusterRequest struct {
	// The unique name of the zone in which to create the cluster.
	// Values are of the form projects/<project>/zones/<zone>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The id to be used when referring to the new cluster within its zone,
	// e.g. just the "test-cluster" section of the full name
	// "projects/<project>/zones/<zone>/clusters/test-cluster".
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id" json:"cluster_id,omitempty"`
	// The cluster to create.
	// The "name", "delete_time", and "current_operation" fields must be left
	// blank.
	Cluster *google_bigtable_admin_cluster_v11.Cluster `protobuf:"bytes,3,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *CreateClusterRequest) Reset()         { *m = CreateClusterRequest{} }
func (m *CreateClusterRequest) String() string { return proto.CompactTextString(m) }
func (*CreateClusterRequest) ProtoMessage()    {}

func (m *CreateClusterRequest) GetCluster() *google_bigtable_admin_cluster_v11.Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

// Metadata type for the operation returned by
// BigtableClusterService.CreateCluster.
type CreateClusterMetadata struct {
	// The request which prompted the creation of this operation.
	OriginalRequest *CreateClusterRequest `protobuf:"bytes,1,opt,name=original_request" json:"original_request,omitempty"`
}

func (m *CreateClusterMetadata) Reset()         { *m = CreateClusterMetadata{} }
func (m *CreateClusterMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateClusterMetadata) ProtoMessage()    {}

func (m *CreateClusterMetadata) GetOriginalRequest() *CreateClusterRequest {
	if m != nil {
		return m.OriginalRequest
	}
	return nil
}

// Metadata type for the operation returned by
// BigtableClusterService.UpdateCluster.
type UpdateClusterMetadata struct {
	// The request which prompted the creation of this operation.
	OriginalRequest *google_bigtable_admin_cluster_v11.Cluster `protobuf:"bytes,1,opt,name=original_request" json:"original_request,omitempty"`
}

func (m *UpdateClusterMetadata) Reset()         { *m = UpdateClusterMetadata{} }
func (m *UpdateClusterMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateClusterMetadata) ProtoMessage()    {}

func (m *UpdateClusterMetadata) GetOriginalRequest() *google_bigtable_admin_cluster_v11.Cluster {
	if m != nil {
		return m.OriginalRequest
	}
	return nil
}

// Request message for BigtableClusterService.DeleteCluster.
type DeleteClusterRequest struct {
	// The unique name of the cluster to be deleted.
	// Values are of the form projects/<project>/zones/<zone>/clusters/<cluster>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteClusterRequest) Reset()         { *m = DeleteClusterRequest{} }
func (m *DeleteClusterRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteClusterRequest) ProtoMessage()    {}

// Request message for BigtableClusterService.UndeleteCluster.
type UndeleteClusterRequest struct {
	// The unique name of the cluster to be un-deleted.
	// Values are of the form projects/<project>/zones/<zone>/clusters/<cluster>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *UndeleteClusterRequest) Reset()         { *m = UndeleteClusterRequest{} }
func (m *UndeleteClusterRequest) String() string { return proto.CompactTextString(m) }
func (*UndeleteClusterRequest) ProtoMessage()    {}

// Metadata type for the operation returned by
// BigtableClusterService.UndeleteCluster.
type UndeleteClusterMetadata struct {
}

func (m *UndeleteClusterMetadata) Reset()         { *m = UndeleteClusterMetadata{} }
func (m *UndeleteClusterMetadata) String() string { return proto.CompactTextString(m) }
func (*UndeleteClusterMetadata) ProtoMessage()    {}
