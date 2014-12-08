var map

function initialize() {
 // loadPage('map');
 var data;
  getLocations(data ,0, 0);  
}

function getLocations(recData, status, stairID)
{
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange=function() {
    if (xmlHttp.readyState==4 && xmlHttp.status==200) {
        var json = xmlHttp.responseText;
        var obj = JSON.parse(json);
        //var obj = {name:'tomte',color:'red'};
        console.log(obj);
        if(status == 0){
          loadMap(obj);  
        } else if(status == 1){
          //pass along locations from LoadMap and comments for location
          createMarker(recData,obj)
        }
        //document.getElementById("loadhere").innerHTML='<p>'+obj[1].position+'</p>';
      }
    else{
      return "TOMTE";
    }
  };
  if(status == 0){
      xmlHttp.open( "GET", "http://79.136.28.106:8888/stairs", false );   
  }else if(status == 1){
      xmlHttp.open( "GET", "http://79.136.28.106:8888/comment/"+stairID, false );
  }else{
    
  }
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
  
    createAllMarkers(locations);
    
    

    google.maps.event.addListener(map, 'click', function(event) {
      createMarker(event.latLng);
    });
}

function createAllMarkers(locations){
  for (var i=0; i < locations.length; i++) {
      console.log(locations[i]);
      pos = locations[i].position.replace('(','').replace(')','');
      var bits = pos.split(/,\s*/);
      console.log(bits[0]+", "+bits[1]);
      latLng = CreateLatLngObject(bits[0], bits[1]);
      locations[i].position = latLng;
      getLocations(locations[i], 1, locations[i].id);
    }
}