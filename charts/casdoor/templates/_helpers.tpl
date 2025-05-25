{{/*
Expand the name of the chart.
*/}}
{{- define "casdoor.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "casdoor.fullname" -}}
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
Create chart name and version as used by the chart label.
*/}}
{{- define "casdoor.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "casdoor.labels" -}}
helm.sh/chart: {{ include "casdoor.chart" . }}
{{ include "casdoor.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "casdoor.selectorLabels" -}}
app.kubernetes.io/name: {{ include "casdoor.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- with .Values.selectorLabels }}
{{- toYaml . | nindent 0 }}
{{- end }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "casdoor.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "casdoor.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Create dataSourceName used in the configmap
*/}}
{{- define "casdoor.dataSourceName" -}}
{{- $options := list }}
{{- range $key, $value := .Values.database.options }}
{{- $options = append $options (printf "%s=%v" $key $value) }}
{{- end }}
{{- $optionsRaw := "" }}
{{- if or (eq .Values.database.driver "mysql") (eq .Values.database.driver "sqlite") }}
{{- $optionsRaw = join "&" $options }}
{{- else }}
{{- $optionsRaw = join " " $options }}
{{- end }}
{{- if eq .Values.database.driver "mysql" -}}
{{- .Values.database.username }}:{{ .Values.database.password }}@tcp({{ .Values.database.host }}:{{ default "3306" .Values.database.port }})/{{ .Values.database.dbName }}?{{ $optionsRaw }}
{{- else if or (eq .Values.database.driver "postgres") (eq .Values.database.driver "cockroachdb") -}}
"user={{ .Values.database.username }} password={{ .Values.database.password }} host={{ .Values.database.host }} port={{ default "5432" .Values.database.port }} dbname={{ .Values.database.dbName }} {{ $optionsRaw }}"
{{- else if eq .Values.database.driver "sqlite" -}}
file:/data/{{ .Values.database.dbName }}.db?{{ $optionsRaw }}
{{- else -}}
file:/data/casdoor.db?cache=shared
{{- end }}
{{- end }}

{{/*
Create dbName used in the configmap
*/}}
{{- define "casdoor.dbName" -}}
{{- if eq .Values.database.driver "mysql" -}}
{{ .Values.database.databaseName }}
{{- else if eq .Values.database.driver "postgres" -}}
{{- else if eq .Values.database.driver "cockroachdb" -}}
{{- else -}}
{{ .Values.database.databaseName }}
{{- end }}
{{- end }}

{{/*
Create EnvfromConfigmap
*/}}
{{- define "casdoor.envFromConfigmap" -}}
{{- range . -}}
- name: {{ .name }}
  valueFrom:
    configMapKeyRef:
      name: {{ .configmapName }}
      key: {{ .key }}
{{- end -}}
{{- end }}

{{/*
Create EnvfromSecret
*/}}
{{- define "casdoor.envFromSecret" -}}
{{- range . -}}
- name: {{ .name }}
  valueFrom:
    secretKeyRef:
      name: {{ .secretName }}
      key: {{ .key }}
{{- end -}}
{{- end -}}

{{/*
Create Envfrom
*/}}
{{- define "casdoor.envFrom" -}}
{{- range . -}}
{{- if eq .type "configmap" -}}
- configMapRef:
    name: {{ .name }}
{{ end }}
{{- if eq .type "secret" -}}
- secretRef:
    name: {{ .name }}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create a image name
*/}}
{{- define "casdoor.image" -}}
{{- printf "%s:%s" (join "/" (list .Values.image.registry .Values.image.repository .Values.image.name)) (.Values.image.tag | default .Chart.AppVersion) }}
{{- end }}

{{/*
Create databae name
*/}}
{{- define "casdoor.database" -}}
{{- if eq .Values.database.driver "sqlite" }}

{{- else if eq .Values.database.driver "mysql" }}

{{- else if eq .Values.database.driver "postgres" }}

{{- else if eq .Values.database.driver "cockroachdb" }}

{{- else }}

{{- end }}
{{- end }}