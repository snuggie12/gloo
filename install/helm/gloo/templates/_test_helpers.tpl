{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}

{{- define "gloo.testcontainer" -}}
- name: nginx
  image: nginx:1.7.9
  ports:
  - containerPort: 80
{{- end }}

{{/*
Test that extraListenersHelper is rendered correctly
*/}}
{{- define "gloo.testlistener" -}}
- name: test_listener
{{- end -}}

{{- define "gloo.testIstioVolume" -}}
- hostPath:
    path: /var/run/sds
  name: sds-uds-path
{{- end }}

{{- define "gloo.testIstioVolumeMount" -}}
- mountPath: /var/run/sds
  name: sds-uds-path
{{- end }}
