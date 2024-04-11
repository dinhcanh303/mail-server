pipeline {
    agent any
    stages {
        stage('Clone'){
            steps {
                git 'https://github.com/dinhcanh303/mail-server.git'
            }
        }
    }
    post {
        always {
            mail bcc: '', body: 'Please check server ?', cc: 'dinhcanhng303@gmail.com', from: '', replyTo: '', subject: 'Jenkins Clone Code', to: 'rapidstars97@gmail.com'
        }
    }
}