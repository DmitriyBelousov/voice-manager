package voiceManager

type Voicer interface{
	ParseCommand()
	ParseName() string
}