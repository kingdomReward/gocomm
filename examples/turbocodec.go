
import (
	"fmt"
	"reflect"
	"wiless/gocomm"
	"wiless/gocomm/chipset"
	"wiless/vlib"
)


func init(){
	fmt.Printf("\n AutoGenerated package customchips")
}


type TurboCodec struct {
	name               string
	isInitialized      bool
	Pins               map[string]chipset.PinInfo
	Modules            map[string]chipset.ModuleInfo
	ModuleNames        map[int]string
	PinNames           map[int]string
}

/// AutoGenerated through script

func (m TurboCodec) InPinCount() int {
	return 2
}

func (m TurboCodec) OutPinCount() int {
	return 2
}

func (m TurboCodec) Pin(pid int) chipset.PinInfo {
	return m.Pins[m.PinNames[pid]]
}

func (m TurboCodec) PinIn(pid int) chipset.PinInfo {
	if pid >= m.InPinCount() {
		fmt.Printf("%d > No of Input Pins %d", pid, m.InPinCount())
		var result chipset.PinInfo
		result.Id = -1
		return result
	}

	return m.Pins[m.PinNames[pid]]

}
func (m TurboCodec) PinByID(pid int) chipset.PinInfo {

	return m.Pins[m.PinNames[pid]]
}

func (m TurboCodec) PinOut(pid int) chipset.PinInfo {
	if pid >= m.OutPinCount() {
		fmt.Printf("%d > No of Output Pins %d", pid, m.OutPinCount())
		var result chipset.PinInfo
		result.Id = -1
		return result
	}
	return m.Pins[m.PinNames[pid+m.InPinCount()]]

}

func (m TurboCodec) PinByName(pinname string) chipset.PinInfo {
	return m.Pins[pinname]
}

func (m TurboCodec) ModulesCount() int {
	return 2
}
func (m TurboCodec) ModuleByName(mname string) chipset.ModuleInfo {
	return m.Modules[mname]
}

func (m TurboCodec) Module(moduleid int) chipset.ModuleInfo {
	return m.ModuleByName(m.ModuleNames[moduleid])
}

func (m TurboCodec) SayHello() {
	fmt.Printf("\n Hi from \n %v", m.Name())
}

func (m TurboCodec) Name() string {
	return "TurboCodec"
}

func (m TurboCodec) IsInitialized() bool {
	return m.isInitialized
}

func (m *TurboCodec) InitializeChip() {
	m.name = "TurboCodec"
	m.InitPins()
	m.InitModules()
}


func (m *TurboCodec) InitPins() {
	totalpins := m.InPinCount() + m.OutPinCount()
	m.Pins = make(map[string]chipset.PinInfo, totalpins)
	m.PinNames = make(map[int]string, totalpins)
	strlist := [4]string{ "uncodedBits","codedBits","decodedBits" }
	for cnt:=0;cnt<len(strlist);cnt++ {
	m.PinNames[cnt]=strlist[cnt]
	}


/// something try begins
	var pinfo chipset.PinInfo

	pinfo.Name = "uncodedBits"
	pinfo.Id = 0
	pinfo.InputPin = true
	pinfo.DataType = reflect.TypeOf((*gocomm.BitChannel)(nil)).Elem()
	
	
	m.Pins["uncodedBits"] = pinfo

	pinfo.Name = "LLRbits"
	pinfo.Id = 1
	pinfo.InputPin = true
	pinfo.DataType = reflect.TypeOf((*gocomm.BitChannel)(nil)).Elem()
	
	
	m.Pins["LLRbits"] = pinfo

	pinfo.Name = "codedBits"
	pinfo.Id = 2
	pinfo.InputPin = false
	pinfo.DataType = reflect.TypeOf((*gocomm.Complex128Channel)(nil)).Elem()
	
	pinfo.CreateChannel()
	
	m.Pins["codedBits"] = pinfo

	pinfo.Name = "decodedBits"
	pinfo.Id = 3
	pinfo.InputPin = false
	pinfo.DataType = reflect.TypeOf((*gocomm.Complex128Channel)(nil)).Elem()
	
	pinfo.CreateChannel()
	
	m.Pins["decodedBits"] = pinfo


/// something try ends

 
}






