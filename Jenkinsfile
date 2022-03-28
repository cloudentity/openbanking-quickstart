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

        stage('CDR Tests') {
            steps {
                script {
                    try {
                        sh 'make run-cdr-local'
                        retry(3) {
                            sh 'make run-cdr-tests-headless'
                        }
                        sh 'make clean'
                    } catch(exc) {
                        failure('Tests failed')
                    }
                }
            }
        }

        stage('OBUK Tests with disabled MFA') {
            steps {
                script {
                    try {
                        sh 'make disable-mfa run-obuk-local'
                        sh 'make run-obuk-tests-headless'
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
                        sh 'make enable-mfa run-obuk-local'
                        sh 'make run-obuk-tests-headless'
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
                        sh 'make run-obbr-local'
                        sh 'make run-obbr-tests-headless'
                        sh 'make clean'
                    } catch(exc) {
                        failure('Tests failed')
                    }
                }
            }
        }
        stage('SaaS Tests') {
            environment {
                SAAS_TENANT_ID = 'amfudxn6-qa-us-east-1-ob-quickstart'
                SAAS_CLIENT_ID = credentials('OPENBANKING_CONFIGURATION_CLIENT_ID')
                SAAS_CLIENT_SECRET = credentials('OPENBANKING_CONFIGURATION_CLIENT_SECRET')
            }
            steps {
                script {
                    try {
                        sh 'make disable-mfa set-saas-configuration run-apps-with-saas'
                        sh 'make run-saas-tests-headless'
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
            sh 'docker-compose -f docker-compose.acp.local.yaml -f docker-compose.obuk.yaml -f docker-compose.obbr.yaml -f docker-compose.cdr.yaml logs > docker-compose.log; true'
            archiveArtifacts(artifacts: 'docker-compose.log', allowEmptyArchive: true)
            sh 'make clean'
            archiveArtifacts(artifacts: 'tests/cypress/screenshots/**/*', allowEmptyArchive: true)
            archiveArtifacts(artifacts: 'tests/cypress/videos/**/*', allowEmptyArchive: true)
        }
    }
}
