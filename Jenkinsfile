pipeline {
    agent any
    
    tools {
        go 'Go_1.21' // Add "Go_1.21" in Jenkins → Global Tool Configuration
    }
    
    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/your/repo.git'
            }
        }
        
        stage('Build') {
            steps {
                sh 'go mod tidy'
                sh 'go build -v ./...'
            }
        }
        
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
    }
    
    post {
        failure {
            emailext(
                to: 'raghavendrakm8@gmail.com',
                subject: "❌ Build FAILED: ${env.JOB_NAME} #${env.BUILD_NUMBER}",
                body: """
                <h2 style='color:red;'>Build Failed!</h2>
                <p><b>Job:</b> ${env.JOB_NAME}</p>
                <p><b>Build Number:</b> ${env.BUILD_NUMBER}</p>
                <p><b>URL:</b> <a href="${env.BUILD_URL}">${env.BUILD_URL}</a></p>
                <hr>
                <p>Please check Jenkins logs for details.</p>
                """
            )
        }
    }
}