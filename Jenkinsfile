pipeline {
    agent any
    tools {
    go 'Gotool'
    }

    stages {
        stage('Go Deps') {
            steps {
                sh 'go version'
                sh 'go mod download'

            }
        }
        stage('Ginkgo') {
                    steps {
                        sh 'ginkgo -v ./address-api-tests ./ginkgo-tests'
                    }
                }
    }
}
