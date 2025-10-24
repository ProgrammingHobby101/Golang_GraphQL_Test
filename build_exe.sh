git add .
git rm -f "bootstrap"
git rm -f "go_lambda_test.zip"
GOOS=linux CGO_ENABLED=0 GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go
git archive --format=zip --output="go_lambda_test.zip" HEAD bootstrap