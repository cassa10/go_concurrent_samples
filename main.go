package main

import (
	"github.com/cassa10/go_concurrent_samples/channels_responses"
	"github.com/cassa10/go_concurrent_samples/high_order"
	"github.com/cassa10/go_concurrent_samples/util"
)

func main() {
	util.DoAllIf(false, []func(){
		channel_responses.Sample1SimpleChannels,
		channel_responses.Sample2GenericChannels,
		channel_responses.Sample3Timeout,
	})
	util.DoAllIf(true, []func(){
		high_order.Sample,
	})

}
