pipeline {
    agent {
        label 'openbanking'
    }
    triggers {
        cron(env.BRANCH_NAME == 'master' ? 'H 5 * * *' : '')
    }
    options {
        skipStagesAfterUnstable()
        timeout(time: 1, unit: 'HOURS')
    }
    environment {
        VERIFY_TEST_RUNNER_TIMEOUT_MS = 80000
        SAAS_TENANT_ID = 'amfudxn6-qa-us-east-1-ob-quickstart'
        SAAS_CLIENT_ID = credentials('OPENBANKING_CONFIGURATION_CLIENT_ID')
        SAAS_CLIENT_SECRET = credentials('OPENBANKING_CONFIGURATION_CLIENT_SECRET')
        SAAS_CLEANUP_CLIENT_ID = credentials('OPENBANKING_CLEANUP_CLIENT_ID')
        SAAS_CLEANUP_CLIENT_SECRET = credentials('OPENBANKING_CLEANUP_CLIENT_SECRET')
        DEBUG = 'true'
    }
    stages {
        stage('Prepare') {
            steps {
                 echo 'BRANCH_NAME: ' + env.BRANCH_NAME
                 sh '''#!/bin/bash
                        echo "127.0.0.1       authorization.cloudentity.com test-docker" | sudo tee -a /etc/hosts
                        cd tests && yarn install
                 '''
                 sh 'docker-compose version'
            }
        }
        stage('Build') {
            steps {
                sh 'rm -f docker-compose.log'
                sh 'make clean'
                sh 'make lint'
                sh 'make stop-runner'
                sh 'make build'
            }
        }

        stage('CDR Tests') {
            steps {
                script {
                    try {
                        sh 'make run-cdr-local'
                        sh 'make run-cdr-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('CDR Tests failed')
                    } finally {
                        sh 'make clean'
                    }
                }
            }
        }
        stage('FDX Tests with disabled MFA') {
            steps {
                script {
                    try {
                        sh 'make disable-mfa run-fdx-local'
                        sh 'make run-fdx-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('FDX Tests with disabled MFA failed')
                    } finally {
                        sh 'make clean'
                    }
                }
            }
        }
        stage('FDX Tests with enabled MFA') {
            steps {
                script {
                    try {
                        sh 'make enable-mfa run-fdx-local'
                        sh 'make run-fdx-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('FDX Tests with enabled MFA failed')
                    } finally {
                        sh 'make clean'
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
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('OBUK Tests with disabled MFA failed')
                    } finally {
                        sh 'make clean'
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
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('OBUK Tests with enabled MFA failed')
                    } finally {
                        sh 'make clean'
                    }
                }
            }
        }
        stage('OBBR Tests with disabled MFA') {
            steps {
                script {
                    try {
                        sh 'make disable-mfa run-obbr-local'
                        sh 'make run-obbr-tests-headless'
                        sh 'make clean'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('OBBR Tests with disabled MFA failed')
                    } finally {
                        sh 'make clean'
                    }
                }
            }
        }
        stage('OBBR Tests with enabled MFA') {
            steps {
                script {
                    try {
                        sh 'make enable-mfa run-obbr-local'
                        sh 'make run-obbr-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('OBBR Tests with enabled MFA failed')
                    } finally {
                        sh 'make clean'
                    }
                }
            }
        }
        stage('SaaS FDX Tests') {
            steps {
                script {
                    try {
                        sh 'make disable-mfa set-saas-configuration run-fdx-saas'
                        sh 'make run-saas-fdx-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('SaaS FDX Tests failed')
                    } finally {
                        sh 'make clean-fdx-saas'
                    }
                }
            }
        }
        stage('SaaS OBUK Tests') {
            steps {
                script {
                    try {
                        sh 'make disable-mfa set-saas-configuration run-obuk-saas'
                        sh 'make run-saas-obuk-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('SaaS OBUK Tests failed')
                    } finally {
                        sh 'make clean-obuk-saas'
                    }
                }
            }
        }
        stage('SaaS OBBR Tests') {
            steps {
                script {
                    try {
                        sh 'make disable-mfa set-saas-configuration run-obbr-saas'
                        sh 'make run-saas-obbr-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('SaaS OBBR Tests failed')
                    } finally {
                        sh 'make clean-obbr-saas'
                    }
                }
            }
        }
    }

    post {
        failure {
            script {
                captureCypressArtifacts()
                if (env.BRANCH_NAME=='master') {
                    sendSlackNotification(currentBuild.result, '#epic-open-banking-improvements', '', true)
                }
            }
        }

        unstable {
            script {
                captureCypressArtifacts()
                if (env.BRANCH_NAME=='master') {
                    sendSlackNotification(currentBuild.result, '#epic-open-banking-improvements', '', true)
                }
            }
        }

        fixed {
            script {
                if (env.BRANCH_NAME=='master') {
                    sendSlackNotification(currentBuild.result, '#epic-open-banking-improvements', '', true)
                }
            }
        }

        cleanup {
            script {
                if (getContext(hudson.FilePath)) {
                    deleteDir()
                }
            }
        }
    }
}

void captureDockerLogs() {
    sh 'rm -rf logs'
    sh 'mkdir logs'
    sh '''#!/bin/bash
        SERVICE_LIST=($(docker ps --format {{.Names}}))
        echo "Service list is ${SERVICE_LIST[*]}"
        for service in "${SERVICE_LIST[@]}"; do
        # Skip null items
        if [ -z "$service" ]; then
            continue
        fi
        echo "Service is $service"
        if [[ $(docker ps | grep "$service" | wc -c) -ne 0 ]]; then
            docker logs "$service" >"logs"/"$service".log 2>&1
        else
            echo "Service $service was not present"
        fi
        done
    '''
    sh 'tar -zcvf docker_logs.tar.gz logs'
    archiveArtifacts(artifacts: 'docker_logs.tar.gz', allowEmptyArchive: true)
}

void captureCypressArtifacts() {
    archiveArtifacts(artifacts: 'tests/cypress/screenshots/**/*', allowEmptyArchive: true)
    archiveArtifacts(artifacts: 'tests/cypress/videos/**/*', allowEmptyArchive: true)
}
