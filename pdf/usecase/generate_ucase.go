package usecase

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/ExPreman/go-svg2pdf/pdf"
	"github.com/jung-kurt/gofpdf"
	"github.com/sirupsen/logrus"
)

type pdfUsecase struct {
	contextTimeout time.Duration
}

func NewPDFUsecase(timeout time.Duration) pdf.PDFUsecase {
	return &pdfUsecase{
		contextTimeout: timeout,
	}
}

func (a *pdfUsecase) GeneratePDF(c context.Context, template string) error {
	var (
		sig gofpdf.SVGBasicType
		err error
	)

	_, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	read, err := ioutil.ReadFile("pdf/template/" + template + ".svg")
	if err != nil {
		logrus.Error(err)
		return err
	}
	// newContents := strings.Replace(string(read), "{{name}}", "Yusuf Septiananda", -1)

	sig, err = gofpdf.SVGBasicParse(read)
	if err != nil {
		logrus.Error(err)
		return err
	}
	scale := 100 / sig.Wd
	scaleY := 30 / sig.Ht
	if scale > scaleY {
		scale = scaleY
	}
	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0
	fmt.Println(sig)
	pdf.SVGBasicWrite(&sig, scale)

	err = pdf.OutputFileAndClose("pdf/template/output.pdf")
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
