// Code generated by "stringer -type SeverityLvl log.go"; DO NOT EDIT.

package l5424

import "strconv"

const _SeverityLvl_name = "EmergencyLvlAlertLvlCritLvlErrorLvlWarnLvlNoticeLvlInfoLvlDebugLvl"

var _SeverityLvl_index = [...]uint8{0, 12, 20, 27, 35, 42, 51, 58, 66}

func (i SeverityLvl) String() string {
	if i >= SeverityLvl(len(_SeverityLvl_index)-1) {
		return "SeverityLvl(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SeverityLvl_name[_SeverityLvl_index[i]:_SeverityLvl_index[i+1]]
}