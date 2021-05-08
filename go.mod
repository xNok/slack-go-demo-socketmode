module xnok/slack-go-demo

go 1.16

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/go-test/deep v1.0.7
	github.com/joho/godotenv v1.3.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/zerolog v1.20.0
	github.com/slack-go/slack v0.8.1
)

replace github.com/slack-go/slack => github.com/xnok/slack v0.8.1-0.20210508192837-9bc8016012b3
