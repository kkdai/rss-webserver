rss-webserver:
==============

 [![GoDoc](https://godoc.org/github.com/kkdai/rss-webserver?status.svg)](https://godoc.org/github.com/kkdai/rss-webserver)  [![Build Status](https://travis-ci.org/kkdai/rss-webserver.svg?branch=master)](https://travis-ci.org/kkdai/rss-webserver)





How to use it
=============

Here is an exist Heroku Server which you can try directly.

```

//Get Github Starred RSS by user "kkdai"
http://githubrss.herokuapp.com/starred?kkdai

//Get Github Follower RSS by user "kkdai"
http://githubrss.herokuapp.com/follower?kkdai


//Get Github Following RSS by user "kkdai"
http://githubrss.herokuapp.com/following?kkdai

```


Publish to Heroku
=============

```
//Login heroku
heroku login

//Using Golang backpack
heroku create -b https://github.com/kr/heroku-buildpack-go.git

//Push your complete code to Heroku
git push heroku master
```


Add "Deploy to Heroku"
=============

- copy [app.json](https://raw.githubusercontent.com/kkdai/rss-webserver/master/app.json)
- Remember must include [buildpack](https://devcenter.heroku.com/articles/app-json-schema#buildpacks).


Installation and Usage
=============


[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

Inspired By
=============

- [Heroku Deploy](https://devcenter.heroku.com/articles/heroku-button)
- [Heroku deploy Go](http://dougblack.io/words/a-restful-micro-framework-in-go.html)


Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

