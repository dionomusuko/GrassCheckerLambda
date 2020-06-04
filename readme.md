# GrassChecker

Githubの草が生えているか23時に通知するサービス

## Getting Started

 1. slackのwebhook URLを取得
 
       下記アドレスにアクセス
    
       https://slack.com/services/new/incoming-webhook
    
       メッセージを送りたいChannelを指定
    
       Webhook URLをメモ
       
 2. AWS Lambdaの関数を作成
 
 3. LambdaのCloudwatch Events/EvventBridgeをトリガーに設定
 
       スケジュール式として
        
       cron(0 14 * * ? *)　を追加
   
 4. Lambdaに環境変数を追加
   
       | キー | 値 |
       |:----|:----|
      |GITHUB_API|https://api.github.com/users/自分のアカウント名/repos|
      |SLACK_API|webhook URL|
         
 5. Lambda用にコンパイル

         `$ GOOS=linux GOARCH=amd64 go build -o main`
 
       zip化する
         
         `$ zip main.zip ./main`
 
 6. lambndaに関数パッケージとしてmain.zipをアップロード
 
       ハンドラはmain
  