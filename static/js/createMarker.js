var currentMark

function createMarker(location) {
  var marker = new google.maps.Marker({
      position: location,
      map: map,
      title: ""+location
  });   
 
  contentString = "<h4>"+marker.title+"</h4>"+
  "<img src='http://media.bymk.se/2014/06/katten_jarvis.jpg' alt='Smileyface' height='300' width='400'><input type='button' class='btn btn-primary' onclick='getJson();' value=Add Location'>";

  var infowindow = new google.maps.InfoWindow({
    content: contentString
  });
  
  google.maps.event.addListener(marker, 'click', function () {
          currentMark = marker;
          map.setCenter(currentMark.position)
          infowindow.open(map, marker);


});
google.maps.event.addListener(infowindow,'closeclick',function(){
   map.setCenter(currentMark.position) 
   //removes the marker
   // then, remove the infowindows name from the array
});

}
function getJson(){
//'{"userID":1,"firstName":"jimmiie","lastName":"van eijsden","idToken":"12983682682"}'
    var response = httpGet();
    document.getElementById("loadhere").innerHTML=response;

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
        return obj;
      }
    else{
      return "TOMTE";
    }
  } 
  xmlHttp.open( "GET", "http://79.136.28.106:8888/users/1", false );
  xmlHttp.send( null );
  
}