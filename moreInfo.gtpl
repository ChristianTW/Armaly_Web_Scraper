<html>
    <head>
    <title>Armaly's Scraper</title>
    </head>
    <body>
    <h1>More Info</h1>
        <!-- https://gowebexamples.com/templates/ -->
        <h2>Feed Results</h2>
        <table>
               <th>Location</th>
               <th>Title</th>
               <th>Company</th>
               <th>Posting Date</th>
                <tr>
                    <td>{{.Location}}</td>
                   <td>{{.Title}}</td>
                   <td>{{.Company}}</td>
                   <td>{{.Date}}</td>
                </tr>
        </table>

        <h2>Description</h2>
        {{.Description}}
    </body>
</html>