package message

import "hc/pkg/packet"

// ErrorResponse sends an error message to the Habbo client.
//
// It's based on the following ugly Lingo code (Login Handler Class.ls:250)
//
//	on handleErr me, tMsg
//		error(me, "Error from server:" && tMsg.content, #handle_error)
//		if tMsg.content contains "login incorrect" then
//		removeConnection(tMsg.connection.getID())
//		me.getComponent().setaProp(#pOkToLogin, 0)
//		if getObject(#session).exists("failed_password") then
//		openNetPage(getText("login_forgottenPassword_url"))
//		me.getInterface().showLogin()
//		return 0
//		else
//		getObject(#session).set("failed_password", 1)
//		me.getInterface().showLogin()
//		executeMessage(#alert, [#Msg: "Alert_WrongNameOrPassword"])
//		end if
//		else
//		if tMsg.content contains "mod_warn" then
//		tDelim = the itemDelimiter
//		the itemDelimiter = "/"
//		tTextStr = tMsg.content.item[2..tMsg.content.item.count]
//		the itemDelimiter = tDelim
//		executeMessage(#alert, [#title: "alert_warning", #Msg: tTextStr, #modal: 1])
//		else
//		if tMsg.content contains "Version not correct" then
//		executeMessage(#alert, [#Msg: "Old client version!!!"])
//		end if
//		end if
//		end if
//		return 1
//	end
type ErrorResponse struct {
	Msg string
}

func (e ErrorResponse) Body() []byte {
	packetWriter := packet.AcquireWriter()
	defer packet.ReleaseWriter(packetWriter)

	//   tMsgs.setaProp(33, #handleErr)
	packetWriter.AppendHeader(33)
	packetWriter.AppendString(e.Msg)

	b, _ := packetWriter.Bytes()
	return b
}
