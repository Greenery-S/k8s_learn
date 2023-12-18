package main

type Response struct {
	Code    int
	Message string
}

type HealthzResponse struct {
	Response
}

type VersionResponse struct {
	Response
	Version string
}
