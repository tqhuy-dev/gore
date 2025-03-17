package utilities

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

type IExcelManager interface {
}

type excelManager struct {
	Name     string
	mapSheet map[string]int
	file     *excelize.File
	mapStyle map[string]int
}

func NewExcelManager(name string) IExcelManager {
	f := excelize.NewFile()
	return &excelManager{
		Name:     name,
		mapSheet: make(map[string]int),
		mapStyle: make(map[string]int),
		file:     f,
	}
}

type OptHeader struct {
	Headers   []interface{}
	ColStart  int
	RowStart  int
	Cell      string
	StyleName string
}

type OptCell struct {
	ColTopLeft     int
	RowTopLeft     int
	ColBottomRight int
	RowBottomRight int
	StyleName      string
	Values         []interface{}
}

func (e *excelManager) SetHeaders(sheetName string, opt OptHeader) error {
	if len(opt.Cell) == 0 || (opt.ColStart == 0 && opt.RowStart == 0) || len(opt.Headers) == 0 {
		return nil
	}
	var err error
	if len(opt.Cell) == 0 {
		opt.Cell, err = excelize.CoordinatesToCellName(opt.ColStart, opt.RowStart)
		if err != nil {
			return err
		}
	}
	if len(opt.StyleName) > 0 && e.mapStyle[opt.StyleName] > 0 {
		if err = e.SetCellStyle(sheetName, OptCell{
			ColTopLeft:     opt.ColStart,
			RowTopLeft:     opt.RowStart,
			ColBottomRight: opt.ColStart + len(opt.Headers) - 1,
			RowBottomRight: opt.RowStart,
			StyleName:      opt.StyleName,
		}); err != nil {
			return err
		}
	}
	if err = e.file.SetSheetRow(sheetName, opt.Cell, opt.Headers); err != nil {
		return err
	}
	return nil
}

func (e *excelManager) NewStyle(styleName string, style excelize.Style) error {
	styleId, err := e.file.NewStyle(&style)
	if err != nil {
		return err
	}
	e.mapStyle[styleName] = styleId
	return nil
}

func (e *excelManager) SetCellStyle(sheetName string, opt OptCell) error {
	topLeftCell, err := excelize.CoordinatesToCellName(opt.ColTopLeft, opt.RowTopLeft)
	if err != nil {
		return err
	}

	bottomRightCell, err := excelize.CoordinatesToCellName(opt.ColBottomRight, opt.RowBottomRight)
	if err != nil {
		return err
	}

	return e.file.SetCellStyle(sheetName, topLeftCell, bottomRightCell, e.mapStyle[opt.StyleName])
}

func (e *excelManager) SetCellValue(sheetName string, opt OptCell) error {
	cell, err := excelize.CoordinatesToCellName(opt.ColTopLeft, opt.RowTopLeft)
	if err != nil {
		return err
	}

	if len(opt.StyleName) > 0 && e.mapStyle[opt.StyleName] > 0 {
		if err = e.SetCellStyle(sheetName, OptCell{
			ColTopLeft:     opt.ColTopLeft,
			RowTopLeft:     opt.RowTopLeft,
			ColBottomRight: opt.ColTopLeft + len(opt.Values) - 1,
			RowBottomRight: opt.RowTopLeft,
			StyleName:      opt.StyleName,
		}); err != nil {
			return err
		}
	}
	if err = e.file.SetSheetRow(sheetName, cell, opt.Values); err != nil {
		return err
	}
	return nil
}

func (e *excelManager) SaveAs() error {
	return e.file.SaveAs(fmt.Sprintf("%s.xlsx", e.Name))
}
