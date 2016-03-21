# Jsonbeat

Jsonbeat is a [Beat](https://www.elastic.co/products/beats) used for
reading json events (on a single line) through STDIN, and outputting
the events through the beats framework.  This is kind of like
filebeats, except Jsonbeat will parse json events, and send them to
the index and as the event-type that you specify.


## Edit the configuration file

## Build and Run the program

### Configure your environment for Google Container Engine

```
# login to google cloud
gcloud auth login

# disable usage reporting to google
gcloud config set disable_usage_reporting true

# install kubectl
gcloud components install -q kubectl

# locate your project id
gcloud projects list

# set default project
gcloud config set core/project <project_id>

# locate your zone and cluster
gcloud container clusters list

# set default zone
gcloud config set compute/zone <zone_id>

# set default cluster
gcloud config set container/cluster <cluster_id>

# configure kubectl in google's environment
gcloud container clusters get-credentials <cluster_id>

# check your config
gcloud config list
```

### App Build and Run Locally

```
make app-build
make app-run
```

### Docker Build and Run Locally

```
make docker-build
make docker-run
```
