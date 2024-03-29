package goStr

var apaMinorWords = map[string]bool{
	"a":     true,
	"an":    true,
	"the":   true,
	"and":   true,
	"but":   true,
	"or":    true,
	"for":   true,
	"nor":   true,
	"on":    true,
	"at":    true,
	"to":    true,
	"by":    true,
	"with":  true,
	"of":    true,
	"in":    true,
	"into":  true,
	"near":  true,
	"over":  true,
	"off":   true,
	"up":    true,
	"down":  true,
	"out":   true,
	"as":    true,
	"about": true,
	"from":  true,
	"via":   true,
}

type CurrencyCode string

const (
	CURRENCY_AED CurrencyCode = "AED" // United Arab Emirates Dirham
	CURRENCY_AFN CurrencyCode = "AFN" // Afghan Afghani
	CURRENCY_ALL CurrencyCode = "ALL" // Albanian Lek
	CURRENCY_AMD CurrencyCode = "AMD" // Armenian Dram
	CURRENCY_ANG CurrencyCode = "ANG" // Netherlands Antillean Guilder
	CURRENCY_AOA CurrencyCode = "AOA" // Angolan Kwanza
	CURRENCY_ARS CurrencyCode = "ARS" // Argentine Peso
	CURRENCY_AUD CurrencyCode = "AUD" // Australian Dollar
	CURRENCY_AWG CurrencyCode = "AWG" // Aruban Florin
	CURRENCY_BAM CurrencyCode = "BAM" // Bosnia-Herzegovina Convertible Mark
	CURRENCY_BBD CurrencyCode = "BBD" // Barbadian Dollar
	CURRENCY_BDT CurrencyCode = "BDT" // Bangladeshi Taka
	CURRENCY_BGN CurrencyCode = "BGN" // Bulgarian Lev
	CURRENCY_BHD CurrencyCode = "BHD" // Bahraini Dinar
	CURRENCY_BIF CurrencyCode = "BIF" // Burundian Franc
	CURRENCY_BMD CurrencyCode = "BMD" // Bermudian Dollar
	CURRENCY_BND CurrencyCode = "BND" // Brunei Dollar
	CURRENCY_BOB CurrencyCode = "BOB" // Bolivian Boliviano
	CURRENCY_BRL CurrencyCode = "BRL" // Brazilian Real
	CURRENCY_BSD CurrencyCode = "BSD" // Bahamian Dollar
	CURRENCY_BTN CurrencyCode = "BTN" // Bhutanese Ngultrum
	CURRENCY_BWP CurrencyCode = "BWP" // Botswana Pula
	CURRENCY_BYN CurrencyCode = "BYN" // Belarusian Ruble
	CURRENCY_BZD CurrencyCode = "BZD" // Belize Dollar
	CURRENCY_CAD CurrencyCode = "CAD" // Canadian Dollar
	CURRENCY_CDF CurrencyCode = "CDF" // Congolese Franc
	CURRENCY_CHF CurrencyCode = "CHF" // Swiss Franc
	CURRENCY_CLP CurrencyCode = "CLP" // Chilean Peso
	CURRENCY_CNY CurrencyCode = "CNY" // Chinese Yuan Renminbi
	CURRENCY_COP CurrencyCode = "COP" // Colombian Peso
	CURRENCY_CRC CurrencyCode = "CRC" // Costa Rican Colon
	CURRENCY_CUP CurrencyCode = "CUP" // Cuban Peso
	CURRENCY_CVE CurrencyCode = "CVE" // Cape Verdean Escudo
	CURRENCY_CZK CurrencyCode = "CZK" // Czech Koruna
	CURRENCY_DJF CurrencyCode = "DJF" // Djiboutian Franc
	CURRENCY_DKK CurrencyCode = "DKK" // Danish Krone
	CURRENCY_DOP CurrencyCode = "DOP" // Dominican Peso
	CURRENCY_DZD CurrencyCode = "DZD" // Algerian Dinar
	CURRENCY_EGP CurrencyCode = "EGP" // Egyptian Pound
	CURRENCY_ERN CurrencyCode = "ERN" // Eritrean Nakfa
	CURRENCY_ETB CurrencyCode = "ETB" // Ethiopian Birr
	CURRENCY_EUR CurrencyCode = "EUR" // Euro
	CURRENCY_FJD CurrencyCode = "FJD" // Fijian Dollar
	CURRENCY_FKP CurrencyCode = "FKP" // Falkland Islands Pound
	CURRENCY_GBP CurrencyCode = "GBP" // British Pound Sterling
	CURRENCY_GEL CurrencyCode = "GEL" // Georgian Lari
	CURRENCY_GGP CurrencyCode = "GGP" // Guernsey Pound
	CURRENCY_GHS CurrencyCode = "GHS" // Ghanaian Cedi
	CURRENCY_GIP CurrencyCode = "GIP" // Gibraltar Pound
	CURRENCY_GMD CurrencyCode = "GMD" // Gambian Dalasi
	CURRENCY_GNF CurrencyCode = "GNF" // Guinean Franc
	CURRENCY_GTQ CurrencyCode = "GTQ" // Guatemalan Quetzal
	CURRENCY_GYD CurrencyCode = "GYD" // Guyanaese Dollar
	CURRENCY_HKD CurrencyCode = "HKD" // Hong Kong Dollar
	CURRENCY_HNL CurrencyCode = "HNL" // Honduran Lempira
	CURRENCY_HRK CurrencyCode = "HRK" // Croatian Kuna
	CURRENCY_HTG CurrencyCode = "HTG" // Haitian Gourde
	CURRENCY_HUF CurrencyCode = "HUF" // Hungarian Forint
	CURRENCY_IDR CurrencyCode = "IDR" // Indonesian Rupiah
	CURRENCY_ILS CurrencyCode = "ILS" // Israeli New Sheqel
	CURRENCY_IMP CurrencyCode = "IMP" // Isle of Man Pound
	CURRENCY_INR CurrencyCode = "INR" // Indian Rupee
	CURRENCY_IQD CurrencyCode = "IQD" // Iraqi Dinar
	CURRENCY_IRR CurrencyCode = "IRR" // Iranian Rial
	CURRENCY_ISK CurrencyCode = "ISK" // Icelandic Krona
	CURRENCY_JEP CurrencyCode = "JEP" // Jersey Pound
	CURRENCY_JMD CurrencyCode = "JMD" // Jamaican Dollar
	CURRENCY_JOD CurrencyCode = "JOD" // Jordanian Dinar
	CURRENCY_JPY CurrencyCode = "JPY" // Japanese Yen
	CURRENCY_KES CurrencyCode = "KES" // Kenyan Shilling
	CURRENCY_KGS CurrencyCode = "KGS" // Kyrgyzstani Som
	CURRENCY_KHR CurrencyCode = "KHR" // Cambodian Riel
	CURRENCY_KID CurrencyCode = "KID" // Kiribati Dollar
	CURRENCY_KMF CurrencyCode = "KMF" // Comorian Franc
	CURRENCY_KRW CurrencyCode = "KRW" // South Korean Won
	CURRENCY_KWD CurrencyCode = "KWD" // Kuwaiti Dinar
	CURRENCY_KYD CurrencyCode = "KYD" // Cayman Islands Dollar
	CURRENCY_KZT CurrencyCode = "KZT" // Kazakhstani Tenge
	CURRENCY_LAK CurrencyCode = "LAK" // Laotian Kip
	CURRENCY_LBP CurrencyCode = "LBP" // Lebanese Pound
	CURRENCY_LKR CurrencyCode = "LKR" // Sri Lankan Rupee
	CURRENCY_LRD CurrencyCode = "LRD" // Liberian Dollar
	CURRENCY_LSL CurrencyCode = "LSL" // Lesotho Loti
	CURRENCY_LYD CurrencyCode = "LYD" // Libyan Dinar
	CURRENCY_MAD CurrencyCode = "MAD" // Moroccan Dirham
	CURRENCY_MDL CurrencyCode = "MDL" // Moldovan Leu
	CURRENCY_MGA CurrencyCode = "MGA" // Malagasy Ariary
	CURRENCY_MKD CurrencyCode = "MKD" // Macedonian Denar
	CURRENCY_MMK CurrencyCode = "MMK" // Myanmar Kyat
	CURRENCY_MNT CurrencyCode = "MNT" // Mongolian Tugrik
	CURRENCY_MOP CurrencyCode = "MOP" // Macanese Pataca
	CURRENCY_MRO CurrencyCode = "MRO" // Mauritanian Ouguiya (pre-2018)
	CURRENCY_MRU CurrencyCode = "MRU" // Mauritanian Ouguiya
	CURRENCY_MUR CurrencyCode = "MUR" // Mauritian Rupee
	CURRENCY_MVR CurrencyCode = "MVR" // Maldivian Rufiyaa
	CURRENCY_MWK CurrencyCode = "MWK" // Malawian Kwacha
	CURRENCY_MXN CurrencyCode = "MXN" // Mexican Peso
	CURRENCY_MYR CurrencyCode = "MYR" // Malaysian Ringgit
	CURRENCY_MZN CurrencyCode = "MZN" // Mozambican Metical
	CURRENCY_NAD CurrencyCode = "NAD" // Namibian Dollar
	CURRENCY_NGN CurrencyCode = "NGN" // Nigerian Naira
	CURRENCY_NIO CurrencyCode = "NIO" // Nicaraguan Cordoba
	CURRENCY_NOK CurrencyCode = "NOK" // Norwegian Krone
	CURRENCY_NPR CurrencyCode = "NPR" // Nepalese Rupee
	CURRENCY_NZD CurrencyCode = "NZD" // New Zealand Dollar
	CURRENCY_OMR CurrencyCode = "OMR" // Omani Rial
	CURRENCY_PAB CurrencyCode = "PAB" // Panamanian Balboa
	CURRENCY_PEN CurrencyCode = "PEN" // Peruvian Nuevo Sol
	CURRENCY_PGK CurrencyCode = "PGK" // Papua New Guinean Kina
	CURRENCY_PHP CurrencyCode = "PHP" // Philippine Peso
	CURRENCY_PKR CurrencyCode = "PKR" // Pakistani Rupee
	CURRENCY_PLN CurrencyCode = "PLN" // Polish Zloty
	CURRENCY_PYG CurrencyCode = "PYG" // Paraguayan Guarani
	CURRENCY_QAR CurrencyCode = "QAR" // Qatari Rial
	CURRENCY_RON CurrencyCode = "RON" // Romanian Leu
	CURRENCY_RSD CurrencyCode = "RSD" // Serbian Dinar
	CURRENCY_RUB CurrencyCode = "RUB" // Russian Ruble
	CURRENCY_RWF CurrencyCode = "RWF" // Rwandan Franc
	CURRENCY_SAR CurrencyCode = "SAR" // Saudi Riyal
	CURRENCY_SBD CurrencyCode = "SBD" // Solomon Islands Dollar
	CURRENCY_SCR CurrencyCode = "SCR" // Seychellois Rupee
	CURRENCY_SDG CurrencyCode = "SDG" // Sudanese Pound
	CURRENCY_SEK CurrencyCode = "SEK" // Swedish Krona
	CURRENCY_SGD CurrencyCode = "SGD" // Singapore Dollar
	CURRENCY_SHP CurrencyCode = "SHP" // Saint Helena Pound
	CURRENCY_SLL CurrencyCode = "SLL" // Sierra Leonean Leone
	CURRENCY_SOS CurrencyCode = "SOS" // Somali Shilling
	CURRENCY_SRD CurrencyCode = "SRD" // Surinamese Dollar
	CURRENCY_SSP CurrencyCode = "SSP" // South Sudanese Pound
	CURRENCY_STN CurrencyCode = "STN" // Sao Tome and Principe Dobra
	CURRENCY_SYP CurrencyCode = "SYP" // Syrian Pound
	CURRENCY_SZL CurrencyCode = "SZL" // Swazi Lilangeni
	CURRENCY_THB CurrencyCode = "THB" // Thai Baht
	CURRENCY_TJS CurrencyCode = "TJS" // Tajikistani Somoni
	CURRENCY_TMT CurrencyCode = "TMT" // Turkmenistani Manat
	CURRENCY_TND CurrencyCode = "TND" // Tunisian Dinar
	CURRENCY_TOP CurrencyCode = "TOP" // Tongan Pa'anga
	CURRENCY_TRY CurrencyCode = "TRY" // Turkish Lira
	CURRENCY_TTD CurrencyCode = "TTD" // Trinidad and Tobago Dollar
	CURRENCY_TWD CurrencyCode = "TWD" // New Taiwan Dollar
	CURRENCY_TZS CurrencyCode = "TZS" // Tanzanian Shilling
	CURRENCY_UAH CurrencyCode = "UAH" // Ukrainian Hryvnia
	CURRENCY_UGX CurrencyCode = "UGX" // Ugandan Shilling
	CURRENCY_USD CurrencyCode = "USD" // United States Dollar
	CURRENCY_UYU CurrencyCode = "UYU" // Uruguayan Peso
	CURRENCY_UZS CurrencyCode = "UZS" // Uzbekistan Som
	CURRENCY_VES CurrencyCode = "VES" // Venezuelan Bolivar
	CURRENCY_VND CurrencyCode = "VND" // Vietnamese Dong
	CURRENCY_VUV CurrencyCode = "VUV" // Vanuatu Vatu
	CURRENCY_WST CurrencyCode = "WST" // Samoan Tala
	CURRENCY_XAF CurrencyCode = "XAF" // CFA Franc BEAC
	CURRENCY_XCD CurrencyCode = "XCD" // East Caribbean Dollar
	CURRENCY_XDR CurrencyCode = "XDR" // Special Drawing Rights
	CURRENCY_XOF CurrencyCode = "XOF" // CFA Franc BCEAO
	CURRENCY_XPF CurrencyCode = "XPF" // CFP Franc
	CURRENCY_YER CurrencyCode = "YER" // Yemeni Rial
	CURRENCY_ZAR CurrencyCode = "ZAR" // South African Rand
	CURRENCY_ZMW CurrencyCode = "ZMW" // Zambian Kwacha
	CURRENCY_ZWL CurrencyCode = "ZWL" // Zimbabwean Dollar
)

