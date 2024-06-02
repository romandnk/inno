package types

import "inno/hw3/formatter"

type ChainFormatter struct {
	formats []formatter.Formatter
}

func (c *ChainFormatter) AddFormatter(f formatter.Formatter) *ChainFormatter {
	c.formats = append(c.formats, f)
	return c
}

func (c *ChainFormatter) Format(str string) string {
	for _, formatter := range c.formats {
		str = formatter.Format(str)
	}
	return str
}
