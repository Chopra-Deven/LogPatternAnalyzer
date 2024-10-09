package store

import (
	"LogPattern/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/snappy"
	cmap "github.com/orcaman/concurrent-map"
	"go.elara.ws/pcre"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

const MatchingScore = 0.5
const PatternPlaceHolder = "****"

type MaskingPattern struct {
	pattern *pcre.Regexp
	value   string
}

var (
	items = cmap.New() // pattern as key and patternId as value

	patternIds = cmap.New() // pluginId and comma-separated patternIds

	patterns = sync.Map{} // patternId with pattern

	patternId int32 // Atomic pattern ID

	patternCounter int32 // Atomic counter

	logger = utils.NewLogger("Log Pattern Cache Store", "cacheStore")

	lookupPatterns []*MaskingPattern
)

func Init() {

	loadPatterns()

	lookupPatterns = []*MaskingPattern{

		{value: "<EMAIL>", pattern: pcre.MustCompile("([\\w+\\.+\\-]+@+[\\w+\\.+\\-]+[\\.\\w]{2,})((?=[^A-Za-z0-9])|$)")},
		{value: "<URL>", pattern: pcre.MustCompile("(((https|http):\\/\\/|www[.])[A-Za-z0-9+&@#\\/%?=~_()-|!:,.;]*[-A-Za-z0-9+&@#\\/%=~_()|])((?=[^A-Za-z0-9])|$)")},
		{value: "<ID>", pattern: pcre.MustCompile("(([0-9a-f]{2,}:){3,}([0-9a-f]{2,}))((?=[^A-Za-z0-9])|$)")},
		{value: "<IP>", pattern: pcre.MustCompile("(\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3})((?=[^A-Za-z0-9])|$)")},
		{value: "<GUID>", pattern: pcre.MustCompile("([\\w]{8}\\b-[\\w]{4}\\b-[\\w]{4}\\b-[\\w]{4}\\b-[\\w]{12})((?=[^A-Za-z0-9])|$)")},
		{value: "<NUM>", pattern: pcre.MustCompile("([\\-\\+]?\\d+)((?=[^A-Za-z0-9])|$)")},
		{value: "<SEQ>", pattern: pcre.MustCompile("([0-9a-f]{6,} ?){3,}((?=[^A-Za-z0-9])|$)")},
		{value: "<SEQ>", pattern: pcre.MustCompile("([0-9A-F]{4} ?){4,}((?=[^A-Za-z0-9])|$)")},
		{value: "<HEX>", pattern: pcre.MustCompile("(0x[a-f0-9A-F]+)((?=[^A-Za-z0-9])|$)")},
		{value: "<CMD>", pattern: pcre.MustCompile("(?<=executed cmd )(\".+?\")")},
	}

}

func DetectPattern(context utils.MotadataMap, tokenizers []*utils.Tokenizer) utils.MotadataMap {

	atomic.AddInt32(&patternCounter, 1)

	message := context.GetStringValue("message")

	plugin := context.GetStringValue("plugin.id") + "-" + context.GetStringValue("event.category")

	context[utils.Status] = utils.StatusFail

	message = mask(message)

	tokenizers[0].Split(message, utils.SpaceSeparator)

	//tokens := strings.Split(message, utils.SpaceSeparator)

	// Check if plugin already has patterns
	if idsRaw, ok := patternIds.Get(plugin); ok {

		if idsRaw != nil {

			if ids, ok := idsRaw.([]int); ok {

				min, max := 0, 0

				for _, id := range ids {

					bufferBytes := bytes.Buffer{}

					if pattern, ok := patterns.Load(id); ok {

						stringPattern := pattern.(string)

						tokenizers[1].Split(stringPattern, utils.SpaceSeparator)

						//patternTokens := strings.Split(stringPattern, utils.SpaceSeparator)

						if tokenizers[1].Counts > tokenizers[0].Counts {

							min = tokenizers[0].Counts

							max = tokenizers[1].Counts

						} else {

							min = tokenizers[1].Counts

							max = tokenizers[0].Counts

						}

						tokens := matchTokens(tokenizers[1].Tokens[:tokenizers[1].Counts], tokenizers[0].Tokens[:tokenizers[0].Counts], min, max)

						if tokens != nil && len(tokens) > 0 {

							for i := 0; i < max; i++ {

								if contains(tokens, i) && i < tokenizers[1].Counts {

									bufferBytes.WriteString(tokenizers[1].Tokens[i] + utils.SpaceSeparator)

								} else {

									bufferBytes.WriteString(PatternPlaceHolder + utils.SpaceSeparator)
								}
							}

							messagePattern := strings.TrimSpace(bufferBytes.String())

							if _, exists := items.Get(messagePattern); !exists {

								if value, exists := items.Get(stringPattern); exists {

									// Key exists, so delete and get the value
									items.Remove(pattern.(string))

									updatePattern(messagePattern, value.(int), plugin, ids)
								}
							}

							if val, exists := items.Get(messagePattern); exists {

								context[utils.Status] = utils.StatusSucceed

								context["pattern.id"] = val.(int)

								if value, ok := patterns.Load(val.(int)); ok {

									context["pattern"] = value.(string)
								}

								return context
							}
						}
					}
				}

			}

		}

	}

	// If no match, add new pattern
	updatePattern(message, int(atomic.AddInt32(&patternId, 1)), plugin, nil)

	if val, exists := items.Get(message); exists {

		context[utils.Status] = utils.StatusSucceed

		context["pattern.id"] = val.(int)

		if value, ok := patterns.Load(val.(int)); ok {

			context["pattern"] = value.(string)
		}

		return context
	}

	return context
}

// MatchTokens function compares tokens and returns the matching indices
func matchTokens(patternTokens, messageTokens []string, min, max int) []int {

	var patternScore float64

	var tokens []int

	for i := 0; i < min; i++ {

		score := getScore(strings.TrimSpace(patternTokens[i]), strings.TrimSpace(messageTokens[i]))

		patternScore += 1 * (score / float64(max))

		if score == 1 {
			tokens = append(tokens, i)
		}
	}

	if (1 - patternScore) <= MatchingScore {
		return tokens
	}

	return nil
}

// UpdatePattern updates the items and patternIds maps
func updatePattern(pattern string, id int, plugin string, ids []int) {

	items.Set(pattern, id)

	patterns.Store(id, pattern)

	var patternIdsSet map[int]struct{}

	if ids != nil && len(ids) > 0 {

		patternIdsSet = make(map[int]struct{})

		for _, patternId := range ids {

			patternIdsSet[patternId] = struct{}{}
		}

		patternIdsSet[id] = struct{}{}

		var idsList []int

		for patternId := range patternIdsSet {

			idsList = append(idsList, patternId)
		}

		patternIds.Set(plugin, idsList)

	} else {

		if idsRaw, ok := patternIds.Get(plugin); ok {

			if idsRaw != nil {

				if ids, ok := idsRaw.([]int); ok {

					patternIds.Set(plugin, append(ids, id))
				}
			}

		} else {

			patternIdsSet = map[int]struct{}{id: {}}

			var idsList []int

			for patternId := range patternIdsSet {
				idsList = append(idsList, patternId)
			}

			patternIds.Set(plugin, idsList)
		}
	}

}

func contains(slice []int, elem int) bool {

	for _, v := range slice {

		if v == elem {
			return true
		}

	}
	return false
}

func getScore(token1, token2 string) float64 {

	if len(token1) > 0 && len(token1) > 0 && !strings.EqualFold(token1, PatternPlaceHolder) && !strings.EqualFold(token2, PatternPlaceHolder) && token1 == token2 {

		return 1

	}

	return 0
}

func mask(message string) string {

	for _, pattern := range lookupPatterns {

		message = pattern.pattern.ReplaceAllString(message, pattern.value)
	}

	return message
}

func Flush(path string) {

	// Combine all maps into a single structure
	data := make(map[string]interface{})

	stats := utils.MotadataMap{}

	// Convert items cmap to a standard map
	itemsMap := make(map[string]interface{})

	for item := range items.IterBuffered() {
		itemsMap[item.Key] = item.Val
	}

	data["items"] = itemsMap

	stats["items"] = itemsMap

	standardMap := make(map[int]interface{})
	// Convert sync.Map to a standard map
	patterns.Range(func(key, value interface{}) bool {
		intKey := key.(int)         // Type assertion to int for the key
		standardMap[intKey] = value // Direct assignment into standard map
		return true
	})

	stats["standards"] = standardMap

	// Convert patternIds cmap to a standard map
	patternIdsMap := make(map[string]interface{})
	for item := range patternIds.IterBuffered() {

		patternIdsMap[item.Key] = item.Val

		stats[item.Key] = len(item.Val.([]int))
	}
	data["patternIds"] = patternIdsMap

	data["last.patternID"] = atomic.LoadInt32(&patternId)

	stats["patternIds"] = patternIdsMap

	stats["total.items"] = items.Count()

	stats["detected.patterns"] = atomic.LoadInt32(&patternCounter)

	stats["last.patternID"] = atomic.LoadInt32(&patternId)

	atomic.StoreInt32(&patternCounter, int32(0))

	// Marshal the combined map to JSON
	bytes, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		logger.Info(utils.MotadataString(fmt.Sprintf("Error marshaling data to JSON while flushing %v	:", err)))
		return
	}

	statBytes, err := json.MarshalIndent(stats, "", "  ")

	if err != nil {
		logger.Info(utils.MotadataString(fmt.Sprintf("Error marshaling stat data to JSON while flushing %v:", err)))
		return
	}

	logger.Info("State JSON : " + utils.MotadataString(string(statBytes)))

	// Compress JSON data using Snappy
	compressedData := snappy.Encode(nil, bytes)

	if err := os.WriteFile(path, compressedData, 0644); err != nil {

		logger.Info(utils.MotadataString(fmt.Sprintf("failed to write aggregation context file : %s , reason : %s", path, err.Error())))

	} else {

		if err = os.Rename(path, path); err != nil {

			logger.Info(utils.MotadataString(fmt.Sprintf("failed to overwrite aggregation context file : %s , reason : %s", path, err.Error())))
		}
	}
}

