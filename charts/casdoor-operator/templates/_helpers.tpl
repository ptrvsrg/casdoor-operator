{{/*
Expand the name of the chart.
*/}}
{{- define "casdoor-operator.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "casdoor-operator.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create a image name
*/}}
{{- define "casdoor-operator.image" -}}
{{- printf "%s:%s" (join "/" (list .Values.image.registry .Values.image.repository .Values.image.name)) (.Values.image.tag | default .Chart.AppVersion) }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "casdoor-operator.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "casdoor-operator.labels" -}}
helm.sh/chart: {{ include "casdoor-operator.chart" . }}
{{ include "casdoor-operator.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "casdoor-operator.selectorLabels" -}}
app.kubernetes.io/name: {{ include "casdoor-operator.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- with .Values.selectorLabels }}
{{- toYaml . | nindent 0 }}
{{- end }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "casdoor-operator.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "casdoor-operator.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
