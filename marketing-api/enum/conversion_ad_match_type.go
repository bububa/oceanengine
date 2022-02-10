package enum

// ConversionAdMatchType 归因方式
type ConversionAdMatchType = int

const (
	// ConversionAdMatchType_CLICK 点击归因
	ConversionAdMatchType_CLICK ConversionAdMatchType = 0
	// ConversionAdMatchType_IMPRESSION 展示归因
	ConversionAdMatchType_IMPRESSION ConversionAdMatchType = 1
	// ConversionAdMatchType_PLAY 有效播放归因
	ConversionAdMatchType_PLAY ConversionAdMatchType = 2
)
