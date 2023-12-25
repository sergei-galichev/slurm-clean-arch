//go:build swag
// +build swag

package http

// Need to install swag: go install github.com/swaggo/swag/cmd/swag@latest
// For project sync: go get github.com/swaggo/swag/cmd/swag

//go:generate swag init --parseDependency --generalInfo delivery.go --output swagger/docs/
