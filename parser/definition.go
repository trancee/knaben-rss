package parser

import "github.com/bzick/tokenizer"

const (
	TUnknown tokenizer.TokenKey = iota

	TAny

	TDot
	TApostrophe
	TParenthesis

	TAlias

	TSeason
	TEpisode

	TTags

	TLanguage

	TResolution
	TQuality

	TNetwork

	TSource

	TAudio
	TChannels

	TCompression

	TGroup
)

var Letters = []rune{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

var Alias = []string{
	"AKA", "aka",
}

var Any = []string{
	"EXTENDED.EDITION", "EXTENDED EDITION", "EXTENDED.CUT", "EXTENDED CUT", "Extended.Cut", "Extended Cut", "EXTENDED", "Extended",

	"ALTERNATIVE.CUT",
	"BW",
	"CHRONO",
	"COLORIZED",
	"CONVERT",
	"DC",
	"DIRFIX",
	// "DUBBED",
	// "DV",
	// "EXTENDED",
	"EXTRAS",
	"FINAL",
	"FS",
	// "HDR",
	// "HDR10Plus",

	"HYBRiD",
	"INTERNAL", "iNTERNAL",
	"LIMITED",
	// "LINE",
	"NFOFIX",
	"OAR",
	"OM",
	"PROOFFIX",
	"PROPER",
	"PURE",
	"RATED",
	"READNFO",
	"REAL",
	"RECUT",
	"REMASTERED",
	"REPACK",
	"RERIP",
	"RESTORED",
	"RETAIL",
	"RM4k", "RM4K", // Remastered to 4K

	"SAMPLEFIX",
	// "SDR",
	"SOURCE.SAMPLE",
	"SUBBED",
	"THEATRICAL",
	"UNCENSORED",
	"UNCUT",
	"UNRATED",
	"WS",
}

var Tags = []string{
	"CBR",  // Constant Bit-Rate
	"VBR",  // Variable Bit-Rate
	"DS4K", // 4K downscaled less than its original resolution
	// "ENG",  // Netflix
	// "ITA",
	"NL",
	"NORDiC",
	"LATINO",                  // Netflix
	"MULTi",                   // The release has a minimum of 2 audio languages
	"MULTiSUBS", "Multi-Subs", // The release has a minimum of 6 subtitle languages
	"PROPER", // A corrected version released by a different group
	"REPACK", // A corrected version released by the same group that issued the original release
	"RM4K",   // 4K remaster in 1080p
	"R1", "R2", "R3",
	"RA", "RB", "RC",
}

var Language = []string{
	"ENGLISH", "English", "ENG",
	"GERMAN", "German", "GER",
	"ITALIAN", "Italian", "ITA",
	"RUSSIAN", "Russian", "RUS",
}

var Resolution = []string{
	"2160p",
	"m-1080p",
	"1080p",
	"1080i",
	"m-720p",
	"720p",
	"480p",
	"480i",
	"mHD",
	"µHD",
	"PAL",
	"NTSC",
}
var Quality = []string{
	"Dolby Vision", "Dolby.Vision", "DolbyVision", "DV",
	"UHD",
	"HDR10Plus", "HDR10+", "HDR10", "HDR",
	"10-Bit", "10-bit", "10Bit", "10bit",
	"SDR",
	"8-Bit", "8-bit", "8Bit", "8bit",
}
var Source = []string{
	"Blu-Ray", "BluRay", "BLURAY", "BD25", "BD50", "BD66", "BD100", "BD5", "BD9",
	"COMPLETE.BLURAY", "BD-Disk",
	"BDMV", "BDISO",
	"BluRay-RiP", "BluRay-Rip", "BLURAYRIP",
	"BR-RiP", "BR-Rip", "BRRIP", "BRRip", "BRip",
	"BD-RiP", "BD-Rip", "BDRIP", "BDRip", "BD-R", "BDR",
	"DVD-9", "DVD-5", "Full-Rip", "DVD-Full", "DVDMux",
	"DVD",
	"DVD-RiP", "DVD-Rip", "DVDRIP", "DVDRip", "DVDR",
	"DDC",
	"WEBDL", "WEB-DL", "WEB DL", "WEB-DLRip",
	"VOD-RiP", "VOD-Rip", "VODRIP", "VODRip", "VOD",
	"WEB-RiP", "WEB-Rip", "WEBRIP", "WEBRip", "WEB Rip", "WebHD", "WEB", "Web",
	"HDRip", "HDRiP",
	"WEB-Cap", "WEBCAP", "WEB Cap",
	"HC", "HD-Rip",
	"VODRip", "VODR",
	"DS-RiP", "DS-Rip", "DSRip", "DSR", "DS",
	"DVB-RiP", "DVB-Rip", "DVBRip", "DVB",
	"SAT-RiP", "SAT-Rip", "SATRip", "SAT",
	"HDTV-RiP", "HDTV-Rip", "HDTVRip", "HDTV",
	"DTVRip", "DTV",
	"PDTV",
	"DTHRip",
	"TV-RiP", "TV-Rip", "TVRip", "TV",
	"WEBSCREENER", "BDSCR", "DVDSCREENER", "DVDSCR", "SCREENER", "SCR",
	"PPVRip", "PPV",
	"PDVD", "PreDVDRip",
	"VHS-RiP", "VHS-Rip", "VHS",
	"TELECINE", "HDTC", "TC",
	"WORKPRINT", "WP",
	"TELESYNC", "HDTS", "TS",
	"CAM-RiP", "CAM-Rip", "HDCAM", "CAM",
	"Remux",
	"iTunes",
}
var Compression = []string{
	"x265",
	"X265",
	"h265",
	"H265",
	"H.265",
	"HEVC",
	"x264",
	"X264",
	"h264",
	"H264",
	"H.264",
	"AVC",
	"AV1",
	"DivX",
	"XviD",
	"Xvid",
}

var Audio = []string{
	"Dolby True-HD", "True-HD", "Dolby TrueHD", "TrueHD",
	"Dolby Atmos", "Atmos",
	"Dolby Digital Plus", "DolbyDigitalPlus", "DD Plus", "DDPlus", "DDP", "DD+",
	"Dolby Digital", "DolbyDigital", "DD",
	"E-AC-3", "E-AC3", "EAC3", "EC-3", "EC3",
	"AC-3", "AC3",
	"AC3D",
	"DTS-HD MA", "DTS-HD HR", "DTS-HD", "DTS-ES", "DTS:X", "DTS",
	"AAC",
	"LPCM", "PCM",
	"FLAC",
	// "MP4",
	"MP3", "MP2",
	"LD",
	"LiNE.DUBBED", "LINE.DUBBED", "Line.Dubbed",
	"LiNE", "LINE", "Line",
	"MD",
	"MiC", "MIC", "Mic",
	"MiC.DUBBED", "MIC.DUBBED", "Mic.Dubbed",
}
var Channels = []string{
	"7.1.4", "7.1", "7.0",
	"6.1", "6.0", "6CH",
	"5.2", "5.1.4", "5.1.2", "5.1", "5.0",
	"4.2", "4.1", "4.0",
	"3.1.2", "3.1", "3.0",
	"2.1", "2.0",
	"1.0",
	// "7.1.4", "7 1 4", "7.1", "7 1", "7.0", "7 0",
	// "6.1", "6 1", "6.0", "6 0", "6CH",
	// "5.2", "5 2", "5.1.4", "5 1 4", "5.1.2", "5 1 2", "5.1", "5 1",
	// "4.2", "4 2", "4.1", "4 1",
	// "3.1.2", "3 1 2", "3.1", "3 1",
	// "2.1", "2 1", "2.0", "2 0",
	// "1.0", "1 0",
}

var Network = []string{
	"ABC",
	"ATVP",
	"AMZN",
	"BBC",
	"BOOM",
	"CBC",
	"CBS",
	"CC",
	"CR",
	"CRAV",
	"CRITERION",
	"CW",
	"DCU",
	"DSNP",
	"DSNY",
	"FBWatch",
	"FREE",
	"FOX",
	"HMAX",
	"HULU",
	"HTSR",
	"iP",
	"iT",
	"LIFE",
	"MTV",
	"MUBI",
	"NBC",
	"NICK",
	"NF",
	"OAR",
	"PCOK",
	"PMTP",
	"RED",
	"TVNZ",
	"STAN",
	"STZ",
}
