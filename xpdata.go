package xpmetrics

//go:generate stringer -type=XPMsg

import "sync"

// XPMsg describes message type sent by XPlane
type XPMsg uint8

const (
	FrameRateMsg XPMsg = iota
	TimesMsg
	SimStatsMsg
	SpeedsMsg
	MachVVIGloadMsg
	WeatherMsg
	AircraftAtmMsg
	SysPressuresMsg
	JoystickMsg
	OtherFlightControlsMsg
	ArtificialStabilityMsg
	FlightControlsMsg
	WingSweepThrustVectoringMsg
	TrimFlatpStatsSpeedBrakesMsg
	GearBrakesMsg
	AngularMomentsMsg
	AngularVelocitiesMsg
	PitchRollMsg
	AngleAttackSideSlipPathsMsg
	MagneticCompassMsg
	LatLonAltMsg
	LocVelocityDistanceMsg
	PlanesLatMsg
	PlanesLonMsg
	PlanesAltMsg
	ThrottleCmdMsg
	ThrottleActualMsg
	EngineFeatherMsg
	PropellerSettingMsg
	MixtureSettingMsg
	CarburatorHeatSettingMsg
	CowlFlatSettingMsg
	MagnetoSettingMsg
	StarterTimeOutMsg
	EnginePowerMsg
	EngineThrustMsg
	EngineTorqueMsg
	EngineRPMMsg
	PropellerRPMMsg
	PropellerPitchMsg
	PropJetWashMsg
	N1Msg
	N2Msg
	MPRMsg
	EPRMsg
	FFMsg
	ITTMsg
	EGTMsg
	CHTMsg
	OilPressureMsg
	OilTemperatureMsg
	FuelPressureMsg
	GeneratorAmpMsg
	BatteryAmpMsg
	BatteryVoltMsg
	ElectricPumpOnOffMsg
	IdleSpeedLowHighMsg
	BatteryOnOffMsg
	GenOnOffMsg
	InverterOnOffMsg
	FADECOnOffMsg
	IgniterOnOffMsg
	FuelWeightsMsg
	PayloadWeightsCGMsg
	AeroForcesMsg
	EngineForcesMsg
	GearVForceMsg
	GearDeploymentMsg
	LiftOverDrageCoeffMsg
	PropellerEfficiencyMsg
	AileronDef1Msg
	AileronDef2Msg
	RollSpoilerDef1Msg
	RollSpoilerDef2Msg
	ElevatorDefMsg
	RudderDefMsg
	YawBrakeDefMsg
	ControlForcesMsg
	TotalVerticalThrustVectorMsg
	TotalLateralThrustVectorMsg
	PitchCyclicDiscTitlsMsg
	RollCyclicDiscTitlsMsg
	PitchCyclicDiscFlappingMsg
	RollCyclicFlappingMsg
	GroundEffectLiftWingsMsg
	GroundEffectDragWingsMsg
	GroundEffectWashWingsMsg
	GroundEffectLiftStabilizersMsg
	GroundEffectDragStabilizersMsg
	GroundEffectWashStabilizersMsg
	GroundEffectLiftPropellersMsg
	GroundEffectDragPropellersMsg
	WingLiftMsg
	WingDragMsg
	StabilizerLiftMsg
	StabilizerDragMsg
	COM12FreqMsg
	NAV12FreqMsg
	NAV12OBSMsg
	NAV1DefMsg
	NAV2DefMsg
	ADF12StatusMsg
	DMEStatusMsg
	GPSStatusMsg
	TransponderStatusMsg
	MarkerStatusMsg
	ElectricalSwitchesMsg
	EFISSwitchesMsg
	AutoPilotFDHUDSwtichesMsg
	AntiIceSwtichesMsg
	AntiIceFuelSwtichesMsg
	ClutchArtStabilitySwitchesMsg
	MiscellaneousSwitchesMsg
	GeneralAnn1Msg
	GeneralAnn2Msg
	EngineAnnMsg
	AutoPilotArmedMsg
	AutoPilotModesMsg
	AutoPilotValuesMsg
	WeaponStatusMsg
	PressStatusMsg
	APUGPUStatusMsg
	RadarStatus
	HydraulicStatusMsg
	ElectricalSolarSystemsMsg
	IcingStatus1Msg
	IcingStatus2Msg
	WarningStatusMsg
	FlightPlanLegsMsg
	HardwareOptionsMsg
	CameraLocationMsg
	GroundLocationMsg
	ClimbStatsMsg
	CruiseStatsMsg
	GearSteeringMsg
	MotionPlatformStatsMsg
)

// XPData the storage for all XPlane values
type XPData struct {
	sync.RWMutex
	m map[uint8][8]float32
}

// NewXPData returns a new XPData ready to store values
func NewXPData() *XPData {
	return &XPData{
		m: make(map[uint8][8]float32),
	}
}

// Insert a new message
func (d *XPData) Insert(didx uint8, fvals [8]float32) {
	d.Lock()
	d.m[didx] = fvals
	d.Unlock()
}

// Position helper to get position
func (d *XPData) Position() (ok bool, lat, lng, alt float32) {
	d.RLock()
	defer d.RUnlock()
	v, ok := d.m[uint8(LatLonAltMsg)]
	if !ok {
		return false, 0, 0, 0
	}
	return true, v[0], v[1], v[2]
}

// Compass helper to get magnetic compass
func (d *XPData) Compass() (lat, lng, alt float32, ok bool) {
	d.RLock()
	defer d.RUnlock()
	v, ok := d.m[uint8(LatLonAltMsg)]
	if !ok {
		return 0, 0, 0, false
	}
	return v[0], v[1], v[2], true
}

// Query return values stored for msg type, return false if no date have been received
func (d *XPData) Query(msg XPMsg) ([8]float32, bool) {
	d.RLock()
	defer d.RUnlock()
	v, ok := d.m[uint8(msg)]
	if !ok {
		return [8]float32{}, false
	}
	return v, true
}
