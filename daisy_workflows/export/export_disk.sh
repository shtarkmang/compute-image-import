#!/bin/bash
# Copyright 2017 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

function serialOutputPrefixedKeyValue() {
  stdbuf -oL echo "$1: <serial-output key:'$2' value:'$3'>"
}

# Verify VM has access to Google APIs
curl --silent --fail "https://www.googleapis.com/discovery/v1/apis" &> /dev/null;
if [[ $? -ne 0 ]]; then
  echo "ExportFailed: Cannot access Google APIs. Ensure that VPC settings allow VMs to access Google APIs either via external IP or Private Google Access. More info at: https://cloud.google.com/vpc/docs/configure-private-google-access"
  exit
fi

BYTES_1GB=1073741824
URL="http://169.254.169.254/computeMetadata/v1/instance/attributes"
GCS_PATH=$(curl -f -H Metadata-Flavor:Google ${URL}/gcs-path)
LICENSES=$(curl -f -H Metadata-Flavor:Google ${URL}/licenses)

mkdir ~/upload

SOURCE_DEVICE=$(readlink -f /dev/disk/by-id/google-disk-image-export)
# Source disk size info.
SOURCE_SIZE_BYTES=$(lsblk $SOURCE_DEVICE --output=size -b | sed -n 2p)
SOURCE_SIZE_GB=$(awk "BEGIN {print int(((${SOURCE_SIZE_BYTES}-1)/${BYTES_1GB}) + 1)}")
serialOutputPrefixedKeyValue "GCEExport" "source-size-gb" "${SOURCE_SIZE_GB}"

echo "GCEExport: Running export tool."
if [[ -n $LICENSES ]]; then
  gce_export -buffer_prefix ~/upload -gcs_path "$GCS_PATH" -disk "$SOURCE_DEVICE" -licenses "$LICENSES" -y
else
  gce_export -buffer_prefix ~/upload -gcs_path "$GCS_PATH" -disk "$SOURCE_DEVICE" -y
fi
GCE_EXPORT_RESULT=$?
if [[ ${GCE_EXPORT_RESULT} -eq 2 ]]; then
  echo "ExportFailed: Failed to export disk source to GCS because default Compute Engine Service account does not have storage.objects.create permission for [Privacy-> ${GCS_PATH}. <-Privacy]."
  exit 1
fi
if [[ ${GCE_EXPORT_RESULT} -ne 0 ]]; then
  echo "ExportFailed: Failed to export disk source to GCS [Privacy-> ${GCS_PATH}. <-Privacy]."
  exit 1
fi

# Exported image size info.
TARGET_SIZE_BYTES=$(gsutil ls -l "${GCS_PATH}" | head -n 1 | awk '{print $1}')
TARGET_SIZE_GB=$(awk "BEGIN {print int(((${TARGET_SIZE_BYTES}-1)/${BYTES_1GB}) + 1)}")
serialOutputPrefixedKeyValue "GCEExport" "target-size-gb" "${TARGET_SIZE_GB}"

echo "ExportSuccess"
sync
