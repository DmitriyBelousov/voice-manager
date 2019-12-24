package voiceManager

import "fmt"

type VoiceManager interface{
	Start()
	Stop()
	Close()
}

//VoiceClient ...
type VoiceClient struct{
	finder Finder
	caller Caller
	voicer Voicer
}

//Start ...
func (vc *VoiceClient) Start(){
	vc.voicer.ParseCommand()
	vc.voicer.ParseName()

	vc.finder.OpenPhoneBook()
	defer 	vc.finder.ClosePoneBook()
	vc.finder.FindContact()

	vc.caller.MakeCall()
	vc.Close()
}

///Stop ...
func (vc *VoiceClient) Stop(){
	vc.caller.CancelCall()
}

//Close ...
func (vc *VoiceClient) Close(){
	fmt.Println("Закрытие голосового менеджера")
}

//New factory func
func New(f Finder, c Caller, v Voicer) VoiceClient{
	return VoiceClient{
		caller: c,
		finder: f,
		voicer: v,
	}
}