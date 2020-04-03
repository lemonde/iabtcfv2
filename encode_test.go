package iabtcf_test

import (
	iabtcf "github.com/SirDataFR/iab-tcf-v2"
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	str := "COxSKBCOxSKCCBcABCENAgCMAPzAAEPAAAqIDaQBQAMgAgABqAR0A2gDaQAwAMgAgANoAAA.IDaQBQAMgAgABqAR0A2g.QF3QAgABAA1A.eEAAAAAAAUA"
	data := &iabtcf.TCData{
		CoreString: &iabtcf.CoreString{
			Version:              2,
			Created:              timeFromDeciSeconds(15859228738),
			LastUpdated:          timeFromDeciSeconds(15859228802),
			CmpId:                92,
			CmpVersion:           1,
			ConsentScreen:        2,
			ConsentLanguage:      "EN",
			VendorListVersion:    32,
			TcfPolicyVersion:     2,
			IsServiceSpecific:    false,
			UseNonStandardStacks: false,
			SpecialFeatureOptIns: map[int]bool{
				1: true,
				2: true,
			},
			PurposesConsent: map[int]bool{
				1:  true,
				2:  true,
				3:  true,
				4:  true,
				5:  true,
				6:  true,
				9:  true,
				10: true,
			},
			PurposesLITransparency: map[int]bool{
				2:  true,
				7:  true,
				8:  true,
				9:  true,
				10: true,
			},
			PurposeOneTreatment: false,
			PublisherCC:         "FR",
			MaxVendorId:         436,
			IsRangeEncoding:     true,
			VendorsConsent:      map[int]bool{},
			NumEntries:          5,
			RangeEntries: []*iabtcf.RangeEntry{
				{
					StartVendorID: 25,
					EndVendorID:   25,
				},
				{
					StartVendorID: 32,
					EndVendorID:   32,
				},
				{
					StartVendorID: 53,
					EndVendorID:   53,
				},
				{
					StartVendorID: 285,
					EndVendorID:   285,
				},
				{
					StartVendorID: 436,
					EndVendorID:   436,
				},
			},
			MaxVendorIdLI:         436,
			IsRangeEncodingLI:     true,
			VendorsLITransparency: map[int]bool{},
			NumEntriesLI:          3,
			RangeEntriesLI: []*iabtcf.RangeEntry{
				{
					StartVendorID: 25,
					EndVendorID:   25,
				},
				{
					StartVendorID: 32,
					EndVendorID:   32,
				},
				{
					StartVendorID: 436,
					EndVendorID:   436,
				},
			},
			NumPubRestrictions: 0,
		},
		DisclosedVendors: &iabtcf.DisclosedVendors{
			SegmentType:     1,
			MaxVendorId:     436,
			IsRangeEncoding: true,
			NumEntries:      5,
			RangeEntries: []*iabtcf.RangeEntry{
				{
					StartVendorID: 25,
					EndVendorID:   25,
				},
				{
					StartVendorID: 32,
					EndVendorID:   32,
				},
				{
					StartVendorID: 53,
					EndVendorID:   53,
				},
				{
					StartVendorID: 285,
					EndVendorID:   285,
				},
				{
					StartVendorID: 436,
					EndVendorID:   436,
				},
			},
		},
		AllowedVendors: &iabtcf.AllowedVendors{
			SegmentType:     2,
			MaxVendorId:     750,
			IsRangeEncoding: true,
			NumEntries:      2,
			RangeEntries: []*iabtcf.RangeEntry{
				{
					StartVendorID: 2,
					EndVendorID:   2,
				},
				{
					StartVendorID: 53,
					EndVendorID:   53,
				},
			},
		},
		PublisherTC: &iabtcf.PublisherTC{
			SegmentType: 3,
			PubPurposesConsent: map[int]bool{
				1: true,
				2: true,
				7: true,
			},
			PubPurposesLITransparency: map[int]bool{},
			NumCustomPurposes:         2,
			CustomPurposesConsent: map[int]bool{
				1: true,
			},
			CustomPurposesLITransparency: map[int]bool{},
		},
	}

	result := data.ToTCString()
	if result != str {
		t.Errorf("Encode() should produce the same string: in = %s, out = %s", str, result)
	}

}
func TestEncodeCoreString(t *testing.T) {
	str := "COxSKBCOxSKCCBcABCENAgCMAPzAAEPAAAqIDaQBQAMgAgABqAR0A2gDaQAwAMgAgANoAAA"
	segment := &iabtcf.CoreString{
		Version:              2,
		Created:              timeFromDeciSeconds(15859228738),
		LastUpdated:          timeFromDeciSeconds(15859228802),
		CmpId:                92,
		CmpVersion:           1,
		ConsentScreen:        2,
		ConsentLanguage:      "EN",
		VendorListVersion:    32,
		TcfPolicyVersion:     2,
		IsServiceSpecific:    false,
		UseNonStandardStacks: false,
		SpecialFeatureOptIns: map[int]bool{
			1: true,
			2: true,
		},
		PurposesConsent: map[int]bool{
			1:  true,
			2:  true,
			3:  true,
			4:  true,
			5:  true,
			6:  true,
			9:  true,
			10: true,
		},
		PurposesLITransparency: map[int]bool{
			2:  true,
			7:  true,
			8:  true,
			9:  true,
			10: true,
		},
		PurposeOneTreatment: false,
		PublisherCC:         "FR",
		MaxVendorId:         436,
		IsRangeEncoding:     true,
		VendorsConsent:      map[int]bool{},
		NumEntries:          5,
		RangeEntries: []*iabtcf.RangeEntry{
			{
				StartVendorID: 25,
				EndVendorID:   25,
			},
			{
				StartVendorID: 32,
				EndVendorID:   32,
			},
			{
				StartVendorID: 53,
				EndVendorID:   53,
			},
			{
				StartVendorID: 285,
				EndVendorID:   285,
			},
			{
				StartVendorID: 436,
				EndVendorID:   436,
			},
		},
		MaxVendorIdLI:         436,
		IsRangeEncodingLI:     true,
		VendorsLITransparency: map[int]bool{},
		NumEntriesLI:          3,
		RangeEntriesLI: []*iabtcf.RangeEntry{
			{
				StartVendorID: 25,
				EndVendorID:   25,
			},
			{
				StartVendorID: 32,
				EndVendorID:   32,
			},
			{
				StartVendorID: 436,
				EndVendorID:   436,
			},
		},
		NumPubRestrictions: 0,
	}

	result := segment.Encode()
	if result != str {
		t.Errorf("Encode() should produce the same string: in = %s, out = %s", str, result)
	}
}

