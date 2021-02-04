#!/bin/groovy

// {{.Name}}

import jenkins.*
import jenkins.model.*
import hudson.*
import hudson.model.*
import env.common

/* Global Variables */
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
}