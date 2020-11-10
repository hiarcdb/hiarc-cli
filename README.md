# hiarc-cli
A Command Line Interface for Hiarc

make build

make compile

## Usage
### Files
```bash
hiarc file create file-1 --name 'file-1.txt' --path '~/Desktop/a-file.txt' --description 'a description' --metadata '{"department": "engineering"}' --storage-service 'aws-us-east-1-bucket-name'
```
```bash
hiarc file get file-1
```
```bash
hiarc file update file-1 --name 'file-1-changed.txt' --description 'a new description' --metadata '{"department": "sales"}'
```
```bash
hiarc file delete file-1
```
```bash
# Only use --name to change the file's name on the local system to which you are downloading
hiarc file download file-1 --path '~/Downloads' --name 'file-1-different-local-name.txt'
```
```bash
hiarc file attach new-key --storage-id 'object-key-in-bucket' --storage-service 'aws-us-east-bucket'
```
```bash
hiarc file copy source-file-key destination-file-key --storage-service 'azure-blob'
```
```bash
hiarc file direct-download file-1
```
```bash
hiarc file direct-upload --expires-in 600 --storage-service 'aws-us-east-bucket'
```
```bash
hiarc file add-version file-1 --name 'file-1.txt' --path './new-a-file-1.txt' --storage-service 'azure-blob'
```
```bash
hiarc file add-user file-1 user-1 CO_OWNER
```
```bash
hiarc file add-group file-1 group-1 READ_ONLY
```
```bash
hiarc file add-classification file-1 classification-1
```
```bash
hiarc file add-retention file-1 retention-1
```
```bash
hiarc file filter file-1 file-2 file-3 --as-user user-1
```
```bash
hiarc file get versions 123
```
```bash
hiarc file get collections 123
```
```bash
hiarc file get retention-policies 123
```
### Users
```bash
hiarc user get user-1
```
```bash
hiarc user get groups user-1
```
```bash
hiarc user get all
```
```bash
hiarc user get current
```
```bash
hiarc user get current groups
```
```bash
hiarc user create user-1 --name "amsxbg" --metadata '{"department": "sales"}'
```
```bash
hiarc user update user-1 --metadata '{"quotaCarrying": true}'
```
```bash
hiarc user find --query '{"prop": "department", "op": "starts with", "value": "sal" }' --query '{"bool": "and"}' --query '{"prop": "quotaCarrying", "op": "=", "value": true}'
```