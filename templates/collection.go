// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package templates

const collectionTemplate = `<!DOCTYPE HTML>
<html lang="{{.LanguageTag}}">
<meta charset="utf-8">
<head>
	<title>{{.Title}}</title>
	<link rel="stylesheet" type="text/css" href="/theme/screen.css">
</head>

<body>

<article>

<header>
<h1>
{{.Title}}
</h1>
</header>

<section>
{{.Description}}
</section>

<section>
{{.Content}}
</section>

<section>
<ul>
{{range .Entries}}
<li>
	<a href="{{.Path}}">{{.Title}}</a>
	<p>{{.Description}}</p>
</li>
{{end}}
</ul>
</section>

</article>

</body>
</html>`
