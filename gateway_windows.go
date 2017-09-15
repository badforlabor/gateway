package gateway

import (
	"net"
	"os/exec"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func DiscoverGateway() (ip net.IP, err error) {
	routeCmd := exec.Command("route", "print", "0.0.0.0")
	output, err := routeCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	str, err := Decode(output)
	if err != nil {
		return nil, err
	}

	return parseWindowsRoutePrint(str)
}
func Decode(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	//defer I.Close()
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	//defer O.Close()
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
