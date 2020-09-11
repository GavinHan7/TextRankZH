package parse

import (
	"github.com/huichen/sego"
)

// TokenizeText function use the given raw text and parses by a Rule object and
// retrieves the parsed text in a Text struct object.
func TokenizeText(rawText string, rule Rule) Text {
	return findSentences(rawText, rule)
}

func findSentences(rawText string, rule Rule) Text {
	text := Text{}

	var sentence string

	for _, chr := range rawText {
		if !rule.IsSentenceSeparator(chr) {
			sentence = sentence + string(chr)
		} else if len(sentence) > 0 {
			sentence = sentence + string(chr)

			text.Append(sentence, findWords(sentence))

			sentence = ""
		}
	}

	if len(sentence) > 0 {
		text.Append(sentence, findWords(sentence))
	}

	return text
}

func findWords(rawSentence string) (words []string) {
	
	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("github.com/huichen/sego/data/dictionary.txt")

	// 分词
	text := []byte(rawSentence)
	segments := segmenter.Segment(text)
  
	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	words = sego.SegmentsToSlice(segments, false) 

	return
}
