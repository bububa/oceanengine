package enum

type ConversionAdMatchType = int

const (
	// 点击归因
	ConversionAdMatchType_CLICK ConversionAdMatchType = 0
	// 展示归因
	ConversionAdMatchType_IMPRESSION ConversionAdMatchType = 1
	// 有效播放归因
	ConversionAdMatchType_PLAY ConversionAdMatchType = 2
)
