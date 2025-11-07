Error Correction
-------------------
502 Bad Gateway - I was getting this error when added "/" home path to the API Gateway, so deleting this endpoint on the API Gateway got rid of the error.

Working command in Windows OS Powershell and Git-Bash
---------------------------------------------------------
Step #1) ./build_start.cmd

AWS CLI upload Setup
--------------------------
NOTE: important to remember that I don't need trustpolicyforlambdaservice.json, instead set my user permissions through AWS's web browser console.
First Goal: login to AWS's CLI.
	Step #1) aws lambda update-function-code --function-name golang_graphQL --zip-file fileb://go_lambda_test3.zip --region us-east-2
NOTE: I get the "region" line from my lambda function's URL.
NOTE: I need to configure my aws cli with my login account info, by running "aws configure", during this type "text" when it asks for "Default output format".
Second Goal: "Add Policy Permission"
	Step #1) go to AWS console, and use the search bar at the top of page to enter "IAM" console page. Then click on the number assigned to the "Roles" text Then on that page select the "Permissions" link/tab and then on the Roles page select my lambda function; "golang_graphQL-role-h4m8zgt8". Then select "AWSLambda_FullAccess", and under the Permissions' tab, select "Lambda" link. Then under the Permissions page, select "UpdateFunctionCode" link, then under Permissions page, click on "Entities attached" and on that page select "Testingaws" checkbox and select "Attach" under "Entities attached" tab. If I set the correct role I will see a success output from the AWS CLI that says my role is;  "Role": "arn:aws:iam::777486915038:role/service-role/golang_graphQL-role-h4m8zgt8",
Third Goal:  aws lambda update-function-code --function-name golang_graphQL --zip-file fileb://go_lambda_test3.zip --region us-east-2 


test sh file (only file runner) try in Git-Bash only.
---------------------------------------------
NOTE;Run Git-Bash in "Administrator mode"  
     Step #1 runs the following code in a "sh file".
Step #1) sh build_exe.sh

git rm -f "bootstrap"
git rm -f "go_lambda_test.zip"
GOOS=linux CGO_ENABLED=0 GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go
git archive --format=zip --output="go_lambda_test.zip" HEAD bootstrap


method of zipping my bootstrap exe
----------------------------------------
git archive --format=zip --output="go_lambda_test.zip" HEAD bootstrap
^^^the above uses the bootstrap of the master branch.

test bat file works in Git-Bash (worked!)
---------------------------------------------
Step #1)  run ./build_start.bat  

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

I got the main code from
-------------------------------- 
- https://github.com/code-heim/go_73_lambda/blob/main/multi_route/main.go
- https://www.youtube.com/watch?v=aAA4tgkv2cI&t=480s

Permissions Info (AWS CLOUD permission)
------------------------------------------
By default, Lambda will create an execution role with permissions to upload logs to Amazon CloudWatch Logs. You can customize this default role later when adding triggers.