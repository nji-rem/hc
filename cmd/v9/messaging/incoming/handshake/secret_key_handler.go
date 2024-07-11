package handshake

import (
	"bytes"
	"context"
	"github.com/panjf2000/gnet/v2"
	"hc/pkg/packet/encoding/base64"
	"hc/pkg/packet/encoding/vl64"
)

func HandleSecretKey(ctx context.Context, c gnet.Conn, packet any) error {
	var buf bytes.Buffer
	buf.WriteString(base64.Encode(257)) // ?

	buf.Write(vl64.Encode(0))

	buf.WriteByte(1)

	c.Write(buf.Bytes())

	buf.Reset()

	buf.WriteString(base64.Encode(8)) // @H
	buf.WriteString("[100,105,110,115,120,125,130,135,140,145,150,155,160,165,170,175,176,177,178,180,185,190,195,200,205,206,207,210,215,220,225,230,235,240,245,250,255,260,265,266,267,270,275,280,281,285,290,295,300,305,500,505,510,515,520,525,530,535,540,545,550,555,565,570,575,580,585,590,595,596,600,605,610,615,620,625,626,627,630,635,640,645,650,655,660,665,667,669,670,675,680,685,690,695,696,700,705,710,715,720,725,730,735,740]")
	buf.WriteByte(2)
	buf.WriteByte(1)

	c.Write(buf.Bytes())

	return nil
}
