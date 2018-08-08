package xlsWrite

import (
	"github.com/tealeg/xlsx"
)

type Report struct {
	isOpen bool
	file   *xlsx.File
}

func All() {
	xlsx.NewFile()
}

func (r *Report) CreateReport(Name string) {
	r.file = xlsx.NewFile()
	r.isOpen = true

}
