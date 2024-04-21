package parser

import (
	"regexp"
	"strings"

	"github.com/bzick/tokenizer"
	"golang.org/x/exp/maps"

	"rss.knaben.eu/entity"
)

// https://scenerules.org/

// https://scenerules.org/html/2020_X265.html

// Named directory arguments formatted inside <> must be included.
// Optional arguments formatted inside [] can be used in some cases.

// Feature.Title.<YEAR>.<TAGS>.[LANGUAGE].<RESOLUTION>.<FORMAT>.<x264|x265>-GROUP
// Weekly.TV.Show.[COUNTRY_CODE].[YEAR].SXXEXX[Episode.Part].[Episode.Title].<TAGS>.[LANGUAGE].<RESOLUTION>.<FORMAT>.<x264|x265>-GROUP
// Weekly.TV.Show.Special.SXXE00.Special.Title.<TAGS>.[LANGUAGE].<RESOLUTION>.<FORMAT>.<x264|x265>-GROUP
// Multiple.Episode.TV.Show.SXXEXX-EXX[Episode.Part].[Episode.Title].<TAGS>.[LANGUAGE].<RESOLUTION>.<FORMAT>.<x264|x265>-GROUP
// Cross.Over.TV.Show.One.SXXEXX[Episode.Part].[Episode.Title]_Show.Two.SXXEXX[Episode.Part].[Episode.Title].<TAGS>.[LANGUAGE].<RESOLUTION>.<FORMAT>.<x264|x265>-GROUP
// Miniseries.Show.PartX.[Episode.Title].<TAGS>.[LANGUAGE].<RESOLUTION>.<FORMAT>.<x264|x265>-GROUP

// https://scenerules.org/html/2014_BLURAY.html

// Movie.Name.YEAR.<PROPER/READ.NFO/REPACK>.<MULTi/SUBS/REGiON>.COMPLETE.BLURAY-GROUP
// TV.Show.SxxDxx.<PROPER/READ.NFO/REPACK>.<MULTi/SUBS/REGiON>.COMPLETE.BLURAY-GROUP
// Music.YEAR.<PROPER/READ.NFO/REPACK>.<MULTi/SUBS/REGiON>.COMPLETE.MBLURAY-GROUP

var rGroup = regexp.MustCompile(`-?([\w\s]*)(\[([\w\s]*)\])?`)
var rSeasonEpisode = regexp.MustCompile(`(S\d{2})(E\d{2})?`)

func (p *Parser) parse() *entity.Release {
	if p.stream == nil {
		return nil
	}

	titles := []string{}
	episodeTitles := []string{}

	isTitle := true
	isEpisodeTitle := false

	release := &entity.Release{}

	for p.stream.IsValid() {
		func() {
			defer p.stream.GoNext()

			token := p.stream.CurrentToken()
			// fmt.Println(token)
			value := strings.ReplaceAll(token.ValueString(), "_", " ")
			if strings.ContainsAny(value, "0123456789") || strings.HasSuffix(value, "BLURAY") {
				value = strings.ReplaceAll(value, " ", ".")
			}
			value = strings.TrimSpace(value)
			// fmt.Println(value)

			if token.Key() > 0 && token.Key() != TLanguage && token.Key() != TAudio && token.Key() != TNetwork {
				isTitle = false
			}

			if token.Is(tokenizer.TokenInteger) {
				value := token.ValueInt64()

				if value > 1900 && value < 2100 {
					release.Year = int64(value)

					isTitle = false

					if p.stream.NextToken().Is(tokenizer.TokenInteger) {
						isEpisodeTitle = true // Only assume we need Episode Title if we have Date format (YYYY MM DD)
					}

					return
				}
			}
			// if number, err := strconv.Atoi(value); err == nil {
			// 	if number > 1900 && number < 2100 {
			// 		release.Year = number

			// 		isTitle = false

			// 		return
			// 	}
			// }

			if token.Is(tokenizer.TokenFloat) {
				panic("FLOAT:" + value)
			}
			if token.Is(tokenizer.TokenString) && p.stream.NextToken().Is(tokenizer.TokenUndef) {
				if value, found := strings.CutPrefix(value, "-"); found {
					m := rGroup.FindStringSubmatch(strings.ReplaceAll(value, "|", " "))
					// fmt.Println(m)
					release.Group = append(release.Group, strings.TrimSpace(m[1]))
				}
			}

			if isTitle {
				if r := rSeasonEpisode.FindStringSubmatch(value); r != nil {
					// fmt.Println(value, r)
					release.Season = r[1]
					release.Episode = r[2]

					isEpisodeTitle = true
				} else {
					titles = append(titles, value)

					return
				}
			} else {
				switch token.Key() {
				case TTags:
					{
						// fmt.Println("TTags", token)
						release.Tags = append(release.Tags, value)
					}

				case TLanguage:
					{
						// fmt.Println("TLanguage", token)
						release.Language = append(release.Language, value)
					}

				case TResolution:
					{
						// fmt.Println("TResolution", token)
						release.Resolution = append(release.Resolution, value)

						isEpisodeTitle = false
					}

				case TNetwork:
					{
						// fmt.Println("TNetwork", token)
						release.Network = append(release.Network, value)

						isEpisodeTitle = false
					}

				case TSource:
					{
						// fmt.Println("TSource", token)
						release.Source = append(release.Source, value)

						isEpisodeTitle = false
					}

				case TAudio:
					{
						// fmt.Println("TAudio", token)
						release.Audio = append(release.Audio, value)

						isEpisodeTitle = false
					}
				case TChannels:
					{
						// fmt.Println("TChannels", token)
						release.Channels = append(release.Channels, value)

						isEpisodeTitle = false
					}

				case TQuality:
					{
						// fmt.Println("TQuality", token)
						release.Quality = append(release.Quality, strings.ReplaceAll(value, ".", " "))

						isEpisodeTitle = false
					}
				case TCompression:
					{
						// fmt.Println("TCompression", token)
						release.Compression = append(release.Compression, value)

						isEpisodeTitle = false
					}
				default:
					{
						// fmt.Println("DEFAULT", token)

						if isEpisodeTitle && !token.Is(tokenizer.TokenInteger) {
							episodeTitles = append(episodeTitles, value)
						} else {
							if r := rSeasonEpisode.FindStringSubmatch(value); r != nil {
								// fmt.Println(value, r)
								release.Season = r[1]
								release.Episode = r[2]

								isEpisodeTitle = true
							}
						}
					}
				}
			}

			isTitle = false
		}()
	}

	release.Title = strings.TrimSpace(strings.Join(titles, " "))
	release.EpisodeTitle = strings.TrimSpace(strings.Join(episodeTitles, " "))

	return release
}

