# Application to synchronize events from Facebook to Google

-------------------------------------------------------------
## Description
-------------------------------------------------------------

Using this application, one can copy or synchronize all events from Facebook to Google. In the demo version, application just fetches events data(name, description, id) from Facebook.

-------------------------------------------------------------
## Build
-------------------------------------------------------------

1. Install and configure go in your system https://golang.org/doc/install
2. Get and save somewhere safe, your Facebook user access token https://developers.facebook.com/tools/explorer . Remember that token has limited period of usage, so don't forget to update it. (A TEMPORARY SOLUTION FOR DEMO VERSION)
3. Copy repository in your GOPATH, cd to the working directory and build the project by using go tool: ***go build***
4. Run an executable file with ***-token <fb_access_token>*** parameter and (optional) ***-group <group_id>***.

-------------------------------------------------------------
## ETC
-------------------------------------------------------------
Application was created for upwork.com project https://www.upwork.com/jobs/Facebook-quot-Page-quot-Everybody-Event-Post-Sync-with-Google-Calendar_~0176ebdb97be9a8c3d

This is just a demo.