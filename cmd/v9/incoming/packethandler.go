package incoming

import (
	"github.com/panjf2000/gnet/v2"
)

type PacketHandler struct {
}

func (p PacketHandler) Handle(c gnet.Conn, buf []byte) error {
	return nil
	// Parse packet structure
	//header, packet, err := p.packetParser.Parse(buf)
	//if err != nil {
	//	return err
	//}
	//
	//if err := p.routeExecutor.ExecutePacket(header, packet); err != nil {
	//	return err
	//}
	//
	//return nil

	//if len(buf) < 5 {
	//	return fmt.Errorf("len(packetHello) is too small, got %d", len(buf))
	//}
	//
	//c.Write([]byte("HELLO"))
	//
	//fmt.Println(string(buf))
	//
	//if string(buf[3:5]) == "CJ" {
	//	var buf bytes.Buffer
	//	buf.WriteString(base64.Encode(257)) // ?
	//
	//	buf.Write(vl64.Encode(0))
	//
	//	buf.WriteByte(1)
	//
	//	c.Write(buf.Bytes())
	//
	//	buf.Reset()
	//
	//	buf.WriteString(base64.Encode(8)) // @H
	//	buf.WriteString("[100,105,110,115,120,125,130,135,140,145,150,155,160,165,170,175,176,177,178,180,185,190,195,200,205,206,207,210,215,220,225,230,235,240,245,250,255,260,265,266,267,270,275,280,281,285,290,295,300,305,500,505,510,515,520,525,530,535,540,545,550,555,565,570,575,580,585,590,595,596,600,605,610,615,620,625,626,627,630,635,640,645,650,655,660,665,667,669,670,675,680,685,690,695,696,700,705,710,715,720,725,730,735,740]")
	//	buf.WriteByte(2)
	//	buf.WriteByte(1)
	//
	//	c.Write(buf.Bytes())
	//
	//	log.Info().Msg("Sent @H")
	//
	//	log.Info().Msgf("Sent CJ resp")
	//	return nil
	//}
}
