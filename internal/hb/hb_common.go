package hb

// #include <stdlib.h>
// #include <hb.h>
import "C"
import (
	"unsafe"
)

// Codepoint holds Unicode codepoints. Also used to hold glyph IDs.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-codepoint-t
type Codepoint = uint32

// Direction is the direction of a text segment or buffer.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-direction-t
type Direction C.hb_direction_t

const (
	DirectionInvalid Direction = C.HB_DIRECTION_INVALID // Initial, unset direction.
	DirectionLTR     Direction = C.HB_DIRECTION_LTR     // Text is set horizontally from left to right.
	DirectionRTL     Direction = C.HB_DIRECTION_RTL     // Text is set horizontally from right to left.
	DirectionTTB     Direction = C.HB_DIRECTION_TTB     // Text is set vertically from top to bottom.
	DirectionBTT     Direction = C.HB_DIRECTION_BTT     // Text is set vertically from bottom to top.
)

// Learn is a data type for scripts. Each Script's value is a Tag corresponding
// to the four-letter values defined by ISO 15924.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-script-t
type Script C.hb_script_t

const (
	ScriptCommon                Script = C.HB_SCRIPT_COMMON                 // Zyyy
	ScriptInherited             Script = C.HB_SCRIPT_INHERITED              // Zinh
	ScriptUnknown               Script = C.HB_SCRIPT_UNKNOWN                // Zzzz
	ScriptArabic                Script = C.HB_SCRIPT_ARABIC                 // Arab
	ScriptArmenia               Script = C.HB_SCRIPT_ARMENIAN               // Armn
	ScriptBengali               Script = C.HB_SCRIPT_BENGALI                // Beng
	ScriptCyrillic              Script = C.HB_SCRIPT_CYRILLIC               // Cyrl
	ScriptDevanagari            Script = C.HB_SCRIPT_DEVANAGARI             // Deva
	ScriptGeorgian              Script = C.HB_SCRIPT_GEORGIAN               // Geor
	ScriptGreek                 Script = C.HB_SCRIPT_GREEK                  // Grek
	ScriptGujarati              Script = C.HB_SCRIPT_GUJARATI               // Gujr
	ScriptGurmukhi              Script = C.HB_SCRIPT_GURMUKHI               // Guru
	ScriptHangul                Script = C.HB_SCRIPT_HANGUL                 // Hang
	ScriptHan                   Script = C.HB_SCRIPT_HAN                    // Hani
	ScriptHebrew                Script = C.HB_SCRIPT_HEBREW                 // Hebr
	ScriptHiragana              Script = C.HB_SCRIPT_HIRAGANA               // Hira
	ScriptKannada               Script = C.HB_SCRIPT_KANNADA                // Knda
	ScriptKatakana              Script = C.HB_SCRIPT_KATAKANA               // Kana
	ScriptLao                   Script = C.HB_SCRIPT_LAO                    // Laoo
	ScriptLatin                 Script = C.HB_SCRIPT_LATIN                  // Latn
	ScriptMalayalam             Script = C.HB_SCRIPT_MALAYALAM              // Mlym
	ScriptOriya                 Script = C.HB_SCRIPT_ORIYA                  // Orya
	ScriptTamil                 Script = C.HB_SCRIPT_TAMIL                  // Taml
	ScriptTelugu                Script = C.HB_SCRIPT_TELUGU                 // Telu
	ScriptThai                  Script = C.HB_SCRIPT_THAI                   // Thai
	ScriptTibetan               Script = C.HB_SCRIPT_TIBETAN                // Tibt
	ScriptBopomofo              Script = C.HB_SCRIPT_BOPOMOFO               // Bopo
	ScriptBraille               Script = C.HB_SCRIPT_BRAILLE                // Brai
	ScriptCanadianSyllabics     Script = C.HB_SCRIPT_CANADIAN_SYLLABICS     // Cans
	ScriptCherokee              Script = C.HB_SCRIPT_CHEROKEE               // Cher
	ScriptEthiopic              Script = C.HB_SCRIPT_ETHIOPIC               // Ethi
	ScriptKhmer                 Script = C.HB_SCRIPT_KHMER                  // Khmr
	ScriptMongolian             Script = C.HB_SCRIPT_MONGOLIAN              // Mong
	ScriptMyanmar               Script = C.HB_SCRIPT_MYANMAR                // Mymr
	ScriptOgham                 Script = C.HB_SCRIPT_OGHAM                  // Ogam
	ScriptRunic                 Script = C.HB_SCRIPT_RUNIC                  // Runr
	ScriptSinhala               Script = C.HB_SCRIPT_SINHALA                // Sinh
	ScriptSyriac                Script = C.HB_SCRIPT_SYRIAC                 // Syrc
	ScriptThaana                Script = C.HB_SCRIPT_THAANA                 // Thaa
	ScriptYi                    Script = C.HB_SCRIPT_YI                     // Yiii
	ScriptDeseret               Script = C.HB_SCRIPT_DESERET                // Dsrt
	ScriptGothic                Script = C.HB_SCRIPT_GOTHIC                 // Goth
	ScriptOldItalic             Script = C.HB_SCRIPT_OLD_ITALIC             // Ital
	ScriptBuhid                 Script = C.HB_SCRIPT_BUHID                  // Buhd
	ScriptHanunoo               Script = C.HB_SCRIPT_HANUNOO                // Hano
	ScriptTagalog               Script = C.HB_SCRIPT_TAGALOG                // Tglg
	ScriptTagbanwa              Script = C.HB_SCRIPT_TAGBANWA               // Tagb
	ScriptCypriot               Script = C.HB_SCRIPT_CYPRIOT                // Cprt
	ScriptLimbu                 Script = C.HB_SCRIPT_LIMBU                  // Limb
	ScriptLinearB               Script = C.HB_SCRIPT_LINEAR_B               // Linb
	ScriptOsmanya               Script = C.HB_SCRIPT_OSMANYA                // Osma
	ScriptShavian               Script = C.HB_SCRIPT_SHAVIAN                // Shaw
	ScriptTaiLe                 Script = C.HB_SCRIPT_TAI_LE                 // Tale
	ScriptUgaritic              Script = C.HB_SCRIPT_UGARITIC               // Ugar
	ScriptBuginese              Script = C.HB_SCRIPT_BUGINESE               // Bugi
	ScriptCoptic                Script = C.HB_SCRIPT_COPTIC                 // Copt
	ScriptGlagolitic            Script = C.HB_SCRIPT_GLAGOLITIC             // Glag
	ScriptKharoshthi            Script = C.HB_SCRIPT_KHAROSHTHI             // Khar
	ScriptNewTaiLue             Script = C.HB_SCRIPT_NEW_TAI_LUE            // Talu
	ScriptOldPersian            Script = C.HB_SCRIPT_OLD_PERSIAN            // Xpeo
	ScriptSylotiNagri           Script = C.HB_SCRIPT_SYLOTI_NAGRI           // Sylo
	ScriptTifinagh              Script = C.HB_SCRIPT_TIFINAGH               // Tfng
	ScriptBalinese              Script = C.HB_SCRIPT_BALINESE               // Bali
	ScriptCuneiform             Script = C.HB_SCRIPT_CUNEIFORM              // Xsux
	ScriptNko                   Script = C.HB_SCRIPT_NKO                    // Nkoo
	ScriptPhagsPa               Script = C.HB_SCRIPT_PHAGS_PA               // Phag
	ScriptPhoenician            Script = C.HB_SCRIPT_PHOENICIAN             // Phnx
	ScriptCarian                Script = C.HB_SCRIPT_CARIAN                 // Cari
	ScriptCham                  Script = C.HB_SCRIPT_CHAM                   // Cham
	ScriptKayahLi               Script = C.HB_SCRIPT_KAYAH_LI               // Kali
	ScriptLepcha                Script = C.HB_SCRIPT_LEPCHA                 // Lepc
	ScriptLycian                Script = C.HB_SCRIPT_LYCIAN                 // Lyci
	ScriptLydian                Script = C.HB_SCRIPT_LYDIAN                 // Lydi
	ScriptOlChiki               Script = C.HB_SCRIPT_OL_CHIKI               // Olck
	ScriptRejang                Script = C.HB_SCRIPT_REJANG                 // Rjng
	ScriptSaurashtra            Script = C.HB_SCRIPT_SAURASHTRA             // Saur
	ScriptSundanese             Script = C.HB_SCRIPT_SUNDANESE              // Sund
	ScriptVai                   Script = C.HB_SCRIPT_VAI                    // Vaii
	ScriptAvestan               Script = C.HB_SCRIPT_AVESTAN                // Avst
	ScriptBamum                 Script = C.HB_SCRIPT_BAMUM                  // Bamu
	ScriptEgyptianHieroglyphs   Script = C.HB_SCRIPT_EGYPTIAN_HIEROGLYPHS   // Egyp
	ScriptImperialAramaic       Script = C.HB_SCRIPT_IMPERIAL_ARAMAIC       // Armi
	ScriptInscriptionalPahlavi  Script = C.HB_SCRIPT_INSCRIPTIONAL_PAHLAVI  // Phli
	ScriptInscriptionalParthian Script = C.HB_SCRIPT_INSCRIPTIONAL_PARTHIAN // Prti
	ScriptJavanese              Script = C.HB_SCRIPT_JAVANESE               // Java
	ScriptKaithi                Script = C.HB_SCRIPT_KAITHI                 // Kthi
	ScriptLisu                  Script = C.HB_SCRIPT_LISU                   // Lisu
	ScriptMeeteiMayek           Script = C.HB_SCRIPT_MEETEI_MAYEK           // Mtei
	ScriptOldSouthArabian       Script = C.HB_SCRIPT_OLD_SOUTH_ARABIAN      // Sarb
	ScriptOldTurkic             Script = C.HB_SCRIPT_OLD_TURKIC             // Orkh
	ScriptSamaritan             Script = C.HB_SCRIPT_SAMARITAN              // Samr
	ScriptTaiTham               Script = C.HB_SCRIPT_TAI_THAM               // Lana
	ScriptTaiViet               Script = C.HB_SCRIPT_TAI_VIET               // Tavt
	ScriptBatak                 Script = C.HB_SCRIPT_BATAK                  // Batk
	ScriptBrahmi                Script = C.HB_SCRIPT_BRAHMI                 // Brah
	ScriptMandaic               Script = C.HB_SCRIPT_MANDAIC                // Mand
	ScriptChakma                Script = C.HB_SCRIPT_CHAKMA                 // Cakm
	ScriptMeroiticCursive       Script = C.HB_SCRIPT_MEROITIC_CURSIVE       // Merc
	ScriptMeroiticHieroglyphs   Script = C.HB_SCRIPT_MEROITIC_HIEROGLYPHS   // Mero
	ScriptMiao                  Script = C.HB_SCRIPT_MIAO                   // Plrd
	ScriptSharada               Script = C.HB_SCRIPT_SHARADA                // Shrd
	ScriptSoraSompeng           Script = C.HB_SCRIPT_SORA_SOMPENG           // Sora
	ScriptTakri                 Script = C.HB_SCRIPT_TAKRI                  // Takr
	ScriptBassaVah              Script = C.HB_SCRIPT_BASSA_VAH              // Bass
	ScriptCaucasianAlbanian     Script = C.HB_SCRIPT_CAUCASIAN_ALBANIAN     // Aghb
	ScriptDuployan              Script = C.HB_SCRIPT_DUPLOYAN               // Dupl
	ScriptElbasan               Script = C.HB_SCRIPT_ELBASAN                // Elba
	ScriptGrantha               Script = C.HB_SCRIPT_GRANTHA                // Gran
	ScriptKhojki                Script = C.HB_SCRIPT_KHOJKI                 // Khoj
	ScriptKhudawadi             Script = C.HB_SCRIPT_KHUDAWADI              // Sind
	ScriptLinearA               Script = C.HB_SCRIPT_LINEAR_A               // Lina
	ScriptMahajani              Script = C.HB_SCRIPT_MAHAJANI               // Mahj
	ScriptManichaean            Script = C.HB_SCRIPT_MANICHAEAN             // Mani
	ScriptMendeKikakui          Script = C.HB_SCRIPT_MENDE_KIKAKUI          // Mend
	ScriptModi                  Script = C.HB_SCRIPT_MODI                   // Modi
	ScriptMro                   Script = C.HB_SCRIPT_MRO                    // Mroo
	ScriptNabataean             Script = C.HB_SCRIPT_NABATAEAN              // Nbat
	ScriptOldNorthArabian       Script = C.HB_SCRIPT_OLD_NORTH_ARABIAN      // Narb
	ScriptOldPermic             Script = C.HB_SCRIPT_OLD_PERMIC             // Perm
	ScriptPahawhHmong           Script = C.HB_SCRIPT_PAHAWH_HMONG           // Hmng
	ScriptPalmyrene             Script = C.HB_SCRIPT_PALMYRENE              // Palm
	ScriptPauCinHau             Script = C.HB_SCRIPT_PAU_CIN_HAU            // Pauc
	ScriptPsalterPahlavi        Script = C.HB_SCRIPT_PSALTER_PAHLAVI        // Phlp
	ScriptSiddham               Script = C.HB_SCRIPT_SIDDHAM                // Sidd
	ScriptTirhuta               Script = C.HB_SCRIPT_TIRHUTA                // Tirh
	ScriptWarangCiti            Script = C.HB_SCRIPT_WARANG_CITI            // Wara
	ScriptAhom                  Script = C.HB_SCRIPT_AHOM                   // Ahom
	ScriptAnatolianHieroglyphs  Script = C.HB_SCRIPT_ANATOLIAN_HIEROGLYPHS  // Hluw
	ScriptHatran                Script = C.HB_SCRIPT_HATRAN                 // Hatr
	ScriptMultani               Script = C.HB_SCRIPT_MULTANI                // Mult
	ScriptOldHungarian          Script = C.HB_SCRIPT_OLD_HUNGARIAN          // Hung
	ScriptSignwriting           Script = C.HB_SCRIPT_SIGNWRITING            // Sgnw
	ScriptAdlam                 Script = C.HB_SCRIPT_ADLAM                  // Adlm
	ScriptBhaiksuki             Script = C.HB_SCRIPT_BHAIKSUKI              // Bhks
	ScriptMarchen               Script = C.HB_SCRIPT_MARCHEN                // Marc
	ScriptOsage                 Script = C.HB_SCRIPT_OSAGE                  // Osge
	ScriptTangut                Script = C.HB_SCRIPT_TANGUT                 // Tang
	ScriptNewa                  Script = C.HB_SCRIPT_NEWA                   // Newa
	ScriptMasaramGondi          Script = C.HB_SCRIPT_MASARAM_GONDI          // Gonm
	ScriptNushu                 Script = C.HB_SCRIPT_NUSHU                  // Nshu
	ScriptSoyombo               Script = C.HB_SCRIPT_SOYOMBO                // Soyo
	ScriptZanabazarQquare       Script = C.HB_SCRIPT_ZANABAZAR_SQUARE       // Zanb
	ScriptDogra                 Script = C.HB_SCRIPT_DOGRA                  // Dogr
	ScriptGunjalaGondi          Script = C.HB_SCRIPT_GUNJALA_GONDI          // Gong
	ScriptHanifiRohingya        Script = C.HB_SCRIPT_HANIFI_ROHINGYA        // Rohg
	ScriptMakasar               Script = C.HB_SCRIPT_MAKASAR                // Maka
	ScriptMedefaidrin           Script = C.HB_SCRIPT_MEDEFAIDRIN            // Medf
	ScriptOldSogdian            Script = C.HB_SCRIPT_OLD_SOGDIAN            // Sogo
	ScriptSogdian               Script = C.HB_SCRIPT_SOGDIAN                // Sogd
	ScriptElymaic               Script = C.HB_SCRIPT_ELYMAIC                // Elym
	ScriptNandinagari           Script = C.HB_SCRIPT_NANDINAGARI            // Nand
	ScriptNyiakengPuachueHmong  Script = C.HB_SCRIPT_NYIAKENG_PUACHUE_HMONG // Hmnp
	ScriptWancho                Script = C.HB_SCRIPT_WANCHO                 // Wcho
	ScriptChorasmian            Script = C.HB_SCRIPT_CHORASMIAN             // Chrs
	ScriptDivesAkuru            Script = C.HB_SCRIPT_DIVES_AKURU            // Diak
	ScriptKhitanSmallScript     Script = C.HB_SCRIPT_KHITAN_SMALL_SCRIPT    // Kits
	ScriptYezidi                Script = C.HB_SCRIPT_YEZIDI                 // Yezi
	ScriptCyproMinoan           Script = C.HB_SCRIPT_CYPRO_MINOAN           // Cpmn
	ScriptOldUyghur             Script = C.HB_SCRIPT_OLD_UYGHUR             // Ougr
	ScriptTangsa                Script = C.HB_SCRIPT_TANGSA                 // Tnsa
	ScriptToto                  Script = C.HB_SCRIPT_TOTO                   // Toto
	ScriptVithkuqi              Script = C.HB_SCRIPT_VITHKUQI               // Vith
	ScriptMath                  Script = C.HB_SCRIPT_MATH                   // Zmth
	ScriptKawi                  Script = C.HB_SCRIPT_KAWI                   // Kawi
	ScriptNagMundari            Script = C.HB_SCRIPT_NAG_MUNDARI            // Nagm
	ScriptInvalid               Script = C.HB_SCRIPT_INVALID                // No script set
)

