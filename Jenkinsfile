pipeline {
    agent any

    tools {
        go 'go-project'
    }

    environment {
        APP_NAME = 'go-project'
    }

    stages {

        stage('Checkout') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/Tech-devops18/go-project.git'
            }
        }

        stage('Go Version') {
            steps {
                sh 'go version'
            }
        }

        stage('Download Dependencies') {
            steps {
                sh 'go mod tidy'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o ${APP_NAME}'
            }
        }

        stage('Test') {
            steps {
                sh '''
                    set +e

                    go test ./... -v -coverprofile=coverage.out > test_output.txt
                    EXIT_CODE=$?

                    go install github.com/jstemmer/go-junit-report@latest
                    $(go env GOPATH)/bin/go-junit-report < test_output.txt > report.xml

                    go tool cover -func=coverage.out
                    go tool cover -html=coverage.out -o coverage.html || true

                    exit $EXIT_CODE
                '''
            }
        }

        stage('Copy package') {
            steps {
                sh '''
                   mkdir -p gobuild/${BUILD_ID}
                   cp ${APP_NAME} gobuild/${BUILD_ID}/${APP_NAME}
                '''
            }
        }

        stage('Archive Binary') {
            steps {
                archiveArtifacts artifacts: 'gobuild/${BUILD_ID}/*', fingerprint: true
            }
        }
    }

    post {
        always {
            junit 'report.xml'

            publishHTML(target: [
                allowMissing: false,
                alwaysLinkToLastBuild: true,
                keepAll: true,
                reportDir: '.',
                reportFiles: 'coverage.html',
                reportName: 'Go Test Coverage'
            ])

            cleanws()
        }

        success {
            echo 'Go build completed successfully'
        }

        failure {
            echo 'Go build failed'
        }
    }
}
