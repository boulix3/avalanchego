// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/json"
)

type Info struct {
	IP             string      `json:"ip"`
	PublicIP       string      `json:"publicIP,omitempty"`
	ID             ids.NodeID  `json:"nodeID"`
	Version        string      `json:"version"`
	LastSent       time.Time   `json:"lastSent"`
	LastReceived   time.Time   `json:"lastReceived"`
	ObservedUptime json.Uint32 `json:"observedUptime"`
	TrackedSubnets []ids.ID    `json:"trackedSubnets"`
}
