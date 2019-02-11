<html>
    <head>
    <title>Armaly's Scraper</title>
    </head>
    <body>
    <h1>Armaly's Scraper</h1>
        <form action="/search" method="post">
            <p>Put in 1 or 2 key words to include or exclude by!</p>
            Key Word #1: <br>
            <input type ="text" name="keyWord1"> <br>
            <input type="radio" name="filter1" value="1" checked>Include
            <input type="radio" name="filter1" value="2" checked>Exclude<br>


            Key Word #2 <br>
             <input type ="text" name="keyWord2"><br>
             <input type="radio" name="filter2" value="1" checked>Include
             <input type="radio" name="filter2" value="2" checked>Exclude<br>

             <input type="submit" value="Submit">
            <input type="reset">
        </form>

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