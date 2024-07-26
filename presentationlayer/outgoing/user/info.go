package user

import "hc/pkg/packet"

type InfoResponse struct {
	Name   string
	Figure string
	Sex    string
	Motto  string
}

func (i InfoResponse) Body() []byte {
	packetWriter := packet.AcquireWriter()
	defer packet.ReleaseWriter(packetWriter)

	packetWriter.AppendHeader(5)
	packetWriter.AppendObject("name", i.Name)
	packetWriter.AppendObject("figure", i.Figure)
	packetWriter.AppendObject("sex", i.Sex)
	packetWriter.AppendObject("customData", i.Motto)
	packetWriter.AppendObject("ph_tickets", "0")
	packetWriter.AppendObject("photo_film", "0")
	packetWriter.AppendObject("ph_figure", "")
	packetWriter.AppendObject("directMail", "0")

	b, _ := packetWriter.Bytes()
	return b
}
