# hiarc-cli
A Command Line Interface for Hiarc

make build

make compile

## Usage
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