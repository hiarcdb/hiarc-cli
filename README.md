# hiarc-cli
A Command Line Interface for Hiarc

make build

make compile

## Usage
### Files
```bash
hiarc file create file-1 --name 'file-1.txt' --path ~/Desktop/a-file.txt --description 'a description' --metadata '{"department": "engineering"}' --storage-service 'aws-us-east-1-bucket-name'
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
hiarc file download file-1 --path ~/Downloads --name 'file-1-different-local-name.txt'
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
hiarc file add-version file-1 --name 'file-1.txt' --path ./new-a-file-1.txt --storage-service 'azure-blob'
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
### Collections
```bash
hiarc collection create collection-1 --name 'collection 1' --description 'a collection of files and children' --metadata '{"department": "marketing"}'
```
```bash
hiarc collection update collection-1 --metadata '{"important": true, "cost": 50000}'
```
```bash
hiarc collection get collection-1
```
```bash
hiarc collection get all
```
```bash
hiarc collection get children collection-1
```
```bash
hiarc collection get files collection-2
```
```bash
hiarc collection get items collection-1
```
```bash
hiarc collection add-user collection-1 user-1 read_only
```
```bash
hiarc collection add-group collection-1 group-1 co_owner
```
```bash
hiarc collection add-file collection-1 file-1
```
```bash
hiarc collection add-child collection-1 collection-2
```
```bash
hiarc collection find --query '{"prop": "department", "op": "starts with", "value": "mark" }' --query '{"bool": "and"}' --query '{"prop": "cost", "op": ">", "value": 1000}'
```
```bash
hiarc collection remove-file collection-1 file-1
```
```bash
hiarc collection delete collection-1
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
### Groups
```bash
hiarc group create group-1 --name "group-1" --metadata '{"department": "sales"}'
```
```bash
hiarc group get group-1
```
```bash
hiarc group get all
```
```bash
hiarc group get all current --as-user user-1
```
```bash
hiarc user get for-user user-1
```
```bash
hiarc user update group-1 --metadata '{"quotaCarrying": true}'
```
```bash
hiarc group add-user group-1 user-1
```
```bash
hiarc group find --query '{"prop": "department", "op": "starts with", "value": "sal" }' --query '{"bool": "and"}' --query '{"prop": "quotaCarrying", "op": "=", "value": true}'
```
```bash
hiarc group delete group-1
```