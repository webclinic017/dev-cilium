analytics_settings(False)

load('ext://uibutton', 'cmd_button', 'location')

allow_k8s_contexts('pod')

local_resource('argocd', cmd = 'echo; argocd app sync argocd --local k/argocd --preview-changes --dry-run --assumeYes 2>/dev/null | awk "/Previewing/ {flag=1;next} /^TIMESTAMP/ {flag=0} flag"; echo', deps = ['k/argocd'])
