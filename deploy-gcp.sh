#!/bin/bash

export PROJECT_ID=savvy-pad-350214
export REGION=us-central1

gcloud builds submit \
  --tag gcr.io/$PROJECT_ID/go-test1 \
  --project $PROJECT_ID

gcloud run deploy go-test1 \
  --image gcr.io/$PROJECT_ID/go-test1 \
  --port 8080 \
  --platform managed \
  --region $REGION \
  --allow-unauthenticated \
  --project $PROJECT_ID
