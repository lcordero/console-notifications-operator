#!/bin/sh

if [ $# -ne 1 ]; then
  echo -e "Ignoring generate k8s and CRDs\n"
else
  echo -e "Running FULL\n"
  # ======= BEFORE TEST ======
  # Generate K8s
  operator-sdk generate k8s
  # Generate CRDs
  operator-sdk  generate crds

  oc delete -f deploy/crds/operator.openshift.io_consolenotifications_crd.yaml
fi


oc create -f deploy/crds/operator.openshift.io_consolenotifications_crd.yaml

# ======== RUN FROM LOCAL ====
oc delete -f deploy/test-notification.yaml
oc create -f deploy/test-notification.yaml

operator-sdk run --local --namespace console-notifications --kubeconfig kubeconfig



# WORK NOTES
# El Operador debe crear el POD y la NOTIFICATION
# El POD haría:
#       Revisar a diario si la alerta ya está vencida
#       Si está vencida debe ir a ConsoleNotification (Del operator) y cambiarle el active a false
#         con eso se activa el Reconcile y se elimina el POD y el NOTIFICATION

