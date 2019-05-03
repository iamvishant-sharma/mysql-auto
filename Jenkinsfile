pipeline {
    agent any

    parameters {
        string(defaultValue: 'uat', name: 'port', description:'port')
    }

    stages {
       stage ('automating...') {
            steps {
                sh 'ssh kubecli@47.91.111.100 -p 2290 ansible -m ping localhost'
            }
        }
    }
}