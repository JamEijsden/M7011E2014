


function placeMarker(location) {
  if(markerPlaced == 0){
    var marker = new google.maps.Marker({
                position:location,
                map: map,
                draggable:true,
                animation: google.maps.Animation.DROP,
                title:"New Marker",
              icon: "http://maps.google.com/mapfiles/ms/micons/blue.png"
            });
  } 
    
}