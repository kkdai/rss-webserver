// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package githubrss

const AtomTmpl string = `
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
  <channel>
    <title>{{ .RssTitle }}</title>
    <link>{{ .RssLink }}</link>
    <atom:link href="{{ .RssLink }}" rel="self" type="application/rss+xml" />
    <description>{{ .RssDesc }}</description>
    <pubDate>{{ .RssNow }}</pubDate>
    <lastBuildDate>{{ .RssNow }}</lastBuildDate>

{{ range .Items }}
      <item>
        <title>{{ .Title }}</title>
        <description>{{ .Desc }}</description>
        <pubDate>{{ .PubDate }}</pubDate>
        <guid>{{ .Link }}</guid>
        <link>{{ .Link }}</link>
      </item>
{{ end }}      

  </channel>
</rss>
`
