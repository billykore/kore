package tpl

func DeploymentTemplate() []byte {
	return []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: monorepo-{{ .ServiceName }}-app
  labels:
    name: monorepo-{{ .ServiceName }}-app
spec:
  replicas: 1
  selector:
    matchLabels:
      name: monorepo-{{ .ServiceName }}-service
  template:
    metadata:
      name: monorepo-{{ .ServiceName }}-service
      labels:
        name: monorepo-{{ .ServiceName }}-service
    spec:
      containers:
        - name: monorepo-{{ .ServiceName }}-app
          image: billykore/monorepo-{{ .ServiceName }}:latest
          imagePullPolicy: Always
          envFrom:
            - configMapRef:
                name: monorepo-{{ .ServiceName }}-env
          resources:
            requests:
              cpu: 100m
              memory: 100Mi
            limits:
              cpu: 300m
              memory: 300Mi
      restartPolicy: Always

---

apiVersion: v1
kind: Service
metadata:
  name: monorepo-{{ .ServiceName }}-service
spec:
  selector:
    name: monorepo-{{ .ServiceName }}-service
  ports:
    - name: "http"
      port: 8000
      targetPort: 8000
    - name: "grpc"
      port: 9000
      targetPort: 9000
  type: ClusterIP
`)
}

func EnvTemplate() []byte {
	return []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  name: monorepo-{{ .ServiceName }}-env
data:
  HTTP_PORT: "8000"
  GRPC_PORT: "9000"
  PROJECT_ID: "todo-list-app-a5058"
  FIRESTORE_SDK: "firebase-sdk.json"
`)
}