func (m *TurboCodec) InitModules() {


	var totalModules int = 2
	 
/// AUTO CODE
/// Begin module i ----------------------------
	var minfo chipset.ModuleInfo
	m.Modules = make(map[string]chipset.ModuleInfo, totalModules)
	m.ModuleNames = make(map[int]string, totalModules)
	
	strlist := [2]string{ "encoder","decoder" }
	for cnt:=0;cnt<len(strlist);cnt++ {
	m.ModuleNames[cnt]=strlist[cnt]
	}
	var temp,otemp []int
	



	minfo.Name = "encoder"
	minfo.Id = 0
	minfo.Desc = ""	
	
	temp=append(temp,m.PinByName("uncodedBits").Id)
	

	otemp=append(otemp,m.PinByName("codedBits").Id)
	
	minfo.InPins = temp
	minfo.OutPins = otemp
	method := reflect.ValueOf(m).MethodByName("encode")
	minfo.Function = method
	minfo.FunctionName = "encode"

	m.Modules["encoder"]=minfo
	/// End module i ----------------------------


	minfo.Name = "decoder"
	minfo.Id = 1
	minfo.Desc = ""	
	
	temp=append(temp,m.PinByName("codedBits").Id)
	

	otemp=append(otemp,m.PinByName("decodedBits").Id)
	
	minfo.InPins = temp
	minfo.OutPins = otemp
	method := reflect.ValueOf(m).MethodByName("decode")
	minfo.Function = method
	minfo.FunctionName = "decode"

	m.Modules["decoder"]=minfo
	/// End module i ----------------------------


/// AUTO CODE
	
m.isInitialized=true
}






func (m *TurboCodec) encode(uncodedBits gocomm.BitChannel ) {
/// Read your data from Input channel(s) [uncodedBits] 
/// And write it to OutputChannels  [codedBits]
/*

///	codedBits:=m.Pins["codedBits"].Channel.(gocomm.<DataType>)
	iters := 1
	for i := 0; i < iters; i++ {
		chData := <-[uncodedBits]	
		iters = chData.MaxExpected	
		/// Do process here with chData

		outData:= encodeFn(chData)
		outData.MaxExpected= ??
		codedBits <- outData
		
	}
	*/

}



func (m *TurboCodec) decode(codedBits gocomm.Complex128Channel ) {
/// Read your data from Input channel(s) [codedBits] 
/// And write it to OutputChannels  [decodedBits]
/*

///	decodedBits:=m.Pins["decodedBits"].Channel.(gocomm.<DataType>)
	iters := 1
	for i := 0; i < iters; i++ {
		chData := <-[codedBits]	
		iters = chData.MaxExpected	
		/// Do process here with chData

		outData:= decodeFn(chData)
		outData.MaxExpected= ??
		decodedBits <- outData
		
	}
	*/

}






func (m *TurboCodec) encodeFn(uncodedBits gocomm.BitChannel ) {
/// Read your data from Input channel(s) [uncodedBits] 
/// And write it to OutputChannels  [codedBits]
/*

///	codedBits:=m.Pins["codedBits"].Channel.(gocomm.<DataType>)
	iters := 1
	for i := 0; i < iters; i++ {
		chData := <-[uncodedBits]	
		iters = chData.MaxExpected	
		/// Do process here with chData

		outData:= m.encodeFn(chData)
		codedBits <- outData
		
	}
	*/

}



func (m *TurboCodec) decodeFn(codedBits gocomm.Complex128Channel ) {
/// Read your data from Input channel(s) [codedBits] 
/// And write it to OutputChannels  [decodedBits]
/*

///	decodedBits:=m.Pins["decodedBits"].Channel.(gocomm.<DataType>)
	iters := 1
	for i := 0; i < iters; i++ {
		chData := <-[codedBits]	
		iters = chData.MaxExpected	
		/// Do process here with chData

		outData:= m.decodeFn(chData)
		decodedBits <- outData
		
	}
	*/

}














