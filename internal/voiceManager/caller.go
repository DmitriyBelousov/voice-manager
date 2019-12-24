package voiceManager

type Caller interface{
	MakeCall()
	CancelCall()
}