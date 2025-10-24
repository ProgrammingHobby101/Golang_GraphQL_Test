test bat file works in Windows PowerShell
---------------------------------------------
Step #1)  run .\build_start.bat    

test#2 bat file
------------------------------
DEL go_lambda_test.zip
sh build_exe.sh
zip myFunction.zip bootstrap



test#2 bat file
------------------------------
DEL go_lambda_test.tar.zip
sh build_exe.sh
 tar -czvf go_lambda_test.tar.zip bootstrap
tar -tzvf go_lambda_test.tar.zip


Best YouTube video tutorial
---------------------------------
https://www.youtube.com/watch?v=aAA4tgkv2cI

Build Golang lambda code into a executable (WORKS! Automated)
---------------------------------------------------------------
Step #1) Open git-bash.
Step #2) Run the following command in git-bash
        ./build_start.bat
Step #3) upload the zip "go_lambda_test.tar.gz" to AWS Lambda Code.

Build Golang lambda code into a executable (MANUAL RUNS) 
--------------------------------------------------
NOTE: may have to run the build_start.bat file in git-bash so that the "-o bootstrap" file get run in a Linux shell(git-bash).
Step#1) Open "Git Bash"
Step#2) run the following command in "git bash"
GOOS=linux CGO_ENABLED=0 GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go
Step#3) zip the bootstrap executable.
Note: I got this code from; https://gemini.google.com/app/eefd04fee68d2bb9?is_sa=1&is_sa=1&android-min-version=301356232&ios-min-version=322.0&campaign_id=bkws&utm_source=sem&utm_source=google&utm_medium=paid-media&utm_medium=cpc&utm_campaign=bkws&utm_campaign=2024enUS_gemfeb&pt=9008&mt=8&ct=p-growth-sem-bkws&gclsrc=aw.ds&gad_source=1&gad_campaignid=20108148196&gbraid=0AAAAApk5BhnQgzvhNwLwljcWmGLCINdfx&gclid=CjwKCAjw3tzHBhBREiwAlMJoUiNA9NAxXs40D0kGk8HuDhj96VW9GfEWdi4-Rq6910_G_YsI_2ElahoCSqMQAvD_BwE  


Permissions Info (AWS CLOUD permission)
------------------------------------------
By default, Lambda will create an execution role with permissions to upload logs to Amazon CloudWatch Logs. You can customize this default role later when adding triggers.