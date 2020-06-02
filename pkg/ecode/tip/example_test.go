package tip_test

import (
	"time"

	"git.atmatrix.org/k12/zero/pkg/ecode/tip"
	xhttp "git.atmatrix.org/k12/zero/pkg/net/http/blademaster"
	"git.atmatrix.org/k12/zero/pkg/net/netutil/breaker"
	xtime "git.atmatrix.org/k12/zero/pkg/time"
)

func ExampleInit() {
	conf := &tip.Config{
		Domain: "172.16.33.248:6401",
		Diff:   xtime.Duration(5 * time.Minute),
		ClientConfig: &xhttp.ClientConfig{
			App: &xhttp.App{
				Key:    "test",
				Secret: "e6c4c252dc7e3d8a90805eecd7c73396",
			},
			Dial:      xtime.Duration(time.Millisecond * 100),
			Timeout:   xtime.Duration(time.Second * 2),
			KeepAlive: xtime.Duration(time.Second * 2),
			Breaker: &breaker.Config{
				Window:  xtime.Duration(time.Millisecond * 10),
				Sleep:   xtime.Duration(time.Second * 10),
				Bucket:  10,
				Ratio:   0.5,
				Request: 100,
			},
		},
	}
	tip.Init(conf)
}
