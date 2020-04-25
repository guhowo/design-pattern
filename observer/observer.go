package observer

import "fmt"

//文件分为两部分，上半部分定义了一些抽象接口，包括observer和subject
type IObserver interface {
	Update(temperature float32, humidity float32, pressure float32)
}

type ISubject interface {
	AttachObserver(o IObserver)
	DetachObserver(o IObserver)
	NotifyAll()
}

type WeatherData struct {
	Temperature float32
	Humidity    float32
	Pressure    float32
	Observers   map[IObserver]bool
}

func NewWealthData() *WeatherData {
	return &WeatherData{
		Observers: make(map[IObserver]bool),
	}
}

func (w *WeatherData) AttachObserver(o IObserver) {
	w.Observers[o] = true
}

func (w *WeatherData) DetachObserver(o IObserver) {
	if _, ok := w.Observers[o]; ok {
		delete(w.Observers, o)
	}
}

//notify all Observers
func (w *WeatherData) NotifyAll() {
	for o := range w.Observers {
		o.Update(w.Temperature, w.Humidity, w.Pressure)
	}
}

func (w *WeatherData) SetMeasurement(temperature float32, humidity float32, pressure float32) {
	w.Temperature = temperature
	w.Humidity = humidity
	w.Pressure = pressure
	w.NotifyAll()
}

//下面定义了concrete observer
type CurrentConditionsDisplay struct {
	Temperature float32
	Humidity    float32
	weathData   ISubject
}

func (ccd *CurrentConditionsDisplay) Update(temperature float32, humidity float32, pressure float32) {
	ccd.Temperature = temperature
	ccd.Humidity = humidity
	// pressure 没用到
	ccd.Display()
}

func (ccd *CurrentConditionsDisplay) Display() {
	fmt.Println("Current conditions: " + fmt.Sprintf("%v", ccd.Temperature) + "F degrees and " + fmt.Sprintf("%v", ccd.Humidity) + "% humidity")
}
