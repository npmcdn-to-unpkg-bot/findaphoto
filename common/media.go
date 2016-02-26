package common

import (
	"time"
)

var MediaIndexName = "media-index"

const MediaTypeName = "media"

type Media struct {
	Signature     string `json:"signature"`
	Filename      string `json:"filename"`
	Path          string `json:"path"`
	LengthInBytes int64  `json:"lengthinbytes"`

	MimeType        string  `json:"mimetype,omitempty"`
	Width           int     `json:"width,omitempty"`
	Height          int     `json:"height,omitempty"`
	DurationSeconds float32 `json:"durationseconds,omitempty"`

	// EXIF info
	ApertureValue   string `json:"aperture,omitempty"`
	ExposureProgram string `json:"exposureprogram,omitempty"`
	ExposureTime    string `json:"exposuretime,omitempty"`
	Flash           string `json:"flash,omitempty"`
	FNumber         string `json:"fnumber,omitempty"`
	FocalLength     string `json:"focallength,omitempty"`
	Iso             string `json:"iso,omitempty"`
	WhiteBalance    string `json:"whitebalance,omitempty"`
	LensInfo        string `json:"lensinfo,omitempty"`
	LensModel       string `json:"lensmodel,omitempty"`
	CameraMake      string `json:"cameramake,omitempty"`
	CameraModel     string `json:"cameramodel,omitempty"`

	// For arrays - see here for mappings & searching: http://stackoverflow.com/questions/26258292/querystring-search-on-array-elements-in-elastic-search
	Keywords []string `json:"keywords,omitempty"`

	// Location
	Location GeoPoint `json:"location,omitempty"`

	// Placename, from the reverse coding of the location
	LocationCountryName string `json:"countryname,omitempty"`
	LocationCountryCode string `json:"countrycode,omitempty"`
	LocationCityName    string `json:"cityname,omitempty"`
	LocationSiteName    string `json:"sitename,omitempty"`
	LocationPlaceName   string `json:"placename,omitempty"`

	// Date related fields
	DateTime  time.Time `json:"datetime"`  // 2009-06-15T13:45:30.0000000-07:00 'round trip pattern'
	Date      string    `json:"date"`      // yyyyMMdd - for aggregating by date
	DayName   string    `json:"dayname"`   // (Wed, Wednesday)
	MonthName string    `json:"monthname"` // (Apr, April)
}

type GeoPoint struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type CandidateFile struct {
	FullPath      string
	AliasedPath   string
	Signature     string
	LengthInBytes int64
	Exif          ExifOutput
}

type ExifOutput struct {
	SourceFile string
	File       ExifOutputFile
	EXIF       ExifOutputExif
	IPTC       ExifOutputIptc
	Quicktime  ExifOutputQuicktime
	Composite  ExifOutputComposite
}

type ExifOutputFile struct {
	MIMEType    string
	ImageHeight int
	ImageWidth  int
}

type ExifOutputExif struct {
	ApertureValue    float32
	CreateDate       string
	DateTimeOriginal string
	ModifyDate       string
	ExposureProgram  string
	ExposureTime     interface{} // Sigh - sometimes a number, sometimes a string - 1 is a number, while "1/200" is a string. Probably an exiftool'ism
	Flash            string
	FNumber          float32
	FocalLength      string
	GPSLatitudeRef   string
	GPSLatitude      string
	GPSLongitudeRef  string
	GPSLongitude     string
	ISO              interface{} // Most cameras use an int, some a string (!)
	LensInfo         string
	LensModel        string
	Make             string
	Model            string
	WhiteBalance     string
}

type ExifOutputQuicktime struct {
	ContentCreateDate string
	CreateDate        string
	ModifyDate        string
	ImageWidth        int
	ImageHeight       int
	Duration          string
}

type ExifOutputComposite struct {
	GPSPosition string
}

type ExifOutputIptc struct {
	Keywords interface{} // Some are []string - others are string. Exiftool seems to be the source
}