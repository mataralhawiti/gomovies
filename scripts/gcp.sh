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
export EMAIL=""
# --
# Cloud Functions
## Create User-managed service account for Cloud Functions
gcloud iam service-accounts create sa-matar-gcp \
    --description="SA Matar GCP" \
    --display-name="cloud functions SA"

## Grant service account an IAM role on the project
gcloud projects add-iam-policy-binding $GCP_PROJECT \
    --member="serviceAccount:sa-matar-gcp@$GCP_PROJECT.iam.gserviceaccount.com" \
    --role="roles/cloudfunctions.developer"

## To allow users to impersonate the service account
gcloud iam service-accounts add-iam-policy-binding \
    sa-matar-gcp@$GCP_PROJECT.iam.gserviceaccount.com \
    --member="user:$EMAIL" \
    --role="roles/iam.serviceAccountUser"