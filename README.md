go-meerkat
==========

go-meerkat is a Go library for accessing the Meerkat API. This document also
provides details of the API itself.

## Overview
The Meerkat API is fairly simple and self-describing. All endpoints here have
been discovered simply by viewing streams on the web and inspecting the requests.

In some cases, the web server will also show what routes are available if an 
invalid route is specified. Examples of these are shown below.

Some endpoints require a version parameter (`v`) to be specified in the URL.

## API Hosts
There are several hosts that are used in accessing the API. So far, I've found
the following.
  * resources.meerkatapp.co
    * Primary endpoint, used for accessing broadcasts and users
  * social.meerkatapp.co
    * Used for taking actions on a user - follow, report, etc
  * social-cdn.meerkatapp.co
    * Social graph information including following/followers lists
  * channels.meerkatapp.co
    * Shows actions on a stream - comments, restreams
  * cdn.meerkatapp.co
    * Hosts playlists
  * static.meerkatapp.co
    * Static content like profile/cover images

## Endpoints
Retrieving http://resources.meerkatapp.co/broadcasts/asdf shows the
following error and we can see what resources are available for broadcasts.

    Requesting "GET /asdf" on servlet "/broadcasts" but only have:
      * GET /
      * GET /:broadcastId/activities
      * GET /:broadcastId/summary
      * GET /:broadcastId/watchers
      * GET /:scheduleId/schedule

A similar example for /users on the same host.

    Requesting "GET /asdf" on servlet "/users" but only have:
      * GET /:uid/privateProfile
      * GET /:uid/profile
      * GET /leaderboard

And another example on social.meerkatapp.co/users/xxx/asdf?v=2

    Requesting "GET /xxx/asdf" on servlet "/users" but only have:
      * DELETE /:uid
      * DELETE /:uid/fans
      * DELETE /:uid/followers
      * GET /:uid/followers
      * GET /:uid/following
      * GET /:uid/profile
      * GET /:uid/sweepSessions
      * GET /i-am-alive
      * OPTIONS /
      * OPTIONS /:uid
      * OPTIONS /:uid/followers
      * OPTIONS /:uid/following
      * OPTIONS /:uid/twitterFriendsSignUp
      * OPTIONS /search
      * POST /
      * POST /:uid/fans
      * POST /:uid/followers
      * POST /:uid/invites
      * POST /:uid/reports
      * POST /:uid/socialLinks
      * POST /:uid/twitterFriends
      * POST /:uid/twitterFriendsSignUp
      * PUT /
      * PUT /:uid
      * PUT /:uid/profile
      * PUT /available
      * PUT /search

## Meerkat Resources

