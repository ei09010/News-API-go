package constants

type Category string

const (
	Business      Category = "Business"
	Entertainment          = "Entertainment"
	Health                 = "Health"
	Science                = "Science"
	Sports                 = "Sports"
	Technology             = "Technology"
)

type Country int

const (
	AE Country = iota
	/// <summary>
	/// Argentina
	/// </summary>
	AR
	/// <summary>
	/// Austria
	/// </summary>
	AT
	/// <summary>
	/// Australia
	/// </summary>
	AU
	/// <summary>
	/// Belgium
	/// </summary>
	BE
	BG
	BR
	/// <summary>
	/// Canada
	/// </summary>
	CA
	CH
	/// <summary>
	/// China
	/// </summary>
	CN
	CO
	CU
	/// <summary>
	/// Czech Republic
	/// </summary>
	CZ
	/// <summary>
	/// Germany
	/// </summary>
	DE
	/// <summary>
	/// Egypt
	/// </summary>
	EG
	/// <summary>
	/// France
	/// </summary>
	FR
	/// <summary>
	/// United Kingdom
	/// </summary>
	GB
	/// <summary>
	/// Greece
	/// </summary>
	GR
	/// <summary>
	/// Hong Kong
	/// </summary>
	HK
	/// <summary>
	/// Hungary
	/// </summary>
	HU
	ID
	/// <summary>
	/// Ireland
	/// </summary>
	IE
	IL
	IN
	/// <summary>
	/// Italy
	/// </summary>
	IT
	/// <summary>
	/// Japan
	/// </summary>
	JP
	/// <summary>
	/// South Korea
	/// </summary>
	KR
	LT
	LV
	MA
	/// <summary>
	/// Mexico
	/// </summary>
	MX
	MY
	NG
	/// <summary>
	/// Netherlands
	/// </summary>
	NL
	/// <summary>
	/// Norway
	/// </summary>
	NO
	/// <summary>
	/// New Zealand
	/// </summary>
	NZ
	PH
	PL
	/// <summary>
	/// Portugal
	/// </summary>
	PT
	RO
	RS
	/// <summary>
	/// Russia
	/// </summary>
	RU
	SA
	SE
	SG
	SI
	SK
	TH
	TR
	TW
	UA
	/// <summary>
	/// United States
	/// </summary>
	US
	VE
	ZA
)

type ErrorCode int

const (
	ApiKeyExhausted ErrorCode = iota
	ApiKeyMissing
	ApiKeyInvalid
	ApiKeyDisabled
	ParametersMissing
	ParametersIncompatible
	ParameterInvalid
	RateLimited
	RequestTimeout
	SourcesTooMany
	SourceDoesNotExist
	SourceUnavailableSortedBy
	SourceTemporarilyUnavailable
	UnexpectedError
	UnknownError
)

type Language string

const (
	/// <summary>
	/// Afrikaans (South Africa)
	/// </summary>
	AF Language = "AF"
	AN
	AR_lang
	AZ
	BG_lang
	BN
	BR_lang
	BS
	CA_lang
	CS
	CY
	DA
	/// <summary>
	/// German
	/// </summary>
	DE_lang
	EL
	/// <summary>
	/// English
	/// </summary>
	EN
	EO
	/// <summary>
	/// Spanish
	/// </summary>
	ES
	ET
	EU
	FA
	FI
	FR_lang
	GL
	HE
	HI
	HR
	HT
	HU_lang
	HY
	ID_lang
	IS
	/// <summary>
	/// Italian
	/// </summary>
	IT_lang
	/// <summary>
	/// Japanese
	/// </summary>
	JP_lang
	JV
	KK
	KO
	LA
	LB
	LT_lang
	LV_lang
	MG
	MK
	ML
	MR
	MS
	/// <summary>
	/// Dutch
	/// </summary>
	NL_lang
	NN
	NO_lang
	OC
	PL_lang
	/// <summary>
	/// Portuguese
	/// </summary>
	PT_lang
	RO_lang
	RU_lang
	SH
	SK_lang
	SL
	SQ
	SR
	SV
	SW
	TA
	TE
	TH_lang
	TL
	TR_lang
	UK
	UR
	VI
	VO
	/// <summary>
	/// Chinese
	/// </summary>
	ZH
)

type SortBy int

const (
	/// <summary>
	/// Sort by publisher popularity
	/// </summary>
	Popularity SortBy = iota
	/// <summary>
	/// Sort by article publish date (newest first)
	/// </summary>
	PublishedAt
	/// <summary>
	/// Sort by relevancy to the Q param
	/// </summary>
	Relevancy
)

type Status int

const (
	/// <summary>
	/// Request was successful
	/// </summary>
	Ok Status = iota
	/// <summary>
	/// Request failed
	/// </summary>
	Error
)
