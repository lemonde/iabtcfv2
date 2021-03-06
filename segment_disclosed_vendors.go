package iabtcfv2

import (
	"encoding/base64"
)

type DisclosedVendors struct {
	SegmentType      int
	MaxVendorId      int
	IsRangeEncoding  bool
	DisclosedVendors map[int]bool
	NumEntries       int
	RangeEntries     []*RangeEntry
}

// Returns true if vendor id is disclosed for validating OOB signaling
func (d *DisclosedVendors) IsVendorDisclosed(id int) bool {
	if d.IsRangeEncoding {
		for _, entry := range d.RangeEntries {
			if entry.StartVendorID <= id && id <= entry.EndVendorID {
				return true
			}
		}
		return false
	}

	return d.DisclosedVendors[id]
}

// Returns structure as a base64 raw url encoded string
func (d *DisclosedVendors) Encode() string {
	bitSize := 20

	if d.IsRangeEncoding {
		bitSize += 12
		entriesSize := len(d.RangeEntries)
		for _, entry := range d.RangeEntries {
			if entry.EndVendorID > entry.StartVendorID {
				entriesSize += 16 * 2
			} else {
				entriesSize += 16
			}
		}
		bitSize += entriesSize
	} else {
		if d.MaxVendorId == 0 {
			for id, _ := range d.DisclosedVendors {
				if id > d.MaxVendorId {
					d.MaxVendorId = id
				}
			}
		}
		bitSize += d.MaxVendorId
	}

	var e = newTCEncoder(make([]byte, bitSize/8))
	if bitSize%8 != 0 {
		e = newTCEncoder(make([]byte, bitSize/8+1))
	}

	e.writeInt(d.SegmentType, 3)
	e.writeInt(d.MaxVendorId, 16)
	e.writeBool(d.IsRangeEncoding)
	if d.IsRangeEncoding {
		e.writeInt(len(d.RangeEntries), 12)
		e.writeRangeEntries(d.RangeEntries)
	} else {
		for i := 0; i < d.MaxVendorId; i++ {
			e.writeBool(d.IsVendorDisclosed(i + 1))
		}
	}

	return base64.RawURLEncoding.EncodeToString(e.bytes)
}
