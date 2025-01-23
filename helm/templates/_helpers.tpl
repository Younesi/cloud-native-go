{{- define "cloud-native-go.fullname" -}}
{{- .Release.Name -}}
{{- end }}

{{- define "cloud-native-go.labels" -}}
app: {{ include "cloud-native-go.fullname" . }}
chart: "{{ .Chart.Name }}-{{ .Chart.Version }}"
appVersion: "{{ .Chart.AppVersion }}"
release: "{{ .Release.Name }}"
{{- end }}