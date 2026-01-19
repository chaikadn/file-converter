package converter

import (
	"io"

	"github.com/chaikadn/file-converter/internal/models"
)

type Converter interface {
	Convert(input io.Reader, output io.Writer, conversionType models.ConversionType) error // TODO: add ctx, convertion options
	SupportedType() models.ConversionType
}
