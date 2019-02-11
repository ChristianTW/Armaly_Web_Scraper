<html>
    <head>
    <title>Armaly's Scraper</title>
    </head>
    <body>
    <h1>Armaly's Scraper Search Results</h1>
        <!-- https://gowebexamples.com/templates/ -->
        <h2>Feed Results</h2>
        <ul>
            <form action="/scraper" method="post">
               {{range $i, $v := .}}
                   <li>{{$v.Title}}  </li> <input type="button" name="MoreInfo" value="{{$i}}">
               {{end}}
            </form>
        </ul>
    </body>
</html>