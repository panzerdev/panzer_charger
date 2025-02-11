package panzer_charger

import (
	"context"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/charger"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/modbus"
)

// MyCustomCharger charger implementation
type MyCustomCharger struct {
	log     *util.Logger
	conn    *modbus.Connection
	current uint16
	wakeup  bool
}

func init() {
	charger.RegisterChargerCtx("my-custom-charger", NewMyCustomChargerFromConfig)
}

// NewMyCustomChargerFromConfig creates a MyCustomCharger charger from generic config
func NewMyCustomChargerFromConfig(ctx context.Context, other map[string]interface{}) (api.Charger, error) {
	cc := modbus.Settings{
		ID: 1,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	return NewMyCustomCharger(ctx, cc.URI, cc.Device, cc.Comset, cc.Baudrate, cc.Protocol(), cc.ID)
}

// NewMyCustomCharger creates MyCustomCharger charger
func NewMyCustomCharger(ctx context.Context, uri, device, comset string, baudrate int, proto modbus.Protocol, slaveID uint8) (api.Charger, error) {
	conn, err := modbus.NewConnection(uri, device, comset, baudrate, proto, slaveID)
	if err != nil {
		return nil, err
	}

	log := util.NewLogger("MyCustomCharger")
	conn.Logger(log.TRACE)

	wb := &MyCustomCharger{
		log:     log,
		conn:    conn,
		current: 60, // assume min current
	}
//....
	return wb, nil
}


// Status implements the api.Charger interface
func (wb *MyCustomCharger) Status() (api.ChargeStatus, error) {
	//....
	return api.StatusA, nil
}

// Enabled implements the api.Charger interface
func (wb *MyCustomCharger) Enabled() (bool, error) {
	//....
	return true, nil
}

// Enable implements the api.Charger interface
func (wb *MyCustomCharger) Enable(enable bool) error {
	//....
	return nil
}

// MaxCurrent implements the api.Charger interface
func (wb *MyCustomCharger) MaxCurrent(current int64) error {
	//....
	return wb.MaxCurrentMillis(float64(current))
}

var _ api.ChargerEx = (*MyCustomCharger)(nil)

// MaxCurrentMillis implements the api.ChargerEx interface
func (wb *MyCustomCharger) MaxCurrentMillis(current float64) error {
	//....
	return nil
}

var _ api.Meter = (*MyCustomCharger)(nil)

// CurrentPower implements the api.Meter interface
func (wb *MyCustomCharger) CurrentPower() (float64, error) {
	//....
	return 0, nil
}

var _ api.MeterEnergy = (*MyCustomCharger)(nil)

// TotalEnergy implements the api.MeterEnergy interface
func (wb *MyCustomCharger) TotalEnergy() (float64, error) {
	//....
	return 0, nil
}

var _ api.PhaseCurrents = (*MyCustomCharger)(nil)

// Currents implements the api.PhaseCurrents interface
func (wb *MyCustomCharger) Currents() (float64, float64, float64, error) {
	//....
	return 0,0,0,nil
}

var _ api.PhaseVoltages = (*MyCustomCharger)(nil)

// Voltages implements the api.PhaseVoltages interface
func (wb *MyCustomCharger) Voltages() (float64, float64, float64, error) {
	//....
	return 0,0,0,nil
}

var _ api.Diagnosis = (*MyCustomCharger)(nil)

// Diagnose implements the api.Diagnosis interface
func (wb *MyCustomCharger) Diagnose() {
	//....
}

var _ api.Resurrector = (*MyCustomCharger)(nil)

// WakeUp implements the api.Resurrector interface
func (wb *MyCustomCharger) WakeUp() error {
	//....
	return nil
}