func normalize(inputs []string) []string {
	outputs := make(map[string]struct{})

	for _, v := range inputs {
		outputs[strings.NewReplacer(".", "_", " ", "_").Replace(v)] = struct{}{}
	}

	return maps.Keys(outputs)
}

var rChannels = regexp.MustCompile(`[^\d]\d([. ])\d(([. ])\d)?[^\d]`)

func sanitize(input string) string {
	normalize := func(input string, inputs []string) string {
		for _, v := range inputs {
			if strings.ContainsAny(v, ". ") {
				// Rebel.Moon.Part.Two.The.Scargiver.2024.1080p.WEBRip.10Bit.DDP5.1.x265-Asiimov
				input = strings.ReplaceAll(input, v, strings.NewReplacer(".", "_", " ", "_").Replace(v)) // Make sure there are spaces around, e.g. Night of the Demons 2 1994 => Night of the Demons 2_1994
			}
		}

		return input
	}

	input = normalize(input, Any)

	input = normalize(input, Tags)

	input = normalize(input, Language)

	input = normalize(input, Resolution)
	input = normalize(input, Quality)
	input = normalize(input, Source)
	input = normalize(input, Compression)

	input = normalize(input, Audio)
	input = normalize(input, Channels)

	input = normalize(input, Network)

	if m := rChannels.FindStringSubmatchIndex((input)); len(m) > 0 {
		// fmt.Println(input)
		// fmt.Println(m, input[m[2]:m[3]])
		input = input[0:m[2]] + "_" + input[m[3]:]
		// fmt.Println(input)
	}

	output := strings.NewReplacer(".", "|", " ", "|", "(", "", ")", "").Replace(input)

	return output
}

type IParser interface {
	Parse(input string) *entity.Release
}

type Parser struct {
	parser *tokenizer.Tokenizer
	stream *tokenizer.Stream
}

func NewParser() IParser {
	parser := tokenizer.New()

	parser.SetWhiteSpaces(append(tokenizer.DefaultWhiteSpaces, byte('|')))

	// parser.DefineTokens(TSeason, []string{"S"})
	// parser.DefineTokens(TEpisode, []string{"E"})
	// parser.AllowKeywordSymbols([]rune{'(', ')', '[', ']', '\'', '.'}, tokenizer.Numbers)
	// parser.DefineTokens(TDot, []string{"."})
	// parser.DefineTokens(TApostrophe, []string{"'"})
	// parser.DefineStringToken(TParenthesis, `(`, `)`)
	// parser.DefineStringToken(TokenQuotedString, `"`, `"`).AddInjection(TokenOpenInjection, TokenCloseInjection)
	// parser.DefineStringToken(TAny, ``, ``).SetEscapeSymbol('.')
	parser.DefineStringToken(TGroup, `-`, `-`)
	// parser.AllowKeywordSymbols([]rune{'-'}, Letters)
	// parser.DefineStringToken(TAlias, `[`, `]`)
	parser.AllowKeywordSymbols([]rune{'\'', ':'}, tokenizer.Numbers) // L'hypothèse // ITV Unwind: Sheep in the Crags

	parser.DefineTokens(TAny, normalize(Any))

	parser.DefineTokens(TAlias, normalize(Alias))

	parser.DefineTokens(TTags, normalize(Tags))

	parser.DefineTokens(TLanguage, normalize(Language))

	// https://en.wikipedia.org/wiki/Pirated_movie_release_types
	// https://de.wikipedia.org/wiki/Releaseformate_(Warez)
	parser.DefineTokens(TResolution, normalize(Resolution))
	parser.DefineTokens(TQuality, normalize(Quality))
	parser.DefineTokens(TSource, normalize(Source))
	parser.DefineTokens(TCompression, normalize(Compression))

	parser.DefineTokens(TAudio, normalize(Audio))
	parser.DefineTokens(TChannels, normalize(Channels))

	parser.DefineTokens(TNetwork, normalize(Network))

	// parser.AllowKeywordUnderscore()

	// parser.AllowKeywordSymbols([]rune{'S', 'E', 'p'}, tokenizer.Numbers)
	// parser.AllowKeywordSymbols([]rune{'S'}, tokenizer.Numbers)
	// parser.AllowKeywordSymbols([]rune{'E'}, tokenizer.Numbers)
	parser.AllowNumbersInKeyword() // S08E05, otherwise S is split from 08E05

	return &Parser{
		parser: parser,
	}
}

func (p *Parser) Parse(input string) *entity.Release {
	input = sanitize(input)

	p.stream = p.parser.ParseString(input)
	defer p.stream.Close()

	return p.parse()
}
