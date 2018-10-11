package flv

type DataType uint8

const (
	Number          DataType = 0
	Boolean         DataType = 1
	String          DataType = 2
	Object          DataType = 3
	Null            DataType = 5
	Undefined       DataType = 6
	Reference       DataType = 7
	ECMAArray       DataType = 8
	ObjectEndMarker DataType = 9
	StrictArray     DataType = 10
	Date            DataType = 11
	LongString      DataType = 12
)

func (p *Parser) parseScriptTag(length uint32) error {
	// TODO: parse script tag content
	// write tag header
	if err := p.doWrite(p.bufTH); err != nil {
		return err
	}
	p.buf.Reset()
	// write body
	if err := p.doCopy(length); err != nil {
		return err
	}
	return nil
}