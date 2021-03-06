// Code generated by "enumer -type FacilityLvl log.go"; DO NOT EDIT

package l5424

import (
	"fmt"
)

const _FacilityLvlName = "FacilityKernelFacilityUserFacilityMailFacilitySystemDFacilitySecurityFacilitySyslogDFacilityPrinterFacilityNewsFacilityUUCPFacilityClockFacilityAuthorizationFacilityFTPFacilityNTPFacilityAuditFacilityAlertFacilityClockDFacilityLocal0FacilityLocal1FacilityLocal2FacilityLocal3FacilityLocal4FacilityLocal5FacilityLocal6FacilityLocal7"

var _FacilityLvlIndex = [...]uint16{0, 14, 26, 38, 53, 69, 84, 99, 111, 123, 136, 157, 168, 179, 192, 205, 219, 233, 247, 261, 275, 289, 303, 317, 331}

func (i FacilityLvl) String() string {
	if i >= FacilityLvl(len(_FacilityLvlIndex)-1) {
		return fmt.Sprintf("FacilityLvl(%d)", i)
	}
	return _FacilityLvlName[_FacilityLvlIndex[i]:_FacilityLvlIndex[i+1]]
}

var _FacilityLvlValues = []FacilityLvl{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}

var _FacilityLvlNameToValueMap = map[string]FacilityLvl{
	_FacilityLvlName[0:14]:    0,
	_FacilityLvlName[14:26]:   1,
	_FacilityLvlName[26:38]:   2,
	_FacilityLvlName[38:53]:   3,
	_FacilityLvlName[53:69]:   4,
	_FacilityLvlName[69:84]:   5,
	_FacilityLvlName[84:99]:   6,
	_FacilityLvlName[99:111]:  7,
	_FacilityLvlName[111:123]: 8,
	_FacilityLvlName[123:136]: 9,
	_FacilityLvlName[136:157]: 10,
	_FacilityLvlName[157:168]: 11,
	_FacilityLvlName[168:179]: 12,
	_FacilityLvlName[179:192]: 13,
	_FacilityLvlName[192:205]: 14,
	_FacilityLvlName[205:219]: 15,
	_FacilityLvlName[219:233]: 16,
	_FacilityLvlName[233:247]: 17,
	_FacilityLvlName[247:261]: 18,
	_FacilityLvlName[261:275]: 19,
	_FacilityLvlName[275:289]: 20,
	_FacilityLvlName[289:303]: 21,
	_FacilityLvlName[303:317]: 22,
	_FacilityLvlName[317:331]: 23,
}

// FacilityLvlString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FacilityLvlString(s string) (FacilityLvl, error) {
	if val, ok := _FacilityLvlNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to FacilityLvl values", s)
}

// FacilityLvlValues returns all values of the enum
func FacilityLvlValues() []FacilityLvl {
	return _FacilityLvlValues
}

// IsAFacilityLvl returns "true" if the value is listed in the enum definition. "false" otherwise
func (i FacilityLvl) IsAFacilityLvl() bool {
	for _, v := range _FacilityLvlValues {
		if i == v {
			return true
		}
	}
	return false
}
