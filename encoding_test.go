package mail

import "testing"

// TestEncoding_String tests the string method of the Encoding object
func TestEncoding_String(t *testing.T) {
	tests := []struct {
		name string
		e    Encoding
		want string
	}{
		{"Encoding: Base64", EncodingB64, "base64"},
		{"Encoding: QP", EncodingQP, "quoted-printable"},
		{"Encoding: None/8bit", NoEncoding, "8bit"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.e.String() != tt.want {
				t.Errorf("wrong string for Encoding returned. Expected: %s, got: %s",
					tt.want, tt.e.String())
			}
		})
	}
}

// TestCharset_String tests the string method of the Charset object
func TestCharset_String(t *testing.T) {
	tests := []struct {
		name string
		c    Charset
		want string
	}{
		{"Charset: US-ASCII", CharsetASCII, "US-ASCII"},
		{"Charset: UTF-7", CharsetUTF7, "UTF-7"},
		{"Charset: UTF-8", CharsetUTF8, "UTF-8"},
		{"Charset: ISO-8859-1", CharsetISO88591, "ISO-8859-1"},
		{"Charset: ISO-8859-2", CharsetISO88592, "ISO-8859-2"},
		{"Charset: ISO-8859-3", CharsetISO88593, "ISO-8859-3"},
		{"Charset: ISO-8859-4", CharsetISO88594, "ISO-8859-4"},
		{"Charset: ISO-8859-5", CharsetISO88595, "ISO-8859-5"},
		{"Charset: ISO-8859-6", CharsetISO88596, "ISO-8859-6"},
		{"Charset: ISO-8859-7", CharsetISO88597, "ISO-8859-7"},
		{"Charset: ISO-8859-9", CharsetISO88599, "ISO-8859-9"},
		{"Charset: ISO-8859-13", CharsetISO885913, "ISO-8859-13"},
		{"Charset: ISO-8859-14", CharsetISO885914, "ISO-8859-14"},
		{"Charset: ISO-8859-15", CharsetISO885915, "ISO-8859-15"},
		{"Charset: ISO-8859-16", CharsetISO885916, "ISO-8859-16"},
		{"Charset: ISO-2022-JP", CharsetISO2022JP, "ISO-2022-JP"},
		{"Charset: ISO-2022-KR", CharsetISO2022KR, "ISO-2022-KR"},
		{"Charset: windows-1250", CharsetWindows1250, "windows-1250"},
		{"Charset: windows-1251", CharsetWindows1251, "windows-1251"},
		{"Charset: windows-1252", CharsetWindows1252, "windows-1252"},
		{"Charset: windows-1255", CharsetWindows1255, "windows-1255"},
		{"Charset: windows-1256", CharsetWindows1256, "windows-1256"},
		{"Charset: KOI8-R", CharsetKOI8R, "KOI8-R"},
		{"Charset: KOI8-U", CharsetKOI8U, "KOI8-U"},
		{"Charset: Big5", CharsetBig5, "Big5"},
		{"Charset: GB18030", CharsetGB18030, "GB18030"},
		{"Charset: GB2312", CharsetGB2312, "GB2312"},
		{"Charset: GBK", CharsetGBK, "GBK"},
		{"Charset: TIS-620", CharsetTIS620, "TIS-620"},
		{"Charset: EUC-KR", CharsetEUCKR, "EUC-KR"},
		{"Charset: Shift_JIS", CharsetShiftJIS, "Shift_JIS"},
		{"Charset: Unknown", CharsetUnknown, "Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c.String() != tt.want {
				t.Errorf("wrong string for Charset returned. Expected: %s, got: %s",
					tt.want, tt.c.String())
			}
		})
	}
}
