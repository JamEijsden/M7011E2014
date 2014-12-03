function getLocations(){
//'{"userID":1,"firstName":"jimmiie","lastName":"van eijsden","idToken":"12983682682"}'
    locations = httpGet();
    document.getElementById("loadhere").innerHTML='<p>'+locations[1].stairname+'</p>';
    return locations;
    

}

function httpGet()
{
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        console.log(obj);
        document.getElementById("loadhere").innerHTML='<p>'+obj[1].position+'</p>';
        return obj;
      }
    else{
      return "TOMTE";
    }
  };
  xmlHttp.open( "GET", "http://79.136.28.106:8888/stairs", false );
  xmlHttp.send( null );
  
}