package tool

import (
	"log"

	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/mmcdole/gofeed"
)

var VOICE_NAME_MAP_WAVENET = map[string]string{
	"cmn-CN": "cmn-CN-Wavenet-A",
	"en-US":  "en-US-Neural2-C",
}

var VOICE_NAME_MAP_STANDARD = map[string]string{
	"cmn-CN": "cmn-CN-Standard-D",
	"en-US":  "en-US-Standard-C",
}

func getSanitizedContentChunks(item *gofeed.Item) (textchunks []string) {
	content := item.Title + "\n\n"

	if len(item.Content) > 0 {
		content += StripHtmlTags(item.Content)
	} else if len(item.Description) > 0 {
		content += StripHtmlTags(item.Description)
	}

	return ChunksByte(content, 5000)
}

func GetSynthesizeSpeechRequests(item *gofeed.Item, lang string, useNaturalVoice bool, speechSpeed float64) []*texttospeechpb.SynthesizeSpeechRequest {
	contentChunks := getSanitizedContentChunks(item)

	if len(lang) == 0 {
		lang = GuessLanguageByUnicode(item.Title)
	}

	lang = GetSanitizedLanguageCode(lang)

	log.Printf("lang: %v\n", lang)

	languageName := VOICE_NAME_MAP_STANDARD[lang]
	if useNaturalVoice {
		languageName = VOICE_NAME_MAP_WAVENET[lang]
	}

	log.Printf("using voice %v for language code %v", languageName, lang)

	synthesizeRequest := make([]*texttospeechpb.SynthesizeSpeechRequest, 0)

	for _, chunk := range contentChunks {

		req := texttospeechpb.SynthesizeSpeechRequest{
			// Set the text input to be synthesized.
			Input: &texttospeechpb.SynthesisInput{
				InputSource: &texttospeechpb.SynthesisInput_Text{Text: chunk},
			},
			// Build the voice request, select the language code ("en-US") and the SSML
			// voice gender ("neutral").
			Voice: &texttospeechpb.VoiceSelectionParams{
				LanguageCode: lang,
				Name:         languageName,
				SsmlGender:   texttospeechpb.SsmlVoiceGender_FEMALE,
			},
			// Select the type of audio file you want returned.
			AudioConfig: &texttospeechpb.AudioConfig{
				AudioEncoding: texttospeechpb.AudioEncoding_MP3,
				SpeakingRate:  speechSpeed,
			},
		}

		synthesizeRequest = append(synthesizeRequest, &req)
	}

	return synthesizeRequest
}
