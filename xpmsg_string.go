// Code generated by "stringer -type=XPMsg"; DO NOT EDIT

package xpmetrics

import "fmt"

const _XPMsg_name = "FrameRateMsgTimesMsgSimStatsMsgSpeedsMsgMachVVIGloadMsgWeatherMsgAircraftAtmMsgSysPressuresMsgJoystickMsgOtherFlightControlsMsgArtificialStabilityMsgFlightControlsMsgWingSweepThrustVectoringMsgTrimFlatpStatsSpeedBrakesMsgGearBrakesMsgAngularMomentsMsgAngularVelocitiesMsgPitchRollMsgAngleAttackSideSlipPathsMsgMagneticCompassMsgLatLonAltMsgLocVelocityDistanceMsgPlanesLatMsgPlanesLonMsgPlanesAltMsgThrottleCmdMsgThrottleActualMsgEngineFeatherMsgPropellerSettingMsgMixtureSettingMsgCarburatorHeatSettingMsgCowlFlatSettingMsgMagnetoSettingMsgStarterTimeOutMsgEnginePowerMsgEngineThrustMsgEngineTorqueMsgEngineRPMMsgPropellerRPMMsgPropellerPitchMsgPropJetWashMsgN1MsgN2MsgMPRMsgEPRMsgFFMsgITTMsgEGTMsgCHTMsgOilPressureMsgOilTemperatureMsgFuelPressureMsgGeneratorAmpMsgBatteryAmpMsgBatteryVoltMsgElectricPumpOnOffMsgIdleSpeedLowHighMsgBatteryOnOffMsgGenOnOffMsgInverterOnOffMsgFADECOnOffMsgIgniterOnOffMsgFuelWeightsMsgPayloadWeightsCGMsgAeroForcesMsgEngineForcesMsgGearVForceMsgGearDeploymentMsgLiftOverDrageCoeffMsgPropellerEfficiencyMsgAileronDef1MsgAileronDef2MsgRollSpoilerDef1MsgRollSpoilerDef2MsgElevatorDefMsgRudderDefMsgYawBrakeDefMsgControlForcesMsgTotalVerticalThrustVectorMsgTotalLateralThrustVectorMsgPitchCyclicDiscTitlsMsgRollCyclicDiscTitlsMsgPitchCyclicDiscFlappingMsgRollCyclicFlappingMsgGroundEffectLiftWingsMsgGroundEffectDragWingsMsgGroundEffectWashWingsMsgGroundEffectLiftStabilizersMsgGroundEffectDragStabilizersMsgGroundEffectWashStabilizersMsgGroundEffectLiftPropellersMsgGroundEffectDragPropellersMsgWingLiftMsgWingDragMsgStabilizerLiftMsgStabilizerDragMsgCOM12FreqMsgNAV12FreqMsgNAV12OBSMsgNAV1DefMsgNAV2DefMsgADF12StatusMsgDMEStatusMsgGPSStatusMsgTransponderStatusMsgMarkerStatusMsgElectricalSwitchesMsgEFISSwitchesMsgAutoPilotFDHUDSwtichesMsgAntiIceSwtichesMsgAntiIceFuelSwtichesMsgClutchArtStabilitySwitchesMsgMiscellaneousSwitchesMsgGeneralAnn1MsgGeneralAnn2MsgEngineAnnMsgAutoPilotArmedMsgAutoPilotModesMsgAutoPilotValuesMsgWeaponStatusMsgPressStatusMsgAPUGPUStatusMsgRadarStatusHydraulicStatusMsgElectricalSolarSystemsMsgIcingStatus1MsgIcingStatus2MsgWarningStatusMsgFlightPlanLegsMsgHardwareOptionsMsgCameraLocationMsgGroundLocationMsgClimbStatsMsgCruiseStatsMsgGearSteeringMsgMotionPlatformStatsMsg"

var _XPMsg_index = [...]uint16{0, 12, 20, 31, 40, 55, 65, 79, 94, 105, 127, 149, 166, 193, 221, 234, 251, 271, 283, 310, 328, 340, 362, 374, 386, 398, 412, 429, 445, 464, 481, 505, 523, 540, 557, 571, 586, 601, 613, 628, 645, 659, 664, 669, 675, 681, 686, 692, 698, 704, 718, 735, 750, 765, 778, 792, 812, 831, 846, 857, 873, 886, 901, 915, 934, 947, 962, 975, 992, 1013, 1035, 1049, 1063, 1081, 1099, 1113, 1125, 1139, 1155, 1183, 1210, 1233, 1255, 1281, 1302, 1326, 1350, 1374, 1404, 1434, 1464, 1493, 1522, 1533, 1544, 1561, 1578, 1590, 1602, 1613, 1623, 1633, 1647, 1659, 1671, 1691, 1706, 1727, 1742, 1767, 1785, 1807, 1836, 1860, 1874, 1888, 1900, 1917, 1934, 1952, 1967, 1981, 1996, 2007, 2025, 2050, 2065, 2080, 2096, 2113, 2131, 2148, 2165, 2178, 2192, 2207, 2229}

func (i XPMsg) String() string {
	if i >= XPMsg(len(_XPMsg_index)-1) {
		return fmt.Sprintf("XPMsg(%d)", i)
	}
	return _XPMsg_name[_XPMsg_index[i]:_XPMsg_index[i+1]]
}
