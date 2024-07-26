package handshake

import (
	"fmt"
	"hc/pkg/packet"
)

type TryLogin struct {
	Username string
	Password string
}

// ParseTryLogin parses the "TRY_LOGIN" packet.
//
// Login Handler Class.ls:307 contains the following:
// tCmds.setaProp("TRY_LOGIN", 4)
//
// Which is used by sendLogin (line 78):
//
//	on sendLogin me, tConnection
//		if objectExists("nav_problem_obj") then
//		removeObject("nav_problem_obj")
//		end if
//		if me.getComponent().isOkToLogin() then
//		tUserName = getObject(#session).get(#userName)
//		tPassword = getObject(#session).get(#Password)
//		if not stringp(tUserName) or not stringp(tPassword) then
//		return removeConnection(tConnection.getID())
//		end if
//		if (tUserName = EMPTY) or (tPassword = EMPTY) then
//		return removeConnection(tConnection.getID())
//		end if
//		return tConnection.send("TRY_LOGIN", [#string: tUserName, #string: tPassword])
//		end if
//		return 1
//	end
//
// With other words, we'll be receiving two strings from the client: username and password.
func ParseTryLogin(body []byte) (TryLogin, error) {
	reader := packet.AcquireReader(body)
	defer packet.ReleaseReader(reader)

	username, err := reader.String()
	if err != nil {
		return TryLogin{}, fmt.Errorf("unable to get username from packet body: %s", err.Error())
	}

	password, err := reader.String()
	if err != nil {
		return TryLogin{}, fmt.Errorf("unable to get password from packet body: %s", err.Error())
	}

	return TryLogin{
		Username: username,
		Password: password,
	}, nil
}
