package epochdisplay

import (
	"html/template"
	"net/http"
)

type TemplateData struct {
	Req *http.Request
}

func (t TemplateData) Clocks() []string {
	return t.Req.URL.Query()["clock"]
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	td := TemplateData{r}
	t := template.New("Main")
	t, _ = t.Parse(mainHTML)
	t.Execute(w, td)
}

const mainHTML = `
<html>
<head>
	<title>Digital Clock</title>
	<script src="static/moment.js"></script>
	<script src="static/moment-timezone-with-data-2010-2020.min.js"></script>
	<style>
	#layout {
	}
	#header {
		display: flex;
		justify-content: flex-end;
		align-items: center;
		flex-direction: row;
		flex-wrap: nowrap;
	}
	#main {
		color:black;font-family: "Courier New", Monospace, Sans-serif;font-size: 12vw;
		display: flex;
		justify-content: center;
		align-items: center;
		flex-direction: row;
		flex-wrap: wrap;
		height: 100%;
	}
    #panel {
		position: absolute;
		right: 0;
		width: 30%;
		top: 0;
		bottom: 0;
		box-shadow: 0px 0px 15px black;
		background: linear-gradient(to bottom, #cb60b3 0%,#ad1283 50%,#de47ac 100%);
		padding: 10px;
		outline: 0;
		transform: scaleX(0.00001);
		transform-origin: 100% 50%;
		-webkit-transform-origin: 100% 50%;
		transition: transform 0.3s ease-in-out;
	}

	#panel:target {
		transform: scaleX(1);
	}

	.panelfont {
		padding: 10px;
		color: white;
		font-family: ‘Palatino Linotype’, ‘Book Antiqua’, Palatino, serif;
	}
	</style>
</head>
<body>
	<script>
	  var zones = []
	  {{ range $index, $timezone := .Clocks}}
	    zones.push({{$timezone}})
	  {{end}}
  	  function digclock() {
		var clocks = document.getElementsByClassName("clock")
		for (var i = 0; i < clocks.length; i++) {
		   var result = "Error"
		   if (zones[i].toLowerCase() == "epoch" || zones[i] == "") {
			   var date = new Date()
			   result = date.getTime()/1000 | 0
		   } else {
			   result = moment().tz(zones[i]).format("HH:mm:ss z")
		   }
		   clocks[i].innerHTML = result
		}
	  }
	  setInterval(function(){digclock()},100)
	</script>
	<div id="layout">
	  <div id="header">
		  <nav><a href="#panel">&#9776;</a></nav>
	  </div>
	  <div id="panel">
		  <nav><a href="#">&#9776;</a></nav>
          <div class="panelfont">
		  Query parameters:<br>
		  <ul>
		    <li>clock=[timezone|epoc|town]</li>
		  </ul>
		  Example:<br>
		  &clock=pdt&clock=epoc&clock=paris
		  </div>
	  </div>
	  <div id="main">
	    {{ range .Clocks }}
	      <div class="clock"></div>
	    {{ end }}
	  </div>
	</div>
</body>
</html>`
