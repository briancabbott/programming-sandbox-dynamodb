// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Provides lexicon name and lexicon content in string format. For more
// information, see Pronunciation Lexicon Specification (PLS) Version 1.0
// (https://www.w3.org/TR/pronunciation-lexicon/).
type Lexicon struct {

	// Lexicon content in string format. The content of a lexicon must be in PLS
	// format.
	Content *string

	// Name of the lexicon.
	Name *string
}

// Contains metadata describing the lexicon such as the number of lexemes, language
// code, and so on. For more information, see Managing Lexicons
// (https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
type LexiconAttributes struct {

	// Phonetic alphabet used in the lexicon. Valid values are ipa and x-sampa.
	Alphabet *string

	// Language code that the lexicon applies to. A lexicon with a language code such
	// as "en" would be applied to all English languages (en-GB, en-US, en-AUS, en-WLS,
	// and so on.
	LanguageCode LanguageCode

	// Date lexicon was last modified (a timestamp value).
	LastModified *time.Time

	// Number of lexemes in the lexicon.
	LexemesCount int32

	// Amazon Resource Name (ARN) of the lexicon.
	LexiconArn *string

	// Total size of the lexicon, in characters.
	Size int32
}

// Describes the content of the lexicon.
type LexiconDescription struct {

	// Provides lexicon metadata.
	Attributes *LexiconAttributes

	// Name of the lexicon.
	Name *string
}

// SynthesisTask object that provides information about a speech synthesis task.
type SynthesisTask struct {

	// Timestamp for the time the synthesis task was started.
	CreationTime *time.Time

	// Specifies the engine (standard or neural) for Amazon Polly to use when
	// processing input text for speech synthesis. Using a voice that is not supported
	// for the engine selected will result in an error.
	Engine Engine

	// Optional language code for a synthesis task. This is only necessary if using a
	// bilingual voice, such as Aditi, which can be used for either Indian English
	// (en-IN) or Hindi (hi-IN). If a bilingual voice is used and no language code is
	// specified, Amazon Polly will use the default language of the bilingual voice.
	// The default language for any voice is the one returned by the DescribeVoices
	// (https://docs.aws.amazon.com/polly/latest/dg/API_DescribeVoices.html) operation
	// for the LanguageCode parameter. For example, if no language code is specified,
	// Aditi will use Indian English rather than Hindi.
	LanguageCode LanguageCode

	// List of one or more pronunciation lexicon names you want the service to apply
	// during synthesis. Lexicons are applied only if the language of the lexicon is
	// the same as the language of the voice.
	LexiconNames []string

	// The format in which the returned output will be encoded. For audio stream, this
	// will be mp3, ogg_vorbis, or pcm. For speech marks, this will be json.
	OutputFormat OutputFormat

	// Pathway for the output speech file.
	OutputUri *string

	// Number of billable characters synthesized.
	RequestCharacters int32

	// The audio frequency specified in Hz. The valid values for mp3 and ogg_vorbis are
	// "8000", "16000", "22050", and "24000". The default value for standard voices is
	// "22050". The default value for neural voices is "24000". Valid values for pcm
	// are "8000" and "16000" The default value is "16000".
	SampleRate *string

	// ARN for the SNS topic optionally used for providing status notification for a
	// speech synthesis task.
	SnsTopicArn *string

	// The type of speech marks returned for the input text.
	SpeechMarkTypes []SpeechMarkType

	// The Amazon Polly generated identifier for a speech synthesis task.
	TaskId *string

	// Current status of the individual speech synthesis task.
	TaskStatus TaskStatus

	// Reason for the current status of a specific speech synthesis task, including
	// errors if the task has failed.
	TaskStatusReason *string

	// Specifies whether the input text is plain text or SSML. The default value is
	// plain text.
	TextType TextType

	// Voice ID to use for the synthesis.
	VoiceId VoiceId
}

// Description of the voice.
type Voice struct {

	// Additional codes for languages available for the specified voice in addition to
	// its default language. For example, the default language for Aditi is Indian
	// English (en-IN) because it was first used for that language. Since Aditi is
	// bilingual and fluent in both Indian English and Hindi, this parameter would show
	// the code hi-IN.
	AdditionalLanguageCodes []LanguageCode

	// Gender of the voice.
	Gender Gender

	// Amazon Polly assigned voice ID. This is the ID that you specify when calling the
	// SynthesizeSpeech operation.
	Id VoiceId

	// Language code of the voice.
	LanguageCode LanguageCode

	// Human readable name of the language in English.
	LanguageName *string

	// Name of the voice (for example, Salli, Kendra, etc.). This provides a human
	// readable voice name that you might display in your application.
	Name *string

	// Specifies which engines (standard or neural) that are supported by a given
	// voice.
	SupportedEngines []Engine
}
