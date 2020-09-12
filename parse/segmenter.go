package parse

import (
	"os"
	"github.com/huichen/sego"
)

var instance *segmenter

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
	text := []byte(content)
	segments := s.handler.Segment(text)
  
	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToSlice函数的注释。
	words = sego.SegmentsToSlice(segments, false)
	return
}