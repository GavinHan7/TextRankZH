package parse

// Rule interface and its methods make possible the polimorf usage of process
// how Rule retrieve tokens from text.
type Rule interface {
	IsSentenceSeparator(rune rune) bool
}

// RuleDefault struct implements the Rule interface. It contains the separator
// characters and can decide a character is separator or not.
type RuleDefault struct {
	sentenceSeparators [9]string
}

// NewRule constructor retrieves a RuleDefault pointer.
func NewRule() *RuleDefault {
	return &RuleDefault{
		[9]string{"?", "!", ";", "？", "！", "。", "；", "…", "\n"},
	}
}

// IsSentenceSeparator method retrieves true when a character is a kind of
// special character and possibly it separates to words from each other.
func (r *RuleDefault) IsSentenceSeparator(rune rune) bool {
	chr := string(rune)

	for _, val := range r.sentenceSeparators {
		if chr == val {
			return true
		}
	}

	return false
}
