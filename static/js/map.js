var map

function initialize() {
 // loadPage('map');
  getLocations();  
}

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


//google.maps.event.addDomListener(window, 'load', initialize);
function CreateLatLngObject(Latitude, Longitude) {
  var latlng = new google.maps.LatLng(parseFloat(Latitude), parseFloat(Longitude));
  return latlng;
}
    
function loadMap(locations){
  var mapCanvas = document.getElementById('main_map_canvas');
    var myLatlng = new google.maps.LatLng(-25.363882,131.044922);
    var mapOptions = {
      zoom: 4,
      center: myLatlng,
      mapTypeId: google.maps.MapTypeId.ROADMAP
    }
    map = new google.maps.Map(mapCanvas, mapOptions);
    document.getElementById("loadhere").innerHTML='<p>'+locations[0].stairname+'</p>';
  
    for (var i=0; i < locations.length; i++) {
      pos = locations[i].position.replace('(','').replace(')','');
      var bits = pos.split(/,\s*/);
      console.log(bits[0]+", "+bits[1]);
      latLng = CreateLatLngObject(bits[0], bits[1]);
      createMarker(latLng);
    }
    
    

    google.maps.event.addListener(map, 'click', function(event) {
      createMarker(event.latLng);
    });
}
