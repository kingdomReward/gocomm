package core

import (
	"fmt"
	"reflect"
	"wiless/gocomm"
	"wiless/gocomm/chipset"
	"wiless/vlib"
)

func init() {
	fmt.Printf("\n AutoGenerated : Filter")
}

type Filter struct {
	name          string
	isInitialized bool
	Pins          map[string]chipset.PinInfo
	Modules       map[string]chipset.ModuleInfo
	ModuleNames   map[int]string
	PinNames      map[int]string
	coeff         vlib.VectorC
	FilterMemory  vlib.VectorC
}

/// AutoGenerated through script

func (m Filter) InPinCount() int {
	return 1
}

func (m Filter) OutPinCount() int {
	return 1
}

func (m Filter) Pin(pid int) chipset.PinInfo {
	return m.Pins[m.PinNames[pid]]
}

func (m Filter) PinIn(pid int) chipset.PinInfo {
	if pid >= m.InPinCount() {
		fmt.Printf("%d > No of Input Pins %d", pid, m.InPinCount())
		var result chipset.PinInfo
		result.Id = -1
		return result
	}

	return m.Pins[m.PinNames[pid]]

}
func (m Filter) PinByID(pid int) chipset.PinInfo {

	return m.Pins[m.PinNames[pid]]
}

func (m Filter) PinOut(pid int) chipset.PinInfo {
	if pid >= m.OutPinCount() {
		fmt.Printf("%d > No of Output Pins %d", pid, m.OutPinCount())
		var result chipset.PinInfo
		result.Id = -1
		return result
	}
	return m.Pins[m.PinNames[pid+m.InPinCount()]]

}

func (m Filter) PinByName(pinname string) chipset.PinInfo {
	return m.Pins[pinname]
}

func (m Filter) ModulesCount() int {
	return 1
}
func (m Filter) ModuleByName(mname string) chipset.ModuleInfo {
	return m.Modules[mname]
}

func (m Filter) Module(moduleid int) chipset.ModuleInfo {
	return m.ModuleByName(m.ModuleNames[moduleid])
}

func (m Filter) SayHello() {
	fmt.Printf("\n Hi from \n %v", m.Name())
}

func (m Filter) Name() string {
	return "Filter"
}

func (m Filter) IsInitialized() bool {
	return m.isInitialized
}

func (m *Filter) InitializeChip() {
	m.name = "Filter"
	m.InitPins()
	m.InitModules()
}

func (f Filter) GetCoeff() (coeff vlib.VectorC) {
	return f.coeff
}

func (f *Filter) SetCoeff(coeff vlib.VectorC) {
	f.coeff = coeff
	Taps := f.coeff.Size()
	f.FilterMemory.Resize(Taps)
}

func (m *Filter) InitPins() {
	totalpins := m.InPinCount() + m.OutPinCount()
	m.Pins = make(map[string]chipset.PinInfo, totalpins)
	m.PinNames = make(map[int]string, totalpins)
	strlist := [2]string{"inputPin0", "outputPin0"}
	for cnt := 0; cnt < len(strlist); cnt++ {
		m.PinNames[cnt] = strlist[cnt]
	}

	/// something try begins
	var pinfo chipset.PinInfo

	pinfo.Name = "inputPin0"
	pinfo.Id = 0
	pinfo.InputPin = true
	pinfo.DataType = reflect.TypeOf((*gocomm.Complex128Channel)(nil)).Elem()

	m.Pins["inputPin0"] = pinfo

	pinfo.Name = "outputPin0"
	pinfo.Id = 1
	pinfo.InputPin = false
	pinfo.DataType = reflect.TypeOf((*gocomm.Complex128Channel)(nil)).Elem()

	pinfo.CreateChannel()

	m.Pins["outputPin0"] = pinfo

	/// something try ends

}

func (m *Filter) InitModules() {

	var totalModules int = 1

	/// AUTO CODE
	/// something try begins
	var minfo chipset.ModuleInfo
	m.Modules = make(map[string]chipset.ModuleInfo, totalModules)
	m.ModuleNames = make(map[int]string, totalModules)

	strlist := [1]string{"filter"}
	for cnt := 0; cnt < len(strlist); cnt++ {
		m.ModuleNames[cnt] = strlist[cnt]
	}
	var temp, otemp []int

	minfo.Name = "filter"
	minfo.Id = 0
	minfo.Desc = ""

	temp = append(temp, m.PinByName("inputPin0").Id)

	otemp = append(otemp, m.PinByName("outputPin0").Id)

	minfo.InPins = temp
	minfo.OutPins = otemp
	m.Modules["filter"] = minfo

	/// AUTO CODE

	m.isInitialized = true
}

func (m *Filter) Filter(inputPin0 gocomm.Complex128Channel) {
	/// Read your data from Input channel(s) [inputPin0]
	/// And write it to OutputChannels  [outputPin0]
	//Taps := m.GetCoeff().Size()
	// fmt.Printf("\n Filters : %v ,%v", Taps, m.Coeff)
	outputPin0 := m.Pins["outputPin0"].Channel.(gocomm.Complex128Channel)
	iters := 1
	for i := 0; i < iters; i++ {
		chData := <-inputPin0
		iters = chData.MaxExpected

		/// Do process here with chData
		outData := m.FilterFn(chData)
		// fmt.Printf("\n %d/%d : Output %v ", i, iters, outData.Ch)

		outputPin0 <- outData
		//outData.MaxExpected = chData.MaxExpected + Taps - 1 /// no use of this line
	}

}

func (m *Filter) FilterFn(sample gocomm.SComplex128Obj) (result gocomm.SComplex128Obj) {

	result = sample /// Carefull if not same ChType
	m.FilterMemory = m.FilterMemory.ShiftLeft(sample.Ch)
	foutput := vlib.DotC(m.coeff, m.FilterMemory)

	result.Ch = foutput
	result.Message = sample.Message + " Filter"
	result.Ts = sample.Ts
	result.TimeStamp = sample.TimeStamp
	result.MaxExpected = sample.MaxExpected

	return result
}