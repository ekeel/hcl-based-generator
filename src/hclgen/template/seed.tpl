#!/bin/groovy

// {{.Name}}

import jenkins.*
import jenkins.model.*
import hudson.*
import hudson.model.*
import env.common

/* Global Variables */
String GitCredentialID = '{{.GitCredentialID}}'
{{range .ConstantStrings}}{{.}}
{{end}}
/* Shared Configuration */
def sharedConfig = [
{{range .SharedConfig.Variables}}	{{.Name}}: '{{.Value}}',
{{end}}]

/* Job Configuration */
def jobs = [
{{range .Jobs}}	[
		jobName: '{{.Name}}',
		scriptPath: '{{.ScriptPath}}',
		description: '{{.Description}}',

		{{range .Variables}}{{.Name}}: '{{.Value}}',
		{{end}}
		{{range .Environments}}{{.Name}}: [
			{{range .Variables}}{{.Name}}: '{{.Value}}',
			{{end}}
		],
		{{end}}
		view: [
			name: '{{.View.Name}}'
			description: '{{.View.Description}}'
		]
	],
{{end}}
]

jobs.each { jobConfig ->
	config = sharedConfig + jobConfig

	pipelineJob(config.jobName) {
		description = config.Description

		definition {
			cpsScm {
				scm {
					git {
						remote {
							name('origin')
							url(config.gitRepo)
							branch(config.gitBranch)
							credentials(GitCredentialID)
						}
						extensions {
							wipeOutWorkspace()
						}
					}
				}
				scriptPath(config.deployScript)
			}
		}
	}

	properties {
		buildDiscarderProperty {
			strategy {
				logRotator {
					numToKeepStr(componentconfig.numToKeep ?: 20)
					daysToKeepStr('')
					artifactDaysToKeepStr('')
					artifactNumToKeepStr('')
				}
			}
		}
	}

	environmentVariables {
		{{/* TODO: add environment variables to config */}}
	}
}