func TestEncodeDisclosedVendors(t *testing.T) {
	str := "IDaQBQAMgAgABqAR0A2g"
	segment := &iabtcf.DisclosedVendors{
		SegmentType:     1,
		MaxVendorId:     436,
		IsRangeEncoding: true,
		NumEntries:      5,
		RangeEntries: []*iabtcf.RangeEntry{
			{
				StartVendorID: 25,
				EndVendorID:   25,
			},
			{
				StartVendorID: 32,
				EndVendorID:   32,
			},
			{
				StartVendorID: 53,
				EndVendorID:   53,
			},
			{
				StartVendorID: 285,
				EndVendorID:   285,
			},
			{
				StartVendorID: 436,
				EndVendorID:   436,
			},
		},
	}

	result := segment.Encode()
	if result != str {
		t.Errorf("Encode() should produce the same string: in = %s, out = %s", str, result)
	}
}

func TestEncodeAllowedVendors(t *testing.T) {
	str := "QF3QAgABAA1A"
	segment := &iabtcf.AllowedVendors{
		SegmentType:     2,
		MaxVendorId:     750,
		IsRangeEncoding: true,
		NumEntries:      2,
		RangeEntries: []*iabtcf.RangeEntry{
			{
				StartVendorID: 2,
				EndVendorID:   2,
			},
			{
				StartVendorID: 53,
				EndVendorID:   53,
			},
		},
	}

	result := segment.Encode()
	if result != str {
		t.Errorf("Encode() should produce the same string: in = %s, out = %s", str, result)
	}
}

func TestEncodePublisherTC(t *testing.T) {
	str := "eEAAAAAAAUA"
	segment := &iabtcf.PublisherTC{
		SegmentType: 3,
		PubPurposesConsent: map[int]bool{
			1: true,
			2: true,
			7: true,
		},
		PubPurposesLITransparency: map[int]bool{},
		NumCustomPurposes:         2,
		CustomPurposesConsent: map[int]bool{
			1: true,
		},
		CustomPurposesLITransparency: map[int]bool{},
	}

	result := segment.Encode()
	if result != str {
		t.Errorf("Encode() should produce the same string: in = %s, out = %s", str, result)
	}
}

func timeFromDeciSeconds(deciseconds int64) time.Time {
	return time.Unix(deciseconds/10, (deciseconds%10)*int64(time.Millisecond*100)).UTC()
}
