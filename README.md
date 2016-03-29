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

**note**: If you occur any problem with connect to [http://githubrss.herokuapp.com](http://githubrss.herokuapp.com), it is not code problem. Because I use Heroku free account and the quota is full because IFTTT query.  XD

Publish to Heroku
=============

- Init git and commit your code.
	- `git init`
	- `git add .`
	- `git commit -m "init project"`
- Login heroku
  - `heroku login`
- Using Golang backpack
  - `heroku create -b https://github.com/kr/heroku-buildpack-go.git`

- Vendoring: For Go 1.6 (refer [here](https://devcenter.heroku.com/articles/go-support#migrating-from-go1-5-godep-workspace-to-go1-6-with-a-vendor-directory))
   - `go get -u github.com/tools/godep`
   - Save your dependency `godep save github.com/kkdai/githubrss`
   - Submit your depdendency 	
	   - `git add .`
	   - `git commit -m "godep"
- Push your complete code to Heroku
  - `git push heroku master`




Add "Deploy to Heroku"
=============

- copy [app.json](https://raw.githubusercontent.com/kkdai/rss-webserver/master/app.json) (refer more spec [here](https://devcenter.heroku.com/articles/app-json-schema#buildpacks))
- Remember must include [buildpack](https://devcenter.heroku.com/articles/app-json-schema#buildpacks). And all buildpack list [here](https://devcenter.heroku.com/articles/buildpacks#officially-supported-buildpacks).


Installation and Usage
=============


[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

Inspired By
=============

- [Github Status RSS feeder](https://github.com/kkdai/githubrss)
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

