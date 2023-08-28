export GOOGLE_CLOUD_PROJECT=<PROJECT_ID>
export AUTH_TOKEN=$(gcloud auth print-access-token)
export INSTANCE_ID=your-instance-name
export REGION=us-central1
export CDAP_ENDPOINT=$(gcloud beta data-fusion instances describe \
    --location=$REGION \
    --format="value(apiEndpoint)" \
  ${INSTANCE_ID})

# --
export GCP_PROJECT=golang-389808
export GCP_DATASET=gomovies
export GCP_TABLE=movies