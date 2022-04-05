// swa is a tiny REST api service with /time endpoint
package main

func main() {
	conf := WebConfig{}
	conf.ParseParams()

	listenAndServe(conf)
}
