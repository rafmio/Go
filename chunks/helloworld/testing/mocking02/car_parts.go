package mocking02

type CrankingCircuit interface {
	RotateShaft() error
}

type CalibrationParams struct {
	FuelLever  int
	Latitude   int
	Longtitude int
}

type Calibrator interface {
	Calibrate(params CalibrationParams) (calibrated bool, err error)
}

type Accelerator interface {
	Accelerate() error
}
