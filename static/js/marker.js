var markerPlaced = 0
function placeMarker(location) {
  if(markerPlaced == 0){
    var marker = new google.maps.Marker({
                position:location,
                map: map,
                draggable:true,
                animation: google.maps.Animation.DROP,
                title:"Stair!",
              icon: "http://maps.google.com/mapfiles/ms/micons/blue.png"
            });
  }
    
    markerPlaced = 1
}