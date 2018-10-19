package pdf

import (
	"context"
)

type PDFUsecase interface {
	GeneratePDF(ctx context.Context, template string) error
}
