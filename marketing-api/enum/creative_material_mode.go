package enum

// CreativeMaterialMode 创意呈现方式
type CreativeMaterialMode string

const (
	// CreativeMaterialMode_CUSTOM_CREATIVE 自定义创意
	CreativeMaterialMode_CUSTOM_CREATIVE CreativeMaterialMode = "CUSTOM_CREATIVE"
	// CreativeMaterialMode_PROGRAMMATIC_CREATIVE 程序化创意
	CreativeMaterialMode_PROGRAMMATIC_CREATIVE CreativeMaterialMode = "PROGRAMMATIC_CREATIVE"
	// CreativeMaterialMode_STATIC_ASSEMBLE 程序化创意
	CreativeMaterialMode_STATIC_ASSEMBLE CreativeMaterialMode = "STATIC_ASSEMBLE"
	// CreativeMaterialMode_INTERVENE 自定义创意
	CreativeMaterialMode_INTERVENE CreativeMaterialMode = "INTERVENE"
)
