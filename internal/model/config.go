package model

type (
	Config struct {
        Id int
		// UseNaturalVoice determines whether to use natural voice from Google, this only have quota of 1
		// Million, whereas the quota for standard voice is 4 Million
		UseNaturalVoice bool `json:"useNaturalVoice"`
		// Speed of synthesized speech
		SpeechSpeed float64 `json:"speechSpeed"`
		// FullTextServiceUrl a full text service to be used when a feed is not
		//in full text
		FullTextServiceUrl string `json:"fullTextServiceUrl,omitempty"`
	}
)

