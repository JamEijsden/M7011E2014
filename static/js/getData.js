function getLocations()
{
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        //var obj = {name:'tomte',color:'red'};
        console.log(obj);
        //document.getElementById("loadhere").innerHTML='<p>'+obj[1].position+'</p>';
        loadMap(obj);
      }
    else{
      return "TOMTE";
    }
  };
  xmlHttp.open( "GET", "http://79.136.28.106:8888/stairs", false );
  xmlHttp.send( null );
}

function getUser(id, data)
{
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        data = JSON.parse(data);
        //var obj = {name:'tomte',color:'red'};
        data.user = obj.userID;
        console.log("GET ID "+obj);
        sendForm(data);
      }
    else{
      return "TOMTE";
    }
  };
  xmlHttp.open( "GET", "http://79.136.28.106:8888/users/"+id, false );
  xmlHttp.send( null );
}