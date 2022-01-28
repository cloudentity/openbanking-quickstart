pipeline {
    agent {
        label 'openbanking'
    }
    environment {
        VERIFY_TEST_RUNNER_TIMEOUT_MS = 80000
    }
    options {
        timeout(time: 1, unit: 'HOURS')
    }
    stages {
        stage('Prepare') {
            steps {
                 sh '''#!/bin/bash
                        echo "127.0.0.1       authorization.cloudentity.com test-docker" | sudo tee -a /etc/hosts
                        cd tests && yarn install
                 '''
            }
        }
        stage('Build') {
            steps {
                sh 'rm -f docker-compose.log'
                sh 'make clean'
                sh 'make build'
            }
        }
        stage('OBUK Tests with disabled MFA') {
            steps {
                script {
                    try {
                        sh 'make disable-mfa run'
                        sh 'make run-tests-headless'
                        sh 'make clean'
                    } catch(exc) {
                        failure('Tests failed')
                    }
                }
            }
        }
        stage('OBUK Tests with enabled MFA') {
            steps {
                script {
                    try {
                        sh 'make enable-mfa run'
                        sh 'make run-tests-headless'
                        sh 'make clean'
                    } catch(exc) {
                        failure('Tests failed')
                    }
                }
            }
        }
        stage('OBBR Tests') {
            steps {
                script {
                    try {
                        sh 'make enable-spec-obbr run'
                        sh 'yarn --cwd tests run cypress run -s cypress/integration/TppTechnicalTests.ts'
                        sh 'make clean'
                    } catch(exc) {
                        failure('Tests failed')
                    }
                }
            }
        }
    }

    post {
        failure {
            sh 'docker-compose logs > docker-compose.log; true'
            archiveArtifacts(artifacts: 'docker-compose.log', allowEmptyArchive: true)
            sh 'make clean'
            archiveArtifacts(artifacts: 'tests/cypress/screenshots/**/*', allowEmptyArchive: true)
            archiveArtifacts(artifacts: 'tests/cypress/videos/**/*', allowEmptyArchive: true)
        }
    }
}
