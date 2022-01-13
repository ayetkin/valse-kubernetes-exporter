{{/*
Expand the name of the chart.
*/}}
{{- define "valse.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "valse.fullname" -}}
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
{{- define "valse.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "valse.labels" -}}
helm.sh/chart: {{ include "valse.chart" . }}
{{ include "valse.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "valse.selectorLabels" -}}
app.kubernetes.io/name: {{ include "valse.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app: {{ include "valse.name" . }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "valse.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "valse.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "valse.configmapName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "valse.fullname" .) .Values.configmap.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "valse.rbacName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "valse.fullname" .) .Values.rbac.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "valse.secretName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "valse.fullname" .) .Values.secret.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}