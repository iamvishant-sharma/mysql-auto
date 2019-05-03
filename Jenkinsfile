pipeline {
    agent any

    parameters {
        string(defaultValue: '2290', name: 'port', description:'port')
    }

    stages {
       stage ('automating...') {
            steps {
                sh 'ssh kubecli@47.91.111.100 -p ${port} pwd'
            }
        }
    }
}