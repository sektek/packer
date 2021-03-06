// Code generated by protoc-gen-goext. DO NOT EDIT.

package metering

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func (m *UsageRecord) SetUuid(v string) {
	m.Uuid = v
}

func (m *UsageRecord) SetSkuId(v string) {
	m.SkuId = v
}

func (m *UsageRecord) SetQuantity(v int64) {
	m.Quantity = v
}

func (m *UsageRecord) SetTimestamp(v *timestamp.Timestamp) {
	m.Timestamp = v
}

func (m *AcceptedUsageRecord) SetUuid(v string) {
	m.Uuid = v
}

func (m *RejectedUsageRecord) SetUuid(v string) {
	m.Uuid = v
}

func (m *RejectedUsageRecord) SetReason(v RejectedUsageRecord_Reason) {
	m.Reason = v
}
