pipeline {
    agent any
    tools {
    go 'Gotool'
    }

    stages {
        stage('Go Deps') {
            steps {
                sh 'go version'
                sh 'make deps'

            }
        }
        stage('Ginkgo') {
                    steps {
                        sh 'make functional-tests'
                    }
                }
    }
}