func loadPatterns() {

	path := utils.CurrentDir + utils.PathSeparator + utils.ConfigDirectory + utils.PathSeparator + "log-patterns"

	compressedData, err := os.ReadFile(path)
	if err != nil {
		logger.Info(utils.MotadataString(fmt.Sprintf("Error reading file: %v", err)))
		return
	}

	jsonData, err := snappy.Decode(nil, compressedData)
	if err != nil {
		logger.Info(utils.MotadataString(fmt.Sprintf("Error decompressing data: %v", err)))
		return
	}

	data := make(map[string]interface{})

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		logger.Info(utils.MotadataString(fmt.Sprintf("Error unmarshaling data: %v", err)))
		return
	}

	if itemsData, ok := data["items"].(map[string]interface{}); ok {

		for key, value := range itemsData {

			items.Set(fmt.Sprint(key), int(value.(float64)))

			patterns.Store(int(value.(float64)), fmt.Sprint(key))
		}
	}

	atomic.StoreInt32(&patternId, int32(data["last.patternID"].(float64)))

	// Populate the `patternIds` cmap
	if patternIdsData, ok := data["patternIds"].(map[string]interface{}); ok {

		for key, value := range patternIdsData {

			var patternIdsList []int

			for _, id := range value.([]interface{}) {

				patternIdsList = append(patternIdsList, int(id.(float64)))
			}

			patternIds.Set(fmt.Sprint(key), patternIdsList)
		}
	}

	logger.Info("Data successfully read from file and maps populated")
}
