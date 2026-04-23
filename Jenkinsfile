pipeline {
    agent any
    tools {
    go 'Gotool'
    }

    stages {
        stage('Go Deps') {
            steps {
                sh 'go --version'

            }
        }
        stage('Ginkgo') {
                    steps {
                        echo 'Hello  from jenkins World'
                    }
                }
    }
}