// UserDataKey is a data structure for holding user-data keys.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-user-data-key-t
type UserDataKey C.hb_user_data_key_t

// Language is a data type for languages. Each Language corresponds to a BCP 47
// language tag.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-language-t
type Language C.hb_language_t

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-feature-t
type Feature struct {
	Tag   Tag
	Value uint32
	Start uint32
	End   uint32
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-variation-t
type Variation struct {
	Tag   Tag
	Value float32
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-tag-t
type Tag [4]byte

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-tag-from-string
func TagFromString(str string) Tag {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	tag := C.hb_tag_from_string(cStr, -1)
	return *(*Tag)(unsafe.Pointer(&tag))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-tag-to-string
func TagToString(tag Tag) string {
	var buf [4]byte
	C.hb_tag_to_string(*(*C.hb_tag_t)(unsafe.Pointer(&tag)), (*C.char)(unsafe.Pointer(&buf)))

	return string(buf[:])
}

// DirectionFromString converts a string to a Direction.
//
// Matching is loose and applies only to the first letter. For examples, "LTR"
// and "left-to-right" will both return DirectionLTR. Unmatched strings will
// return DirectionInvalid
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-direction-from-string
func DirectionFromString(str string) Direction {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	return Direction(C.hb_direction_from_string(cStr, -1))
}

// DirectionToString converts a Direction to a string.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-direction-to-string
func DirectionToString(direction Direction) string {
	return C.GoString(C.hb_direction_to_string(C.hb_direction_t(direction)))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-script-from-iso15924-tag
func ScriptFromISO15924Tag(tag Tag) Script {
	return Script(C.hb_script_from_iso15924_tag(*(*C.hb_tag_t)(unsafe.Pointer(&tag))))
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-script-to-iso15924-tag
func ScriptToISO15924Tag(script Script) Tag {
	tag := C.hb_script_to_iso15924_tag(C.hb_script_t(script))
	return *(*Tag)(unsafe.Pointer(&tag))
}

// ScriptFromString converts a string str representing an ISO 15924 script tag to
// a corresponding Script. Shorthand for TagFromString then ScriptFromISO15924Tag.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-script-from-string
func ScriptFromString(str string) Script {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	return Script(C.hb_script_from_string(cStr, -1))
}

// ScriptGetHorizontalDirection fetches the Direction of a script when it is set
// horizontally. All right-to-left scripts will return DirectionRTL. All
// left-to-right scripts will return DirectionLTR. Scripts that can be written
// either horizontally or vertically will return DirectionInvalid. Unknown scripts
// will return DirectionLTR.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-script-get-horizontal-direction
func ScriptGetHorizontalDirection(script Script) Direction {
	return Direction(C.hb_script_get_horizontal_direction(C.hb_script_t(script)))
}

// LanguageFromString converts str representing a BCP 47 language tag to the
// corresponding Language.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-language-from-string
func LanguageFromString(str string) Language {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	return Language(C.hb_language_from_string(cStr, -1))
}

// LanguageToString converts a Language to a string.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-language-to-string
func LanguageToString(language Language) string {
	return C.GoString(C.hb_language_to_string(language))
}

// LanguageGetDefault fetches the default language from current locale.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-language-get-default
func LanguageGetDefault() Language {
	return Language(C.hb_language_get_default())
}

// LanguageMatches checks whether a second language tag is the same or a more
// specific version of the provided language tag. For example, "fa_IR.utf8" is a
// more specific tag for "fa" or for "fa_IR".
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-language-matches
func LanguageMatches(language, specific Language) bool {
	return C.hb_language_matches(language, specific) == 1
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-feature-from-string
func FeatureFromString(str string) (feature Feature, ok bool) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	ok = C.hb_feature_from_string(cStr, -1, (*C.hb_feature_t)(unsafe.Pointer(&feature))) == 1
	return
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-feature-to-string
func FeatureToString(feature *Feature) string {
	buf := make([]byte, 128)
	C.hb_feature_to_string((*C.hb_feature_t)(unsafe.Pointer(feature)), (*C.char)(unsafe.Pointer(&buf[0])), 128)
	return string(buf)
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-variation-from-string
func VariationFromString(str string) (variation Variation, ok bool) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	ok = C.hb_variation_from_string(cStr, -1, (*C.hb_variation_t)(unsafe.Pointer(&variation))) == 1
	return
}

// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-variation-to-string
func VariationToString(variation *Variation) string {
	buf := make([]byte, 128)
	C.hb_variation_to_string((*C.hb_variation_t)(unsafe.Pointer(variation)), (*C.char)(unsafe.Pointer(&buf[0])), 128)
	return string(buf)
}

// DestroyFunc is a method type for destroying user-data callbacks.
//
// Learn more: https://harfbuzz.github.io/harfbuzz-hb-common.html#hb-destroy-func-t
type DestroyFunc C.hb_destroy_func_t
