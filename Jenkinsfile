def rtDocker;
def buildInfo;

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
    parameters {
        booleanParam(name: 'RUN_XRAY_SCAN', defaultValue: false, description: 'Check this option if you want to run Xray Scan for vulnerabilities in artifacts.')
    }
    environment {
        COMPOSE_HTTP_TIMEOUT = 120 
        VERIFY_TEST_RUNNER_TIMEOUT_MS = 80000
        SAAS_TENANT_ID = 'amfudxn6-qa-us-east-1-ob-quickstart'
        SAAS_CLIENT_ID = credentials('OPENBANKING_CONFIGURATION_CLIENT_ID')
        SAAS_CLIENT_SECRET = credentials('OPENBANKING_CONFIGURATION_CLIENT_SECRET')
        SAAS_CLEANUP_CLIENT_ID = credentials('OPENBANKING_CLEANUP_CLIENT_ID')
        SAAS_CLEANUP_CLIENT_SECRET = credentials('OPENBANKING_CLEANUP_CLIENT_SECRET')
        NOTIFICATION_CHANNEL = credentials('OPENBANKING_NOTIFICATION_CHANNEL')
        JENKINS_AUTH_USER = credentials('JENKINS_AUTH_USER')
        JENKINS_AUTH_TOKEN = credentials('JENKINS_AUTH_TOKEN')
        DEBUG = 'true'
    }
    stages {
        stage('Prepare') {
            steps {
                script{
                    if (env.BRANCH_NAME.startsWith('PR-')) {
                        abortPreviousRunningBuilds()
                    }
                    if (params.RUN_XRAY_SCAN == true) {
                        artifactory = initArtifactoryServer()
                        rtServer = artifactory[0]
                        rtDocker = artifactory[1]
                        buildInfo = artifactory[2]
                    }
                }
                sh '''#!/bin/bash
                        echo "127.0.0.1       authorization.cloudentity.com test-docker" | sudo tee -a /etc/hosts
                        echo "127.0.0.1       mock-data-recipient" | sudo tee -a /etc/hosts
                        cd tests && yarn install
                '''
                sh 'docker-compose version'
                sh "docker rm -f \$(docker ps -aq) || true"
                 
                retry(3) {
                    sh "make run-tests-verify"
                }
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
        stage("Xray Scan") {
            when {
                expression {
                    params.RUN_XRAY_SCAN == true
                }
            }
            steps {
                script {
                    sh 'make pull-docker-images'
                    sh 'docker images'
                    sh 'make retag-docker-images'
                    dockerList = sh(
                        script: """
                                make -s list-docker-images | sed \"s/cloudentity\\//docker.cloudentity.io\\//g\"
                                """,
                        returnStdout: true
                    ).trim()
                    images = dockerList.split("\n")
                    pushCommits(rtDocker, buildInfo, images, "")
                    buildInfo.env.collect()
                    rtServer.publishBuildInfo buildInfo
                    scanConfig = [
                        'buildName'   : buildInfo.name,
                        'buildNumber' : buildInfo.number,
                        'failBuild'   : false
                    ]
                    scanResult = rtServer.xrayScan scanConfig
                    if (scanResult.foundVulnerable) {
                        scanresult = scanResult.toString()
                        writeFile(file: '/tmp/scanresult.json', text: scanresult)
                        env.XRAY_SCAN_TABLE = sh(
                            script: './scripts/format_xray_result.sh',
                            returnStdout: true
                        ).trim()
                        env.VULNERABILITIES = true
                        currentBuild.result = 'UNSUCCESSFUL'
                    }
                }
            }
        }
        stage('Unit tests') {
            steps {
                sh 'make test'
            }
        }
        stage('CDR Tests') {
            steps {
                script {
                    sh 'make clean'
                    try {
                        sh 'make run-cdr-local'
                        sh 'make run-cdr-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('CDR Tests failed')
                    }
                }
            }
        }
        stage('FDX Tests with disabled MFA') {
            steps {
                script {
                    sh 'make clean'
                    try {
                        sh 'make disable-mfa run-fdx-local'
                        sh 'make run-fdx-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('FDX Tests with disabled MFA failed')
                    }
                }
            }
        }
        stage('FDX Tests with enabled MFA') {
            steps {
                script {
                    sh 'make clean'
                    try {
                        sh 'make enable-mfa run-fdx-local'
                        sh 'make run-fdx-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('FDX Tests with enabled MFA failed')
                    }
                }
            }
        }
        stage('OBUK Tests with disabled MFA') {
            steps {
                script {
                    sh 'make clean'
                    try {
                        sh 'make disable-mfa run-obuk-local'
                        sh 'make run-obuk-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('OBUK Tests with disabled MFA failed')
                    }
                }
            }
        }
        stage('OBUK Tests with enabled MFA') {
            steps {
                script {
                    sh 'make clean'
                    try {
                        sh 'make enable-mfa run-obuk-local'
                        sh 'make run-obuk-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('OBUK Tests with enabled MFA failed')
                    }
                }
            }
        }
        stage('OBBR Tests with disabled MFA') {
            steps {
                script {
                    sh 'make clean'
                    try {
                        sh 'make disable-mfa run-obbr-local'
                        sh 'make run-obbr-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('OBBR Tests with disabled MFA failed')
                    }
                }
            }
        }
        stage('OBBR Tests with enabled MFA') {
            steps {
                script {
                    sh 'make clean'
                    try {
                        sh 'make enable-mfa run-obbr-local'
                        sh 'make run-obbr-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('OBBR Tests with enabled MFA failed')
                    }
                }
            }
        }
        stage('SaaS FDX Tests') {
            steps {
                script {
                    sh 'make clean-saas'
                    try {
                        sh 'make disable-mfa run-fdx-saas'
                        sh 'make run-saas-fdx-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('SaaS FDX Tests failed')
                    }
                }
            }
        }
        stage('SaaS OBUK Tests') {
            steps {
                script {
                    sh 'make clean-saas'
                    try {
                        sh 'make disable-mfa run-obuk-saas'
                        sh 'make run-saas-obuk-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('SaaS OBUK Tests failed')
                    }
                }
            }
        }
        stage('SaaS OBBR Tests') {
            steps {
                script {
                    sh 'make clean-saas'
                    try {
                        sh 'make disable-mfa run-obbr-saas'
                        sh 'make run-saas-obbr-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('SaaS OBBR Tests failed')
                    }
                }
            }
        }
        stage('SaaS CDR Tests') {
            steps {
                script {
                    sh 'make clean-saas'
                    try {
                        sh 'make disable-mfa run-cdr-saas'
                        sh 'make run-saas-cdr-tests-headless'
                    } catch(exc) {
                        captureDockerLogs()
                        unstable('SaaS CDR Tests failed')
                    }
                }
            }
        }
    }

    post {
        always {
            sh "make clean-saas"
        }
        unsuccessful {
            script {
                if (env.VULNERABILITIES) {
                    env.SLACK_MESSAGE = """
                                        |:warning: JFrog Xray scan found security vulnerabilities
                                        |${BUILD_URL}
                                        |```
                                        |${XRAY_SCAN_TABLE}
                                        |```
                                        |To view the vulnerabilities: click on the build link, then 'Xray Scan Report', then Violations tab.
                                        |Based on your judgement, fix the vulnerabilities or configure them to be ignored.
                                        """.stripMargin()
                    sendSlackNotification('UNSUCCESSFUL', '#proj-openbanking', "${SLACK_MESSAGE}", false, false, '', true)
                }
            }
        }
        failure {
            script {
                captureCypressArtifacts()
                if (env.BRANCH_NAME=='master') {
                    sendSlackNotification(currentBuild.result, NOTIFICATION_CHANNEL, '', true)
                }
            }
        }

        unstable {
            script {
                captureCypressArtifacts()
                if (env.BRANCH_NAME=='master') {
                    sendSlackNotification(currentBuild.result, NOTIFICATION_CHANNEL, '', true)
                }
            }
        }

        fixed {
            script {
                if (env.BRANCH_NAME=='master') {
                    sendSlackNotification(currentBuild.result, NOTIFICATION_CHANNEL, '', true)
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
