package parse

import (
	"os"
	"strings"
	"github.com/huichen/sego"
)

var instance *segmenter
var allowSpeechTags = [15]string{"an", "i", "j", "l", "n", "nr", "nrfg", "ns", "nt", "nz", "t", "v", "vd", "vn", "eng"}

type segmenter struct {
	handler sego.Segmenter
}

func SegmenterInit() *segmenter {
	if instance == nil {
		instance = new(segmenter)

		// 载入词典
		dictPath := os.Getenv("SEGO_DICTIONARY_PATH")
		instance.handler.LoadDictionary(dictPath)
	}
	return instance
}

/**
 * 分词
 */
func (s *segmenter) Cut(content string) (words []string) {
	r := strings.NewReplacer("\t", "", "\n", "", " ", "")
	content = r.Replace(content)
	segments := s.handler.Segment([]byte(content))
  
	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	str := sego.SegmentsToString(segments, false)


	// 过滤词性列表
	for _, word := range strings.Split(str, " ") {
		token := strings.Split(word, "/")
		if len(token) < 2 {
			continue
		}
		if(s.isSpeech(token[1])) {
			words = append(words, token[0])
		}
	}
	return words
}

func (s *segmenter) isSpeech(pos string) bool {
	for _, v := range allowSpeechTags {
		if pos == v {
			return true
		}
	}
	return false
}