package observer

import "testing"

//具体的subject: concrete subject
func TestWeatherData_Observer(t *testing.T) {
	weatherData := NewWealthData()
	wd := &CurrentConditionsDisplay{weathData: weatherData}
	weatherData.AttachObserver(wd)
	weatherData.SetMeasurement(80, 65, 30.4)
}
