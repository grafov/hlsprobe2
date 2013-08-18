// Templates for webpages
package main

const ReportMainPageTemplate =`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>HLS Probe II Reports</title>
<meta name="description" content="Creating a table with Twitter Bootstrap. Learn how to use Twitter Bootstrap toolkit to create Tables with examples.">
<link href="/css/bootstrap.min.css" rel="stylesheet">
<link href="/css/custom.css" rel="stylesheet">
</head>
<body><h1>Reports are:</h1>
 <ul>
  <li><a href="rprt/3hours">Errors for all streams for last 3 hours</a></li>
  <li><a href="rprt/last">Last streams errors (complete)</a></li>
  <li><a href="rprt/last-critical">Last streams errors (critical only)</a></li>
 </ul>
</body>
</html>
`

const Report3HoursTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>Errors per 3 hours for all groups :: HLS Probe II</title>
<meta name="description" content="Creating a table with Twitter Bootstrap. Learn how to use Twitter Bootstrap toolkit to create Tables with examples.">
<link href="/css/bootstrap.min.css" rel="stylesheet">
<link href="/css/custom.css" rel="stylesheet">
<script src="/js/bootstrap.min.js"></script>
</head>
<body><h1>Errors per 3 hours for all groups</h1>
<table class="table table-bordered table-condensed">
        <thead>
          <tr>
            <th rowspan="2">Group</th>
            <th rowspan="2">Name</th>
            <th colspan="6">Last hour</th>
            <th colspan="6">Two hours ago</th>
            <th colspan="6">Three hours ago</th>
          </tr>
          <tr>
            <th rel="tooltip" title="Slow response">SR</th>
            <th rel="tooltip" title="Bad status">BS</th>
            <th rel="tooltip" title="Bad URI">BU</th>
            <th rel="tooltip" title="Timeout on read">RT</th>
            <th rel="tooltip" title="Timeout on connect">CT</th>
            <th rel="tooltip" title="Slow response">SR</th>
            <th rel="tooltip" title="Bad status">BS</th>
            <th rel="tooltip" title="Bad URI">BU</th>
            <th rel="tooltip" title="Timeout on read">RT</th>
            <th rel="tooltip" title="Timeout on connect">CT</th>
            <th rel="tooltip" title="Slow response">SR</th>
            <th rel="tooltip" title="Bad status">BS</th>
            <th rel="tooltip" title="Bad URI">BU</th>
            <th rel="tooltip" title="Timeout on read">RT</th>
            <th rel="tooltip" title="Timeout on connect">CT</th>
            <th rel="tooltip" title="Connection refused">CR</th>
          </tr>
        </thead>
        <tbody>
        {{#TableData}}
          <tr>
            <td>{{group}}</td>
            <td><a href="{{uri}}">{{name}}</a></td>
            <td{{#sw-severity}} class="{{sw-severity}}"{{/sw-severity}} rel="tooltip" title="Slow response">{{sw}}</td>
            <td{{#bs-severity}} class="{{bs-severity}}"{{/bs-severity}} rel="tooltip" title="Bad status">{{bs}}</td>
            <td{{#bu-severity}} class="{{bu-severity}}"{{/bu-severity}} rel="tooltip" title="Bad URI">{{bu}}</td>
            <td{{#rt-severity}} class="{{rt-severity}}"{{/rt-severity}} rel="tooltip" title="Timeout on read">{{rt}}</td>
            <td{{#ct-severity}} class="{{ct-severity}}"{{/ct-severity}} rel="tooltip" title="Timeout on connect">{{ct}}</td>
            <td{{#cr-severity}} class="{{cr-severity}}"{{/cr-severity}} rel="tooltip" title="Slow response">{{cr}}</td>
            <td>{{sw2}}</td>
            <td>{{bs2}}</td>
            <td>{{bs2}}</td>
            <td>{{rt2}}</td>
            <td>{{ct2}}</td>
            <td>{{cr2}}</td>
            <td>{{sw3}}</td>
            <td>{{bs3}}</td>
            <td>{{bu3}}</td>
            <td>{{rt3}}</td>
            <td>{{ct3}}</td>
            <td>{{cr3}}</td>
          </tr>
        {{/TableData}}
        </tbody>
</table>
Generated by <a href="http://github.com/grafov/hlsprobe2">HLS Probe II</a>
</body>
</html>
`

const ReportLastTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>All groups last errors :: HLS Probe II</title>
<meta name="description" content="Creating a table with Twitter Bootstrap. Learn how to use Twitter Bootstrap toolkit to create Tables with examples.">
<link href="/css/bootstrap.min.css" rel="stylesheet">
<link href="/css/custom.css" rel="stylesheet">
<script src="/js/bootstrap.min.js"></script>
</head>
<body><h1>Last errors</h1>
<table class="table table-bordered table-condensed">
        <thead>
          <tr>
            <th>Group</th>
            <th>Name</th>
            <th>Status</th>
            <th>Length</th>
            <th>Request Duration</th>
            <th>Last Checked</th>
            <th>Error</th>
            <th>Last Hour</th>
          </tr>
        </thead>
        <tbody>
        {{#TableData}}
          <tr class="{{severity}}">
            <td>{{group}}</td>
            <td><a href="{{uri}}">{{name}}</a></td>
            <td>{{status}}</td>
            <td>{{contentlength}}</td>
            <td>{{elapsed}}</td>
            <td>{{started}}</td>
            <td>{{error}}</td>
            <td>{{totalerrs}}</td>
          </tr>
        {{/TableData}}
        </tbody>
</table>
Generated by <a href="http://github.com/grafov/hlsprobe2">HLS Probe II</a>
</body>
</html>
`

const ReportGroupLastTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{#Vars}}{{group}} :: {{/Vars}}HLS Probe II</title>
<meta name="description" content="Creating a table with Twitter Bootstrap. Learn how to use Twitter Bootstrap toolkit to create Tables with examples.">
<link href="/css/bootstrap.min.css" rel="stylesheet">
<link href="/css/custom.css" rel="stylesheet">
<script src="/js/bootstrap.min.js"></script>
</head>
<body>{{#Vars}}<h1>{{group}} last errors</h1>{{/Vars}}
<table class="table table-bordered table-condensed sortered">
        <thead>
          <tr>
            <th>Name</th>
            <th>Status</th>
            <th>Length</th>
            <th>Request Duration</th>
            <th>Last Checked</th>
            <th>Error</th>
            <th>Last Hour</th>
          </tr>
        </thead>
        <tbody>
        {{#TableData}}
          <tr class="{{severity}}">
            <td><a href="{{uri}}">{{name}}</a></td>
            <td>{{status}}</td>
            <td>{{contentlength}}</td>
            <td>{{elapsed}}</td>
            <td>{{started}}</td>
            <td>{{error}}</td>
            <td>{{totalerrs}}</td>
          </tr>
        {{/TableData}}
        </tbody>
</table>
Generated by <a href="http://github.com/grafov/hlsprobe2">HLS Probe II</a>
</body>
</html>
`