var CurrencyFormats = map[CurrencyCode]string{
	CURRENCY_AED: "د.إ",
	CURRENCY_AFN: "؋",
	CURRENCY_ALL: "L",
	CURRENCY_AMD: "֏",
	CURRENCY_ANG: "ƒ",
	CURRENCY_AOA: "Kz",
	CURRENCY_ARS: "$",
	CURRENCY_AUD: "$",
	CURRENCY_AWG: "ƒ",
	CURRENCY_BAM: "КМ",
	CURRENCY_BBD: "$",
	CURRENCY_BDT: "৳",
	CURRENCY_BGN: "лв",
	CURRENCY_BHD: ".د.ب",
	CURRENCY_BIF: "FBu",
	CURRENCY_BMD: "$",
	CURRENCY_BND: "$",
	CURRENCY_BOB: "Bs.",
	CURRENCY_BRL: "R$",
	CURRENCY_BSD: "$",
	CURRENCY_BTN: "Nu.",
	CURRENCY_BWP: "P",
	CURRENCY_BYN: "Br",
	CURRENCY_BZD: "BZ$",
	CURRENCY_CAD: "$",
	CURRENCY_CDF: "FC",
	CURRENCY_CHF: "CHF",
	CURRENCY_CLP: "$",
	CURRENCY_CNY: "¥",
	CURRENCY_COP: "$",
	CURRENCY_CRC: "₡",
	CURRENCY_CUP: "$",
	CURRENCY_CVE: "$",
	CURRENCY_CZK: "Kč",
	CURRENCY_DJF: "Fdj",
	CURRENCY_DKK: "kr",
	CURRENCY_DOP: "RD$",
	CURRENCY_DZD: "د.ج",
	CURRENCY_EGP: "£",
	CURRENCY_ERN: "Nfk",
	CURRENCY_ETB: "Br",
	CURRENCY_EUR: "€",
	CURRENCY_FJD: "$",
	CURRENCY_FKP: "£",
	CURRENCY_GBP: "£",
	CURRENCY_GEL: "₾",
	CURRENCY_GGP: "£",
	CURRENCY_GHS: "₵",
	CURRENCY_GIP: "£",
	CURRENCY_GMD: "D",
	CURRENCY_GNF: "FG",
	CURRENCY_GTQ: "Q",
	CURRENCY_GYD: "$",
	CURRENCY_HKD: "$",
	CURRENCY_HNL: "L",
	CURRENCY_HRK: "kn",
	CURRENCY_HTG: "G",
	CURRENCY_HUF: "Ft",
	CURRENCY_IDR: "Rp",
	CURRENCY_ILS: "₪",
	CURRENCY_IMP: "£",
	CURRENCY_INR: "₹",
	CURRENCY_IQD: "ع.د",
	CURRENCY_IRR: "﷼",
	CURRENCY_ISK: "kr",
	CURRENCY_JEP: "£",
	CURRENCY_JMD: "$",
	CURRENCY_JOD: "د.ا",
	CURRENCY_JPY: "¥",
	CURRENCY_KES: "KSh",
	CURRENCY_KGS: "сом",
	CURRENCY_KHR: "៛",
	CURRENCY_KID: "$",
	CURRENCY_KMF: "CF",
	CURRENCY_KRW: "₩",
	CURRENCY_KWD: "د.ك",
	CURRENCY_KYD: "$",
	CURRENCY_KZT: "〒",
	CURRENCY_LAK: "₭",
	CURRENCY_LBP: "ل.ل",
	CURRENCY_LKR: "රු",
	CURRENCY_LRD: "$",
	CURRENCY_LSL: "L",
	CURRENCY_LYD: "ل.د",
	CURRENCY_MAD: "د.م.",
	CURRENCY_MDL: "lei",
	CURRENCY_MGA: "Ar",
	CURRENCY_MKD: "ден",
	CURRENCY_MMK: "K",
	CURRENCY_MNT: "₮",
	CURRENCY_MOP: "P",
	CURRENCY_MRO: "UM",
	CURRENCY_MRU: "UM",
	CURRENCY_MUR: "₨",
	CURRENCY_MVR: "ރ.",
	CURRENCY_MWK: "MK",
	CURRENCY_MXN: "$",
	CURRENCY_MYR: "RM",
	CURRENCY_MZN: "MT",
	CURRENCY_NAD: "$",
	CURRENCY_NGN: "₦",
	CURRENCY_NIO: "C$",
	CURRENCY_NOK: "kr",
	CURRENCY_NPR: "₨",
	CURRENCY_NZD: "$",
	CURRENCY_OMR: "ر.ع.",
	CURRENCY_PAB: "B/.",
	CURRENCY_PEN: "S/.",
	CURRENCY_PGK: "K",
	CURRENCY_PHP: "₱",
	CURRENCY_PKR: "₨",
	CURRENCY_PLN: "zł",
	CURRENCY_PYG: "₲",
	CURRENCY_QAR: "ر.ق",
	CURRENCY_RON: "lei",
	CURRENCY_RSD: "РСД",
	CURRENCY_RUB: "₽",
	CURRENCY_RWF: "RF",
	CURRENCY_SAR: "ر.س",
	CURRENCY_SBD: "$",
	CURRENCY_SCR: "₨",
	CURRENCY_SDG: "ج.س.",
	CURRENCY_SEK: "kr",
	CURRENCY_SGD: "$",
	CURRENCY_SHP: "£",
	CURRENCY_SLL: "Le",
	CURRENCY_SOS: "S",
	CURRENCY_SRD: "$",
	CURRENCY_SSP: "£",
	CURRENCY_STN: "Db",
	CURRENCY_SYP: "£",
	CURRENCY_SZL: "E",
	CURRENCY_THB: "฿",
	CURRENCY_TJS: "ЅМ",
	CURRENCY_TMT: "m",
	CURRENCY_TND: "د.ت",
	CURRENCY_TOP: "T$",
	CURRENCY_TRY: "₺",
	CURRENCY_TTD: "$",
	CURRENCY_TWD: "NT$",
	CURRENCY_TZS: "TSh",
	CURRENCY_UAH: "₴",
	CURRENCY_UGX: "USh",
	CURRENCY_USD: "$",
	CURRENCY_UYU: "$",
	CURRENCY_UZS: "UZS",
	CURRENCY_VES: "Bs",
	CURRENCY_VND: "₫",
	CURRENCY_VUV: "Vt",
	CURRENCY_WST: "T",
	CURRENCY_XAF: "FCFA",
	CURRENCY_XCD: "$",
	CURRENCY_XDR: "SDR",
	CURRENCY_XOF: "CFA",
	CURRENCY_XPF: "₣",
	CURRENCY_YER: "﷼",
	CURRENCY_ZAR: "R",
	CURRENCY_ZMW: "ZK",
	CURRENCY_ZWL: "$",
}
