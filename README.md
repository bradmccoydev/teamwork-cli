# teamwork-cli
This application is dedicated to all those who hate logging in to do timesheets.

This application is written in Go and using Github actions for CI/CD

# Dependencies
For the CLI I am using Cobra - https://github.com/spf13/cobra
For the Config I am using Viper - https://github.com/spf13/viper
For logging I am Using zerolog - https://github.com/rs/zerolog

To download locally
go get -u github.com/spf13/cobra
go get -u github.com/spf13/viper
go get -u github.com/rs/zerolog/log

# Arguments
addtime "https://search365.teamwork.com/" "twp_63AhzDGh681Wb88QWJLw5X7LXJaY" "370940" "one" "two"
# Golang
To Test Locally:

``` go run main.go ```

To build the binary:

``` go build main.go ```

To Run the binary:
``` ./main ```

# Docker
docker build . -t bradmccoydev/teamwork-cli
