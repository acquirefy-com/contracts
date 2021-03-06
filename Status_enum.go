// Code generated by go-enum
// DO NOT EDIT!

package contracts

import (
	"fmt"
)

const (
	// StatusPending is a Status of type Pending
	StatusPending Status = iota
	// StatusSuccessful is a Status of type Successful
	StatusSuccessful
	// StatusDecline is a Status of type Decline
	StatusDecline
	// StatusFailed is a Status of type Failed
	StatusFailed
	// StatusExpired is a Status of type Expired
	StatusExpired
)

const _StatusName = "pendingsuccessfuldeclinefailedexpired"

var _StatusMap = map[Status]string{
	0: _StatusName[0:7],
	1: _StatusName[7:17],
	2: _StatusName[17:24],
	3: _StatusName[24:30],
	4: _StatusName[30:37],
}

// String implements the Stringer interface.
func (x Status) String() string {
	if str, ok := _StatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Status(%d)", x)
}

var _StatusValue = map[string]Status{
	_StatusName[0:7]:   0,
	_StatusName[7:17]:  1,
	_StatusName[17:24]: 2,
	_StatusName[24:30]: 3,
	_StatusName[30:37]: 4,
}

// ParseStatus attempts to convert a string to a Status
func ParseStatus(name string) (Status, error) {
	if x, ok := _StatusValue[name]; ok {
		return x, nil
	}
	return Status(0), fmt.Errorf("%s is not a valid Status", name)
}