### User Search
Meerkat just added [just added](http://thenextweb.com/apps/2015/03/18/meerkat-makes-it-easier-to-follow-users-via-the-web/) user search. 
The endpoint is https://social.meerkatapp.co/users/search?v=0.1 and requires a content-type header.

`curl 'https://social.meerkatapp.co/users/search?v=2' -X PUT -H 'Content-Type: application/json' -d '{"username":"redbull"}'`

    {
      "result": [
        "550099452400006f00a5277f"
      ]
    }

### Broadcast
A Meerkat broadcast http://resources.meerkatapp.co/broadcasts/36aab946-661a-47af-a249-5d395659fd2c/summary

Broadcasts are a Meerkat video. They contain various information about the 
video, including broadcaster, statistics, location and tweet ID.

    {
      "result": {
        "broadcaster": {
          "displayName": "Red Bull",
          "name": "redbull",
          "image": "https://static.meerkatapp.co/users/550099452400006f00a5277f/profile-tmb",
          "id": "550099452400006f00a5277f",
          "profile": "https://resources.meerkatapp.co/users/550099452400006f00a5277f/profile"
        },
        "location": "Aspen, United States",
        "likesCount": 79,
        "coverImages": [
          "https://static.meerkatapp.co/broadcasts/36aab946-661a-47af-a249-5d395659fd2c/compressed1"
        ],
        "activities": [
          {
            "content": "c:2015-03-14T18:31:12.899Z,54f4c98f310000b1016d9745,ryanmabbott,Snow sounds crispy. Chill.",
            "profile": "https://static.meerkatapp.co/users/54f4c98f310000b1016d9745/profile"
          },
          ...
        ],
        "commentsCount": 46,
        "restreamsCount": 39,
        "endTime": 1426359163751,
        "watchersCount": 384,
        "cover": "https://static.meerkatapp.co/broadcasts/36aab946-661a-47af-a249-5d395659fd2c/cover",
        "id": "36aab946-661a-47af-a249-5d395659fd2c",
        "fans": [],
        "influencers": [
          "l:5504814922000022000afb4e",
          "l:54eda5683f0000ed01af56e1",
          "h:DoublePipe",
          ...
        ],
        "caption": "Behind the scenes at #DoublePipe doubles jam. Catch the show on @redbulltv",
        "status": "ended",
        "tweetId": 576812556188151808,
        "place": ""
      },
      "followupActions": {
        "playlist": "http://cdn.meerkatapp.co/broadcast/36aab946-661a-47af-a249-5d395659fd2c/live.m3u8",
        "encore": "https://social.meerkatapp.co/users/550099452400006f00a5277f/fans?v=2",
        "activities": "https://resources.meerkatapp.co/broadcasts/36aab946-661a-47af-a249-5d395659fd2c/activities",
        "likes": "https://channels.meerkatapp.co/broadcasts/36aab946-661a-47af-a249-5d395659fd2c/likes",
        "delete": "https://channels.meerkatapp.co/broadcasts/36aab946-661a-47af-a249-5d395659fd2c",
        "restreams": "https://channels.meerkatapp.co/broadcasts/36aab946-661a-47af-a249-5d395659fd2c/restreams",
        "watchers": "https://resources.meerkatapp.co/broadcasts/36aab946-661a-47af-a249-5d395659fd2c/watchers",
        "comments": "https://channels.meerkatapp.co/broadcasts/36aab946-661a-47af-a249-5d395659fd2c/comments"
      }
    }

There's some sort of list of broadcast here, maybe trending?: http://channels.meerkatapp.co/broadcasts/

### Profile
A Meerkat user profile https://resources.meerkatapp.co/users/550099452400006f00a5277f/profile?v=2

User information including name, list of streams, and various stats.
Included also is a _score_ of some sort.

One thing to note is that most `endTime`'s are in milliseconds since epoch,
with the exception of scheduled broadcasts. In those cases, they are seconds.

    {
      "result": {
        "info": {
          "id": "550099452400006f00a5277f",
          "username": "redbull",
          "displayName": "Red Bull",
          "twitterId": "17540485",
          "privacy": "public",
          "bio": "#givesyouwings"
        },
        "stats": {
          "streams": [
            {
              "id": "60dce137-5c49-4040-b650-6460df235bf7",
              "endTime": 1426104542049
            },
            {
              "id": "36aab946-661a-47af-a249-5d395659fd2c",
              "endTime": 1426359163751
            }
          ],
          "streamsCount": 2,
          "followingCount": 903,
          "followersCount": 3831,
          "score": 1680
        }
      },
      "followupActions": {
        "profileThumbImage": "https://static.meerkatapp.co/users/550099452400006f00a5277f/profile-tmb",
        "profileImage": "https://static.meerkatapp.co/users/550099452400006f00a5277f/profile",
        "followers": "https://social-cdn.meerkatapp.co/users/550099452400006f00a5277f/followers?v=2",
        "following": "https://social-cdn.meerkatapp.co/users/550099452400006f00a5277f/following?v=2",
        "report": "https://social.meerkatapp.co/users/550099452400006f00a5277f/reports?v=2",
        "follow": "https://social.meerkatapp.co/users/550099452400006f00a5277f/followers?v=2"
      }
    }