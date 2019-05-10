pipeline {
    agent any

    parameters {
        string(defaultValue: '2290', name: 'port', description:'port')
        string(defaultValue: 'db_name', name: 'db_name', description:'Database name to be created')
        string(defaultValue: 'username', name: 'username', description:'username to be created')
        string(defaultValue: 'password', name: 'password', description:'password to be created')
        string(defaultValue: 'select,insert', name: 'privileges', description:'privileges to be given to the user')
    }

    stages {
       stage ('automating...') {
            steps {
                sh "ssh kubecli@47.91.111.100 -p ${port} ansible-playbook site.yaml -e 'login_user=root login_password=7fpJ3ffm9AHYPOe563U9 db_name=${db_name} username=${username} password=${password} priv=s${privileges}'"
            }
        }
    }
}
