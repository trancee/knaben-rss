package entity_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"rss.knaben.eu/entity"
	"rss.knaben.eu/parser"
)

var releases = map[string]*entity.Release{
	// https://www.xrel.to/tv-nfo/3058807/The-West-Wing-S06E21-German-1080p-WebHD-h264-FKKTV.html
	`The.West.Wing.S06E21.German.1080p.WebHD.h264-FKKTV`:                         {Title: "The West Wing", Season: "S06", Episode: "E21", Language: []string{"German"}, Source: []string{"WebHD"}, Resolution: []string{"1080p"}, Compression: []string{"h264"}, Group: []string{"FKKTV"}},
	`Thor.and.Loki.Blood.Brothers.2011.iNTERNAL.COMPLETE.BLURAY-REFRACTiON`:      {Title: "Thor and Loki Blood Brothers", Year: 2011, Source: []string{"COMPLETE.BLURAY"}, Group: []string{"REFRACTiON"}},
	`Percy.Jackson.and.the.Olympians.S01E02.2160p.Dolby.Vision.And.HDR10.DV.MP4`: {Title: "Percy Jackson and the Olympians", Season: "S01", Episode: "E02", Resolution: []string{"2160p"}, Quality: []string{"Dolby Vision", "HDR10", "DV"}},

	`So.Help.Me.Todd.S02E05.2160p.WEB.H265-SuccessfulCrab[TGx]`:                       {Title: "So Help Me Todd", Season: "S02", Episode: "E05", Resolution: []string{"2160p"}, Source: []string{"WEB"}, Compression: []string{"H265"}, Group: []string{"SuccessfulCrab"}},
	`BMF.S03E07.2160p.WEB.H265-LAZYCUNTS[TGx]`:                                        {Title: "BMF", Season: "S03", Episode: "E07", Resolution: []string{"2160p"}, Source: []string{"WEB"}, Compression: []string{"H265"}, Group: []string{"LAZYCUNTS"}},
	`The.Gentleman.In.Moscow.S01E02.2160p.AMZN.WEB-DL.DDP5.1.HEVC-NTb ...`:            {Title: "The Gentleman In Moscow", Season: "S01", Episode: "E02", Resolution: []string{"2160p"}, Network: []string{"AMZN"}, Source: []string{"WEB-DL"}, Audio: []string{"DDP"}, Channels: []string{"5.1"}, Compression: []string{"HEVC"}, Group: []string{"NTb"}},
	`Loot.S02E02.HDR.2160p.WEB.h265-ETHEL[TGx]`:                                       {Title: "Loot", Season: "S02", Episode: "E02", Resolution: []string{"2160p"}, Source: []string{"WEB"}, Quality: []string{"HDR"}, Compression: []string{"h265"}, Group: []string{"ETHEL"}},
	`All.The.Light.We.Cannot.See.S01.2160p.NF.WEB-DL.HDR.ENG.LATINO.H265-BEN.THE.MEN`: {Title: "All The Light We Cannot See", Season: "S01", Tags: []string{"LATINO"}, Language: []string{"ENG"}, Resolution: []string{"2160p"}, Quality: []string{"HDR"}, Network: []string{"NF"}, Source: []string{"WEB-DL"}, Compression: []string{"H265"}, Group: []string{"BEN THE MEN"}},
	`Fire.Country.S02E03.See.You.Next.Apocalypse.2160p.PMTP.WEB-DL.DD ...`:            {Title: "Fire Country", Season: "S02", Episode: "E03", EpisodeTitle: "See You Next Apocalypse", Resolution: []string{"2160p"}, Network: []string{"PMTP"}, Source: []string{"WEB-DL"}, Audio: []string{"DD"}},
	`Halo.S02E05.2160p.DV.HDR.DDP5.1.Atmos.x265.MP4-BEN.THE.MEN`:                      {Title: "Halo", Season: "S02", Episode: "E05", Resolution: []string{"2160p"}, Quality: []string{"DV", "HDR"}, Audio: []string{"DDP", "Atmos"}, Channels: []string{"5.1"}, Compression: []string{"x265"}, Group: []string{"BEN THE MEN"}},
	`Halo.S02E05.Aleria.1080p.AMZN.WEB-DL.DDP5.1.Atmos.H.264-GP-TV-NLsubs`:            {Title: "Halo", Season: "S02", Episode: "E05", EpisodeTitle: "Aleria", Resolution: []string{"1080p"}, Network: []string{"AMZN"}, Source: []string{"WEB-DL", "TV"}, Audio: []string{"DDP", "Atmos"}, Channels: []string{"5.1"}, Compression: []string{"H.264"}, Group: []string{"NLsubs"}},
	`South Park S23 2160p AMZN WEB-DL DD+ 5.1 HEVC-MZABI`:                             {Title: "South Park", Season: "S23", Resolution: []string{"2160p"}, Network: []string{"AMZN"}, Source: []string{"WEB-DL"}, Audio: []string{"DD+"}, Channels: []string{"5.1"}, Compression: []string{"HEVC"}, Group: []string{"MZABI"}},
	`Our.Flag.Means.Death.S02E04.DV.HDR.2160p.WEB.H265-SuccessfulCrab ...`:            {Title: "Our Flag Means Death", Season: "S02", Episode: "E04", Resolution: []string{"2160p"}, Quality: []string{"DV", "HDR"}, Source: []string{"WEB"}, Compression: []string{"H265"}, Group: []string{"SuccessfulCrab"}},
	`Ahsoka.S01.COMPLETE.2160p.Dolby.Vision.Profile.Five`:                             {Title: "Ahsoka", Season: "S01", EpisodeTitle: "COMPLETE", Resolution: []string{"2160p"}, Quality: []string{"Dolby Vision"}},
	`Foundation.S01.S01.COMPLETE.2160p.10bit.HDR.WEBRip.6CH.x265.HEVC-PSA`:            {Title: "Foundation", Season: "S01", EpisodeTitle: "S01 COMPLETE", Resolution: []string{"2160p"}, Quality: []string{"10bit", "HDR"}, Source: []string{"WEBRip"}, Channels: []string{"6CH"}, Compression: []string{"x265", "HEVC"}, Group: []string{"PSA"}},
	`Deadloch 2023 S01 S01 Web 2160p HDR10+ x265 5.1 AAC English SRT`:                 {Title: "Deadloch", Year: 2023, Season: "S01", EpisodeTitle: "S01", Language: []string{"English"}, Resolution: []string{"2160p"}, Quality: []string{"HDR10+"}, Source: []string{"Web"}, Audio: []string{"AAC"}, Channels: []string{"5.1"}, Compression: []string{"x265"}},
	`BMF.S03E06.2160p.WEB.H265-LAZYCUNTS[TGx]`:                                        {Title: "BMF", Season: "S03", Episode: "E06", Resolution: []string{"2160p"}, Source: []string{"WEB"}, Compression: []string{"H265"}, Group: []string{"LAZYCUNTS"}},

	`Movie Title 2017 HC 720p HDRiP DD5.1 x264-LEGi0N`:                                                                 {Title: "Movie Title", Year: 2017, Resolution: []string{"720p"}, Source: []string{"HC", "HDRiP"}, Compression: []string{"x264"}, Audio: []string{"DD"}, Channels: []string{"5.1"}, Group: []string{"LEGi0N"}},
	`L'hypothèse.du.movie.volé.AKA.The.Hypothesis.of.the.Movie.Title.1978.1080p.CINET.WEB-DL.AAC2.0.x264-Cinefeel.mkv`: {Title: "L'hypothèse du movie volé AKA The Hypothesis of the Movie Title", Year: 1978, Resolution: []string{"1080p"}, Source: []string{"WEB-DL"}, Audio: []string{"AAC"}, Channels: []string{"2.0"}, Compression: []string{"x264"}, Group: []string{"Cinefeel mkv"}},
	`Das.A.Team.Der.Film.Extended.Cut.German.720p.BluRay.x264-ANCIENT`:                                                 {Title: "Das A Team Der Film", Language: []string{"German"}, Resolution: []string{"720p"}, Source: []string{"BluRay"}, Compression: []string{"x264"}, Group: []string{"ANCIENT"}},
	`Die.fantastische.Reise.des.Dr.Dolittle.2020.German.DL.LD.1080p.WEBRip.x264-PRD`:                                   {Title: "Die fantastische Reise des Dr Dolittle", Year: 2020, Language: []string{"German"}, Resolution: []string{"1080p"}, Source: []string{"WEBRip"}, Audio: []string{"LD"}, Compression: []string{"x264"}, Group: []string{"PRD"}},
	`The.Good.German.2006.GERMAN.720p.HDTV.x264-RLsGrp`:                                                                {Title: "The Good German", Year: 2006, Language: []string{"GERMAN"}, Resolution: []string{"720p"}, Source: []string{"HDTV"}, Compression: []string{"x264"}, Group: []string{"RLsGrp"}},
	`Kick.Movie.2.2013.German.DTS.DL.720p.BluRay.x264-Pate`:                                                            {Title: "Kick Movie 2", Year: 2013, Language: []string{"German"}, Resolution: []string{"720p"}, Source: []string{"BluRay"}, Audio: []string{"DTS"}, Compression: []string{"x264"}, Group: []string{"Pate"}},
	`Drop.Movie.1994.German.AC3D.DL.720p.BluRay.x264-KLASSiGERHD`:                                                      {Title: "Drop Movie", Year: 1994, Language: []string{"German"}, Resolution: []string{"720p"}, Source: []string{"BluRay"}, Audio: []string{"AC3D"}, Compression: []string{"x264"}, Group: []string{"KLASSiGERHD"}},
	`Movie.Aufbruch.nach.Pandora.Extended.2009.German.DTS.720p.BluRay.x264-SoW`:                                        {Title: "Movie Aufbruch nach Pandora", Year: 2009, Language: []string{"German"}, Resolution: []string{"720p"}, Source: []string{"BluRay"}, Audio: []string{"DTS"}, Compression: []string{"x264"}, Group: []string{"SoW"}},
	`World.Movie.Z.2.EXTENDED.2013.German.DL.1080p.BluRay.AVC-XANOR`:                                                   {Title: "World Movie Z 2", Year: 2013, Language: []string{"German"}, Resolution: []string{"1080p"}, Source: []string{"BluRay"}, Compression: []string{"AVC"}, Group: []string{"XANOR"}},
	`1776.1979.EXTENDED.720p.BluRay.X264-AMIABLE`:                                                                      {Title: "1776", Year: 1979, Resolution: []string{"720p"}, Source: []string{"BluRay"}, Compression: []string{"X264"}, Group: []string{"AMIABLE"}},

	`Evil Dead II 1987 REPACK 2160p HYBRiD ITA UHD BluRay DV HDR FLAC 2 0 HEVC-RAIMI`: {Title: "Evil Dead II", Year: 1987, Language: []string{"ITA"}, Resolution: []string{"2160p"}, Quality: []string{"UHD", "DV", "HDR"}, Source: []string{"BluRay"}, Audio: []string{"FLAC"}, Channels: []string{"2.0"}, Compression: []string{"HEVC"}, Group: []string{"RAIMI"}},

	`Penn and Teller Fool Us S10E19 720p x265-T0PAZ`:                                                   {Title: "Penn and Teller Fool Us", Season: "S10", Episode: "E19", Resolution: []string{"720p"}, Compression: []string{"x265"}, Group: []string{"T0PAZ"}},
	`ITV Unwind: Sheep in the Crags 1080p HDTV x265-MVGroup`:                                           {Title: "ITV Unwind: Sheep in the Crags", Resolution: []string{"1080p"}, Source: []string{"HDTV"}, Compression: []string{"x265"}, Group: []string{"MVGroup"}},
	`Special Ops Lioness S01 720p x265-T0PAZ`:                                                          {Title: "Special Ops Lioness", Season: "S01", Resolution: []string{"720p"}, Compression: []string{"x265"}, Group: []string{"T0PAZ"}},
	`Dateline NBC 2024 04 19 Evil Walked Through the Door 1080p HEVC x265-MeGusta`:                     {Title: "Dateline", Year: 2024, EpisodeTitle: "Evil Walked Through the Door", Resolution: []string{"1080p"}, Network: []string{"NBC"}, Compression: []string{"HEVC", "x265"}, Group: []string{"MeGusta"}},
	`S W A T 2017 S07E09 1080p x265-ELiTE`:                                                             {Title: "S W A T", Year: 2017, Season: "S07", Episode: "E09", Resolution: []string{"1080p"}, Compression: []string{"x265"}, Group: []string{"ELiTE"}},
	`Dune Part Two 2024 English 1080p DS4K WEBRip iTunes x265 HEVC 10bit DDP 5.1 Atmos ESub M3GAN-MCX`: {Title: "Dune Part Two", Year: 2024, Tags: []string{"DS4K"}, Language: []string{"English"}, Resolution: []string{"1080p"}, Quality: []string{"10bit"}, Source: []string{"WEBRip", "iTunes"}, Audio: []string{"DDP", "Atmos"}, Channels: []string{"5.1"}, Compression: []string{"x265", "HEVC"}, Group: []string{"MCX"}},

	`Late Night with the Devil (2023) 2160p WEBRip 5.1 10Bit x265 -YTS`:                 {Title: "Late Night with the Devil", Year: 2023, Resolution: []string{"2160p"}, Quality: []string{"10Bit"}, Source: []string{"WEBRip"}, Channels: []string{"5.1"}, Compression: []string{"x265"}, Group: []string{"YTS"}},
	`Rebel Moon Part Two The Scargiver 2024 1080p WEBRip x265-DH`:                       {Title: "Rebel Moon Part Two The Scargiver", Year: 2024, Resolution: []string{"1080p"}, Source: []string{"WEBRip"}, Compression: []string{"x265"}, Group: []string{"DH"}},
	`Rebel.Moon.Part.Two.The.Scargiver.2024.1080p.WEBRip.10Bit.DDP5.1.x265-Asiimov`:     {Title: "Rebel Moon Part Two The Scargiver", Year: 2024, Resolution: []string{"1080p"}, Quality: []string{"10Bit"}, Source: []string{"WEBRip"}, Audio: []string{"DDP"}, Channels: []string{"5.1"}, Compression: []string{"x265"}, Group: []string{"Asiimov"}},
	`Rebel Moon   Part Two: The Scargiver 2024 1080p WEBRip 10Bit DDP 5.1 x265-Asiimov`: {Title: "Rebel Moon Part Two: The Scargiver", Year: 2024, Resolution: []string{"1080p"}, Quality: []string{"10Bit"}, Source: []string{"WEBRip"}, Audio: []string{"DDP"}, Channels: []string{"5.1"}, Compression: []string{"x265"}, Group: []string{"Asiimov"}},
	`Death Walks on High Heels 1971 aka La morte cammina con i tacchi alti Arrow 1080p BluRay x265 HEVC 10bit AAC 1 0 Dual Commentary Luciano Ercoli Frank Wolff Nieves Navarro Simon Andreu George Rigaud Jose Manuel Martin Luciano Rossi Claudie Lange-hq`: {Title: "Death Walks on High Heels", Year: 1971, Resolution: []string{"1080p"}, Quality: []string{"10bit"}, Source: []string{"BluRay"}, Audio: []string{"AAC"}, Channels: []string{"1.0"}, Compression: []string{"x265", "HEVC"}, Group: []string{"hq"}},
	`Night of the Demons 2 1994 SF 1080p BluRay x265 HEVC 10bit AAC 2 0 Commentary Brian Trenchard Smith Amelia Kinkade Cristi Harris Darin Heames Robert Jayne Merle Kennedy Rod McCary Johnny Moran Rick Peters Jennifer Rhodes Christine Taylor-hq`:        {Title: "Night of the Demons 2", Year: 1994, Resolution: []string{"1080p"}, Quality: []string{"10bit"}, Source: []string{"BluRay"}, Audio: []string{"AAC"}, Channels: []string{"2.0"}, Compression: []string{"x265", "HEVC"}, Group: []string{"hq"}},
	`Yuva 2024 Kannada 1080p DS4K WEBRip AMZN x265 HEVC 10bit DDP 5.1 ESub M3GAN-MCX`: {Title: "Yuva", Year: 2024, Tags: []string{"DS4K"}, Resolution: []string{"1080p"}, Quality: []string{"10bit"}, Network: []string{"AMZN"}, Source: []string{"WEBRip"}, Audio: []string{"DDP"}, Channels: []string{"5.1"}, Compression: []string{"x265", "HEVC"}, Group: []string{"MCX"}},
	// `Les Heroines Du Mal 1979 1080p BluRay DUAL DD 2 0 X265-S`:                        {Title: "XXX"},
	// `The Abyss 1989 Ext DC RM4k 1080p BluRay x265 HEVC 10bit AAC 7 1 James Cameron Ed Harris Mary Elizabeth Mastrantonio Michael Biehn Leo Burmester Todd Graff John Bedford Lloyd Kimberly Scott Chris Elliott JC Quinn Pierce Oliver Brewer Dick Warlock-hq`: {Title: "XXX"},
	// `Booksmart (2019) 2160p WEBRip 5.1 10Bit x265 -YTS`:                      {Title: "XXX"},
	// `The Game Plan 2007 1080p BluRay DD+ 5.1 x265-edge2020`:                  {Title: "XXX"},
	// `The Greatest Showman 2017 1080p BluRay DD+ 7.1 x265-edge2020`:           {Title: "XXX"},
	// `The Grinch 2018 1080p BluRay DD+ 7.1 x265-edge2020`:                     {Title: "XXX"},
	// `The Great Outdoors 1988 1080p BluRay DD+ 5.1 x265-edge2020`:             {Title: "XXX"},
	// `Guardians of the Galaxy Vol. 3 2023 1080p BluRay DD+ 7.1 x265-edge2020`: {Title: "XXX"},
	// `Gunpowder Milkshake 2021 1080p BluRay DD+ 5.1 x265-edge2020`:            {Title: "XXX"},
	// `The Good Dinosaur 2015 1080p BluRay DD+ 7.1 x265-edge2020`:              {Title: "XXX"},
	// `The Godfather 1972 1080p BluRay DD+ 5.1 x265-edge2020`:                  {Title: "XXX"},
	// `The Godfather Part II 1974 1080p BluRay DD+ 5.1 x265-edge2020`:          {Title: "XXX"},
	// `The Godfather Part III 1990 1080p BluRay DD+ 5.1 x265-edge2020`:         {Title: "XXX"},
	// `The Goonies 1985 1080p BluRay DD+ 5.1 x265-edge2020`:                    {Title: "XXX"},
	// `The Great Gatsby 2013 1080p BluRay DD+ 5.1 x265-edge2020`:               {Title: "XXX"},
	// `The Great Mouse Detective 1986 1080p BluRay DD+ 5.1 x265-edge2020`:      {Title: "XXX"},
	// `The Green Mile 1999 1080p BluRay DD+ 5.1 x265-edge2020`:                 {Title: "XXX"},
	// `Grumpy Old Men 1993 1080p BluRay DD+ 2.0 x265-edge2020`:                 {Title: "XXX"},
	// `Grudge Match 2013 1080p BluRay DD+ 5.1 x265-edge2020`:                   {Title: "XXX"},
	// `Grown Ups 2 2013 1080p BluRay DD+ 5.1 x265-edge2020`:                    {Title: "XXX"},
	// `Grown Ups 2010 1080p BluRay DD+ 5.1 x265-edge2020`:                      {Title: "XXX"},
	// `Groundhog Day 1993 1080p BluRay DD+ 7.1 x265-edge2020`:                  {Title: "XXX"},
	// `Green Lantern: First Flight 2009 1080p BluRay DD+ 5.1 x265-edge2020`:    {Title: "XXX"},
	// `Grumpier Old Men 1995 1080p BluRay DD+ 5.1 x265-edge2020`:               {Title: "XXX"},
	// `Guardians of the Galaxy Vol. 2 2017 1080p BluRay DD+ 7.1 x265-edge2020`: {Title: "XXX"},
	// `Guardians of the Galaxy 2014 1080p BluRay DD+ 7.1 x265-edge2020`:        {Title: "XXX"},
	// `Hawk the Slayer 1980 RiffTrax 1080p BluRay x265 HEVC 10bit AAC 2 0 Terry Marcel John Jack Palance Bernard Bresslaw W Morgan Sheppard Patricia Quinn Cheryl Campbell Annette Crosbie Catriona MacColl Harry Andrews Roy Kinnear Patrick-Magee`: {Title: "XXX"},
	// `Greyhound 2020 1080p WEBRip DD+ 5.1 Atmos x265-edge2020`:                               {Title: "XXX"},
	// `Problemista (2023) 1080p WEBRip 5.1 10Bit x265 -YTS`:                                   {Title: "XXX"},
	// `Pleasantville (1998 ITA/ENG) [1080p x265] [Paso77]`:                                    {Title: "XXX"},
	// `Rebel Moon (Parte dos): La guerrera que deja marcas (2024) [WEB-DL 1080p X265 bi`:      {Title: "XXX"},
	// `Rebel Moon Part Two The Scargiver 2024 1080p NF WEBRip DDP 5.1 x265 10bit-GalaxyRG265`: {Title: "XXX"},
	// `Late Night with the Devil (2023) 1080p WEBRip 5.1 10Bit x265 -YTS`:                     {Title: "XXX"},
	// `Sweet Dreams (2024) 2160p WEBRip 5.1 10Bit x265 -YTS`:                                  {Title: "XXX"},
	// `Rebel Moon Part Two The Scargiver 2024 1080p WEBRip 10Bit DDP 5.1 x265 Asiimov-mkv`:    {Title: "XXX"},
	// `Rebel.Moon.Part.Two.The.Scargiver.2024.1080p.WEBRip.10Bit.DDP5.1.x265 Asiimov.mkv`:     {Title: "XXX"},
	// `The Two Popes (2019) 2160p WEBRip 5.1 10Bit x265 -YTS`:                                 {Title: "XXX"},
	// `Wild Things 1998 Unrated Remastered 1080p BluRay HEVC x265 5.1 BONE`:                   {Title: "XXX"},
	// `Green Lantern 2011 Extended 1080p BluRay DD+ 5.1 x265-edge2020`:                        {Title: "XXX"},
	// `Green Lantern: Emerald Knights 2011 1080p BluRay DD+ 5.1 x265-edge2020`:                {Title: "XXX"},
	// `Gotta Kick It Up! 2002 1080p UPSCALED AAC 2.0 x265-edge2020`:                           {Title: "XXX"},
	// `Grease 1978 1080p BluRay DD+ 5.1 x265-edge2020`:                                        {Title: "XXX"},
	// `Goosebumps 2015 1080p BluRay DD+ 7.1 x265-edge2020`:                                    {Title: "XXX"},
	// `Goosebumps 2: Haunted Halloween 2018 1080p BluRay DD+ 5.1 x265-edge2020`:               {Title: "XXX"},
	// `GoodFellas 1990 1080p BluRay DD+ 5.1 x265-edge2020`:                                    {Title: "XXX"},
	// `Gran Torino 2008 1080p BluRay DD+ 5.1 x265-edge2020`:                                   {Title: "XXX"},
	// `You Were Never Really Here (2017) 2160p WEBRip 5.1 10Bit x265 -YTS`:                    {Title: "XXX"},
	// `Star Trek First Contact (1996) 2160p BRRip 5.1 10Bit x265 -YTS`:                        {Title: "XXX"},
	// `The Legend of Bagger Vance (2000 ITA/ENG) [1080p x265] [Paso77]`:                       {Title: "XXX"},
	// `Immaculate (2024) 2160p WEBRip 5.1 10Bit x265 -YTS`:                                    {Title: "XXX"},
	// `Carmen from Kawachi 1966 RM4K 1080p BluRay x265 HEVC FLAC SARTRE Kawachi-Karumen`:      {Title: "XXX"},
	// `Carmen from Kawachi (1966) RM4K 1080p BluRay x265 HEVC FLAC SARTRE [Kawachi Karumen]`:  {Title: "XXX"},
	// `Rest in Pieces 1987 aka Descanse en piezas RM4k VS 1080p BluRay x265 HEVC 10bit AAC 2 0 Commentary Jose Ramon Larraz Scott Thompson Baker Lorin Jean Vail Dorothy Malone Jack Taylor Patty Shepard Jeffrey Segal Fernando Bilbao Carole James-hq`: {Title: "XXX"},
	// `Klaus & Barroso (2024) 1080p WEBRip 5.1 10Bit x265 -YTS`:                                {Title: "XXX"},
	// `Ghostbusters Frozen Empire (2024) [1080p] [WEBRip] [x265] [10bit]`:                      {Title: "XXX"},
	// `Sumotherhood 2023 1080p WEB-DL HEVC x265 5.1 BONE`:                                      {Title: "XXX"},
	// `The Inseparables 2023 1080p WEB-DL HEVC x265 5.1 BONE`:                                  {Title: "XXX"},
	// `Whos That Girl 1987 1080p BluRay HEVC x265 5.1 BONE`:                                    {Title: "XXX"},
	// `Los que se quedan (2023) [BDRip 1080p X265 10bits]`:                                     {Title: "XXX"},
	// `The Lorax (2012) 2160p WEBRip 5.1 10Bit x265 -YTS`:                                      {Title: "XXX"},
	// `Ghostbusters Frozen Empire (2024) 1080p WEBRip 10Bit x265 -YTS`:                         {Title: "XXX"},
	// `The Dead Pool (1988 ITA/ENG) [1080p x265] [Paso77]`:                                     {Title: "XXX"},
	// `In The Eye of the Storm The Political Odyssey of Yanis Varoufakis 2024 720p-x265`:       {Title: "XXX"},
	// `In The Eye of the Storm   The Political Odyssey of Yanis Varoufakis (2024) [720p x265]`: {Title: "XXX"},
	// `Immaculate (2024) (1080p AMZN WEB-DL x265 HEVC 10bit EAC3 5.1 Ghost) [QxR]`:             {Title: "XXX"},
	// `La revancha de los novatos (1984) [BDRip 1080p][x265]`:                                  {Title: "XXX"},
	// `Dune Part Two (2024) 2160p WEBRip 5.1 10Bit x265 -YTS`:                                  {Title: "XXX"},
	// `The Hidden Web (2023) 2160p WEBRip 5.1 10Bit x265 -YTS`:                                 {Title: "XXX"},
	// `The Masque of the Red Death 1964 Extended RM4k SC 1080p BluRay x265 HEVC 10bit AAC 2 0 Commentary V2 Roger Corman Vincent Price Hazel Court Jane Asher David Weston Nigel Green Patrick Magee Paul Whitsun Jones Robert Brown Edgar Allan-Poe`: {Title: "XXX"},
	// `Dune Part 2 2024 1080p WEB-DL HEVC x265 10 Bit DDP5.1 Subs KINGDOM RG`:                             {Title: "XXX"},
	// `Road House (2024) [1080p] [WEBRip] [x265] [10bit] [5.1]`:                                           {Title: "XXX"},
	// `Wild Things (1998) 2160p BRRip 5.1 10Bit x265 -YTS`:                                                {Title: "XXX"},
	// `Dune Part Two (2024) 1080p WEBRip 5.1 10Bit x265 -YTS`:                                             {Title: "XXX"},
	// `We Wish You a Married Christmas 2022 1080p AMZN WEBRip DDP 5.1 x265 10bit-GalaxyRG265`:             {Title: "XXX"},
	// `About Schmidt (2002) (1080p BluRay x265 HEVC 10bit AAC 5.1 Tigole)`:                                {Title: "XXX"},
	// `Sleeping Dogs (2024) [1080p] [WEBRip] [x265] [10bit] [5.1]`:                                        {Title: "XXX"},
	// `El camino: Una pel&iacute;cula de Breaking bad (2019) [BDRip 1080p 10 bit x265]`:                   {Title: "XXX"},
	// `El camino: Una película de Breaking bad (2019) [BDRip 1080p 10 bit x265]`:                          {Title: "XXX"},
	// `Immaculate (2024) [1080p] [WEBRip] [x265] [10bit]`:                                                 {Title: "XXX"},
	// `The Tall Man (2012 ITA/ENG) [1080p x265] [Paso77]`:                                                 {Title: "XXX"},
	// `The French Kissers 2009 1080p BluRay x265 HEVC 10bit AAC 5 1 French Tigole-QxR`:                    {Title: "XXX"},
	// `About Schmidt (2002) (1080p BluRay x265 HEVC 10bit AAC 5.1 Tigole) [QxR]`:                          {Title: "XXX"},
	// `Thanksgiving (2023) (2160p BluRay x265 HEVC 10bit HDR AAC 5.1 Tigole) [QxR]`:                       {Title: "XXX"},
	// `Immaculate (2024) 1080p WEBRip 5.1 10Bit x265 -YTS`:                                                {Title: "XXX"},
	// `Jimmy Carr Natural Born Killer 2024 1080p WEB-DL HEVC x265 5.1 BONE`:                               {Title: "XXX"},
	// `Spring Breakers (2012) 2160p WEBRip 5.1 10Bit x265 -YTS`:                                           {Title: "XXX"},
	// `Sharknado 2 The Second One 2014 RiffTrax Live 720p 10bit WEBRip x265-budgetbits`:                   {Title: "XXX"},
	// `Interrogation AKA Przesluchanie 1989 EN subs 720p 10bit BluRay x265-budgetbits`:                    {Title: "XXX"},
	// `Infernal Affairs (2002) 2160p BRRip 5.1 10Bit x265 -YTS`:                                           {Title: "XXX"},
	// `Dune Part Two 2024 Hybrid 2160p iT WEBRip DV HDR10 DDP 5.1 Atmos x265 PrimeX-ProtonMovies`:         {Title: "XXX"},
	// `Dune: Part Two 2024 Hybrid 2160p iT WEBRip DD+ 5.1 Atmos DV HDR10+ x265-PrimeX`:                    {Title: "XXX"},
	// `Cesar 1936 Criterion 1080p BluRay x265 HEVC FLAC SARTRE The Marseille Trilogy-FIXED`:               {Title: "XXX"},
	// `Criterion 1080p BluRay x265 HEVC FLAC SARTRE [The Marseille Trilogy] [FIXED]`:                      {Title: "XXX"},
	// `Cesar (1936) Criterion 1080p BluRay x265 HEVC FLAC SARTRE [The Marseille Trilogy] [FIXED]`:         {Title: "XXX"},
	// `Sudden Impact (1983 ITA/ENG) [1080p x265] [Paso77]`:                                                {Title: "XXX"},
	// `Dune Part Two 2024 1080p AMZN WEB-DL x265 HEVC 10bit EAC3 5.1 Silence-QxR`:                         {Title: "XXX"},
	// `Dune   Part Two (2024) (1080p AMZN WEB-DL x265 HEVC 10bit EAC3 5.1 Silence) [QxR]`:                 {Title: "XXX"},
	// `Dune: Parte Dos (2024) [WEB-DL 1080p X265 10bits]`:                                                 {Title: "XXX"},
	// `Premalu 2024 Malayalam 1080p DS4K WEBRip HS x265 HEVC 10bit DDP 5.1 Atmos ESub M3GAN-MCX`:          {Title: "XXX"},
	// `All That Sex (2023) 2160p WEBRip 5.1 10Bit x265 -YTS`:                                              {Title: "XXX"},
	// `Premalu (2024) Malayalam (1080p DS4K WEBRip HS x265 HEVC 10bit DDP5.1 Atmos ESub   M3GAN)   [MCX]`: {Title: "XXX"},
	// `Breathe 2024 1080p WEB-DL HECV x265 10 Bit DDP5.1 Subs KINGDOM RG`:                                 {Title: "XXX"},
	// `Boomerang (1992 ITA/ENG) [1080p x265] [Paso77]`:                                                    {Title: "XXX"},
	// `Dune.Part.Two.2024.1080p.WEBRip.10Bit.DDP5.1.x265-Asiimov`:                                         {Title: "XXX"},
	// `Kung Fu Panda 4 (2024) [1080p] [WEBRip] [x265] [10bit] [5.1]`:                                      {Title: "XXX"},
	// `The Enforcer (1976 ITA/ENG) [1080p x265] [Paso77]`:                                                 {Title: "XXX"},
	// `The Taste of Things 2023 1080p BluRay x265 HEVC 10bit AAC 5 1 French Tigole-QxR`:                   {Title: "XXX"},
	// `Lisa Frankenstein (2024) (1080p BluRay x265 HEVC 10bit AAC 5.1 Tigole) [QxR]`:                      {Title: "XXX"},
	// `The Taste of Things (2023) (1080p BluRay x265 HEVC 10bit AAC 5.1 French Tigole) [QxR]`:             {Title: "XXX"},
	// `The Beekeeper (2024) (2160p BluRay x265 HEVC 10bit HDR AAC 5.1 Tigole) [QxR]`:                      {Title: "XXX"},
	// `Amar Singh Chamkila (2024) 1080p WEBRip 5.1 10Bit x265 -YTS`:                                       {Title: "XXX"},
	// `To Die For (1995) 2160p BRRip 5.1 10Bit x265 -YTS`:                                                 {Title: "XXX"},
	// `Good Morning, Vietnam 1987 1080p BluRay DD+ 5.1 x265-edge2020`:                                     {Title: "XXX"},
	// `Good Luck Charlie, It's Christmas! 2011 1080p WEBRip DD+ 5.1 x265-edge2020`:                        {Title: "XXX"},
	// `Good Burger 1997 1080p BluRay DD+ 5.1 x265-edge2020`:                                               {Title: "XXX"},
	// `Good Boys 2019 1080p BluRay DD+ 5.1 x265-edge2020`:                                                 {Title: "XXX"},
	// `Good Will Hunting 1997 1080p BluRay DD+ 5.1 x265-edge2020`:                                         {Title: "XXX"},
	// `Legend of the Lost Locket 2024 1080p WEB-DL HEVC x265 5.1 BONE`:                                    {Title: "XXX"},
	// `Supergirl 1984 DC 1080p BluRay HEVC x265 BONE`:                                                     {Title: "XXX"},
	// `The Way We Were (1973 ITA/ENG) [1080p x265] [Paso77]`:                                              {Title: "XXX"},
	// `The Unseen Crisis - Vaccine Stories You Were Never Told 2023 720p x265`:                            {Title: "XXX"},
	// `Damaged (2024) 2160p WEBRip 5.1 10Bit x265 -YTS`:                                                   {Title: "XXX"},
	// `GoldenEye 1995 1080p BluRay DD+ 5.1 x265-edge2020`:                                                 {Title: "XXX"},
	// `The Unseen Crisis   Vaccine Stories You Were Never Told 2023 720p x265`:                            {Title: "XXX"},
	// `Goldfinger 1964 1080p BluRay DD+ 5.1 x265-edge2020`:                                                {Title: "XXX"},
	// `The Hills Have Eyes 1977 Ext RM4k Arrow 1080p BluRay x265 HEVC 10bit AAC 1 0 Commentary Wes Craven Susan Lanier Robert Houston Martin Speer Dee Wallace John Steadman James Whitworth Virginia Vincent Lance Gordon Michael Berryman Janus Blythe 70s-hq`: {Title: "XXX"},
	// `Castle of Blood 1964 aka Danza macabra Uncut RM4k Severin 1080p BluRay x265 HEVC 10bit AAC 2 0 Dual Commentary Antonio Margheriti Sergio Corbucci Barbara Steele Georges Riviere Margrete Robsahm Arturo Dominici Silvano Tranquilli Edgar Allan-Poe`:     {Title: "XXX"},
	// `Willow (1988) 2160p WEBRip 5.1 10Bit x265 -YTS`:                                                                {Title: "XXX"},
	// `Fuks 2 (2024) 1080p WEBRip 5.1 10Bit x265 -YTS`:                                                                {Title: "XXX"},
	// `Cesar 1936 Criterion 1080p BluRay x265 HEVC FLAC SARTRE The Marseille-Trilogy`:                                 {Title: "XXX"},
	// `Dayo 2024 1080p Tagalog WEB-DL HEVC x265 5.1 BONE`:                                                             {Title: "XXX"},
	// `Father Frost AKA Jack Frost AKA Morozko 1964 RiffTrax MST3K quadruple audio 720p 10bit BluRay x265-budgetbits`: {Title: "XXX"},
	// `Riddle of Fire (2023) 1080p WEBRip 5.1 10Bit x265 -YTS`:                                                        {Title: "XXX"},
	// `Pacific Rim Uprising (2018) 2160p BRRip 5.1 10Bit x265 -YTS`:                                                   {Title: "XXX"},
	// `Confidential Assignment 2 International 2022 1080p BluRay x265 HEVC 10bit EAC3 5.1 Korean REX-PxL`:             {Title: "XXX"},
	// `The.Creator.2023.1080p.UHD.BluRay.DV.HDR10.x265.DD.5.1 SM737 [ProtonMovies]`:                                   {Title: "XXX"},
	// `Magnum Force (1973 ITA/ENG) [1080p x265] [Paso77]`:                                                             {Title: "XXX"},
	// `T2 Trainspotting (2017) 2160p BRRip 5.1 10Bit x265 -YTS`:                                                       {Title: "XXX"},
	// `Baghead 2023 1080p WEB-DL HECV x265 10 Bit DDP 5.1 Subs KINGDOM-RG`:                                            {Title: "XXX"},
	// `The Greatest Hits (2024) 2160p WEBRip 5.1 10Bit x265 -YTS`:                                                     {Title: "XXX"},
	// `Baghead 2023 1080p WEB-DL HEVC x265 10 Bit DDP 5.1 Subs KINGDOM-RG`:                                            {Title: "XXX"},
	// `Frankenstein Created Woman 1967 SF 1080p BluRay x265 HEVC 10bit AAC 2 0 Commentary HeVK Terence Fisher Peter Cushing Susan Denberg Thorley Walters Robert Morris Duncan Lamont Blythe Barry Warren Derek Fowlds Alan MacNaughtan Madden Philip Ray hq-60s`: {Title: "XXX"},

	//// CHECK
	// `Forbidden Tales (Joone, Digital Playground) 2001 DVDRip`: {Title: "Forbidden Tales (Joone, Digital Playground)", Year: 2001, Source: []string{"DVDRip"}},
	// `Der.Film.deines.Lebens.German.2011.PAL.DVDR-ETM`:                                                                  {Title: "Der Film deines Lebens", Year: 2011, Resolution: []string{"PAL"}, Source: []string{"DVDR"}},
	// `(500).Days.Of.Movie.(2009).DTS.1080p.BluRay.x264.NLsubs`:                                                          {Title: "(500) Days Of Movie", Year: 2009, Resolution: []string{"1080p"}, Source: []string{"BluRay"}, Compression: []string{"x264"}, Audio: []string{"DTS"}},
	// `The.Man.from.U.N.C.L.E.2015.1080p.BluRay.x264-SPARKS`:                                                             {Title: "The Man from U.N.C.L.E.", Year: 2015, Resolution: []string{"1080p"}, Source: []string{"BluRay"}, Compression: []string{"x264"}, Group: []string{"SPARKS"}},
}

func TestReleases(t *testing.T) {
	parser := parser.NewParser()

	for input, want := range releases {
		release := parser.Parse(input)

		if !cmp.Equal(release, want /*, cmpopts.IgnoreFields(entity.Release{}, "titles", "episodeTitles")*/) {
			t.Errorf("%s\n%s\n%s\n%s\n\n", input, want.JSON(), release.JSON(), release.Render())
		}
	}
}
